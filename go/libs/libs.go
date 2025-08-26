package libs

import "strconv"

func CrivoEratostenes(n int64) (int64, bool) { // isprime de O(sqrt(n))
	if n <= 1 {
		return n, false
	}
	if n <= 3 {
		return n, true
	}
	if n%2 == 0 || n%3 == 0 {
		return n, false
	}

	// Só testa até √n
	for i := int64(5); i*i <= n; i += 6 {
		if n%i == 0 || n%(i+int64(2)) == 0 {
			return n, false
		}
	}
	return n, true
}

func CheckPalindrome(n int) (int, bool) {
	str := strconv.Itoa(n)
	reverse := ""
	for i := len(str) - 1; i >= 0; i-- {
		reverse += string(str[i])
	}
	if reverse == str {
		return n, true
	}
	return n, false
}

func CollatzLength(n int) int {
	count := 0
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		count++
	}
	return count
}
