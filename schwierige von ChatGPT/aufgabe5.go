package main

func IsPalindromeRecursive(s string) bool {

	if len(s) == 0 {
		return true
	}
	if len(s) == 1 {
		return true
	}
	if s[0] == s[len(s)-1] {
		return IsPalindromeRecursive(s[1 : len(s)-1])
	}
	return false
}
