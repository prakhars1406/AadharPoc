package database

import (
	"Aadhar_POC/config"
	"Aadhar_POC/model"
	"github.com/google/uuid"
)

const COLLECTION_NAME = "aadharDetails"

type AadharMongoClient interface {
	GetAadharDetails() (model.AadharDetails, error)
	InsertAadharDetails(model.AadharDetails) (string, error)
}

func (client *MongoClientImpl) GetAadharDetails() (model.AadharDetails, error) {

	return model.AadharDetails{}, nil
}

func (client *MongoClientImpl) InsertAadharDetails(aadharDetails model.AadharDetails) (string, error) {
	c := client.session.DB(config.DATABASE_NAME).C(COLLECTION_NAME)

	uuid := uuid.New()
	aadharDetails.Id = uuid.String()

	err := c.Insert(aadharDetails)

	return uuid.String(), err
}