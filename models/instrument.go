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
	Title        string `json:"title" binding:"required"`
	TypeID       int    `json:"type_id" binding:"required"`
	Hash         string `json:"hash" binding:"required"`
	Source       string `json:"source" binding:"required"`
	InstrumentID string `json:"instrument_id" binding:"required"`
	Date         string `json:"date" binding:"required"`
	Shorttext    string `json:"shorttext" binding:"required"`
	Fulltext     string `json:"fulltext" binding:"required"`
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
	InstrumentID   string `json:"instrument_id"`
	InstrumentName string `json:"instrument_name"`
	Description    string `json:"description"`
	Type           string `json:"type"`
	Ticker         string `json:"ticker"`
	Price          int    `json:"price"`
	Site           string `json:"site"`
	Currency       string `json:"currency"`
	Logo           string `json:"logo"`
}
