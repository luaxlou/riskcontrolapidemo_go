package riskcontrolapidemo_go

import (
	"log"
	"testing"
)

func TestGetBalance(t *testing.T) {

	b := GetBalance(MERCHANT_KEY)

	log.Println(b.Available)
	log.Println(b.Overdraft)
}

func TestGetBills(t *testing.T) {

	bs := GetBills(MERCHANT_KEY, 0, 20, "")
	log.Println("total bills", bs.Count)

	for _, b := range bs.Data {
		log.Println(b)
	}
}
