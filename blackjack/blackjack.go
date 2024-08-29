package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	res := 0
	switch {
	case card == "ace":
		res = 11
	case card == "two":
		res = 2
	case card == "three":
		res = 3
	case card == "four":
		res = 4
	case card == "five":
		res = 5
	case card == "six":
		res = 6
	case card == "seven":
		res = 7
	case card == "eight":
		res = 8
	case card == "nine":
		res = 9
	case card == "ten":
		res = 10
	case card == "jack":
		res = 10
	case card == "queen":
		res = 10
	case card == "king":
		res = 10
	}

	return res
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

	hand := ParseCard(card1) + ParseCard(card2)
	dealerHand := ParseCard(dealerCard)
	decision := ""

	switch {
	case card1 == "ace" && card1 == card2:
		decision = "P"
	case hand == 21:
		decision = "W"
	case hand >= 17 && hand <= 20:
		decision = "S"
	case hand >= 12 && hand <= 16 && dealerHand < 7:
		decision = "S"
	case hand >= 12 && hand <= 16 && dealerHand >= 7:
		decision = "H"
	case hand <= 11:
		decision = "H"
	}

	return decision

}
