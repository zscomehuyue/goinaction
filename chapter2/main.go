package main

import (
	"log"
	"os"

	_ "./matchers"
	"./search"
	//"github.com/Talaran/goinaction/chapter2/search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {

	//test 
	// Perform the search for the specified term.
	search.Run("president")
}
