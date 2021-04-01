package dto

import "Project/models"

type Questions struct {
	Id       int
	Question string
	Options  []models.Options
}
