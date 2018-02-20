package dbutil_test

import (
	"fmt"
	db "github.com/ChrHan/go-sqlite-utility/dbutil"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"log"
	"os"
	"testing"
)

const DB_FILENAME2 = "test2.db"

type Dbutil2Suite struct {
	suite.Suite
	dbutil *db.Dbutil
}

func TestDbutil2Suite(t *testing.T) {
	suite.Run(t, &Dbutil2Suite{})
}

func (dc *Dbutil2Suite) SetupSuite() {
	os.Remove(DB_FILENAME2)
	dc.dbutil = db.New(DB_FILENAME2)
	fmt.Println("Executing SetupSuite")
}

func (dc *Dbutil2Suite) Test1SelectNoPrepare() {
	//	dc.dbutil = db.New(DB_FILENAME2)
	result, err := dc.dbutil.Select()
	if err != nil {
		log.Println(err.Error())
	}
	resultCount := dc.dbutil.SelectCount()
	assert.Nil(dc.T(), result, "Result should not be Nil")
	assert.Equal(dc.T(), 0, resultCount, "Result should be 0")
}
