package main

import (
	"article-app/config"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	config.ReadConfig()
	code := m.Run()
	os.Exit(code)
}
