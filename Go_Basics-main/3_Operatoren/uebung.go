package operatoren

import "fmt"

// Aufgaben:
// X1. Addiere intvar1 mit intvar2 und speichere das Ergebnis in der neuen Variable intvar3
// X2. Gebe das Ergebnis der Multiplikation von intvar1 und intvar2 im Terminal aus
// X3. Gebe das Ergebnis der Division von intvar3 durch intvar1 im Terminal aus
// X4. Gebe den Rest der Diviosion von intvar2 und intvar3 im Terminal aus
// X5. Addiere auf intvar3 1 auf und gebe es anschließend in der Konsole aus
// X6. Dividiere nun 6 durch 4 und gebe es mit Kommastelle im Terminal aus
// X7. Speicher in der Variable string1 "Hallo" und in string2 "Welt"
// X8. Gebe nun die Zwei Variablen im Temrinal aus, einmal mit und einmal ohne Leerzeichen
// 9. Gebe nun "Den Wert: x hat die Variable stringvar2" im Terminal aus und lass für x den Wert von stringvar2 einsetzten
// Im Test kannst du dir anschauen wie der Output aussehen soll, darunter findest du wieder eine lösung
func Uebung() {

	intvar1 := 2
	intvar2 := 6
	intvar3 := intvar1 + intvar2
	string1 := "Hallo"
	string2 := "Welt"

	fmt.Println(intvar3)
	fmt.Println(intvar3/intvar1)
	fmt.Println(intvar2%intvar1)
	intvar3++
	fmt.Println(intvar3)
	fmt.Println(6.0/4.0)
	fmt.Println(string1, string2)
	fmt.Println(string1 + string2)
	fmt.Printf("Den Wert: %s hat die Variable stringvar2", string2)
}
