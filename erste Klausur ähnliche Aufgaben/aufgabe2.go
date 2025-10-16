package main

func IncludeBetween(list []string, x string, y string) []string {
	firstpos := -1
	lastpos := -1

	for i := 0; i < len(list); i++ {
		if list[i] == x {
			firstpos = i
		}
		if list[i] == y {
			lastpos = i
		}
	}

	if firstpos != -1 && lastpos != -1 && firstpos < lastpos {
		return list[firstpos+1 : lastpos]
	}
	return []string{}
}

/*
firstpos := -1
	lastpos := -1

	for i := 0; i < len(list); i++ {
		if list[i] == x {
			firstpos = i
		}
		if list[i] == y {
			lastpos = i
		}
	}

	if firstpos != -1 && lastpos != -1 && firstpos < lastpos {
		return list[firstpos+1 : lastpos]
	}
	return []string{}
*/