package config

import "testing"

func TestGetMongoDBConnection(t *testing.T) {
	ReadConfig()
	if !UseStubDB() {
		dbConnection := GetMongoDBConnection()
		if dbConnection == nil {
			t.Fatal("Unable to get a mongodb connection")
		}
	}
}
