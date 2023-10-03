package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

/* "bufio"
"log"
"os" */ /* "reflect"
"strconv" */ /*  "strings" */ /* "unicode/utf8" */ /* "math/rand"
"time" */ /* "math" */ /* "bufio"
"log"
"math/rand"
"os"
"strconv"
"strings" */ /* "bufio"
"errors"
"fmt"
"io"
"log"
"os"
"strconv" */

var pl = fmt.Println

func main() {

	/* pl("What's your name")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("hello", name)
	} else {
		log.Fatal(err)
	}
	*/

	// pl(reflect.TypeOf(25))

	/*
		v1 := "50000"
		v2, err := strconv.Atoi(v1)
		pl(v2, err, reflect.TypeOf(v2))
	*/
	/*
		v3 := 50000
		v4 := strconv.Itoa(v3)
		pl(v4, reflect.TypeOf(v4))
	*/
	/*
		v5 := "0.0005"
		v6, err := strconv.ParseFloat(v5, 64)
		pl(v6, err, reflect.TypeOf(v6))
	*/

	/*
		s1 := "A word"
		replacer := strings.NewReplacer("A", "Another")
		s2 := replacer.Replace(s1)
		pl(s1, "=>", s2)
		pl(len(s2))

		runes := []rune(s2)
		pl(len(runes))

		s3 := string(runes)
		pl(len(s3))

		pl("Contains 'Another':", strings.Contains(s3, "Another"))
		pl("Index of 'o':", strings.Index(s3, "o"))
		pl("Replace:", strings.Replace(s3, "o", "0", -1), "<=", s3)

		s3 = "\nOther words\n"
		pl(s3)
		s3 = strings.TrimSpace(s3)
		pl(s3)
	*/

	/*
		s4 := "a-b-c-d"
		delimiter := "-"
		prefix := "a-"
		pl("Split", s4, "by", delimiter, "=>", strings.Split(s4, delimiter))
		pl(s4, "ToUpper", "=>", strings.ToUpper(s4))
		pl(s4, "HasPrefix", prefix, "=>", strings.HasPrefix(s4, prefix))
	*/

	/* 	runeStr := "rune string"
	   	pl("Rune Count for", runeStr, "=>", utf8.RuneCountInString(runeStr))

	   	for i, rune := range runeStr {
	   		pl(i, rune)
	   		fmt.Printf("%d : %#U : %c\n", i, rune, rune)
	   	}
	*/

	/* 	now := time.Now()
	   	pl(now)
	   	pl(now.Year(), now.Month())

	   	//deprecated to use timestamp as seed for rand
	   	//	seedSecs := time.Now().Unix()
	   	//	rand.Seed(int64(seedSecs)) // seed the random number generator with current unix timestamp in seconds (seconds since Jan 1

	   	seed := int64(42) // Create a local random generator with the specified seed
	   	rng := rand.New(rand.NewSource(seed))

	   	randNum := rng.Intn(50) + 1
	   	pl("Random:", randNum) */

	/* 	pl(math.Abs(-40))
	   	pl(math.Pow(4, 2))
	   	pl(math.Sqrt(25))
	   	pl(math.Log2(8))
	   	pl(math.Log(math.Exp(2))) */

	//	fmt.Printf("%s %d %c %f %t %o %x\n", "string", 100, 'A', 3.14, true, 9, 17)
	//  fmt.Printf("%4f %.2f %4.f %f\n", math.Pi, math.Pi, math.Pi, math.Pi)
	//  fmt.Printf("%9f %.2f %9.f %f\n", math.Pi, math.Pi, math.Pi, math.Pi)

	/*
	   	pl("------\tWelcome to the Number Guessing Game\t-------")
	   	pl("\tPlease guess a number between 1 and 50")
	   	seed := int64(42) // Create a local random generator with the specified seed
	   	rng := rand.New(rand.NewSource(seed))
	   	randNum := rng.Intn(50) + 1
	   	// var guessNum int
	   loopOfGame:
	   	for {
	   		fmt.Print("Please input your guess:")
	   		reader := bufio.NewReader(os.Stdin)
	   		guessStr, err := reader.ReadString('\n')
	   		if err != nil {
	   			log.Fatal(err)
	   		}
	   		guessNum, err := strconv.Atoi(strings.TrimSpace(guessStr))

	   		if err != nil {
	   			log.Fatal(err)
	   		}

	   		switch {
	   		case guessNum > randNum:
	   			pl("Your guess is too high!")
	   		case guessNum < randNum:
	   			pl("Your guess is too low!")
	   		case guessNum == randNum:
	   			pl("Bingo! Your guess is correct!")
	   			break loopOfGame
	   		}
	   	}
	*/
	/*
		file, err := os.Create("data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		arrayPrimesInt := []int{2, 3, 5, 7, 11}
		arrayPrimesStr := []string{}
		for _, i := range arrayPrimesInt {
			arrayPrimesStr = append(arrayPrimesStr, strconv.Itoa(i))
		}
		for _, PrimeStr := range arrayPrimesStr {
			_, err := io.WriteString(file, " "+PrimeStr+"\n")
			if err != nil {
				errors.New("Error from writing strings into file")
			}
		}

		file, err = os.Open("data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file) // create a new scanner for the file
		for scanner.Scan() {
			pl("Prime: ", scanner.Text())
		}
	*/

	_, err := os.Stat("data.txt")
	if errors.Is(err, os.ErrNotExist) {
		pl("File doesn't exist")
	}

	file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := file.WriteString(" 13\n"); err != nil {
		log.Fatal(err)
	}

}
