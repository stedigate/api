package api

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/pushgate/core/internal/data"
)

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) viewPaymentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.ResponseErrorNotFound(w, r, nil, nil)
		return
	}

	payment, err := app.models.Payments.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.ResponseErrorNotFound(w, r, nil, nil)
		default:
			app.ResponseErrorInternalServerError(w, r, map[string]any{"error": err.Error()}, nil)
		}
		return
	}

	app.ResponseSuccessOk(w, r, map[string]any{"payment": payment}, nil)
}

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) createPaymentHandler(w http.ResponseWriter, r *http.Request) {

	// --------------------------------------------------------------------------------
	var input CreatePaymentRequest
	jsonErr := app.readJSON(w, r, &input)
	if jsonErr != nil {
		app.ResponseErrorBadRequest(w, r, map[string]any{"json": jsonErr.Error()}, nil)
		return
	}

	// --------------------------------------------------------------------------------
	ok, validateErr := input.validate()
	if !ok {
		errs := map[string]interface{}{}
		for key, value := range validateErr {
			errs[key] = value
		}
		app.ResponseErrorFailedValidation(w, r, errs, nil)
		return
	}

	// --------------------------------------------------------------------------------
	payment := &data.Payment{
		Amount:         input.Amount,
		Currency:       input.Currency,
		WalletAddress:  input.PayoutAddress,
		WalletCurrency: input.PayoutCurrency,
		UserID:         1,
		ExtraID:        input.OrderId,
		Description:    input.Description,
		PayCurrency:    "usdttrc20",
		PayAmount:      input.Amount,
		ExpiredAt:      time.Now().Add(time.Hour * 24),
	}
	err := app.models.Payments.Insert(payment)
	if err != nil {
		app.ResponseErrorBadRequest(w, r, map[string]any{
			"error": err.Error(),
		}, nil)
		return

	}

	// --------------------------------------------------------------------------------
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/payments/%d", payment.ID))

	payload := map[string]any{
		"payment": payment,
	}
	app.ResponseSuccessCreated(w, r, payload, headers)
}

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) listPaymentHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}

	app.ResponseSuccessOk(w, r, data, nil)
}

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) deletePaymentHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}

	app.ResponseSuccessEmpty(w, r, data, nil)
}
