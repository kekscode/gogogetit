package main

import (
	"io/ioutil"
	"log"

	proto "github.com/golang/protobuf/proto"
)

func save(m proto.Message) {
	out, merr := proto.Marshal(m)
	if merr != nil {
		log.Fatalln("Failed to encode", merr)
	}

	ferr := ioutil.WriteFile("store", out, 0644)
	if ferr != nil {
		log.Fatalln("Failed to save", ferr)
	}
}

func main() {
	c := &Collection{}
	c.Targets = append(c.Targets, &Target{Url: "test"})
	c.Targets = append(c.Targets, &Target{Url: "test2"})
	save(c)
}
