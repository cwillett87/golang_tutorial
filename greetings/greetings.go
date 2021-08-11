package greetings

import (
	"fmt"
	"errors"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error){
	// If no name was given, return an error witha message.
	if name ==""{
		return "", errors.New("empty name")
	}

	//If name was recieved, return a value that embeds the name
	//in a greetings message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}