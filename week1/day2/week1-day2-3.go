package main

import (
	"fmt"
	"sync"
)
func withdraw(wg *sync.WaitGroup, m *sync.Mutex,bb *int,amt int){
	if *bb < amt{
		fmt.Printf("Amount asked is more than bank balance... Failure, didn't withdraw\n")
		wg.Done()
		return

	}
	m.Lock()
	*bb = *bb - amt
	fmt.Printf("Withdrew %d and bank balance is %d\n",amt,*bb)
	m.Unlock()

	wg.Done()
}
func deposit(wg *sync.WaitGroup, m *sync.Mutex,bb *int,amt int){
	m.Lock()
	*bb = *bb + amt
	fmt.Printf("Deposited %d and bank balance is %d\n",amt,*bb)
	m.Unlock()

	wg.Done()
}
func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	var BankBalance int  = 500
	w.Add(1)
	go withdraw(&w, &m,&BankBalance,100)

	w.Add(1)
	go deposit(&w, &m,&BankBalance,100)

	w.Add(1)
	go withdraw(&w, &m,&BankBalance,100)

	w.Add(1)
	go deposit(&w, &m,&BankBalance,150)

	w.Add(1)
	go withdraw(&w, &m,&BankBalance,1000)

	w.Add(1)
	go deposit(&w, &m,&BankBalance,100)

	w.Wait()

}
