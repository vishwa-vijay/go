package main

import (
	"fmt"

	soap "github.com/hooklift/gowsdl/soap"
	myservice "github.com/vishwa-vijay/go/gowsdl/myservice"
)

func main() {
	client := soap.NewClient("http://www.dneonline.com/calculator.asmx")
	calculator := myservice.NewCalculatorSoap(client)
	mrequest := myservice.Multiply{IntA: 10, IntB: 20}

	mresposne, _ := calculator.Multiply(&mrequest)
	fmt.Println(mresposne.MultiplyResult)

}
