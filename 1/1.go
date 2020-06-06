package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile("[0-9]+")

	fmt.Print("Give Me Sentences : ")
	inputUser, _ := reader.ReadString('\n')

	output := re.FindAllString(inputUser, -1)

	for _, val := range output {
		fmt.Print(val)
	}
}
