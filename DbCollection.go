package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

type DbCollection struct {
	collection               []*sql.DB
	numberOfConnections      int
	user, passwd, addr, name string
}

func (d *DbCollection) init() {
	for z := 0; z < d.numberOfConnections; z++ {
		db := d.createConnection()
		d.collection = append(d.collection, db)
	}
}

func

type DbBuilder struct {
	dbCollection *DbCollection
}

func (b *DbBuilder) withAddress(addr string) *DbBuilder {
	b.dbCollection.addr = addr
	return b
}
func (b *DbBuilder) withName(name string) *DbBuilder {
	b.dbCollection.name = name
	return b
}
func (b *DbBuilder) withUser(user string) *DbBuilder {
	b.dbCollection.user = user
	return b
}
func (b *DbBuilder) withPass(passwd string) *DbBuilder {
	b.dbCollection.passwd = passwd
	return b
}
func (b *DbBuilder) withConnections(numberOfConnections int) *DbBuilder {
	b.dbCollection.numberOfConnections = numberOfConnections
	return b
}
func (b *DbBuilder) make() *DbCollection {
	return b.dbCollection
}

func (d *DbCollection) createConnection() *sql.DB {

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

	return db
}
