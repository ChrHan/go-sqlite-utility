# SQLite Go Wrapper

A `Go` Library to wrap `CRUD` operations to `sqlite3` database.

Currently programmed only for this table specification:

    CREATE TABLE PRODUCTS (
      id  INT PRIMARY KEY,
      product_name  VARCHAR(20)
    )

# Prerequisite

Install `Go`

# Usage

1. Import this repo on your Go implementation

    import "github.com/ChrHan/golang-sqlite-wrapper/dbutil"

1. Initialize dbutil to local variable with `SQLite` database filename

    db := dbutil.New(db_filename)

1. Valid functions on dbutil are:
  - Select()
  
    Returns `select * from products`

  - SelectCount()

    Returns `select count(1) from products`
  
  - Insert(id, product_name)

    Performs `insert into products (id, product_name) values (id, product_name)`

  - Delete(id)

    Performs `delete from products where id = id` 

  - Update(id, product_name)

    Performs `update products set product_name = product_name where id = id` 

# Test

1. `cd dbutil`
2. Run `go test`
