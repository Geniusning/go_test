package defs

type UserInfo struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

type VideoInfo struct {
	ID           string
	AuthorID     int
	Name         string
	DisplayCtime string
}

type Comment struct {
	ID      string
	VideoID string
	Author  string
	Content string
}
