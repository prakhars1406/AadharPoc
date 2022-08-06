package database

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	dataStoreLDBClient LevelDbClient = nil
)

type LevelDBClientImpl struct {
	db *leveldb.DB
}

func LevelDBConnector() LevelDbClient {
	db, err := leveldb.OpenFile("./database.db", nil)
	defer db.Close()
	if err != nil {
		fmt.Println("Error in connection: ", err)
		return nil
	}

	dataStoreLDBClient = &LevelDBClientImpl{db: db}

	return dataStoreLDBClient
}
