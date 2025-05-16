// به نام خدا
// mylib is a package for give you good options
package myLib

import (
	"log"
)

// a err handler
func  Try(err error ,errText... string) {
	var errTextDone string
	
	// add all errText to one
	for _ , errT := range errText{
		errTextDone += errT + " "
	}
	
	// handle the err
	if err != nil{
		log.Fatal("\n \n", err , " | ", errTextDone , "\n ")
	}
}


// اللهم عجل لولیک الفرج
