package tickertype

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// ErrNotSupported is an error for an unsupported asset type
	ErrNotSupported = errors.New("received unsupported asset type")
)

// Item stores the asset type
type Item string

// Items stores a list of assets types
type Items []Item

// Const vars for asset package
const (
	Csv      = Item("csv")
	DataBase = Item("database")
)

var supported = Items{
	Csv,
	DataBase,
}

// returns an Item to string
func (a Item) String() string {
	return string(a)
}

// Strings converts an ticker type array to a string array
func (a Items) Strings() []string {
	var tickerType []string
	for x := range a {
		tickerType = append(tickerType, string(a[x]))
	}
	return tickerType
}

// JoinToString joins an tickerType type array and converts it to a string
// with the supplied separator
func (a Items) JoinToString(separator string) string {
	return strings.Join(a.Strings(), separator)
}

// New takes an input matches to relevant package tickerType
func New(input string) (Item, error) {
	input = strings.ToLower(input)
	for i := range supported {
		if string(supported[i]) == input {
			return supported[i], nil
		}
	}
	return "", fmt.Errorf("%w %v, only supports %v",
		ErrNotSupported,
		input,
		supported)
}
