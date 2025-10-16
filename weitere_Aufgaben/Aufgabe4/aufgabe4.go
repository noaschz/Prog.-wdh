package aufgabe4

//ShortestAbcWithIndex: Erstellen Sie eine Funktion ShortestAbcWithIndex, 
//die das kürzeste Element und dessen Index zurückgibt, das mit "abc" beginnt.

func ShortestAbcWithIndex(list []string) (string, int) {


	shortestpos := -1
	shortestlen := 1000
	index := -1


		for pos, val := range list {
			currentlen := len(val) 
			if currentlen >= 3 && val[:3] == "abc" &&
			currentlen < shortestlen {
				shortestlen = currentlen
				shortestpos = pos
				index = pos
			}
			
		}
	if shortestpos != -1 {
		return list[shortestpos], index
	}

	return "", -1
}






/*
if currentlen >= 3 && val[:3] == "abc" && currentlen < shortestlen{
			shortestlen = currentlen
			shortestpos = pos
			index = pos
		}
	}

	if shortestpos != -1 {
		return list[shortestpos], index
	}

	return "", -1
*/