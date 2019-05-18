package defs

//request
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}

//response
type SignedUp struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}



// Data model
type VideoInfo struct {
	VideoId string
	AuthorId int
	Name string
	CreateTime	string
}

type Comment struct {
	Id string
	VideoId string
	AuthorName string
	Content string
}

// 简单的session
type SimpleSession struct {
	Username string // login name
	TTL int64 // 过期时间
}