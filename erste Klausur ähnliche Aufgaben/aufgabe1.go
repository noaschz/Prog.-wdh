package main

func ShortestXYZ(list []string) string {
	shortestpos := -1
	shortestlen := 1000

	for pos, val := range list {
		currentLen := len(val)
		if currentLen >= 3 && val[:3] == "xyz" && currentLen < shortestlen {
			shortestlen = currentLen
			shortestpos = pos
		}
	}

	if shortestpos != -1 {
		return list[shortestpos]
	}

	return ""
}

/*
shortestslen := 10000
	shortestpos := 10000

	for i := 0; i < len(list); i++ {
		actual := list[i]
		if len(actual) >= 3 && actual[0:3] == "xyz" {
			currentlen := len(actual)
			if currentlen < shortestslen {
				shortestslen = currentlen
				shortestpos = i
			}
		}
	}
	if shortestpos != 10000 {
		return list[shortestpos]
	} else {
		return ""
	}
*/