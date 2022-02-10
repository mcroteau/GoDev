package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type DbOrigin struct {
	collection               []*sql.DB
	numberOfConnections      int
	user, passwd, addr, name string
}

func (d *DbOrigin) init() {
	for z := 0; len(d.collection) < d.numberOfConnections; z++ {
		db := d.createConnection()
		d.collection = append(d.collection, db)
	}
}

func (d *DbOrigin) getConnection() *sql.DB {
	if len(d.collection) > 0 {
		db := d.collection[0]
		go d.init()
		return db
	}
	go d.init()
	time.Sleep(1 * time.Second)
	return d.getConnection()
}

type DbBuilder struct {
	DbOrigin
}

func (b *DbBuilder) build() DbOrigin {
	return b.DbOrigin
}
func (b *DbBuilder) withConnections(connections int) *DbBuilder {
	b.DbOrigin.numberOfConnections = connections
	return b
}
func (b *DbBuilder) withName(name string) *DbBuilder {
	b.DbOrigin.name = name
	return b
}
func (b *DbBuilder) withAddr(addr string) *DbBuilder {
	b.DbOrigin.addr = addr
	return b
}
func (b *DbBuilder) withUser(user string) *DbBuilder {
	b.DbOrigin.user = user
	return b
}
func (b *DbBuilder) withPasswd(passwd string) *DbBuilder {
	b.DbOrigin.passwd = passwd
	return b
}

func (d *DbOrigin) createConnection() *sql.DB {

	cfg := mysql.Config{
		User:                 d.user,
		Passwd:               d.passwd,
		Net:                  "tcp",
		Addr:                 d.addr,
		DBName:               d.name,
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Print(".")

	return db
}
