package defs

type UserInfo struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

//response
type SignedUp struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
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

type SimpleSession struct {
	Username string
	TTL      int64
}
