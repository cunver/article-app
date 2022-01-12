package config

import "testing"

func TestGetMongoDBConnection(t *testing.T) {
	dbConnection := GetMongoDBConnection()
	if dbConnection == nil {
		t.Fatal("Unable to get a mongodb connection")
	}
}
