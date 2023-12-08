package main

type HandAndBid struct {
	hand string
	bid  int
}

const (
	HIGH_CARD = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

func getHandType(cardCount map[uint8]int) int {
	switch len(cardCount) {
	case 1:
		return FIVE_OF_A_KIND
	case 2:
		for _, v := range cardCount {
			if v == 4 {
				return FOUR_OF_A_KIND
			}
		}
		return FULL_HOUSE
	case 3:
		for _, v := range cardCount {
			if v == 3 {
				return THREE_OF_A_KIND
			}
		}
		return TWO_PAIR
	case 4:
		return ONE_PAIR
	default:
		return HIGH_CARD
	}
}
