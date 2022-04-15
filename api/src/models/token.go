package models

import "time"

type Token struct {
	Token     string    `bson:"token" json:"token"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	ExpiresOn time.Time `bson:"expiresOn" json:"expiresOn"`
}
