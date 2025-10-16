package aufgabe14

//Intersection: Erstellen Sie eine Funktion Intersection, 
//die die Schnittmenge der beiden Listen zurückgibt
// KA nr 6 
func Intersection(a, b []int) []int {

if len(a) == 0 {
	return b
}

if len(b) == 0{
	return a
}

result := []int{}

for _, el1 := range a {
	if Contains(el1, b) {
		result = append(result, el1)
	}
}


return result

}

func Contains(e int, l []int) bool {
	for _, el := range l {
		if el == e {
			return true
		}
	}

	return false
}

/*
func Intersection(a, b []int) []int {
    result := []int{}
        if len(a) == 0 {
            return b}
        if len(a) == 0 {
            return b}
        for _, el1 := range a {
            if Contains(el1,b){
                result = append(result, el1)}}
    return result}
func Contains (e int , l []int) bool {
    for _, el := range l {
        if el == e {
            return true}}
return false}

*/