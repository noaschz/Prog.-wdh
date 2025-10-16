package byterechnen

// Dieser Funktion soll ein Zeichen übergeben werden, und einen String zurückgeben
// Ist dieses Zeichen ein Großbuchstabe, soll dieser in einen Kleinbuchstaben umgewandelt werden
// Ist dieses Zeichen ein Kleinbuchstabe, soll dieser in einen Großbuchstaben umgewandelt werden
// Ist dieses Zeichen kein Buchstabe, soll ein Leerstring zurückgegeben werden
func UebungBuchstabe(x byte) string {

	if x >= 'A' && x <= 'Z' {
		return string(x + 32)
	} else if x >= 'a' && x <= 'z' {
		return string(x - 32)

	}
	return "kein buchstabe"

}

// Dieser Funktion soll ein Zeichen übergeben werden, und eine Zahl zurückgeben
// Ist dieses Zeichen eine Zahl, soll diese Zahl als Int zurückgegeben werden
// Ist dieses Zeichen keine Zahl, soll -1 zurückgegeben werden
func UebungZahl(x byte) int {
	if x >= '0' && x <= '9' {
		return int(x - 48)
	}
	return -1
}
