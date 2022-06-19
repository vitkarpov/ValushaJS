package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/vitkarpov/ValushaJS/backend/db"
	"github.com/vitkarpov/ValushaJS/backend/models"
)

func TestGetComments(t *testing.T) {
	gin.SetMode(gin.TestMode)

	conn, mock, err := sqlmock.New()
	var dbInstance db.Database
	dbInstance.Conn = conn
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer conn.Close()

	timestamp := time.Now().Format(time.UnixDate)

	rows := sqlmock.NewRows([]string{"id", "parent_id", "author_name", "comment", "created_at"}).
		AddRow(1, -1, "Viktor", "hello", timestamp).
		AddRow(2, -1, "Well", "hello", timestamp)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM comments ORDER BY created_at DESC")).WillReturnRows(rows)

	router := gin.Default()
	router.Use(db.SetContextMiddlware(dbInstance))
	GetCommentsHandler(router)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/comments", nil)
	req.Header.Add("Accept", "application/json")

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected error code %v got %v", http.StatusOK, w.Code)
	}
	data, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	expected, _ := json.Marshal(models.GetComments{
		Error: "",
		Comments: []models.Comment{
			{ID: 1, ParentID: -1, AuthorName: "Viktor", Comment: "hello", CreatedAt: timestamp},
			{ID: 2, ParentID: -1, AuthorName: "Well", Comment: "hello", CreatedAt: timestamp},
		},
	})
	if err != nil {
		t.Fatalf("expected error to be nil got %v", err)
	}

	responseJson := string(data)
	expectedJson := string(expected)

	if responseJson != expectedJson {
		t.Errorf("expected body to be\n %v \n got \n %v", expectedJson, responseJson)
	}
}
