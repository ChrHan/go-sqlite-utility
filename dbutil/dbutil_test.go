package dbutil_test

import (
	"os"
	"testing"

	db "github.com/ChrHan/go-sqlite-utility/dbutil"
	"github.com/icrowley/fake"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const DB_FILENAME = "test.db"

type DbutilSuite struct {
	suite.Suite
	dbutil *db.Dbutil
}

func TestDbutilSuite(t *testing.T) {
	suite.Run(t, &DbutilSuite{})
}

func (dc *DbutilSuite) SetupSuite() {
	os.Remove(DB_FILENAME)
	dc.dbutil = db.New(DB_FILENAME)
	dc.dbutil.Prepare()
}

func (dc *DbutilSuite) Test1Select() {
	result, _ := dc.dbutil.Select()
	resultCount := dc.dbutil.SelectCount()
	assert.NotNil(dc.T(), result, "Result should not be Nil")
	assert.Equal(dc.T(), 0, resultCount, "Result should be 0")
}

func (dc *DbutilSuite) Test2Insert() {
	dc.dbutil.DeleteAll()
	id := fake.Digits()
	product_name := fake.Product()
	dc.dbutil.Insert(id, product_name)
	result, _ := dc.dbutil.Select()
	resultCount := dc.dbutil.SelectCount()
	assert.NotNil(dc.T(), result, "Result should not be Nil")
	assert.NotEqual(dc.T(), 0, resultCount, "Result should NOT be 0")
	product_name_selected := dc.dbutil.SelectOne(id)
	assert.Equal(dc.T(), product_name, product_name_selected, "Result should be equal to inserted value")
}

func (dc *DbutilSuite) Test3InsertDelete() {
	dc.dbutil.DeleteAll()
	id := fake.Digits()
	product_name := fake.Product()
	dc.dbutil.Insert(id, product_name)
	dc.dbutil.Insert(fake.Digits(), fake.Product())
	dc.dbutil.Delete(id)
	result, _ := dc.dbutil.Select()
	resultCount := dc.dbutil.SelectCount()
	assert.NotNil(dc.T(), result, "Result should not be Nil")
	assert.Equal(dc.T(), 1, resultCount, "Result should be 1")
}

func (dc *DbutilSuite) Test4InsertUpdateDelete() {
	dc.dbutil.DeleteAll()
	id := fake.Digits()
	product_name := fake.Product()
	dc.dbutil.Insert(id, product_name)
	new_product_name := fake.Product()
	dc.dbutil.Update(id, new_product_name)
	result, _ := dc.dbutil.Select()
	resultCount := dc.dbutil.SelectCount()
	resultUpdated := dc.dbutil.SelectOne(id)
	assert.NotNil(dc.T(), result, "Result should not be Nil")
	assert.Equal(dc.T(), 1, resultCount, "Result should be 1")
	assert.Equal(dc.T(), new_product_name, resultUpdated, "Result should be 1")
}

func (dc *DbutilSuite) Test5InsertDouble() {
	dc.dbutil.DeleteAll()
	id := fake.Digits()
	product_name := fake.Product()
	dc.dbutil.Insert(id, product_name)
	err := dc.dbutil.Insert(id, product_name)
	assert.NotNil(dc.T(), err, "err should not be Nil")
	errorMessageExpected := "UNIQUE constraint failed: products.id"
	assert.Equal(dc.T(), errorMessageExpected, err.Error(), "err should be equal to "+errorMessageExpected)
	result, err := dc.dbutil.Select()
	resultCount := dc.dbutil.SelectCount()
	assert.NotNil(dc.T(), result, "Result should not be Nil")
	assert.NotEqual(dc.T(), 0, resultCount, "Result should NOT be 0")
	assert.Equal(dc.T(), 1, resultCount, "Result should be 1")
	product_name_selected := dc.dbutil.SelectOne(id)
	assert.Equal(dc.T(), product_name, product_name_selected, "Result should be equal to inserted value")
}
