package connection

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"openceptor.eu/config"

	_ "github.com/lib/pq"
)

var lockDb = &sync.Mutex{}

type database struct {
	dsn string
	db *sql.DB
}

var databaseInstance *database

func buildDbDsn(c *config.Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Database.Host, c.Database.Port, c.Database.Username, c.Database.Password, c.Database.Database)
}

func openDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func pingDb() {
	for range time.Tick(1 * time.Second) {
		err := GetDbInstance(nil).Ping()

		if err != nil {
			db, err := openDatabase(databaseInstance.dsn)
			if err != nil {
				panic(err)
			}

			databaseInstance.db = db
		}
	}
}

func GetDbInstance(c *config.Config) *sql.DB {
	if databaseInstance != nil {
		return databaseInstance.db;
	}

	lockDb.Lock()
	defer lockDb.Unlock()

	if databaseInstance != nil {
		fmt.Println("databaseInstance already exist, re-use (after lock)")

		return databaseInstance.db;
	}

	fmt.Println("databaseInstance not exist, create it !")
	databaseInstance = &database{}

	databaseInstance.dsn = buildDbDsn(c)

	db, err := openDatabase(databaseInstance.dsn)
	if err != nil {
		panic(err)
	}

	databaseInstance.db = db

	databaseInstance.db.SetMaxOpenConns(c.Database.MaxOpenConns)
	databaseInstance.db.SetMaxIdleConns(c.Database.MaxIdleConns)

	//defer databaseInstance.db.Close()

	go pingDb()
	
	return databaseInstance.db
}
