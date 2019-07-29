package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/vancluever/depextract"
)

func main() {
	if len(os.Args) > 2 {
		log.Fatalf("usage: %s FILE", os.Args[0])
	}

	file := os.Args[1]
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	result, err := depextract.ParseGopkgLock(file, data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(strings.Join(result, " "))
}
