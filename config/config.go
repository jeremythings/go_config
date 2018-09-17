package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

var configuration interface{}
var configurationFilepath string
var configFile = "config.json"

func findConfig() (string, error) {
	var err error
	var filePath string

	if configurationFilepath != "" {
		filePath = configurationFilepath
		_, existerr := os.Stat(filePath)
		if existerr != nil {
			return "", existerr
		}
		return filePath, nil
	}

	// look where program running
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	filePath = path.Join(dir, configFile)
	// fmt.Printf("Looking in %s for config file\n", filePath)
	if _, existerr := os.Stat(filePath); existerr == nil {
		return filePath, nil
	}

	// look where program running/config
	dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	filePath = path.Join(dir, "config", configFile)
	// fmt.Printf("Looking in %s for config file\n", filePath)
	if _, existerr := os.Stat(filePath); existerr == nil {
		return filePath, nil
	}

	// look in cwd
	dir, err = os.Getwd()
	if err != nil {
		return "", err
	}
	filePath = path.Join(dir, configFile)
	// fmt.Printf("Looking in %s for config file\n", filePath)
	if _, existerr := os.Stat(filePath); existerr == nil {
		return filePath, nil
	}

	// look in cwd/config
	dir, err = os.Getwd()
	if err != nil {
		return "", err
	}
	filePath = path.Join(dir, "config", configFile)
	// fmt.Printf("Looking in %s for config file\n", filePath)
	if _, existerr := os.Stat(filePath); existerr == nil {
		return filePath, nil
	}

	// look in os.cwd/../config
	dir, err = os.Getwd()
	if err != nil {
		return "", err
	}
	filePath = path.Join(dir, "../config", configFile)
	// fmt.Printf("Looking in %s for config file\n", filePath)
	if _, existerr := os.Stat(filePath); existerr == nil {
		return filePath, nil
	}

	// look in config in current folder using runtime caller
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", fmt.Errorf("failed to get cwd")
	}
	filePath = path.Join(path.Dir(filename), "../config", configFile)
	// fmt.Printf("Looking in %s for config file\n", filePath)
	if _, existerr := os.Stat(filePath); existerr == nil {
		return filePath, nil
	}

	return "", fmt.Errorf("Config file not found")

}

// SetConfigName - sets the name or the path for the configuration
func SetConfigName(path string) {
	// if path is a full path then set configurationFilepath
	dir := filepath.Dir(path)
	filename := filepath.Base(path)
	configFile = filename
	if dir != "" {
		configurationFilepath = path
	}
}

// LoadConfig - builds config data
func LoadConfig(configurationin interface{}) {
	filepath, err := findConfig()
	if err != nil {
		fmt.Printf("Cannot find config file: %s\n", err)
		os.Exit(1)
	}
	configurationFilepath = filepath
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Cannot open config file: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration = configurationin
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
}

// SaveConfig - writes config to file
func SaveConfig() {
	file, err := os.Create(configurationFilepath)
	if err != nil {
		fmt.Printf("Cannot open config file to write: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
}
