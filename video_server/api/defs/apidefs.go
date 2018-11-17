package defs

//requsts
type UserCredential struct {
	Username string `json:"username"`
	pwd      string `json:pwd`
}

//DB model
type VideoInfo struct {
	Id          string
	AuthorId    int
	Name        string
	DisplayTime string
}

type Comment struct {
	Id      string
	VideoId string
	Author  string
	Content string
}

//Session DB
type SimpleSession struct {
	Username string
	TTL      int64
}
