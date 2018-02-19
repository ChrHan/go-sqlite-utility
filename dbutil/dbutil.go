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
func (d *Dbutil) Prepare() *sql.DB {
	db, err := sql.Open("sqlite3", fmt.Sprintf("./%s", d.Filename))
	_, err = db.Query("select id, name from products")
	if err != nil {
		db.Exec("create table products (id int primary key, product_name varchar(20))")
	}
	return db
}

// Select performs `select * from products` and returns *sql.Rows
func (d *Dbutil) Select() *sql.Rows {
	db := d.Prepare()
	rows, err := db.Query("select id, product_name from products")
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
	return rows
}

// SelectCount performs `select count(1) from products` and returns int
func (d *Dbutil) SelectCount() int {
	db := d.Prepare()
	var result string
	var intResult int
	err := db.QueryRow("select count(1) from products").Scan(&result)
	if err != nil {
		log.Fatal(err)
	}
	intResult, err = strconv.Atoi(result)
	db.Close()
	return intResult
}

// Insert performs `insert into products values (id, product_name) and returns error if any
func (d *Dbutil) Insert(id string, productName string) error {
	db := d.Prepare()
	queryString := fmt.Sprintf("insert into products (id, product_name) values (%s, '%s')", id, productName)
	_, err := db.Exec(queryString)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}

// Update performs `update products set product_name = product_name where id = id` and returns error if any
func (d *Dbutil) Update(id string, productName string) error {
	db := d.Prepare()
	queryString := fmt.Sprintf("update products set product_name = '%s' where id = %s", productName, id)
	_, err := db.Exec(queryString)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}

// SelectOne performs `select * from products where id = id` and returns product_name
func (d *Dbutil) SelectOne(id string) string {
	db := d.Prepare()
	var result string
	err := db.QueryRow(fmt.Sprintf("select product_name from products where id = %s", id)).Scan(&result)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
	return result
}

// Delete performs `delete from products where id = id` and returns error if any
func (d *Dbutil) Delete(id string) error {
	db := d.Prepare()
	queryString := fmt.Sprintf("delete from products where id = %s", id)
	_, err := db.Exec(queryString)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}

// DeleteAll performs `delete from products` and returns error if any
func (d *Dbutil) DeleteAll() error {
	db := d.Prepare()
	queryString := "delete from products"
	_, err := db.Exec(queryString)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}
