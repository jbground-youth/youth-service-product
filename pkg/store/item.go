package store

import (
	"gorm.io/gorm"
	"time"
)

type Item struct {
	gorm.Model
	Id       int64     `gorm:"primary_key" gorm:"column:id"`
	CreateAt time.Time `gorm:"column:created_at"`
	UpdateAt time.Time `gorm:"column:updated_at"`
}

func GetAllItems() {

}

func GetItemById(id int) (*Item, error) {
	var item Item
	err := db.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func CreateItem() error {
	return nil
}

func UpdateItem() {

}

func DeleteItem() {

}
