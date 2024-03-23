package validator

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/truemail-rb/truemail-go"
	"reflect"
	"regexp"
	"strings"
)

var (
	// EmailRX is a regex for sanity checking the format of email addresses.
	// The regex pattern used is taken from  https://html.spec.whatwg.org/#valid-e-mail-address.
	EmailRX      = `([-!#-'*+/-9=?A-Z^-~]+(\.[-!#-'*+/-9=?A-Z^-~]+)*|"([]!#-[^-~ \t]|(\\[\t -~]))+")@([0-9A-Za-z]([0-9A-Za-z-]{0,61}[0-9A-Za-z])?(\.[0-9A-Za-z]([0-9A-Za-z-]{0,61}[0-9A-Za-z])?)*|\[((25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])(\.(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])){3}|IPv6:((((0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}):){6}|::((0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}):){5}|[0-9A-Fa-f]{0,4}::((0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}):){4}|(((0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}):)?(0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}))?::((0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}):){3}|(((0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}):){0,2}(0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}))?::((0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}):){2}|(((0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}):){0,3}(0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}))?::(0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}):|(((0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}):){0,4}(0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}))?::)((0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}):(0|[1-9A-Fa-f][0-9A-Fa-f]{0,3})|(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])(\.(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])){3})|(((0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}):){0,5}(0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}))?::(0|[1-9A-Fa-f][0-9A-Fa-f]{0,3})|(((0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}):){0,6}(0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}))?::)|(?!IPv6:)[0-9A-Za-z-]*[0-9A-Za-z]:[!-Z^-~]+)])`
	TronWalletRX = `^T[a-zA-Z0-9]{33}$`
)

// Validator struct type contains a map of validation errors.
type Validator struct {
	Errors   map[string]string
	Validate *validator.Validate
}

// New is a helper which creates a new Validator instance with an empty errors map.
func New() *Validator {
	v := validator.New(validator.WithRequiredStructEnabled())
	_ = v.RegisterValidation("trc_addr", func(fl validator.FieldLevel) bool {
		regex := regexp.MustCompile(TronWalletRX)
		return regex.MatchString(fl.Field().String())
	})
	_ = v.RegisterValidation("email_addr", func(fl validator.FieldLevel) bool {
		regex := regexp.MustCompile(EmailRX)
		return regex.MatchString(fl.Field().String())
	})

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	_ = v.RegisterValidation("email_dns", EmailDnsValidation)

	return &Validator{Errors: make(map[string]string), Validate: v}
}

func EmailDnsValidation(fl validator.FieldLevel) bool {
	field := fl.Field()

	configuration, err := truemail.NewConfiguration(
		truemail.ConfigurationAttr{
			WhitelistedDomains: []string{"sharebuy.com"},
			VerifierEmail:      "ramin.farmani@gmail.com",
		},
	)

	if err != nil {
		return false
	}

	return truemail.IsValid(field.String(), configuration)
}

func (v *Validator) Check(r interface{}) (bool, map[string]string) {
	err := v.Validate.Struct(r)
	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return false, map[string]string{"json": "invalid json"}
		}

		collectedErrors := v.grabValidationErrors(err)

		return false, collectedErrors
	}

	return true, nil
}

func (v *Validator) grabValidationErrors(err error) map[string]string {
	collectedErrors := map[string]string{}
	for _, err := range err.(validator.ValidationErrors) {
		switch err.Tag() {
		case "required":
			collectedErrors[err.Field()] = fmt.Sprintf("cannot be blank")
		case "email":
			collectedErrors[err.Field()] = fmt.Sprintf("must be a valid email address")
		case "eth_addr":
			collectedErrors[err.Field()] = fmt.Sprintf("must  be a valid Ethereum address")
		case "trc_addr":
			collectedErrors[err.Field()] = fmt.Sprintf("must  be a valid Tron address")
		case "len":
			collectedErrors[err.Field()] = fmt.Sprintf("must be exactly %v characters long", err.Param())
		case "lt":
			collectedErrors[err.Field()] = fmt.Sprintf("must be less than %v", err.Param())
		case "lte":
			collectedErrors[err.Field()] = fmt.Sprintf("must be less than or equal to %v", err.Param())
		case "gt":
			collectedErrors[err.Field()] = fmt.Sprintf("must be greater than %v", err.Param())
		case "gte":
			collectedErrors[err.Field()] = fmt.Sprintf("must be greater than or equal to %v", err.Param())
		default:
			collectedErrors[err.Field()] = fmt.Sprintf("'%v' must satisfy '%s' '%v' criteria", err.Value(), err.Tag(), err.Param())
		}
	}
	return collectedErrors
}
