package models

type Comment struct {
	ID         int    `json:"id"`
	ParentID   int    `json:"parentId"`
	AuthorName string `json:"authorName"`
	Comment    string `json:"comment"`
	CreatedAt  string `json:"createdAt"`
}

type GetComments struct {
	Error    string    `json:"error"`
	Comments []Comment `json:"comments"`
}
