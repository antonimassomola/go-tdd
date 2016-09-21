package hello

func Hello(i string) string {
	o := "Hello "
	o += i
	o += "!"
	return o
}

func Sum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}