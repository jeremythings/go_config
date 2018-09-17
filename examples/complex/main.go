package main

// example with nested configuration structures
// also with a set of methods that themselves call the
// configuration methods

import (
	"fmt"
	"time"

	"../../config"
)

var configuration Configuration

// Database config structire
type Database struct {
	Dbpath string `json:"dbpath"` // path where database is stored
}

// SetDbpath database set path setting
func (obj *Database) setDBPath(dbpath string) {
	obj.Dbpath = dbpath
	fmt.Printf("Configuration struct %s\n", configuration.Database.Dbpath)
	fmt.Printf("database struct %s\n", obj.Dbpath)
}

// GetDbpath database get path setting
func (obj Database) getDBPath() string {
	fmt.Printf("Configuration struct %s\n", configuration.Database.Dbpath)
	fmt.Printf("database struct %s\n", obj.Dbpath)
	return configuration.Database.Dbpath
	// return obj.Dbpath
}

// Identity config structure
type Identity struct {
	Sourcename string `json:"sourcename"` // friendly name of source
	SourceID   string `json:"sourceid"`   // uuid used to identify source of data in service
}

// SetSourceID set identity source id
func (obj *Identity) setSourceID(sourceid string) {
	obj.SourceID = sourceid
}

// GetSourceID get identity source id
func (obj Identity) getSourceID() string {
	return obj.SourceID
}

// SetSourcename identity set source name
func (obj *Identity) setSourcename(sourcename string) {
	obj.Sourcename = sourcename
}

// GetSourcename identity get source name
func (obj Identity) getSourcename() string {
	return obj.Sourcename
}

// Configuration - structure to hold the configuration
type Configuration struct {
	Database Database `json:"database"`
	Identity Identity `json:"identity"`
}

// GetDBPath - get database path
func GetDBPath() string {
	return configuration.Database.getDBPath()
}

// SetDBPath - set database path
func SetDBPath(dbpath string) {
	configuration.Database.setDBPath(dbpath)
}

func main() {

	// initialise the configuration
	configuration = Configuration{}
	// load from the JSON file
	config.LoadConfig(&configuration)
	// print the configuration
	fmt.Printf("Configuration retrieved %+v\n", configuration)
	// get a value
	fmt.Printf("Database path %s\n", GetDBPath())
	// set a value
	SetDBPath("/mount/" + time.Now().Format("2006-01-02_15:04:05"))
	// get a value
	fmt.Printf("Database path %s\n", GetDBPath())
	// save the configuration
	config.SaveConfig()

}
