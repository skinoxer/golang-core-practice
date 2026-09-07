package main

import (
	"math/rand"
	"strings"
)

type CriditCard struct {
	CardNumber string
	ExpireDate string
	Cvv2       string
	BankName   string
}

func main() {
	cards := []CriditCard{
		{CardNumber: "6037997365498552", ExpireDate: "05/04", Cvv2: "254", BankName: "melli"},
		{CardNumber: "6037997365432131", ExpireDate: "02/08", Cvv2: "414", BankName: "melat"},
		{CardNumber: "6037997365491236", ExpireDate: "08/13", Cvv2: "131", BankName: "pasargad"},
	}
	for _, card := range cards {
		if card.ExpireDate < "03/02" {
			println(strings.ToTitle("your card is Expire"))
			continue
		}
		remmainAmmount := getAccountRemain(card.CardNumber, card.ExpireDate)

		println("your card is :", card.CardNumber, "your ammout is :", remmainAmmount)
	}
}
func getAccountRemain(cardNumber string, expireDate string) int {
	min := 1_000_000
	max := 10_000_000

	return rand.Intn(max-min+1) + min
}

// for{
// infinit
// break
// }
// for j := 0; j < 10; j++ {
// 	if j%2 != 0 {
// 		continue
// 	}
// 	 fmt.Println(j)
// }

// list := []int{12, 45, 84, 21, 6, 9, 2, 824, 23}

// for key, value := range list {
// 	fmt.Println("key: ", key, "value: ", value)
