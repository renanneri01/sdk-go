package main

import (
	"fmt"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/httpclient"
	"github.com/mercadopago/sdk-go/pkg/mp"
	"github.com/mercadopago/sdk-go/pkg/paymentmethod"
)

type myRequester struct{}

func (*myRequester) Do(req *http.Request) (*http.Response, error) {
	// my own Do logic
	return nil, nil
}

func main() {
	mp.SetAccessToken("TEST-640110472259637-071923-a761f639c4eb1f0835ff7611f3248628-793910800")

	pmc := paymentmethod.NewClient()

	// can be a http.Client from standard library:
	// standardLibClient := &http.Client{}
	// or can be a custom requester
	myOwnRequester := &myRequester{}

	res, err := pmc.List(
		httpclient.WithCallRequester(myOwnRequester), // sdk will use that requester
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range res {
		fmt.Println(v)
	}
}
