package model

type User struct {
	Username string   `bson:"username" json:"username"`
	Password string   `bson:"password" json:"password"`
	Genre    []string `bson:"genre" json:"genre`
	// JoinedDateTimestamp primitive.Timestamp `bson:"joinedOn"`
	GenreScore map[string]int `bson:"genrescores" json:"genrescores"`
	// Posts               []Post              `bson:"posts"`
}
