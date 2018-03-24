package endpoints

import (
	"time"
	"strconv"

	"github.com/jinzhu/gorm"
)

type BorrowReturnData struct {
    UserEmail string
    ObjectName string
}


type UserObjectData struct {
	UserData User
	ObjectData Object
}

func (e *Endpoints) UserReturnObject(db *gorm.DB, data BorrowReturnData) error {
    return nil
}

func (e *Endpoints) UserTakeObject(db *gorm.DB, data BorrowReturnData) error {
	var userTemp = User{}
	var objTemp = Object{}

	now := time.Now()
	retDate := now.AddDate(0,0,2)
	user := e.GetUserByMail(db, data.UserEmail)
	obj := e.GetObjectByName(db, data.ObjectName)
	i, _ := strconv.Atoi(obj.ObjectId)

	user.UserHasObject = 1
	user.UserObjectId = i

	obj.ObjectIsTaken = 1
	obj.ObjectDateBorrow = now
	obj.ObjectDateReturn = retDate

	db.Model(&userTemp).Where("user_id = ?", user.UserId).Updates(User{UserObjectId: user.UserObjectId, UserHasObject: user.UserHasObject})
	db.Model(&objTemp).Where("object_id = ?", obj.ObjectId).Updates(Object{ObjectIsTaken: obj.ObjectIsTaken, ObjectDateBorrow: obj.ObjectDateBorrow, ObjectDateReturn: obj.ObjectDateReturn})
	return nil
}
