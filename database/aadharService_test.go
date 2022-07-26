package database

import (
	"testing"
)

func BenchmarkGetAadharDetails(b *testing.B){
	dataStoreClient := MongoConnector()
	for i:=0;i<b.N;i++{
		dataStoreClient.GetAadharDetails("1234")
	}
}