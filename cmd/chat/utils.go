package main

import (
	"encoding/json"
	"os"
)

//StringInSlice looksup a string in a slice returns true if found
func StringInSlice(s []string, srch string) (bool, error) {
	for _, val := range s {
		if val == srch {
			return true, nil
		}
	}
	return false, nil
}

//GetConfigFromJSON pass a *struct via an interface to have its vars intialised from json in filename.
func GetConfigFromJSON(filename string, configuration interface{}) (err error) {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		return err
	}
	return err
}
