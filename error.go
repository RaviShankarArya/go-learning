package main

import (
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")

	if err != nil {
		// log.Fatalln("Fatal Error Occured:", err)
		// log.Panicln("Panic Error Occured:", err)
		panic(err)
	}
}
