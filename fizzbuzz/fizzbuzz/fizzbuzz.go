package fizzbuzz

import "strconv"

func Run(intList ...int) []string {
	var strSlice []string
	for _, i := range intList {
		if i%3 == 0 && i%5 == 0 {
			strSlice = append(strSlice, "FizzBuzz")
		} else if i%3 == 0 {
			strSlice = append(strSlice, "Fizz")
		} else if i%5 == 0 {
			strSlice = append(strSlice, "Buzz")
		} else {
			strSlice = append(strSlice, strconv.Itoa(i))
		}
	}
	return strSlice
}
