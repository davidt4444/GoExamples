package service

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
	if value == nil {
		nt.Time, nt.Valid = time.Time{}, false
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		nt.Time, nt.Valid = v, true
		return nil
	case []byte:
		t, err := time.Parse("2006-01-02 15:04:05", string(v))
		if err != nil {
			return err
		}
		nt.Time, nt.Valid = t, true
		return nil
	}
	return fmt.Errorf("Failed to scan time: %v", value)
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

// Post represents the structure of the post in the database
type Post struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	CreatedAt   NullTime `json:"createdAt"`
	Author      *string  `json:"author,omitempty"`
	Category    *string  `json:"category,omitempty"`
	UpdatedAt   NullTime `json:"updatedAt,omitempty"`
	LikesCount  int      `json:"likesCount"`
	AuthorId    *int     `json:"authorId,omitempty"`
	IsPublished bool     `json:"isPublished"`
	Views       int      `json:"views"`
}

var (
	// Database connection is established here
	db *sql.DB
)

func init() {
	var err error
	// Initialize database connection
	data, err := os.ReadFile("../../aws-resources/localhost-mac-go.txt")

	// db, err = sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/dbname")
	db, err = sql.Open("mysql", string(data))
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}

// CreatePost adds a new post to the database
func CreatePost(post *Post) (int64, error) {
	result, err := db.Exec("INSERT INTO posts (title, content, author, category, authorId, isPublished, createdAt) VALUES (?, ?, ?, ?, ?, ?, ?)",
		post.Title, post.Content, post.Author, post.Category, post.AuthorId, post.IsPublished, post.CreatedAt.Time)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// GetAllPosts retrieves all posts from the database
func GetAllPosts() ([]Post, error) {
	var posts []Post
	rows, err := db.Query("SELECT id, title, content, createdAt, author, category, updatedAt, likesCount, authorId, isPublished, views FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		var createdAt, updatedAt NullTime
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &createdAt, &post.Author, &post.Category, &updatedAt, &post.LikesCount, &post.AuthorId, &post.IsPublished, &post.Views); err != nil {
			return nil, err
		}
		post.CreatedAt = createdAt
		post.UpdatedAt = updatedAt
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

// GetPost retrieves a post by its ID
func GetPost(id int) (*Post, error) {
	var post Post
	var createdAt NullTime
	err := db.QueryRow("SELECT id, title, content, createdAt, author, category, updatedAt, likesCount, authorId, isPublished, views FROM posts WHERE id = ?", id).Scan(
		&post.ID, &post.Title, &post.Content, &createdAt, &post.Author, &post.Category, &post.UpdatedAt, &post.LikesCount, &post.AuthorId, &post.IsPublished, &post.Views,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("post with id %d not found", id)
		}
		return nil, err
	}
	post.CreatedAt = createdAt
	return &post, nil
}

// UpdatePost updates an existing post
func UpdatePost(id int, post *Post) error {
	_, err := db.Exec("UPDATE posts SET title = ?, content = ?, author = ?, category = ?, authorId = ?, isPublished = ? WHERE id = ?",
		post.Title, post.Content, post.Author, post.Category, post.AuthorId, post.IsPublished, id)
	if err != nil {
		return err
	}
	return nil
}

// DeletePost removes a post from the database
func DeletePost(id int) error {
	result, err := db.Exec("DELETE FROM posts WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no post found with id %d", id)
	}
	return nil
}
