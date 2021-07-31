package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Username            string              `bson:"username"`
	Password            string              `bson:"password"`
	Genre               []string            `bson:"genre"`
	JoinedDateTimestamp primitive.Timestamp `bson:"joinedOn"`
	GenreScore          map[string]int      `bson:"genrescores"`
	Posts               []Post              `bson:"posts"`
}
