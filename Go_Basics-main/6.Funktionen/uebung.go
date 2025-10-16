package funktionen

import "fmt"

// In dieser Übung bekommt die Funktion Uebung eine Liste von Integern übergeben, mit der sie arbeiten sollen
// 1. Erstelle eine Funktion, welcher eine int Variable übergeben wird und prüft ob die Übergebene Zahl gerade ist
// Ist dies der Fall, soll die Funktion true zurückgeben, ansonsten false, also braucht die Funktion den Rückgabewert bool
// 2. Übergeben sie mittels einer Schleife jede Zahl der intlist der Funktion aus 1. und geben sie die Ergebnisse in eiener Zeile aus, ohne Leerzeichen
// 3. Erstelle eine Funktion, welcher eine int Variable übergeben wird und den Wert mit 2 multipliziert und zurückgibt
// 4. Übergeben sie mittels einer Schleife jede Zahl der intlist der Funktion aus 3. und ersetzte den jewiligen Wert der intlist mit dem Ergebnis der Funktion
// 5. Erstelle eine Funktion, welcher eine int Liste übergeben wird und diese mittels einer Schleife im Terminal in einer neuen Zeile ausgibt
// 6. Übergebe der Funktion aus 5. die intlist

func Uebung(intlsit []int) {

	
	for i := 0; i <len(intlsit); i++ {
		fmt.Print(gerade(intlsit[i]))
	
	}


	for i :=0; i < len(intlsit); i++ {
		intlsit[i] = multiplizieren(intlsit[i])

	}
	
	return ausgabe(intlsit)

}

func gerade(zahl int) bool {
	return zahl%2 == 0
}


func multiplizieren(zahl int) int {

	return zahl * 2

}

func ausgabe(intlsit []int) {
	fmt.Print()
	for i :=0; i < len(intlsit); i++ {
		fmt.Println(intlsit[i]) 
	}
}

