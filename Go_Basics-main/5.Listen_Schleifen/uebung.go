package listen_schleifen

import "fmt"

// X1. Gebe mittels einer Schleife nur die geraden Zahlen der intlist in einer Zeile aus (Tipp benutzte If und Modulo)
// X2. Gebe in einer neuen Zeile mittels einer Schleife aus ob die Zahlen der intliste gleich, kleiner oder größer 100 sind (If-Else)
// XAusgabe: = für gleich, < für kleiner > für größer
// 3. Erzeuge eine bytelist welche die selbe Länge der intlist hat
// 4. Füge mittels einer Schleife die Werte aus der intlist als bytes in die bytelist
// 5. Gebe mittels Schleife in einer neuen Zeile die bytelist im Terminal aus, die Zeichen sollen lesbar sein und in einer Zeile





func Uebung() {
	intlist := []int{72, 97, 108, 108, 111, 32, 87, 101, 108, 116}

	for _, i := range intlist {
		if i%2 == 0 {
			fmt.Print(i)
		}
	}

	fmt.Println("")
	for _, i := range intlist {
		if i == 100 {
			fmt.Print("=")
		
			} else if  i < 100{
				fmt.Print("<")
			} else {
			if i > 100 {
				fmt.Print(">")
			}
		}

		}

bytelist := make([]byte, len(intlist))

for i := 0; i < len(bytelist); i++ {
	bytelist[i] = byte(intlist[i])
}

fmt.Println("")
for _, b := range bytelist {
	fmt.Print(string(b))
}

}

	