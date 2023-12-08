package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	var handAndBids []HandAndBid
	for {
		var handAndBid HandAndBid
		n, err := fmt.Scanf("%s %d\n", &handAndBid.hand, &handAndBid.bid)
		if n == 0 || err != nil {
			break
		}
		handAndBids = append(handAndBids, handAndBid)
	}
	cards := map[uint8]int{'J': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'Q': 12, 'K': 13, 'A': 14}
	compareCards := func(a, b string) int {
		for i := 0; i < 5; i++ {
			if cards[a[i]] != cards[b[i]] {
				return cmp.Compare(cards[a[i]], cards[b[i]])
			}
		}
		return 0
	}
	getCardCount := func(s string) map[uint8]int {
		uniqueChars := make(map[uint8]int)
		for i := 0; i < len(s); i++ {
			uniqueChars[s[i]]++
		}
		jCount, ok := uniqueChars['J']
		if ok {
			if jCount < 5 {
				maxCountCard, maxCount := uint8(0), 0
				for card, count := range uniqueChars {
					if card != 'J' {
						if count > maxCount {
							maxCount = count
							maxCountCard = card
						}
					}
				}
				uniqueChars[maxCountCard] += jCount
				delete(uniqueChars, 'J')
			}
		}
		return uniqueChars
	}
	compareHands := func(a, b HandAndBid) int {
		aHandType := getHandType(getCardCount(a.hand))
		bHandType := getHandType(getCardCount(b.hand))
		if aHandType != bHandType {
			return cmp.Compare(aHandType, bHandType)
		}
		return compareCards(a.hand, b.hand)
	}
	slices.SortFunc(handAndBids, compareHands)
	sum := 0
	for i := 0; i < len(handAndBids); i++ {
		sum += handAndBids[i].bid * (i + 1)
	}
	fmt.Printf("%d\n", sum)
}
