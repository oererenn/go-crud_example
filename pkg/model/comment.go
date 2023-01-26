package model

type Comment struct {
	ID        string `json:"id,omitempty"`
	UserId    int64  `json:"userId,omitempty"`
	UserName  string `json:"userName,omitempty"`
	Content   string `json:"content,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

type CommentDTO struct {
	ID       string `json:"id,omitempty"`
	UserId   int64  `json:"userId,omitempty"`
	Content  string `json:"content,omitempty"`
	UserName string `json:"userName,omitempty"`
}

type CreateCommentDTO struct {
	UserId   int64  `json:"userId,omitempty"`
	Content  string `json:"content,omitempty"`
	UserName string `json:"userName,omitempty"`
}
