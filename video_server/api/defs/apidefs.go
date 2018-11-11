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
