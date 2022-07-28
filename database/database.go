package database

import {
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
}

var (
	db leveldb.DB = nil
)

func LevelDBConnector() (leveldb.DB) {
	db, err := leveldb.OpenFile("./database.db", nil)
	if err != nil {fmt.Println("Error in connection: ", err)
	return nil
	}
	
	return db
}