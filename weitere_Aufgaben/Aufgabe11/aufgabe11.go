package aufgabe11

//LessThan: Erstellen Sie eine Methode LessThan für den Datentyp Card, 
//die prüft, ob c nach den Skat-Regeln kleiner ist als other.

func (c Card) LessThan(other Card) bool {

	if c.Suit == other.Suit && c.Rank < other.Rank{
        return true
    }
return false
}


type Card struct {
    Suit int
    Rank int
}

/*
if c.Rank != other.Rank {
        return c.Rank < other.Rank
    }
    // Wenn die Werte gleich sind, vergleiche die Farben
    return c.Suit < other.Suit

}

type Card struct {
	Suit int
	Rank int
}
*/