package models

type Prices struct {
	Price float64 `gorm:"column:price" json:"price"`
	Date  string  `gorm:"column:date" json:"date"`
	//	Event *New    `gorm:"foreignKey:date;references:date" json:"event"`
	Slug    string `gorm:"column:slug" json:"slug"`
	Title   string `gorm:"column:title" json:"title"`
	TypeID  int    `json:"typeId"`
	Hash    string `json:"hash"`
	EventId string `json:"event_id"`

	Source    string `json:"source"`
	Shorttext string `json:"shorttext"`
}

type PricesDashBord struct {
	Price float64 `gorm:"column:price" json:"price"`
	Date  string  `gorm:"column:date" json:"date"`
	Name  string  `gorm:"column:name" json:"name"`
}

// "price, prices.date, events.title, events.slug, events.hash, events.type_id, events.event_id,source,instrument_id,shorttext"
