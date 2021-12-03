package i

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// I could create a text file and store the input,
// read from the file and split.
// but I chose to use an online tool to convert the blob of input into a comma seperated string
func SaveTheElves() error {
	var inc int
	var prev = 0

	// Read the file.
	dat, err := ioutil.ReadFile("day1/input.txt")
	if err != nil {
		return err
	}

	// Split the string by the delimiter
	s := strings.Split(string(dat), "\n")

	// Range over slice
	for _, val := range s {
		num, err := strconv.Atoi(val)
		if err != nil {
			return err
		}

		if num > prev {
			inc++
		}

		prev = num
	}

	fmt.Println("SONAR-SWEEP-DAY1", inc -1)
	return nil
}
