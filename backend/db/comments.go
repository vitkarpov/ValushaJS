package db

import (
	"github.com/vitkarpov/ValushaJS/backend/models"
)

func (db Database) GetComments() *models.GetComments {
	response := &models.GetComments{}
	rows, err := db.Conn.Query("SELECT * FROM comments ORDER BY CreatedAt DESC")
	response.Error = err.Error()
	if err != nil {
		return response
	}
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.ID, &comment.AuthorName, &comment.Comment, &comment.ParentID, &comment.CreatedAt)
		response.Error = err.Error()
		if err != nil {
			return response
		}
		response.Comments = append(response.Comments, comment)
	}
	return response
}
