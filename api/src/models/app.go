package models

import "time"

type App struct {
	Id          string    `bson:"id" json:"id"`
	Description string    `bson:"description" json:"description"`
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt" json:"updatedAt"`
}
