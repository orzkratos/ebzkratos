package ebzmust_test

import (
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/ebzkratos"
	"github.com/orzkratos/ebzkratos/internal/tests"
	"github.com/orzkratos/ebzkratos/must/ebzmust"
)

func TestDone(t *testing.T) {
	{
		var ebz *ebzkratos.Ebz
		ebzmust.Done(ebz)
	}

	tests.ExpectPanic(t, func() {
		erk := errors.InternalServer("SERVER_ERROR", "database connection failed")
		ebzmust.Done(ebzkratos.New(erk))
	})
}

func TestMust(t *testing.T) {
	{
		var ebz *ebzkratos.Ebz
		ebzmust.Must(ebz)
	}

	tests.ExpectPanic(t, func() {
		erk := errors.BadRequest("BAD_REQUEST", "invalid transaction")
		ebzmust.Must(ebzkratos.New(erk))
	})
}
