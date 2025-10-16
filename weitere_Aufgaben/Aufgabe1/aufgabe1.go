package aufgabe1

//LongestAbc: Erstellen Sie eine Funktion LongestAbc, die das längste
// Element zurückgibt, das mit "abc" beginnt
//KA NR. 1
func LongestAbc(list []string) string {

	longestpos := -1
	longestlen := -1
	
	for pos , val := range list {
		currentlen := len(val)

		if currentlen >= 3 && val[:3] == "abc" &&
		currentlen > longestlen{
			longestlen = currentlen
			longestpos = pos
		}
	}

		if longestpos != -1{
			return list[longestpos]
		}

		
		return ""
	
}


	/*
	longestpos := -1
	longestlen := -1
	
	for pos , val := range list {
		currentlen := len(val)

		if currentlen >= 3 && val[:3] == "abc" && currentlen > longestlen{
			longestlen = currentlen
			longestpos = pos

		}
	}

if longestpos != -1 {
	return list[longestpos]
}

		return ""
	*/



