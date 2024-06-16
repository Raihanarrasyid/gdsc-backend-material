package helper

type RegisterUserRequest struct {
	Username string `json:"username"`
}

type AddNoteRequest struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int64  `json:"user_id"`
}