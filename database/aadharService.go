package database

import (
	"Aadhar_POC/model"
)

type AadharMongoClient interface {
	GetAadharDetails() (model.AadharDetails, error)
}

func (client *MongoClientImpl) GetAadharDetails() (model.AadharDetails, error) {

	return model.AadharDetails{}, nil
}