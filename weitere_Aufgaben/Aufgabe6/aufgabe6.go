package aufgabe6

//ExcludeStringsOutside: Erstellen Sie eine Funktion ExcludeStringsOutside, 
//die alle Elemente außerhalb von first und last zurückgibt.

func ExcludeStringsOutside(list []string, start string, end string) []string {

	firstpos := 0
	lastpos := 0

	for pos, s := range list {

		if s == start{
			firstpos = pos
		}

		if s == end {
			lastpos = pos
		}
	}

	if firstpos >= lastpos {
		return []string{}
	}

	return append(list[:firstpos], list[lastpos +1 :]...)

	}


/*





	firstpos := -1
	lastpos := -1
	
	for pos, s := range list{
	
		if s == start {
			firstpos = pos
		}
	
		if s == end {
			lastpos = pos
		}
	}
	
		if firstpos >= lastpos {
			return []string{}
		}
	
		return append(list[:firstpos], list[lastpos +1:]...)
*/