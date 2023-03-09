package models

// список сотрудников

// type User struct {
// 	ID       uint
// 	Nickname string
// 	Stacks   Stack `json:"Stack" gorm:"foreignKey:UsersID"`
// }

// данные которые приходят с фронта
type InstrumentInput struct {
	//	EventID      string `json:"event_id"`
	Site               string `form:"site" binding:"required"`
	InstrumentName     string `form:"instrument_name" binding:"required"`
	InstrumentFullName string `form:"instrument_full_name" binding:"required"`
	Description        string `form:"description" binding:"required"`
	InstrumentID       string `form:"instrument_id" binding:"required"`
	IndustryID         string `form:"industry_id" binding:"required"`
	Logo               string
}

// "instrumentId": 11411,
// "name": "Биткоин",
// "type": "crypto",
// "ticker": "btc",
// "price": 130,
// "change": "-10",
// "currency": "$"
// для записи в базу
type Instrument struct {
	InstrumentID        string `json:"instrument_id"`
	InstrumentName      string `json:"instrument_name"`
	Instrument_FullName string `json:"instrument_full_name"`
	INSTRUMENT_CATEGORY string `json:"INSTRUMENT_CATEGORY"`
	LIST_SECTION        string `json:"LIST_SECTION"`

	CURRENCY_MOEX string  `json:"CURRENCY_MOEX"`
	Description   string  `json:"description"`
	Type          string  `json:"type"`
	Ticker        string  `json:"ticker"`
	Price         float64 `json:"price"`
	Mark          string  `json:"mark"`
	Isin          string  `json:"isin"`
	Site          string  `json:"site"`
	Currency      string  `json:"currency"`
	Logo          string  `json:"logo"`
}

type PricesInstrument struct {
	Instruments Instrument `json:"instrument"`
	Prices      []Prices   `json:"price"`
}

// "price, prices.date, events.title, events.slug, events.hash, events.type_id, events.event_id,source,instrument_id,shorttext"
