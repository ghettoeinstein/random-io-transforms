package main

import (
	"bufio"
	"fmt"
	_ "io"
	"math/rand"
	"os"
	"strings"
	"time"
)


var count map[string]int

func SetupCount() {
	count = map[string]int{}
}

func SetupRandom() {
	SEED := time.Now().Unix()

	rand.Seed(SEED)

}

func transform() bool {
	max := 2
	min := 0
	result := rand.Intn(max-min) + min
	fmt.Println(result)
	if result == 0 {
		return false
	} else {
		return true
	}
}

func PrettyPrint(p map[string]int) {
	fmt.Println("Current Count")
	fmt.Println("-------------")

	for k, _ := range p {
		titleCaseKey := strings.ToTitle(k)
		fmt.Printf("%v: (%d)\n", titleCaseKey, p[k])

	}

	fmt.Println("")
}
func reverse(s string) string {
	length := len(s)
	result := make([]rune, length)
	for _, r := range s {
		length--
		result[length] = r
	}

	//randomly transform the result
	if transform() {
		if transform() {
			count["upper"]++
			PrettyPrint(count)
			return strings.ToUpper(string(result))

		} else {
			count["lower"]++
			PrettyPrint(count)
			return strings.ToLower(string(result))
		}

	} else {
		count["none"]++
		PrettyPrint(count)

		return string(result)

	}
	return string(result)
}

func CheckForQuit(s string) {
	lower := strings.ToLower(s)

	if lower == "quit" {
		os.Exit(1)
	}

}

func PrintWelcome() {
	fmt.Fprintln(os.Stdout, "Pid is: ", os.Getpid())

}

func main() {

	SetupRandom()
	SetupCount()
	PrintWelcome()

	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		CheckForQuit(scanner.Text())
		fmt.Println(reverse(scanner.Text()))
	}

	reversed := reverse("Hello")
	fmt.Println(string(reversed))

}
