package main

// Simple flat configuration file

import (
	"fmt"
	"time"

	"../../config"
)

// variable to hold configuration
var configuration Configuration

// Configuration - structure defining the configuration
type Configuration struct {
	LastRan  string `json:"lastran"`  // date last ran
	Username string `json:"username"` // username
	ID       string `json:"id"`       // id
}

// setLastRan set last ran
func (obj *Configuration) setLastRan(lastran string) {
	obj.LastRan = lastran
}

// getLastRan get last ran
func (obj Configuration) getLastRan() string {
	return configuration.LastRan
}

// setID set id
func (obj *Configuration) setID(id string) {
	obj.ID = id
}

// getID get id
func (obj Configuration) getID() string {
	return obj.ID
}

// setUsername  set Username
func (obj *Configuration) setUsername(username string) {
	obj.Username = username
}

// getUsername get Username
func (obj Configuration) getUsername() string {
	return obj.Username
}

func main() {

	// initialise the configuration
	configuration = Configuration{}
	// load from the JSON file
	config.LoadConfig(&configuration)
	// get a value
	fmt.Printf("Last Ran %s\n", configuration.getLastRan())
	// change a value
	configuration.setLastRan(time.Now().Format("2006-01-02_15:04:05"))
	fmt.Printf("Last Ran %s\n", configuration.getLastRan())
	// save the value`
	config.SaveConfig()

}
