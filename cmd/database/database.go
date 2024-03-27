package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"petshop/proto/dog"
	"petshop/proto/person"
)

var (
	db  *gorm.DB
	err error
)

func GetDb() (*gorm.DB, error) {
	if db == nil {
		db, err = gorm.Open(sqlite.Open("petshop.db"), &gorm.Config{})
		if err != nil {
			return nil, err
		}

		err = createDatabaseIfNotExists()
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

func createDatabaseIfNotExists() error {
	if _, err = os.Stat("petshop.db"); os.IsNotExist(err) {
		_, err = GetDb()
		if err != nil {
			return err
		}

		err = db.AutoMigrate(&person.Person{})
		if err != nil {
			return err
		}

		err = db.AutoMigrate(&dog.Dog{})
		if err != nil {
			return err
		}
	}
	return nil
}
