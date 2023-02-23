package controllers

import (
	"fmt"
)

func UpdatePrice(Secid string, Prevprice float64) {
	db.Table("instruments").Where("ticker = ?", Secid).Update("price", Prevprice)
}

func SaveParice(prices Prices) {

	//	err := db.Table("prices").Create(Prices).Error
	//err := db.Table("prices").Clauses(clause.Insert{Modifier: "IGNORE"}).Create(prices)
	if db.Table("prices").Set("gorm:insert_modifier", "IGNORE").Create(prices).Error != nil {
		fmt.Print("Should ignore duplicate user insert by insert modifier:IGNORE ")
	}

	//	fmt.Print(prices)

}

// для записи в базу
type Prices struct {
	//	EventID      int    `gorm:"column:event_id;not_null;primary_key;auto_increment"`
	Price float64 `gorm:"column:price;"`
	Date  string  `gorm:"column:date"`
	Name  string  `gorm:"column:name"`
	Value float64 `gorm:"column:value"`
}
