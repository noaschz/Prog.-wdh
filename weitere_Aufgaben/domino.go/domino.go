package aufgabe10

// Domino: Definiert einen Dominostein mit zwei Werten.
// IsValidDominoChain: Prüft, ob die gegebene Liste von Dominosteinen
// eine gültige Kette bildet. Eine Kette ist gültig, wenn der rechte Wert eines
// Steins mit dem linken Wert des nächsten Steins übereinstimmt.



func IsValidDominoChain(dominos []Domino) bool {
    for i:= 0; i < len(dominos)-1; i++{
 if dominos[i].right != dominos[i+1].left{
    return false
 }
}
   return true
}

type Domino struct{
    left int
    right int 
}




/*
   for i := 0; i < len(dominos)-1; i ++{
    if dominos[i].Right != dominos[i+1].Left{
        return false
    } 
  }
return true

}

type Domino struct {
    Left int
    Right int
}

*/