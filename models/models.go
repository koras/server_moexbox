package models

// список сотрудников

// type User struct {
// 	ID       uint
// 	Nickname string
// 	Stacks   Stack `json:"Stack" gorm:"foreignKey:UsersID"`
// }

type New struct {
	EventID   int    `json:"eventId"`
	TypeID    int    `json:"typeId"`
	Hash      string `json:"hash"`
	Source    string `json:"source"`
	Slug      string `gorm:"column:slug" json:"slug"`
	Date      string `gorm:"column:date" json:"date"`
	Title     string `gorm:"column:title" json:"title"`
	Shorttext string `json:"shorttext"`
	Fulltext  string `json:"fulltext"`
}

type EventInstrument struct {
	Data       New         `json:"data"`
	Instrument Instrument  `json:"instrument"`
	PriceDate  []PriceDate `json:"date"`
}

type NewsInstrument struct {
	Data       New        `json:"data"`
	Instrument Instrument `json:"instrument"`
	PriceDate  []Prices   `json:"price"`
}

type EventsInstrument struct {
	Data       []News      `json:"data"`
	Instrument Instrument  `json:"instrument"`
	PriceDate  []PriceDate `json:"date"`
}

type News struct {
	EventID   int    `json:"event_id"`
	TypeID    int    `json:"typeId"`
	Hash      string `json:"hash"`
	Source    string `json:"source"`
	Slug      string `json:"slug"`
	Date      string `json:"date"`
	Title     string `json:"title"`
	Shorttext string `json:"shorttext"`
	Fulltext  string `json:"fulltext"`
}

type SitemapNews struct {
	TypeID   int    `json:"typeId"`
	Slug     string `json:"slug"`
	Date     string `json:"date"`
	Title    string `json:"title"`
	CreateAt string `json:"create_at"`
	//	Shorttext string `json:"shorttext"`
	//	Fulltext  string `json:"fulltext"`
	Ticker string `json:"ticker"`
}

// данные которые приходят с фронта
type EventInput struct {
	//	EventID      string `json:"event_id"`
	Title        string `json:"title"  binding:"required,min=25,max=250"`
	TypeID       int    `json:"typeId"  binding:"required,min=1"`
	Hash         string `json:"hash"`
	Source       string `json:"source" binding:"required,min=5,max=250"`
	InstrumentID string `json:"instrument_id"`
	Date         string `json:"date" binding:"required,min=10,max=11"`
	Shorttext    string `json:"shorttext" binding:"required,min=100,max=250"`
	Fulltext     string `json:"fulltext" binding:"required,min=255,max=3500"`
}

// для записи в базу
type Events struct {
	//	EventID      string `json:"event_id"`
	Title         string `json:"title"`
	TypeID        int    `json:"type_id"`
	UserId        int    `json:"user_id"`
	Hash          string `json:"hash"`
	EventID       int    `json:"event_id"`
	Source        string `json:"source"`
	Slug          string `json:"slug"`
	InstrumentID  string `json:"instrument_id"`
	Date          string `json:"date"`
	Shorttext     string `json:"shorttext"`
	ParentEventId string `json:"parent_event_id"`
	Fulltext      string `json:"fulltext"`
	Published     string `json:"published"`
}

// то что возвращаем
type EventResult struct {
	Hash         string `json:"hash"`
	Slug         string `json:"slug"`
	Status       string `json:"status"`
	InstrumentID string `json:"instrument_id"`
}

// то что возвращаем
type PriceDate struct {
	Date string `gorm:"column:date"`
}

// для записи в базу
type EventsResult struct {
	//	EventID      int    `gorm:"column:event_id;not_null;primary_key;auto_increment"`
	Title        string `gorm:"column:title;"`
	TypeID       int    `gorm:"column:type_id"`
	UserId       int    `gorm:"column:user_id"`
	Hash         string `gorm:"column:hash"`
	Source       string `gorm:"column:source"`
	Slug         string `gorm:"column:slug"`
	InstrumentID string `gorm:"column:instrument_id"`
	Date         string `gorm:"column:date"`
	Shorttext    string `gorm:"column:shorttext"`
	Fulltext     string `gorm:"column:fulltext"`
	Published    string `gorm:"column:published"`
}
