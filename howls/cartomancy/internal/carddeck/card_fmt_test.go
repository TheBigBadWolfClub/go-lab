package carddeck

import (
	"fmt"
	"testing"
)

// check the layout resulting of implementing some fmt functions
func TestCard_GoString(t *testing.T) {
	card := Card{
		CardID: A.cardID(),
		SuitID: SPADES.suiteID(),
	}
	fmt.Printf("%s\n", card)
	fmt.Printf("%q\n", card)
	fmt.Printf("%#v\n", card)
	fmt.Printf("%v\n", card)
}
