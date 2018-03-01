package endpoints

import "github.com/jinzhu/gorm"

type BorrowReturnData struct {
    UserMail string
    ObjectName string
}

func (e *Endpoints) userReturnObject(db *gorm.DB, data BorrowReturnData) error {
    return nil
}

func (e *Endpoints) UserTakeObject(db *gorm.DB, data BorrowReturnData) error {
    return nil
}
