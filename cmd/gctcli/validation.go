package main

import (
	"errors"
	"github.com/thrasher-corp/gocryptotrader/exchanges/tickertype"
	"strings"

	"github.com/thrasher-corp/gocryptotrader/exchanges/asset"
)

var (
	errInvalidPair  = errors.New("invalid currency pair supplied")
	errInvalidAsset = errors.New("invalid asset supplied")
	errInvalidType  = errors.New("invalid type supplied")
)

func validPair(pair string) bool {
	return strings.Contains(pair, pairDelimiter)
}

func validAsset(i string) bool {
	_, err := asset.New(i)
	return err == nil
}

func validtickerType(i string) bool {
	_, err := tickertype.New(i)
	return err == nil
}
