package average_salary_excluding_the_minimum_and_maximum_salary

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func average(salary []int) float64 {
	lastIndex := len(salary) - 1
	sortedSalary := quickSort(salary, 0, lastIndex)
	count := len(sortedSalary[1:lastIndex])
	total := 0
	for _, salary := range sortedSalary[1:lastIndex] {
		total = total + salary
	}
	return float64(total) / float64(count)
}
