package subtract_the_product_and_sum_of_digits_of_an_integer

func subtractProductAndSum(n int) int {
	digits := 1
	productOfDigits := 0
	sumOfDigits := 0
	for {
		if digits > n {
			break
		}
		num := (n / digits) % 10
		digits = digits * 10
		sumOfDigits = sumOfDigits + num
		if digits == 10 && productOfDigits == 0 {
			productOfDigits = num
			continue
		}
		productOfDigits = productOfDigits * num
	}
	return productOfDigits - sumOfDigits
}
