package aufgabe3

//AllAbc: Erstellen Sie eine Funktion AllAbc, 
//die alle Elemente zurückgibt, die mit "abc" beginnen.

func AllAbc(list []string) []string {

result := []string{}

 for _, val := range list {
	if len(val) >= 3 && val[:3] == "abc" {
		result = append(result, val)
	}
 }

 return result
}





/*
result := []string{}

	for _ , val := range list {
		if len(val) >= 3 && val[:3] == "abc" {
			result = append(result, val)
		}
	}


	return result
*/