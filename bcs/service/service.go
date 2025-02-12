package service

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Post represents the structure of the post in the database
type Post struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	CreatedAt   time.Time  `json:"createdAt"`
	Author      *string    `json:"author,omitempty"`
	Category    *string    `json:"category,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	LikesCount  int        `json:"likesCount"`
	AuthorId    *int       `json:"authorId,omitempty"`
	IsPublished bool       `json:"isPublished"`
	Views       int        `json:"views"`
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
	result, err := db.Exec("INSERT INTO posts (title, content, author, category, authorId, isPublished) VALUES (?, ?, ?, ?, ?, ?)",
		post.Title, post.Content, post.Author, post.Category, post.AuthorId, post.IsPublished)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// GetPost retrieves a post by its ID
func GetPost(id int) (*Post, error) {
	post := &Post{}
	err := db.QueryRow("SELECT id, title, content, createdAt, author, category, updatedAt, likesCount, authorId, isPublished, views FROM posts WHERE id = ?", id).Scan(
		&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.Author, &post.Category, &post.UpdatedAt, &post.LikesCount, &post.AuthorId, &post.IsPublished, &post.Views,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("post with id %d not found", id)
		}
		return nil, err
	}
	return post, nil
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
