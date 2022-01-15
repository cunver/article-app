package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {

	ReadConfig()
	assert.NotEmptyf(t, GetDBConfig(), "DB Config parameter can not be empty")
	assert.NotEmptyf(t, GetServerConfig(), "Server config parameters can not be empty")
	assert.Greater(t, GetMaxRecordPerPage(), uint32(0), "Application parameter maxrecordperpage must be greater than 0, got %v", GetMaxRecordPerPage())

}
