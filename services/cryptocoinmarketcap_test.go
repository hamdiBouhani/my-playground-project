package services

import "testing"

func TestGiveCurrentPriceOfBitcoin(t *testing.T) {
	err := give_current_price_of_bitcoin()
	if err != nil {
		t.Error("Expected nil, got ", err)
		t.Fail()
	}
}
