package database

import (
	"github.com/yushengguo557/magellanic-l/internal/models"
	"log"
)

// DataAccessInterface data access interface
//
// 数据访问接口
type DataAccessInterface interface {
	SetupDatabase() error
	LookupUserByEmail(id string) *models.User
	LookupUserByID(email string) *models.User
}

func NewDatabase() (*DataAccessInterface, error) {
	var database DataAccessInterface = &Database{}

	var err = database.SetupDatabase()
	if err != nil {
		log.Fatalln("set up database, err: ", err)
	}

	return &database, nil
}

type Database struct{}

func (db *Database) SetupDatabase() error {
	return nil
}
