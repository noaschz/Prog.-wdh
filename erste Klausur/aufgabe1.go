package own


func LongestAbc(list []string) string {

	longestpos := -1
	longestlen := -1

	for pos, val := range list {
		currentlen := len(val)
		if currentlen > 3 && val[:3] == "abc" && 
		currentlen > longestlen{
				longestlen = currentlen
				longestpos = pos
		}

	}

	if longestpos != -1 {
		return list[longestpos]
	}
	
return ""

}









/*
	longestlen := -1
	longestpos := -1

	for i := 0; i < len(list); i++ {
		actual := list[i]
		if len(actual) >= 3 && actual[0:3] == "abc" {
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
*/