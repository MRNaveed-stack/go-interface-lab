// interface is not about binding similar things together, it's about minimum surface area needed to get job done
package main

import (
	"errors"
	"fmt"
)

// Bad approach
// Junior way (Hard Coded Dependency )

type EmailNotifier struct{}

func (e EmailNotifier) Send(msg string) {
	fmt.Println("Sending email: ", msg)
}

// Here logic is tied to email. if we want to send SMS later? we have to rewrite this
func WelcomeUser(e EmailNotifier) {
	e.Send("Welcome to the platform")
}

// Best Approach
// The Architect Way (Interface-driven)
type Notifier interface {
	Notify(msg string)
}

func WelcomeUSER(n Notifier) {
	n.Notify("Welcome! We do not care how this is sent")
}

// Let's implement a real life example
type PaymentProcessor interface {
	ProcessPayment(amount float64) error
}

type StripeService struct {
	APIKey string
}

func (s StripeService) ProcessPayment(amount float64) error {
	if s.APIKey == "" {
		return errors.New("stripe authentication failed")
	}
	fmt.Printf("[Stripe] Successfully changed credit card: $%.2f\n", amount)
	return nil
}

type PayPalService struct {
	MerchantID string
}

func (p PayPalService) ProcessPayment(amount float64) error {
	fmt.Printf("[PayPal] Successfully redirected user and charged: $%.2f\n")
	return nil
}

type OrderManager struct {
	Processor PaymentProcessor
}

func (om OrderManager) Checkout(orderID string, totalAmount float64) {
	fmt.Printf("Processing order #%s--\n", orderID)
	err := om.Processor.ProcessPayment(totalAmount)
	if err != nil {
		fmt.Printf("Checkout failed for order %s: %v\n", orderID, err)
		return
	}
	fmt.Printf("Checkout successfull! Order %s is now being shipped. \n\n", orderID)

}

func main() {
	stripeGateway := StripeService{APIKey: "sk_live_51Nx... Wait, this is a secret!"}
	orderSystemWithStripe := OrderManager{Processor: stripeGateway}
	orderSystemWithStripe.Checkout("98765", 149.99)

	// Scenario B: Customer switches to PayPal at checkout
	// Notice that OrderManager logic remains 100% untouched
	paypalGateway := PayPalService{MerchantID: "merchant_abc_123"}
	orderSystemWithPayPal := OrderManager{Processor: paypalGateway}
	orderSystemWithPayPal.Checkout("43210", 45.50)
}
