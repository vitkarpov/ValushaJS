package models

type Comment struct {
	ID         string `json:"id"`
	AuthorName string `json:"authorName"`
	Comment    string `json:"comment"`
	ParentID   string `json:"parentId"`
	CreatedAt  string `json:"createdAt"`
}

type GetComments struct {
	Error    string    `json:"error"`
	Comments []Comment `json:"comments"`
}
