package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Configurations exported
type Configurations struct {
	Server      ServerConfigurations
	Database    DatabaseConfigurations
	Application ApplicationConfigurations
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Port int
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	DBServer    string
	DBPort      string
	DBName      string
	MaxPoolSize int
}

type ApplicationConfigurations struct {
	UseStubDB        bool
	MaxRecordPerPage int
}

var configuration Configurations

func ReadConfig() {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	viper.SetDefault("database.dbname", "local")

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// Reading variables using the model
	fmt.Println("Reading variables using the model..")
	fmt.Println("Server Port is\t\t", configuration.Server.Port)
	fmt.Println("Database Server is\t", configuration.Database.DBServer)
	fmt.Println("Database Port is\t", configuration.Database.DBPort)
	fmt.Println("Database Name is\t", configuration.Database.DBName)
	fmt.Println("Database MaxPoolSize is\t", configuration.Database.MaxPoolSize)
	fmt.Println("Application UseStubDB is\t", configuration.Application.UseStubDB)
	fmt.Println("Application Pagination parameter MaxRecordPerPage is\t", configuration.Application.MaxRecordPerPage)

}

func GetServerConfig() ServerConfigurations {
	return configuration.Server
}

func GetDBConfig() DatabaseConfigurations {
	return configuration.Database
}

func GetMaxRecordPerPage() uint32 {
	return uint32(configuration.Application.MaxRecordPerPage)
}

func UseStubDB() bool {
	return configuration.Application.UseStubDB
}

func init() {
	ReadConfig()
}
