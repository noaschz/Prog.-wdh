package aufgabe12

//EqualTo: Erstellen Sie eine Methode EqualTo für den Datentyp Card, 
//die prüft, ob c nach den Skat-Regeln gleich other ist

func (c Card) EqualTo(other Card) bool {

	return c.Suit==other.Rank && c.Rank == other.Rank

}

type Card struct {
	Suit int
	Rank int

}
