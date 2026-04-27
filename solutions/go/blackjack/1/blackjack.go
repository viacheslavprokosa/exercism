// Package blackjack implements a blackjack game.
package blackjack

// Cards is the map whose keys are the card names and values are the card values
var cards = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
	"other": 0,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return cards[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	dealer := ParseCard(dealerCard)
	me := ParseCard(card1) + ParseCard(card2)
	switch {
	case me == 22:
		return "P"
	case me == 21 && dealer < 10:
		return "W"
	case me == 21 && dealer > 9:
		return "S"
	case me >= 17 && me <= 20:
		return "S"
	case (me > 11 && me < 17) && dealer < 7:
		return "S"
	case (me > 11 && me < 17) && dealer > 6:
		return "H"
	case me < 12:
		return "H"
	default:
		return ""
	}
}
