package main 

func main(){
  cards := deck{"Ace of Diamonds", newCard() } //slice of type string
  cards = append(cards, "Six of Spades") //take card slice and add to the end of the cards

  cards.print()

}

func newCard() string {
  return "Five of Diamonds"
}


