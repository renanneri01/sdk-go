package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/credential"
	"github.com/mercadopago/sdk-go/pkg/header"
	"github.com/mercadopago/sdk-go/pkg/paymentmethod"
)

func main() {
	cdt, err := credential.New("TEST-640110472259637-071923-a761f639c4eb1f0835ff7611f3248628-793910800")
	if err != nil {
		fmt.Println(err)
		return
	}

	pmc := paymentmethod.NewClient(cdt)

	headers := http.Header{}
	headers.Add("X-Idempotency-Key", "123999")
	headers.Add("Some-Key", "some_value")

	ctx := context.Background()
	// This will return a child context decorated with your custom headers.
	// They will be forwarded wherever this context is used in any of the APIs.
	ctx = header.Context(ctx, headers)
	res, err := pmc.List(
		ctx,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range res {
		fmt.Println(v)
	}
}
