package utils

import (
	"io/ioutil"
	"log"
)

func ReadFileAsString(path string) string {
	bytes, e := ioutil.ReadFile(path)
	if e != nil {
		log.Fatalln(e)
	}
	return string(bytes)
}
