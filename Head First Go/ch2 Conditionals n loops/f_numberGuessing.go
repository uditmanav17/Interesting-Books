package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// target := rand.Intn(100) + 1 // gives same number each time
	// fmt.Println(target)
	// using rand.seed to overcome this
	seconds := time.Now().Unix() // to get time in seconds
	// fmt.Println(seconds)
	rand.Seed(seconds) // set seed
	target := rand.Intn(100) + 1
	// fmt.Println("A number b/w 1 and 100 is :", target)
	success := false
	for guesses := 10; guesses > 0; guesses-- {
		fmt.Println("Yoy have", guesses, " guess remaining!")
		// get input from user
		reader := bufio.NewReader(os.Stdin) // create reader object
		fmt.Println("Make a Guess: ")
		input, err := reader.ReadString('\n') // take user input
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)  // remove spaces
		guess, err := strconv.Atoi(input) // convert string to int
		if err != nil {
			log.Fatal(err)
		}

		// comparing guess with user input
		if guess < target {
			fmt.Println("Guess too LOW")
		} else if guess > target {
			fmt.Println("Guess too HIGH")
		} else {
			success = true
			fmt.Println("Perfect!\nYou WON!")
			break
		}
	}

	if !success {
		fmt.Println("Sorry, you couldn't guess the number :(\nIt was", target)
	}

}
