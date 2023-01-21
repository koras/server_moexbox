package models

// список сотрудников

// type User struct {
// 	ID       uint
// 	Nickname string
// 	Stacks   Stack `json:"Stack" gorm:"foreignKey:UsersID"`
// }

type New struct {
	EventID  int    `json:"event_id"`
	Event    string `json:"event"`
	Type     string `json:"type"`
	TypeID   int    `json:"typeId"`
	Hash     string `json:"hash"`
	Source   string `json:"source"`
	URL      string `json:"url"`
	TitleURL string `json:"title_url"`

	InstrumentID int    `json:"instrumentId"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Change       string `json:"change"`
	Ticker       string `json:"ticker"`
	Currency     string `json:"currency"`

	Date     string `json:"date"`
	Title    string `json:"title"`
	Text     string `json:"text"`
	Fulltext string `json:"fulltext"`
	Link     string `json:"link"`
}

// данные которые приходят с фронта
type EventInput struct {
	//	EventID      string `json:"event_id"`
	Title        string `form:"title" binding:"required"`
	TypeID       int    `form:"type_id" binding:"required"`
	Hash         string `form:"hash" binding:"required"`
	Source       string `form:"source" binding:"required"`
	InstrumentID int    `form:"instrument_id" binding:"required"`
	Date         string `form:"date" binding:"required"`
	Shorttext    string `form:"shorttext" binding:"required"`
	Fulltext     string `form:"fulltext" binding:"required"`
}

// для записи в базу
type Events struct {
	//	EventID      string `json:"event_id"`
	Title        string `json:"title"`
	TypeID       int    `json:"type_id"`
	UserId       int    `json:"user_id"`
	Hash         string `json:"hash"`
	Source       string `json:"source"`
	Slug         string `json:"slug"`
	InstrumentID int    `json:"instrument_id"`
	Date         string `json:"date"`
	Shorttext    string `json:"shorttext"`
	Fulltext     string `form:"fulltext"`
}

// то что возвращаем
type EventResult struct {
	Hash         string `json:"hash"`
	Slug         string `json:"slug"`
	InstrumentID int    `json:"instrument_id"`
}
