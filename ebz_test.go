package ebzkratos_test

import (
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/ebzkratos"
	"github.com/stretchr/testify/require"
)

func TestNewEbz(t *testing.T) {
	erk := errors.InternalServer("SERVER_ERROR", "database connection failed")
	ebz := ebzkratos.NewEbz(erk)
	require.NotNil(t, ebz)
	require.NotNil(t, ebz.Erk)

	t.Log(ebz.Erk.String())
}

func TestNew(t *testing.T) {
	erk := errors.BadRequest("BAD_REQUEST", "invalid parameter: user_id")
	ebz := ebzkratos.New(erk)
	require.NotNil(t, ebz)
	require.NotNil(t, ebz.Erk)

	t.Log(ebz.Erk.String())
}

func TestNewEbz_NotImplementErrorInterface(t *testing.T) {
	erk := errors.InternalServer("SERVER_ERROR", "database connection failed")
	ebz := ebzkratos.NewEbz(erk)
	require.NotNil(t, ebz)
	require.NotNil(t, ebz.Erk)

	{
		var erx interface{} = ebz
		err, ok := erx.(error) // 不要实现 error 接口，而且注意一定不要实现，否则会加重开发者的心智负担
		require.False(t, ok)
		require.Nil(t, err)
	}

	{
		var erx interface{} = ebz.Erk
		err, ok := erx.(error) // 已经实现 error 接口
		require.True(t, ok)
		require.Error(t, err)
	}

	t.Log(ebz.Erk.String())
}

func TestAs(t *testing.T) {
	{
		var erk = errors.BadRequest("BAD_REQUEST", "invalid input")
		var err error = erk
		// t.Log(erk != nil) // true
		// t.Log(err != nil) // true
		// 具体原因请看这里 https://go.dev/doc/faq#nil_error 因为类型和值都为nil的才是nil否则不是

		res, ok := ebzkratos.As(err)
		require.True(t, ok)
		t.Log(res)
		require.NotNil(t, res)
	}

	{
		var erk *errors.Error
		var err error = erk
		// t.Log(erk != nil) // false
		// t.Log(err != nil) // true
		// 具体原因请看这里 https://go.dev/doc/faq#nil_error 因为类型和值都为nil的才是nil否则不是

		res, ok := ebzkratos.As(err)
		require.True(t, ok)
		t.Log(res)
		require.Nil(t, res)
	}
}

func TestIs(t *testing.T) {
	ebz1 := ebzkratos.NewEbz(errors.BadRequest("BAD_REQUEST", "invalid input-1"))
	ebz2 := ebzkratos.NewEbz(errors.BadRequest("BAD_REQUEST", "invalid input-2"))
	require.True(t, ebzkratos.Is(ebz1, ebz2))

	require.True(t, errors.Is(ebz1.Erk, ebz1.Erk)) // 还是相等
}

func TestFrom(t *testing.T) {
	ebz1 := ebzkratos.NewEbz(errors.InternalServer("SERVER_ERROR", "database error"))
	var err error = ebz1.Erk
	ebz2 := ebzkratos.From(err)
	require.True(t, ebzkratos.Is(ebz1, ebz2))
}
