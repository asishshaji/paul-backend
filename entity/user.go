package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	UserId              primitive.ObjectID  `bson:"_id"`
	Username            string              `bson:"username"`
	Password            string              `bson:"password"`
	Genre               []Genre             `bson:"genre"`
	JoinedDateTimestamp primitive.Timestamp `bson:"joinedOn"`
	GenreScore          []string            `bson:"scores"`
	Posts               []Post              `bson:"posts"`
}
