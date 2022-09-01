package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	x, y, num := 0, 0, 0
	var num1 []int

	input, err := (bufio.NewReader(os.Stdin).ReadString('\n'))
	if err != nil {
		log.Fatal(err)
	}
	input2, err := (bufio.NewReader(os.Stdin).ReadString('\n'))
	if err != nil {
		log.Fatal(err)
	}
	for i, val := range strings.Split(strings.TrimSuffix(strings.Trim(strings.TrimSuffix(strings.Trim(input, " "), "\r\n"), " "), "\n"), " ") {
		num1[i], err = strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
	}

	for i, val := range strings.Split(strings.TrimSuffix(strings.Trim(strings.TrimSuffix(strings.Trim(input2, " "), "\r\n"), " "), "\n"), " ") {
		num, err = strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		if num > num1[i] {
			y++
		}
		if num < num1[i] {
			x++
		}
	}
	fmt.Println("[{},{}}]", x, y)
}
