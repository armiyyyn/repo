func isPalindrome(x int) bool {
	res := true
	fmt.Println("enter the num")
	var x int
	fmt.Scan(&x)
	fmt.Println("enter the same number in reverse")
	var m int
	fmt.Scan(&m)
	if x == m {
		res = true
	} else {
		res = false
	}
	fmt.Println(res)
}
