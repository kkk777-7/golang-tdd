package slices

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersAry ...[]int) []int {
	var sumAry []int
	for _, numbers := range numbersAry {
		sumAry = append(sumAry, Sum(numbers))
	}
	return sumAry
}

func SumAllTails(numbersAry ...[]int) []int {
	var sumTailAry []int
	var num int
	for _, numbers := range numbersAry {
		if len(numbers) == 0 {
			num = 0
		} else {
			num = Sum(numbers[1:])
		}

		sumTailAry = append(sumTailAry, num)
	}
	return sumTailAry
}
