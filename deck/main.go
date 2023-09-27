package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Card struct {
	Suit  string
	Value string
}

type Deck struct {
	Cards []Card
}

type By func(c1, c2 *Card) bool

func (by By) Sort(cards []Card) {
	cs := &cardSorter{
		cards: cards,
		by:    by,
	}
	sort.Sort(cs)
}

func (s *cardSorter) Len() int {
	return len(s.cards)
}

func (s *cardSorter) Swap(i, j int) {
	s.cards[i], s.cards[j] = s.cards[j], s.cards[i]
}

func (s *cardSorter) Less(i, j int) bool {
	return s.by(&s.cards[i], &s.cards[j])
}

type cardSorter struct {
	cards []Card
	by    func(c1, c2 *Card) bool
}

func createDeck() Deck {
	suits := []string{"S", "H", "D", "C"}
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9"}

	deck := Deck{}
	for _, suit := range suits {
		for _, value := range values {
			deck.Cards = append(deck.Cards, Card{Suit: suit, Value: value})
		}
	}
	return deck
}

func NewDeck() Deck {
	return createDeck()
}

func (d *Deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range d.Cards {
		j := r.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

func main() {
	deck := NewDeck()
	deck.Shuffle()
	fmt.Println(deck)

	// Sort by suit
	suit := func(c1, c2 *Card) bool {
		return c1.Suit < c2.Suit
	}
	By(suit).Sort(deck.Cards)
	fmt.Println(deck)
}
