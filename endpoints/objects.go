package endpoints

import (
	"github.com/jinzhu/gorm"
	"time"
	"errors"
)

type Endpoints struct {
	db      *gorm.DB
}

type Object struct {
	ObjectId string 		`gorm:"primary_key;unique"`
	ObjectName string		`gorm:"type:text"`
	ObjectIsTaken int		`gorm:"type:int"`
	ObjectDateBorrow time.Time	`gorm:"type:date"`
	ObjectDateReturn time.Time	`gorm:"type:date"`
}

//NewEndpoint create a new EndPoints object
func (e *Endpoints) NewEndpoints(db *gorm.DB) *Endpoints {
	return &Endpoints{db: db}
}

func (e *Endpoints) AddObject(db *gorm.DB, objectName string) (Object, error) {
	object := Object{
		ObjectName: objectName,
		ObjectIsTaken: 0,
	}
	db.Create(&object)
	return object, nil
}

func (e *Endpoints) GetObjects(db *gorm.DB) []Object {
	var objects []Object
    	db.Find(&objects)
	return objects
}

func (e *Endpoints) GetObjectByName(db *gorm.DB, objName string) Object {
    var obj = Object{}
    db.Where("object_name = ?", objName).First(&obj)
    return obj
}

func (e *Endpoints) GetTakenObject(db *gorm.DB) []Object {
    var objects []Object
    db.Where("object_is_taken = ?", 1).Find(&objects)
    return objects
}

func (e *Endpoints) GetNotTakenObject(db *gorm.DB) []Object {
    var objects []Object
    db.Where("object_is_taken = ?", 0).Find(&objects)
    return objects
}


func (e *Endpoints) DeleteObject(db *gorm.DB, name string) error {
	var obj = Object{}
	err := db.Where("object_name = ?", name).First(&obj).Error
	if err != nil {
		return errors.New("Could'nt remove object")
	}
	id := obj.ObjectId
	return db.Delete(&obj, id).Error
}
