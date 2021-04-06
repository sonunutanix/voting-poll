package dao

type CreatePoll struct {
	Question string
	Options  []string
}

type OptionIdUserId struct {
	OptionId int ` binding:"required"`
	UserId int ` binding:"required"`
}

type QuestionIdUserId struct {
	QuestionId int ` binding:"required"`
	UserId int ` binding:"required"`
}
