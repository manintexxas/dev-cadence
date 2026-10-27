package main

import (
		"fmt"
		"os"
)

type level string

const (
	Info	level = "[INFO]"
	Error	level = "[ERROR]"
	Warning level = "[WARNING]"
)

var (
	directory = getDir()
	filename = "App.log"
)

func getDir() string {
	directory, err := os.Getwd()
	if err != nil {
		fmt.Printf("")
	}

	return directory
}


