package database

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

// var (
// 	db leveldb.DB = nil
// )

type LevelDBClientImpl struct {
	db *leveldb.DB
}

func LevelDBConnector() *leveldb.DB {
	db, err := leveldb.OpenFile("./database.db", nil)
	defer db.Close()
	if err != nil {
		fmt.Println("Error in connection: ", err)
		return nil
	}

	return db
}
