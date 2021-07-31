package model

type Post struct {
	PostId           string
	PostContent      string
	ImageUrl         string
	CreatedTimestamp string
	PostScore        int64
	PostUpVotes      int64
	PostDownVotes    int64
	Genres           []Genre
	Votes            []Vote
}
