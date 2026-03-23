package arrays

func Sum(arr []int) int {
	temp := 0
	
	for _, number := range arr{
		temp += number
	}

	return temp
}

func SumAll(numbersToSum ... []int) (sums []int) {
	for _, slice := range numbersToSum {
		sums = append(sums, Sum(slice))
	}
	return
}

func SumAllTails(numbersToSum ... []int) (sums []int) {
	for _, slice := range numbersToSum {
		if len(slice) == 0{
			sums = append(sums, 0)
		}else {
			sums = append(sums, Sum(slice[1:]))
		}
	}
	return
}