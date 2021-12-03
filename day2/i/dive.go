package i

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Dive() error{
	//Read the input
	// retrive each line and split into two to get the command and the value

	var horizontal = 0
	var vertical = 0
	// Read the file.
	dat, err := ioutil.ReadFile("day2/input.txt")
	if err != nil {
		return err
	}

	// Split the string by the delimiter
	s := strings.Split(string(dat), "\n")
	//s = s[:len(s)-1]


	for _, val := range s {
		// for each line retrieve the command & the value
		// e.g - forward 9 => ["forward", "9"]

		sl := strings.Split(val, " ")
		num, err := strconv.Atoi(sl[1])
		if err != nil {
			return err
		}
		switch cmd := sl[0]; cmd {
		case "forward":
			horizontal += num
		case "down":
			vertical += num
		case "up":
			vertical -= num
		}
	}

	fmt.Println("Dive-Day2", horizontal*vertical)
	return nil
}
