// Always design system using dependency inversion(the 'D' in SOLID)
package main

import (
	"errors"
	"fmt"
)

type PaymentProcessor interface {
	Charge(amount int) error
}

type StripeClient struct {
	APIKey string
}

func (s *StripeClient) Charge(amount int) error {
	fmt.Printf("[Stripe] Charging $%d using API Key: %s...\n", amount, s.APIKey[:4]+"-xxxx")
	return nil
}

type PayPalClient struct {
	AccountEmail string
}

func (p *PayPalClient) Charge(amount int) error {
	if amount > 500 {
		return errors.New("PayPal limits exceeded for single transaction")

	}
	fmt.Printf("[PayPal] Charging $%d from account: %s\n", amount, p.AccountEmail)
	return nil
}

// we have buisness logic that process an order but it needs to charge a card
// We do not want our order code tied to stripe or paypal directly
// so this code simulation is solution for this problem
type OrderService struct {
	payment PaymentProcessor
}

func NewOrderService(pp PaymentProcessor) *OrderService {
	return &OrderService{payment: pp}
}

func (os *OrderService) Checkout(orderID string, total int) {
	fmt.Printf("Processing checkout for order #%s...\n", orderID)

	err := os.payment.Charge(total)
	if err != nil {
		fmt.Printf("Checkout failed for order #%s: %v\n", orderID, err)
		return
	}
	fmt.Printf("Checkout success for order #%s\n\n", orderID)
}

func main() {
	stripe := &StripeClient{APIKey: "sk_live_98235235227"}
	paypal := &PayPalClient{AccountEmail: "billing@company.com"}

	OrderServiceWithStripe := NewOrderService(stripe)
	OrderServiceWithStripe.Checkout("101", 150)

	orderServiceWithPayPal := NewOrderService(paypal)
	orderServiceWithPayPal.Checkout("102", 600)
}
