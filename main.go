package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(50)

	log.Println("Guess the Number - Game")

	for {
		log.Printf("Please guess a number between 1 and 50: ")
		guessedNum, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guessedNum = strings.TrimSpace(guessedNum)
		guessedNumI, err := strconv.Atoi(guessedNum)
		if err != nil {
			log.Fatal(err)
		}
		if guessedNumI < randNum {
			log.Println("Try a higher number!")
		}
		if guessedNumI > randNum {
			log.Println("Try a lower number!")
		}
		if guessedNumI == randNum {
			log.Printf("You guessed it! The random number was indeed %d! :-)\n", guessedNumI)
			break
		}

	}

}
