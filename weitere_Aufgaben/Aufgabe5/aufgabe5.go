package aufgabe5

//IncludeStringsBetween: Erstellen Sie eine Funktion IncludeStringsBetween, 
//die alle Elemente zwischen first und last zurückgibt
// KA NR 2
func IncludeStringsBetween(list []string, start, end string) []string {

	firstpos := 0
	lastpos := 0
	
	for pos, val := range list {
		if val == start {
			firstpos = pos
		}

		if val == end {
			lastpos = pos
		}
	}

	if firstpos >= lastpos {
		return []string{}
	}


	return list[firstpos +1 : lastpos]
}







/*
firstpos := 0
	lastpos := 0

	for pos , s := range list{

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


return list[firstpos +1 : lastpos]
*/
