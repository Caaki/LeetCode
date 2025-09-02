package main

func main(){}


func isPowerOfFour(n int) bool {
    if n == 1 {
        return true
    }
	if n >=4 && 1073741824%n == 0 && n % 3 == 1 {
		
        return true
	}

	return false
}
