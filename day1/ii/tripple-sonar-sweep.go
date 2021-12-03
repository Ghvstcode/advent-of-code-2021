package ii

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)


func SaveTheThreeElves() error {
	var result []int
	var grouped []int
	var inc int
	var prev = 0

	// Read the file.
	dat, err := ioutil.ReadFile("day1/input.txt")
	if err != nil {
		return err
	}


	// Split the string by the delimiter
	s := strings.Split(string(dat), "\n")

	// convert slice of string to slice of int
	// This isn't the most efficient way to go about this, btw.
	for _, val := range s {
		num, err := strconv.Atoi(val)
		if err != nil {
			return err
		}
		result = append(result, num)
	}

	//group numbers by sum of three
	for i := 0; i < len(result); i++ {
		// check to make sure the 2nd and 3rd numbers after the index
		//are not greater than the length of the slice
		if (i+2 < len(result)) && (i+1 < len(result)) {
			grouped = append(grouped, result[i]+result[i+1]+result[i+2])
		}
	}

	// Range over slice of three and do the comparison similar to what we did
	//in the first section
	for _, num := range grouped {
		if num > prev {
			inc++
		}

		prev = num
	}

	fmt.Println("TRIPPLE-SONAR-SWEEP-DAY1(2)", inc-1)
	return nil
}
