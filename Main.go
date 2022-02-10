package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"
)

var db *sql.DB

type Album struct {
	ID     int
	Title  string
	Artist string
	Price  float32
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	builder := &DbBuilder{}
	dbOrigin := builder.withConnections(123).
		withAddr("127.0.0.1:3306").
		withName("album_store").
		withUser("mike").
		withPasswd("password").
		build()

	err := dbOrigin.init()
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	for z := 0; z < 1000; z++ {
		db := dbOrigin.getConnection()
		title := "1234, " + strconv.Itoa(z) + " groove!"
		addAlbum(db, Album{
			Title:  title,
			Artist: "Rappin 4 Tay",
			Price:  49.99,
		})
		//fmt.Printf("ID of added album: %v\n", albID)
	}
	duration := time.Since(start)
	println(duration.String())

}

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func addAlbum(db *sql.DB, alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
