package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Enter a name please")
		return
	}
	name := os.Args[1]
	name = strings.ToUpper(name)
	moods := [...]string{
		"happy", "sad", "awe", "terrible", "loving", "angry",
	}
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(moods))
	fmt.Printf("%s is %s\n", name, moods[n])
}
