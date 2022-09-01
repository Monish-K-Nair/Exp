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
	num, sum := 0, 0
	input, err := (bufio.NewReader(os.Stdin).ReadString('\n'))
	if err != nil {
		log.Fatal(err)
	}
	for _, val := range strings.Split(strings.TrimSuffix(strings.Trim(strings.TrimSuffix(strings.Trim(input, " "), "\r\n"), " "), "\n"), " ") {
		num, err = strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		sum += num
	}
	fmt.Println(sum)
}
