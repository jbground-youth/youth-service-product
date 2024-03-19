package store

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Id       int64     `gorm:"primary_key" gorm:"column:id"`
	CreateAt time.Time `gorm:"column:created_at"`
	UpdateAt time.Time `gorm:"column:updated_at"`
}

func GetAllProducts() {

}

func GetProduct() {

}

func CreateProduct() {

}

func UpdateProduct() {

}

func DeleteProduct() {

}

func CleanAllProducts() error {

	return nil
}
