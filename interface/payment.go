package _interface

import (
	"fmt"
)

type PaymentProcessor interface {
	Pay(amount float64) string
}

type CreditCard struct {
	CardNumber string
}

func (cc CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using Credit Card ending in %s", amount, cc.CardNumber[len(cc.CardNumber)-4:])
}

type PayPal struct {
	Email string
}

func (pp PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using PayPal account %s", amount, pp.Email)
}

func ProcessPayment(p PaymentProcessor, amount float64) {
	receipt := p.Pay(amount)
	fmt.Println(receipt)
}

type MockPayment struct{}

func (m MockPayment) Pay(amount float64) string {
	return "Payment mocked for testing"
}

func PaymentExample() {
	cc := CreditCard{CardNumber: "1234567890123456"}
	pp := PayPal{Email: "user@example.com"}

	ProcessPayment(cc, 100.00)
	ProcessPayment(pp, 59.99)

	mock := MockPayment{}
	ProcessPayment(mock, 0.00)
}
