package endpoints

import (
	"github.com/jinzhu/gorm"
	"errors"
	"fmt"
)

type User struct {
	UserId string			`gorm:"primary_key;unique"`
	UserFirstName string		`gorm:"type:text"`
	UserLastName string		`gorm:"type:text"`
	UserEmail string		`gorm:"type:text"`
	UserObjectId int		`gorm:"type:int, default:0"`
	UserHasObject int		`gorm:"type:int, default:0"`
	UserPromo int			`gorm:"type:int, default:1994"`
}

type UserPost struct {
	UserFirstName string
	UserLastName string
	UserEmail string
	UserPromo int
}

type  UserObject struct {
	UserData User
	ObjectData Object
}

func (e *Endpoints) GetUsers(db *gorm.DB) []User {
	var users []User
	db.Find(&users)
	return users
    }

func (e *Endpoints) GetUserByMail(db *gorm.DB, mail string) User {
	var user = User{}
	db.Where("user_email = ?", mail).First(&user)
	return user
}

func (e *Endpoints) GetUserHasObject(db *gorm.DB) []User {
	var u []User
	db.Where("user_has_object = ?",  1).Find(&u)
	return u
}

func (e *Endpoints) AddUser(db *gorm.DB, userData UserPost) (UserPost, error) {
	var user User

	user.UserFirstName = userData.UserFirstName
	user.UserLastName = userData.UserLastName
	user.UserEmail = userData.UserEmail
	user.UserPromo = userData.UserPromo
		db.Create(&user)
	return userData, nil
}

func (e *Endpoints) DeleteUser(db *gorm.DB, mail string) error {
	var user = User{}
	err := db.Where("user_email = ?", mail).First(&user).Error
	if err != nil {
		return errors.New("Could not find user to delete")
	}
	id := user.UserId;
	return db.Delete(&user, id).Error
}

func (e *Endpoints) GetUserObjectData(db *gorm.DB) ([]UserObject, error) {
	var dataObjUsr = []UserObject{}
	var userObj = UserObject{}

	users := e.GetUserHasObject(db)
	if len(users) == 0 {
		return nil, fmt.Errorf("no users got object")
	}
	for _, user := range users {
		obj := e.GetObjectById(db, user.UserObjectId)
		userObj.UserData = user
		userObj.ObjectData = obj
		dataObjUsr = append(dataObjUsr, userObj)
	}
	return dataObjUsr, nil
}
