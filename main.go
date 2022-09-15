package main

import "fmt"

var subTotal int
var discount int
var getDiscount int
var total int

func DumbWMerchBerkah() {
	discount = 25
	getDiscount = subTotal - (subTotal * discount / 100)

	if getDiscount >= 20000 {
		getDiscount = 20000
	}
}

func DumbMerchMurahBanget() {
	discount = 50
	getDiscount = subTotal - (subTotal * discount / 100)

	if getDiscount >= 45000 {
		getDiscount = 45000
	}
}

func main() {
	fmt.Println("Enter Sub Total Payment:")
	fmt.Scanf("%d", &subTotal)
	if subTotal >= 50000 && subTotal < 70000 {
		DumbWMerchBerkah()
		fmt.Println("Your Voucher Discount: DumbWMerchBerkah")
	} else if subTotal >= 70000 {
		DumbMerchMurahBanget()
		fmt.Println("Your Voucher Discount: DumbMerchMurahBanget")
	} else {
		getDiscount = 0
	}
	total = subTotal - getDiscount
	fmt.Println("Discount :", discount, "%")
	fmt.Println("Discount Price :", getDiscount)
	fmt.Println("Total Payment: Rp. ", total)
}
