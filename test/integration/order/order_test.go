package integration

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/mercadopago/sdk-go/pkg/cardtoken"
	"github.com/mercadopago/sdk-go/pkg/order"
	"github.com/mercadopago/sdk-go/test"
)

var (
	cfg             = test.Config()
	orderClient     = order.NewClient(cfg)
	cardTokenClient = cardtoken.NewClient(cfg)
)

func TestOrder(t *testing.T) {
	t.Run("should_create_order", func(t *testing.T) {
		ctx := context.Background()
		token, err := test.GenerateCardToken(ctx, cardTokenClient)
		if err != nil {
			t.Error("fail to generate card token", err)
		}

		request := order.Request{
			Type:              "online",
			TotalAmount:       "1000.00",
			ExternalReference: "ext_ref_1234",
			Transactions: &order.TransactionRequest{
				Payments: []order.PaymentRequest{
					{
						Amount: "1000.00",
						PaymentMethod: order.PaymentMethodRequest{
							ID:           "master",
							Token:        token,
							Type:         "credit_card",
							Installments: 1,
						},
					},
				},
			},
			Payer: order.PayerRequest{
				Email: fmt.Sprintf("test_user_%s@testuser.com", uuid.New().String()[:7]),
			},
		}

		resource, err := orderClient.Create(ctx, request)
		if resource == nil || resource.ID == "" {
			t.Error("order can't be nil")
		}

		if err != nil {
			t.Errorf(err.Error())
		}
	})
}

func TestGetOrder(t *testing.T) {
	t.Run("should_create_and_get_order", func(t *testing.T) {
		ctx := context.Background()
		token, err := test.GenerateCardToken(ctx, cardTokenClient)
		if err != nil {
			t.Error("fail to generate card token", err)
		}

		request := order.Request{
			Type:              "online",
			TotalAmount:       "1000.00",
			ExternalReference: "ext_ref_12345",
			Transactions: &order.TransactionRequest{
				Payments: []order.PaymentRequest{
					{
						Amount: "1000.00",
						PaymentMethod: order.PaymentMethodRequest{
							ID:           "master",
							Token:        token,
							Type:         "credit_card",
							Installments: 1,
						},
					},
				},
			},
			Payer: order.PayerRequest{
				Email: fmt.Sprintf("test_user_%s@testuser.com", uuid.New().String()[:7]),
			},
		}

		resource, err := orderClient.Create(ctx, request)
		if resource == nil || resource.ID == "" {
			t.Error("order can't be nil")
		}

		if err != nil {
			t.Errorf(err.Error())
		}

		resource, err = orderClient.Get(ctx, resource.ID)
		if resource == nil || resource.ID == "" {
			t.Error("order can't be nil")
		}

		if err != nil {
			t.Errorf(err.Error())
		}

	})
}

func TestTransaction(t *testing.T) {
	t.Run("should_create_transaction", func(t *testing.T) {
		ctx := context.Background()
		token, err := test.GenerateCardToken(ctx, cardTokenClient)
		if err != nil {
			t.Error("fail to generate card token", err)
		}
		request := order.Request{
			Type:              "online",
			ProcessingMode:    "manual",
			TotalAmount:       "1000.00",
			ExternalReference: "ext_1234",
			Payer: order.PayerRequest{
				Email: fmt.Sprintf("test_user_%s@testuser.com", uuid.New().String()[:7]),
			},
		}

		resource, err := orderClient.Create(ctx, request)
		if resource == nil || resource.ID == "" {
			t.Error("order can't be nil")
		}

		if err != nil {
			t.Errorf(err.Error())
		}

		requestTransaction := order.TransactionRequest{
			Payments: []order.PaymentRequest{
				{
					Amount: "1000.00",
					PaymentMethod: order.PaymentMethodRequest{
						ID:           "master",
						Token:        token,
						Type:         "credit_card",
						Installments: 1,
					},
				},
			},
		}

		resp, err := orderClient.CreateTransaction(ctx, resource.ID, requestTransaction)
		if resp == nil || resp.Payments[0].ID == "" {
			t.Error("transaction can't be nil")
		}

		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
