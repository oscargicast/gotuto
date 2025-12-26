/*
==================== ABSTRACT FACTORY (CORRECT) ====================

                   +-----------------------------+
                   |     AbstractFactory         |
                   |-----------------------------|
                   | +CreateProductA() A         |
                   | +CreateProductB() B         |
                   +---------------+-------------+
                                   |
                 ------------------------------------------------
                 |                                              |
+----------------------------------+      +----------------------------------+
|       ConcreteFactoryA            |      |       ConcreteFactoryB            |
|----------------------------------|      |----------------------------------|
| +CreateProductA() A               |      | +CreateProductA() A               |
| +CreateProductB() B               |      | +CreateProductB() B               |
+-------------------+---------------+      +-------------------+---------------+
                    |                                              |
        ----------------------------                  ----------------------------
        |                          |                  |                          |
+-------------------------+  +-------------------------+  +-------------------------+  +-------------------------+
| ConcreteProductA1       |  | ConcreteProductB1       |  | ConcreteProductA2       |  | ConcreteProductB2       |
|-------------------------|  |-------------------------|  |-------------------------|  |-------------------------|
| implements ProductA     |  | implements ProductB     |  | implements ProductA     |  | implements ProductB     |
+-----------+-------------+  +-----------+-------------+  +-----------+-------------+  +-----------+-------------+
            ^                            ^                            ^                            ^
            |                            |                            |                            |
+-------------------------+  +-------------------------+  +-------------------------+  +-------------------------+
| ProductA (interface)    |  | ProductB (interface)    |  | ProductA (interface)    |  | ProductB (interface)    |
+-------------------------+  +-------------------------+  +-------------------------+  +-------------------------+

Client code depends ONLY on:
- AbstractFactory
- ProductA interface
- ProductB interface

===================================================================
*/

package main

import "fmt"

// ======== PRODUCTS (interfaces) ========

type PaymentsClient interface {
	CreatePayment(amount int)
	Capture(paymentID string)
}

type RefundsClient interface {
	Refund(paymentID string)
}

// ======== ABSTRACT FACTORY ========

type PaymentProviderFactory interface {
	Payments() PaymentsClient
	Refunds() RefundsClient
}

// ======== STRIPE IMPLEMENTATION ========

type StripeFactory struct{}

func (StripeFactory) Payments() PaymentsClient {
	return StripePayments{}
}

func (StripeFactory) Refunds() RefundsClient {
	return StripeRefunds{}
}

type StripePayments struct{}

func (StripePayments) CreatePayment(amount int) {
	fmt.Println("[Stripe] create payment:", amount)
}

func (StripePayments) Capture(paymentID string) {
	fmt.Println("[Stripe] capture payment:", paymentID)
}

type StripeRefunds struct{}

func (StripeRefunds) Refund(paymentID string) {
	fmt.Println("[Stripe] refund payment:", paymentID)
}

// ======== MERCADOPAGO IMPLEMENTATION ========

type MercadoPagoFactory struct{}

func (MercadoPagoFactory) Payments() PaymentsClient {
	return MercadoPagoPayments{}
}

func (MercadoPagoFactory) Refunds() RefundsClient {
	return MercadoPagoRefunds{}
}

type MercadoPagoPayments struct{}

func (MercadoPagoPayments) CreatePayment(amount int) {
	fmt.Println("[MercadoPago] create payment:", amount)
}

func (MercadoPagoPayments) Capture(paymentID string) {
	fmt.Println("[MercadoPago] capture payment:", paymentID)
}

type MercadoPagoRefunds struct{}

func (MercadoPagoRefunds) Refund(paymentID string) {
	fmt.Println("[MercadoPago] refund payment:", paymentID)
}

// ======== FACTORY SELECTOR ========

type PaymentProvider int

const (
	ProviderStripe PaymentProvider = iota
	ProviderMercadoPago
)

func NewPaymentProviderFactory(provider PaymentProvider) PaymentProviderFactory {
	switch provider {
	case ProviderStripe:
		return StripeFactory{}
	case ProviderMercadoPago:
		return MercadoPagoFactory{}
	default:
		panic("unsupported provider")
	}
}

func main() {
	factory := NewPaymentProviderFactory(ProviderMercadoPago)

	payments := factory.Payments()
	refunds := factory.Refunds()

	payments.CreatePayment(1000)
	payments.Capture("pay_123")

	refunds.Refund("pay_123")
}
