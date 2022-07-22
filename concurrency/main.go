package main

import (
	"fmt"
)

func receiveDispatches(channel <-chan DispatchNotification) {
	for details := range channel {
		fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
	}
	fmt.Println("Channel has been closed")
}

func main() {
	fmt.Println("main function started")
	CalcStoreTotal(Products)
	fmt.Println("CalcStoreTotal function complete")
	//time.Sleep(time.Second * 3)

	dispatchChannel := make(chan DispatchNotification, 100)
	//var sendOnlyChannel chan<- DispatchNotification = dispatchChannel
	//var receiveOnlyChannel <-chan DispatchNotification = dispatchChannel
	go DispatchOrders(chan<- DispatchNotification(dispatchChannel))
	receiveDispatches((<-chan DispatchNotification)(dispatchChannel))

	//go DispatchOrders(dispatchChannel)
	//for {
	//	if details, open := <-dispatchChannel; open {
	//		fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
	//	} else {
	//		fmt.Println("Channel has been closed")
	//		break
	//	}
	//}

	fmt.Println("main function complete")
}
