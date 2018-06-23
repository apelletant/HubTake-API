package endpoints

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//Plastic definition of Plastic model
type Plastic struct {
	PlasticID    int    `gorm:"primary_key;unique"`
	PlasticColor string `gorm:"type:text"`
	PlasticPrice string `gorm:"type:int"`
}

//AddPlastic ...
func (e *Endpoints) AddPlastic(db *gorm.DB, plastic Plastic) (Plastic, error) {
	err := db.Create(&plastic).Error
	if err != nil {
		return plastic, err
	}
	return plastic, nil
}

//UpdatePlastic ...
func (e *Endpoints) UpdatePlastic(db *gorm.DB, plastic Plastic) (Plastic, error) {
	var plas Plastic
	err := db.Where("plastic_color = ?", plastic.PlasticColor).Find(&plas).Error
	if err != nil {
		return plas, fmt.Errorf("db.Where: %v", err)
	}
	plastic.PlasticID = plas.PlasticID
	err2 := db.Model(&plastic).Where("plastic_id = ?", plastic.PlasticID).Update("plastic_price", plastic.PlasticPrice).Error
	if err2 != nil {
		return plastic, fmt.Errorf("db.Where: %v", err2)
	}
	return plastic, nil
}

//GetPlastics ...
func (e *Endpoints) GetPlastics(db *gorm.DB) ([]Plastic, error) {
	var pla []Plastic
	if err := db.Find(&pla).Error; err != nil {
		return nil, err
	}
	return pla, nil
}

//GetPlasticByColor ...
func (e *Endpoints) GetPlasticByColor(db *gorm.DB, color string) (Plastic, error) {
	var pla Plastic
	if err := db.Where("plastic_color = ?", color).First(&pla).Error; err != nil {
		return pla, err
	}
	return pla, nil
}
