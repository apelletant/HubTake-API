package endpoints

import (
	"strconv"

	"github.com/jinzhu/gorm"
)

//Command definition of Command model
type Command struct {
	CommandID        int    `gorm:"primary_key;unique"`
	CommandModel     string `gorm:"type:text"`
	CommandIDUser    int    `gorm:"type:int"`
	CommandIDPlastic int    `gorm:"type:int"`
	CommandPrice     int    `gorm:"type:int"`
	CommandLength    int    `gorm:"type:int"`
}

//CommandDataComplete definition of CommandDataComplete model
type CommandDataComplete struct {
	CommandID           int    `gorm:"primary_key;unique"`
	CommandModel        string `gorm:"type:text"`
	CommandUserName     string `gorm:"type:int"`
	CommandUserLastName string `gorm:"type:int"`
	CommandPlasticColor string `gorm:"type:int"`
	CommandPrice        int    `gorm:"type:int"`
	CommandLength       int    `gorm:"type:int"`
}

//UserCommand definition of UserCommand model
type UserCommand struct {
	Model  string
	Price  int
	Length int
	Color  string
}

//AddCommand definition of addCommand body
type AddCommand struct {
	Email        string
	PlasticColor string
	URLModel     string
	Length       int
}

func (e *Endpoints) AddCommand(db *gorm.DB, cmdData AddCommand) (AddCommand, error) {
	var cmd Command

	user := e.GetUserByMail(db, cmdData.Email)
	pla, err := e.GetPlasticByColor(db, cmdData.PlasticColor)
	if err != nil {
		return cmdData, err
	}
	cmd.CommandModel = cmdData.URLModel
	cmd.CommandIDUser, err = strconv.Atoi(user.UserId)
	if err != nil {
		return cmdData, err
	}
	cmd.CommandIDUser = pla.PlasticID
	price, err := strconv.Atoi(pla.PlasticPrice)
	if err != nil {
		return cmdData, err
	}
	cmd.CommandPrice = cmdData.Length * price
	cmd.CommandLength = cmdData.Length
	err = db.Create(&cmd).Error
	if err != nil {
		return cmdData, err
	}
	return cmdData, nil
}

func (e *Endpoints) GetCommands(db *gorm.DB) ([]Command, error) {
	var commands []Command
	var commandData []CommandDataComplete
	err := db.Find(&commands).Error
	if err != nil {
		return nil, err
	}

	for _, command := range commands {
		plastic, err := getPlasticByID(db, command.CommandIDPlastic)
		if err != nil {
			return nil, err
		}
		user, err := getUserByID(db, command.CommandIDUser)
		if err != nil {
			return nil, err
		}

		plasPrice, err := strconv.Atoi(plastic.PlasticPrice)
		if err != nil {
			return nil, err
		}

		cmdPrice := command.CommandLength * plasPrice
		cmd := CommandDataComplete{
			CommandID:           command.CommandID,
			CommandModel:        command.CommandModel,
			CommandUserName:     user.UserFirstName,
			CommandUserLastName: user.UserLastName,
			CommandPlasticColor: plastic.PlasticColor,
			CommandPrice:        cmdPrice,
			CommandLength:       command.CommandLength,
		}
		commandData = append(commandData, cmd)
	}
	return commands, nil
}

func (e *Endpoints) GetCommandCustomer(db *gorm.DB, customerName, customerLastName string) ([]UserCommand, error) {
	var user User
	var cmd []Command
	var plastic Plastic
	var userCmd []UserCommand
	err := db.Where("user_first_name = ? AND user_last_name = ?", customerName, customerLastName).Find(&user).Error
	if err != nil {
		return nil, err
	}
	err = db.Where("command_user_id = ?", user.UserId).Find(&cmd).Error
	if err != nil {
		return nil, err
	}
	for _, command := range cmd {
		err := db.Where("plastic_id = ?", command.CommandIDPlastic).Find(&plastic).Error
		if err != nil {
			return nil, err
		}
		userCmdData := UserCommand{
			Model:  command.CommandModel,
			Price:  command.CommandPrice,
			Length: command.CommandLength,
			Color:  plastic.PlasticColor,
		}
		userCmd = append(userCmd, userCmdData)
	}
	return userCmd, nil
}

//DeleteCommand delete the given command
func (e *Endpoints) DeleteCommand(db *gorm.DB, commandID int) error {
	var cmd Command
	err := db.Where("command_id = ?", commandID).First(&cmd).Error
	if err != nil {
		return err
	}
	id := cmd.CommandID
	err = db.Delete(&cmd, id).Error
	if err != nil {
		return err
	}
	return nil
}
