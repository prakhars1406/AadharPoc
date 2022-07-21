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
)

type AadharMongoClient interface {
	GetAadharDetails(id string)(model.AadharDetails,error)
	InsertAadharDetails(model.AadharDetails) (string, error)
}

func (client *MongoClientImpl) GetAadharDetails(id string)(model.AadharDetails,error) {
	session := client.session.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, config.STRONG_MODE)
	c := session.DB(config.DATABASE_NAME).C(fmt.Sprintf(config.COLLECTION_NAME))
	result := bson.M{}
	err := c.Find(bson.M{"id": strings.ToLower(id)}).Select(bson.M{}).One(&result)
	if err != nil {
		return model.AadharDetails{},errors.New("<mongo> Unable to query collection")
	}
	if len(result) == 0 {
		return model.AadharDetails{},errors.New("<mongo> no aadhar details found found")
	}
	var aadharDetails model.AadharDetailsMongoResponse
	bsonBytes, _ := bson.Marshal(result)
	err = bson.Unmarshal(bsonBytes, &aadharDetails)
	if err == nil {
		return model.AadharDetails(aadharDetails),nil
	} else{
		return model.AadharDetails{},err
	}
}

func (client *MongoClientImpl) InsertAadharDetails(aadharDetails model.AadharDetails) (string, error) {
	c := client.session.DB(config.DATABASE_NAME).C(config.COLLECTION_NAME)
	uuid := uuid.New()
	aadharDetails.Id = uuid.String()
	err := c.Insert(aadharDetails)
	return uuid.String(), err
}