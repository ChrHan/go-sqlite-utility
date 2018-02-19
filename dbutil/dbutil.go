package dbutil

import (
	"database/sql"
	"fmt"
	"strconv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/prometheus/log"
)

type Dbutil struct {
	Filename string
}

func New(filename string) *Dbutil {
	return &Dbutil{
		Filename: filename,
	}
}

func (d *Dbutil) Prepare() *sql.DB {
	db, err := sql.Open("sqlite3", fmt.Sprintf("./%s", d.Filename))
	_, err = db.Query("select id, name from products")
	if err != nil {
		db.Exec("create table products (id int primary key, product_name varchar(20))")
	}
  return db
}

func (d *Dbutil) Select() *sql.Rows {
  db := d.Prepare()
	rows, err := db.Query("select id, product_name from products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
	return rows
}

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

func (d *Dbutil) Insert(id string, product_name string) {
  db := d.Prepare()
	query_string := fmt.Sprintf("insert into products (id, product_name) values (%s, '%s')", id, product_name)
	_, err := db.Exec(query_string)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}

func (d *Dbutil) Update(id string, product_name string) {
  db := d.Prepare()
	query_string := fmt.Sprintf("update products set product_name = '%s' where id = %s", product_name, id)
	_, err := db.Exec(query_string)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}

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

func (d *Dbutil) Delete(id string) {
  db := d.Prepare()
	query_string := fmt.Sprintf("delete from products where id = %s", id)
	_, err := db.Exec(query_string)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}

func (d *Dbutil) DeleteAll() {
  db := d.Prepare()
	query_string := "delete from products"
	_, err := db.Exec(query_string)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}
