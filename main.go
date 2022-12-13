package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	userDetails := &Details{
		Name: "kalyan",
		Age:  22,
	}
	data, err := proto.Marshal(userDetails)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)

	getUser := &Details{}
	err = proto.Unmarshal(data, getUser)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Age of the user: ", getUser.GetAge())
	fmt.Println("Name of the User: ", getUser.GetName())
}
