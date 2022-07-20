package database

import (
	"Aadhar_POC/config"
	"gopkg.in/mgo.v2"
	"time"
)

var (
	dataStoreClient MongoClient          = nil
)

type MongoClientImpl struct {
	mongoServer string
	session     *mgo.Session
}

func MongoConnector() MongoClient {
	if dataStoreClient == nil {
		dataStore := config.DATASTORE
		if dataStore == config.MONGO {
			mongoServer := config.MONGO_SERVER
			dialInfo := mgo.DialInfo{
				Addrs:     []string{mongoServer},
				Direct:    false,
				Timeout:   1 * time.Second,
				FailFast:  false,
				PoolLimit: 100,
			}
			session, err := mgo.DialWithInfo(&dialInfo)
			if err != nil {
				// utility.SendNotification(err, fmt.Sprintf("Failed to connect to mongo with configuration [%v]", dialInfo), configs.MONGO_SERVICE)
				return nil
			}
			dataStoreClient = &MongoClientImpl{mongoServer: mongoServer, session: session}
		}

	}
	return dataStoreClient
}