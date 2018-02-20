package dbutil

import (
	"database/sql"
	"fmt"
	"github.com/prometheus/log"
	"strconv"
)

// Dbutil stores filename for SQLite3 database
type Dbutil struct {
	Filename string
}

// New is a function to initialize Dbutil
func New(filename string) *Dbutil {
	return &Dbutil{
		Filename: filename,
	}
}

// Prepare sets up database inside Dbutil, creates product table if not found
func (d *Dbutil) Prepare() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", fmt.Sprintf("./%s", d.Filename))
	log.Println("Prepare #-2")
	if err != nil {
		log.Println("Prepare #-1")
		log.Println(err.Error())
		return nil, err
	}
	log.Println("Prepare #0")
	_, err = db.Query("select id, product_name from products")
	log.Println("Prepare #1")
	if err != nil {
		log.Println("Prepare #2")
		db.Exec("create table products (id int primary key, product_name varchar(20))")
		log.Println("Prepare #3")
		return nil, nil
	}
	return db, nil
}

// Select performs `select * from products` and returns *sql.Rows
func (d *Dbutil) Select() (*sql.Rows, error) {
	log.Println("Select #-2 - util")
	db, err := d.Prepare()
	log.Println("Select #-1 - util")
	if err != nil {
		log.Println("Select #0")
		log.Println(err.Error())
		log.Println("Select #1")
		return nil, err
	}
	rows, err := db.Query("select id, product_name from products")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	db.Close()
	return rows, nil
}

// SelectCount performs `select count(1) from products` and returns int
func (d *Dbutil) SelectCount() int {
	db, err := d.Prepare()
	var result string
	var intResult int
	err = db.QueryRow("select count(1) from products").Scan(&result)
	if err != nil {
		log.Println(err)
	}
	intResult, err = strconv.Atoi(result)
	db.Close()
	return intResult
}

// Insert performs `insert into products values (id, product_name) and returns error if any
func (d *Dbutil) Insert(id string, productName string) error {
	db, err := d.Prepare()
	queryString := fmt.Sprintf("insert into products (id, product_name) values (%s, '%s')", id, productName)
	_, err = db.Exec(queryString)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}

// Update performs `update products set product_name = product_name where id = id` and returns error if any
func (d *Dbutil) Update(id string, productName string) error {
	db, err := d.Prepare()
	queryString := fmt.Sprintf("update products set product_name = '%s' where id = %s", productName, id)
	_, err = db.Exec(queryString)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}

// SelectOne performs `select * from products where id = id` and returns product_name
func (d *Dbutil) SelectOne(id string) string {
	db, err := d.Prepare()
	var result string
	err = db.QueryRow(fmt.Sprintf("select product_name from products where id = %s", id)).Scan(&result)
	if err != nil {
		log.Println(err)
	}
	db.Close()
	return result
}

// Delete performs `delete from products where id = id` and returns error if any
func (d *Dbutil) Delete(id string) error {
	db, err := d.Prepare()
	queryString := fmt.Sprintf("delete from products where id = %s", id)
	_, err = db.Exec(queryString)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}

// DeleteAll performs `delete from products` and returns error if any
func (d *Dbutil) DeleteAll() error {
	db, err := d.Prepare()
	queryString := "delete from products"
	_, err = db.Exec(queryString)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}
