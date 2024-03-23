package data

import (
	"errors"
	"github.com/pushgate/core/pkg/postgresql"
)

// ErrRecordNotFound Define a custom ErrRecordNotFound error. We'll return this from our Get() method when
// looking up a movie that doesn't exist in our database.
var (
	ErrRecordNotFound = errors.New("record not found")
)

// Models Create a struct which wraps the MovieModel. We'll add other models to this,
// like a UserModel and PermissionModel, as our build progresses.
type Models struct {
	DB       postgresql.DB
	Payments interface {
		Insert(payment *Payment) error
		Get(id int64) (*Payment, error)
		Update(payment *Payment) error
		Delete(id int64) error
	}
}

// NewModels For ease of use, we also add a New() method which returns a Models struct containing
// the initialized MovieModel.
func NewModels(db postgresql.DB) Models {
	return Models{
		DB:       db,
		Payments: PaymentModel{DB: db},
	}
}

// NewMockModels For ease of use, we also add a New() method which returns a Models struct containing
// the initialized MovieModel.
func NewMockModels() Models {
	return Models{
		Payments: MockPaymentModel{},
	}
}
