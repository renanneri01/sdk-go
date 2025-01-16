package order

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

var (
	createResponseJSON, _ = os.Open("../../resources/mocks/order/create_full_order_response.json")
	createResponse, _     = io.ReadAll(createResponseJSON)
)

func TestCreate(t *testing.T) {
	type fields struct {
		cfg *config.Config
	}
	type args struct {
		ctx     context.Context
		request Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "should_fail_to_send_request",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							return nil, fmt.Errorf("error")
						},
					},
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should_return_response",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							stringReader := strings.NewReader(string(createResponse))
							stringReaderCloser := io.NopCloser(stringReader)
							return &http.Response{Body: stringReaderCloser}, nil
						},
					},
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want: &Response{
				ID:                "01HRYFWNYRE1MR1E60MW3X0T2P",
				Type:              "online",
				TotalAmount:       "1000.00",
				ExternalReference: "ext_ref_1234",
				CountryCode:       "ARG",
				Status:            "processed",
				StatusDetail:      "accredited",
				CaptureMode:       "automatic",
				ProcessingMode:    "automatic",
				Description:       "some description",
				Marketplace:       "NONE",
				MarketplaceFee:    "10.00",
				ExpirationTime:    "P3D",
				Transactions: TransactionResponse{
					Payments: []PaymentResponse{
						{
							ID:          "01HRYFXQ53Q3JPEC48MYWMR0TE",
							ReferenceID: "123456789",
							Status:      "processed",
							Amount:      "1000.00",
							PaymentMethod: PaymentMethodRequest{
								ID:                  "master",
								Type:                "credit_card",
								Token:               "677859ef5f18ea7e3a87c41d02c3fbe3",
								StatementDescriptor: "LOJA X",
								Installments:        1,
							},
						},
					},
					Refunds: []RefundResponse{},
				},
				Payer: PayerRequest{
					Email:     "{email}",
					FirstName: "John",
					LastName:  "Doe",
					Identification: &IdentificationRequest{
						Type:   "CPF",
						Number: "00000000000",
					},
					Phone: &PhoneRequest{
						AreaCode: "55",
						Number:   "99999999999",
					},
					Address: &AddressRequest{
						StreetName:   "Av. das Nações Unidas",
						StreetNumber: "99",
					},
				},
				Items: []ItemsRequest{
					{
						ID:          "item_id",
						Title:       "Some item title",
						UnitPrice:   "1000.00",
						Description: "Some item description",
						CategoryID:  "category_id",
						Type:        "item_type",
						PictureUrl:  "https://mysite.com/img/item.jpg",
						Quantity:    1,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.cfg,
			}
			got, err := c.Create(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}
