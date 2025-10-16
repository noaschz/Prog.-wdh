package main

func LongestPrefixMatch(list []string, prefix string) string {

	longestlen := -1
	longestpos := -1
	prelen := len(prefix)

	for i := 0; i < len(list); i++ {
		actual := list[i]
		if len(actual) >= prelen && actual[0:prelen] == prefix {
			currentlen := len(actual)
			if currentlen > longestlen {
				longestlen = currentlen
				longestpos = i
			}
		}
	}
	if longestpos != -1 {
		return list[longestpos]
	} else {
		return ""
	}
}
