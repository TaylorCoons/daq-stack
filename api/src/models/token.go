package models

import "time"

type Token struct {
	Key       string    `bson:"key" json:"key"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	ExpiresOn time.Time `bson:"expiresOn" json:"expiresOn"`
}
