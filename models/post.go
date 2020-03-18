package models

import (
	"github.com/jinzhu/gorm"
	// postgress db driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// import sqlite3 driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Quote struct
type Post struct {
	gorm.Model
	Text string
	User string
}

// QuoteManager struct
type PostManager struct {
	db *DB
}

// NewPostManager - Create a quote manager that can be used for retrieving quotes
func NewPostManager(db *DB) (*PostManager, error) {

	db.AutoMigrate(&Post{})

	postmgr := PostManager{}

	postmgr.db = db
	
	return &postmgr, nil
}

// AddPost - Creates a user and hashes the password
func (state *PostManager) AddPost(text string, userName string) *Post {
	post := &Post{
		Text: text,
		User: userName,
	}
	state.db.Create(&post)
	return post
}

func (state *PostManager) GetPosts() []Post {
	var posts []Post
	state.db.Find(&posts)
	return posts
}
