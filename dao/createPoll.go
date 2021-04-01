package dao

type CreatePoll struct {
	Question string
	Options  []string
}

type OptionIdUserId struct {
	OptionId int ` binding:"required"`
	UserId int ` binding:"required"`
}
