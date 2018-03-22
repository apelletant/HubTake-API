package endpoints

import "github.com/jinzhu/gorm"

type BorrowReturnData struct {
    UserMail string
    ObjectName string
}

func (e *Endpoints) UserReturnObject(db *gorm.DB, data BorrowReturnData) error {
    return nil
}

func (e *Endpoints) UserTakeObject(db *gorm.DB, data BorrowReturnData) error {
    return nil
}

/*
func (e *Endpoints) UserObjectData(db *gorm.DB, email string) (ep.Object{} int, error) {
	return nil, 1, nil
}*/
