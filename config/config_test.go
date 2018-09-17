package config

import (
	"os"
	"path"
	"testing"
)

// Database config structire
type Database struct {
	Dbpath string `json:"dbpath"` // path where database is stored
}

// Identity config structure
type Identity struct {
	Sourcename string `json:"sourcename"` // friendly name of source
	SourceID   string `json:"sourceid"`   // uuid used to identify source of data in service
}

// Configuration - structure to hold the configuration
type Configuration struct {
	Database Database `json:"database"`
	Identity Identity `json:"identity"`
}

func TestNotFindConfig(t *testing.T) {
	filepath, err := findConfig()
	if err != nil {
		if err.Error() != "Config file not found" {
			t.Errorf("Error %s", err)
		}
	}
	if filepath != "" {
		t.Errorf("filepath %s should be empty", filepath)
	}
}

func TestFindConfig(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Errorf("Error %s", err)
	}
	SetConfigName(path.Join(dir, "../examples/complex/config.json"))

	filepath, err := findConfig()
	if err != nil {
		t.Errorf("Error %s", err)
	}
	if filepath == "" {
		t.Errorf("filepath %s is empty", filepath)
	}
}

func TestFindConfigName(t *testing.T) {
	SetConfigName("config_test.json")

	filepath, err := findConfig()
	if err != nil {
		t.Errorf("Error %s", err)
	}
	if filepath == "" {
		t.Errorf("filepath %s is empty", filepath)
	}
}

func TestLoadConfig(t *testing.T) {
	SetConfigName("config_test.json")
	var configuration Configuration
	configuration = Configuration{}
	LoadConfig(&configuration)
	if configuration.Database.Dbpath != "/tmp/fsdb" {
		t.Error("Invalid Configuration")
	}
}

func TestSaveConfig(t *testing.T) {
	SetConfigName("config_test.json")
	var configuration Configuration
	configuration = Configuration{}
	LoadConfig(&configuration)
	if configuration.Database.Dbpath != "/tmp/fsdb" {
		t.Error("Invalid Configuration")
	}
	configuration.Database.Dbpath = "/mount/bill"
	SaveConfig()
	LoadConfig(&configuration)
	if configuration.Database.Dbpath != "/mount/bill" {
		t.Error("Invalid Configuration")
	}
	configuration.Database.Dbpath = "/tmp/fsdb"
	SaveConfig()
	LoadConfig(&configuration)
	if configuration.Database.Dbpath != "/tmp/fsdb" {
		t.Error("Invalid Configuration")
	}
}
