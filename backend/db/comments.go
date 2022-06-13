package db

import (
	"log"

	"github.com/vitkarpov/ValushaJS/backend/models"
)

func (db Database) GetComments() *models.GetComments {
	response := &models.GetComments{}
	rows, err := db.Conn.Query("SELECT * FROM comments ORDER BY created_at DESC")
	response.Error = err.Error()
	if err != nil {
		return response
	}
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.ID, &comment.ParentID, &comment.AuthorName, &comment.Comment, &comment.CreatedAt)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		response.Comments = append(response.Comments, comment)
	}
	return response
}
