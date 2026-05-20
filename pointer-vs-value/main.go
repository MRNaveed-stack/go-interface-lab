package main

import "fmt"

// Pointer receiver
//Pointer  use much memory as compared to Value pass
// So use pointer only when you need to change the data throughout your application
// type BankAccount interface {
// 	Deposit(amount float64)
// 	GetBalance() float64
// }

// type Wallet struct {
// 	Balance float64
// }

// func (w *Wallet) Deposit(amount float64) {
// 	w.Balance += amount
// }
// func (w *Wallet) GetBalance() float64 {
// 	return w.Balance
// }

// func main() {
// 	// if we write this:
// 	// var account BankAccount = Wallet{Balance : 100.0}
// 	// This above line will give compiler error

// 	// So we must pass the pointer to satisfy the interface
// 	var account BankAccount = &Wallet{Balance: 100.0}
// 	account.Deposit(50.0)
// 	fmt.Printf("Pointer Example Balance: $%.2f\n", account.GetBalance())
// }

// In this example we will use Value Receiver

// 1. Define the Interface
type Greeter interface {
	Greet() string // This action only reads data to build a string
}

// 2. Define the Struct
type User struct {
	Name string
}

// 3. Implement using a Value Receiver (User)
// Go passes a copy of User to this method. The original cannot be altered.
func (u User) Greet() string {
	// Even if we tried to change u.Name here, it would only change on the copy.
	return "Hello, my name is " + u.Name
}

func main() {
	// NO TRAP HERE: Value receivers accept both values AND pointers!

	// Success Option A: Passing a plain Value
	var g1 Greeter = User{Name: "Alice"}
	fmt.Println("Value Type:", g1.Greet())

	// Success Option B: Passing a Pointer
	var g2 Greeter = &User{Name: "Bob"}
	fmt.Println("Pointer Type:", g2.Greet())
}
