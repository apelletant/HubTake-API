package endpoints

import (
    "github.com/jinzhu/gorm"
)

type User struct {
    UserId string 		`gorm:"primary_key;unique"`
    UserFirstName string	`gorm:"type:varchar"`
    UserLastName string		`gorm:"type:varchar"`
    UserObjectId int		`gorm:"type:int"`
    UserHasObject int		`gorm:"type:int"`
}

func (e *Endpoints) GetUsers(db *gorm.DB) []User {
    var users []User
    db.Find(&users)
    return users
    }

func (e *Endpoints) GetUserByMail(db *gorm.DB, mail string) User {
    var user = User{}
    db.Where("user_mail = ?", mail).First(&user)
    return user
}

func (e *Endpoints) GetUserHasObject(db *gorm.DB) []User {
    var u []User
    db.Where("user_object_id = ?",  1).Find(&u)
    return u
}

//POST USER
func (e *Endpoints) AddUser(db *gorm.DB, userData User) (User, error) {
    return userData, nil
}
