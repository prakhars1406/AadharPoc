package database

import (
	"Aadhar_POC/config"
	"Aadhar_POC/model"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"sync"
)

type AadharMongoClient interface {
	GetAadharDetails(id string,getAadharDetailChannel chan model.AadharDetails,getAadharWaitGroup *sync.WaitGroup)
	InsertAadharDetails(model.AadharDetails) (string, error)
}

func (client *MongoClientImpl) GetAadharDetails(id string,getAadharDetailChannel chan model.AadharDetails,getAadharWaitGroup *sync.WaitGroup) {
	session := client.session.Copy()
	defer session.Close()
	defer getAadharWaitGroup.Done()
	session.SetMode(mgo.Monotonic, config.STRONG_MODE)
	c := session.DB(config.DATABASE_NAME).C(fmt.Sprintf(config.COLLECTION_NAME))
	result := bson.M{}
	err := c.Find(bson.M{"id": strings.ToLower(id)}).Select(bson.M{}).One(&result)
	if err != nil {
		getAadharDetailChannel<-model.AadharDetails{Error:errors.New("<mongo> Unable to query collection") }
		return
	}
	if len(result) == 0 {
		getAadharDetailChannel<-model.AadharDetails{Error:errors.New("no aadhar details found found") }
		return
	}
	var aadharDetails model.AadharDetailsMongoResponse
	bsonBytes, _ := bson.Marshal(result)
	err = bson.Unmarshal(bsonBytes, &aadharDetails)
	if err == nil {
		getAadharDetailChannel<-model.AadharDetails(aadharDetails)
		return
	} else{
		getAadharDetailChannel<-model.AadharDetails{Error:err }
		return
	}
}

func (client *MongoClientImpl) InsertAadharDetails(aadharDetails model.AadharDetails) (string, error) {
	session := client.session.Copy()
	defer session.Close()

	c := session.DB(config.DATABASE_NAME).C(config.COLLECTION_NAME)

	uuid := uuid.New()
	aadharDetails.Id = uuid.String()

	err := c.Insert(aadharDetails)

	return uuid.String(), err
}