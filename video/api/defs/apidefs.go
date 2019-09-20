package defs

/**
	用户信息
 */
type UserCredential struct {
	UserName string `json:"user_name"`
	Pwd string `json:"pwd"`
}

/**
	视频信息
 */
type VideoInfo struct {
	Id string
	AuthorID int
	Name string
	DisplayCtime string
}

/**
	评论信息
 */
type Comment struct {
	ID string
	VideoID string
	AuthorName string
	Content string
}

/**
	session信息
 */
type Session struct {
	SessionID string
	TTL int64
	LoginName string
}

/**
	成功添加用户的返回信息
 */
type SignedUp struct {
	Success bool
	SessionId string
}