package main

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

//func receiveDispatches(channel <-chan DispatchNotification) {
//	for details := range channel {
//		fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
//	}
//	fmt.Println("Channel has been closed")
//}

func enumerateProducts(channel chan<- *Product) {
	for _, p := range ProductList {
		select {
		case channel <- p:
			fmt.Println("Sent product:", p.Name)
		default:
			fmt.Println("Discarding product:", p.Name)
			time.Sleep(time.Second)
		}
	}
	close(channel)
	color.HiMagenta("EnumerateProducts functions complete")
}

//func enumerateProducts(channel1, channel2 chan<- *Product) {
//	for _, p := range ProductList {
//		select {
//		case channel1 <- p:
//			fmt.Println("Send via channel 1")
//		case channel2 <- p:
//			fmt.Println("Send via channel 2")
//		}
//	}
//	close(channel1)
//	close(channel2)
//}

func main() {
	color.Green("Main function started")
	color.Cyan("CalcStoreTotal function started")
	CalcStoreTotal(Products)
	color.Cyan("CalcStoreTotal function complete")

	dispatchChannel := make(chan DispatchNotification, 100)
	//var sendOnlyChannel chan<- DispatchNotification = dispatchChannel
	//var receiveOnlyChannel <-chan DispatchNotification = dispatchChannel
	color.HiYellow("DispatchOrders functions started")
	go DispatchOrders(chan<- DispatchNotification(dispatchChannel))
	//receiveDispatches((<-chan DispatchNotification)(dispatchChannel))

	productChannel := make(chan *Product, 5)
	color.HiMagenta("EnumerateProducts functions started")
	go enumerateProducts(productChannel)

	openChannel := 2

	for {
		select {
		case details, ok := <-dispatchChannel:
			if ok {
				fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
			} else {
				color.HiRed("Dispatch channel has been closed")
				dispatchChannel = nil
				openChannel--
			}
		case product, ok := <-productChannel:
			if ok {
				fmt.Println("Product:", product.Name)
			} else {
				color.HiRed("Product channel has been closed")
				productChannel = nil
				openChannel--
			}
		default:
			if openChannel == 0 {
				goto alldone
			}
			fmt.Println("-- No message ready to be received")
			time.Sleep(time.Millisecond * 500)
		}
	}
alldone:
	fmt.Println("All values received")

	//go DispatchOrders(dispatchChannel)
	//for {
	//	if details, open := <-dispatchChannel; open {
	//		fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
	//	} else {
	//		fmt.Println("Channel has been closed")
	//		break
	//	}
	//}

	// for enumerateProducts function with 2 channels
	//c1 := make(chan *Product, 2)
	//c2 := make(chan *Product, 2)
	//
	//go enumerateProducts(c1, c2)
	//time.Sleep(time.Second)
	//
	//for p := range c1 {
	//	fmt.Println("channel 1 received product:", p.Name)
	//}
	//for p := range c2 {
	//	fmt.Println("channel 2 received product:", p.Name)
	//}

	color.Green("Main function complete")
}
