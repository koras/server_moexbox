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
	Slug      string `json:"slug"`
	Date      string `json:"date"`
	Title     string `json:"title"`
	Shorttext string `json:"shorttext"`
	Fulltext  string `json:"fulltext"`
}

type EventInstrument struct {
	Data       New        `json:"data"`
	Instrument Instrument `json:"instrument"`
}

type EventsInstrument struct {
	Data       []News     `json:"data"`
	Instrument Instrument `json:"instrument"`
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

// данные которые приходят с фронта
type EventInput struct {
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

// для записи в базу
type Events struct {
	//	EventID      string `json:"event_id"`
	Title        string `json:"title"`
	TypeID       int    `json:"type_id"`
	UserId       int    `json:"user_id"`
	Hash         string `json:"hash"`
	EventID      int    `json:"event_id"`
	Source       string `json:"source"`
	Slug         string `json:"slug"`
	InstrumentID string `json:"instrument_id"`
	Date         string `json:"date"`
	Shorttext    string `json:"shorttext"`
	Fulltext     string `form:"fulltext"`
	Published    string `form:"published"`
}

// то что возвращаем
type EventResult struct {
	Hash         string `json:"hash"`
	Slug         string `json:"slug"`
	InstrumentID string `json:"instrument_id"`
}

// для записи в базу
type EventsResult struct {
	//	EventID      string `json:"event_id"`
	Title        string `json:"title"`
	TypeID       int    `json:"type_id"`
	UserId       int    `json:"user_id"`
	EventID      int    `json:"event_id"`
	Hash         string `json:"hash"`
	Source       string `json:"source"`
	Slug         string `json:"slug"`
	InstrumentID string `json:"instrument_id"`
	Date         string `json:"date"`
	Shorttext    string `json:"shorttext"`
	Fulltext     string `form:"fulltext"`
	Published    int
}
