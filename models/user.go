package models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}

type Polls struct {
	Id       uint   `json:"id"`
	Question string `json:"question"`
}

type Options struct {
	Id     uint   `json:"id"`
	Option string `json:"option"`
	PollId uint   `json:"pollId"`
	Votes  int    `gorm:"default:0"`
}

type OptionUser struct {
    Id     uint   `json:"id"`
	OptionId int `json:"optionId"`
	UserId int   `json:"userId"`
}
