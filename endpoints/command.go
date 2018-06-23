package endpoints

import (
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

type CommandDataComplete struct {
	CommandIDz          int    `gorm:"primary_key;unique"`
	CommandModel        string `gorm:"type:text"`
	CommandUserName     string `gorm:"type:int"`
	CommandUserLastName string `gorm:"type:int"`
	CommandPlasticColor string `gorm:"type:int"`
	CommandPrice        int    `gorm:"type:int"`
	CommandLength       int    `gorm:"type:int"`
}

func addCommand(db *gorm.DB) error {
	return nil
}

func getCommands(db *gorm.DB) ([]Command, error) {
	var com []Command
	err := db.Find(&com).Error
	if err != nil {
		return nil, err
	}
	return com, nil
}

func getCommandCustomer(db *gorm.DB) {
}

func updateCommand(db *gorm.DB) {
}

func getUnpaidCommand(db *gorm.DB) {
}
