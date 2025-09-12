package main

import (
	"log"
	"os"
	"strconv"
)

func lookup(envString string, defaultValue string) string {
	value, ok := os.LookupEnv(envString)
	if !ok {
		log.Printf("Optional env variable %s was not found, using default value %s \n", envString, defaultValue)
		return defaultValue
	}
	return value
}

func lookupRequired(envString string) string {
	value, ok := os.LookupEnv(envString)
	if !ok {
		log.Panicf("Required env variable %s was not found, please make sure it is present in your .env file \n", envString)
	}
	return value
}

func lookupBool(envString string, defaultValue bool) bool {
	l := lookup(envString, strconv.FormatBool(defaultValue))
	v, err := strconv.ParseBool(l)
	if err != nil {
		log.Printf("Error parsing env variable %s: expected bool, received %s \n", envString, l)
		return defaultValue
	}
	return v
}
