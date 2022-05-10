package main

import (
	"bufio"
	"fmt"
	"morse/converter"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Morse code converter, use characters that are found in the International Morse table.")
	fmt.Print("Type in the message you want to see in Morse code: ")
	message, _ := reader.ReadString('\n')
	converter.Converter(strings.TrimSpace(strings.ToUpper(message)))

}
