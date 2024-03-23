package api

import (
	"github.com/stedigate/core/pkg/validator"
)

type CreatePaymentRequest struct {
	Amount         float64 `json:"amount" validate:"required,gt=0,lte=1000"`
	Currency       string  `json:"currency" validate:"required,eq=usd"`
	WebhookUrl     string  `json:"webhook_url" validate:"http_url"`
	OrderId        string  `json:"order_id" validate:"alphanum,max=52"`
	Description    string  `json:"description" validate:"printascii,max=255"`
	PayoutAddress  string  `json:"payout_address" validate:"trc_addr"`
	PayoutCurrency string  `json:"payout_currency" validate:"oneof=usdttrc20 usdterc20,required_with=PayoutAddress"`
}

func (r *CreatePaymentRequest) validate() (bool, map[string]string) {
	validate := validator.New()

	return validate.Check(r)
}
