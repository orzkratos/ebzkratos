package ebzkratos_test

import (
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/ebzkratos"
	"github.com/stretchr/testify/require"
)

// TestNewEbz verifies basic Ebz creation from Kratos error
// Tests that NewEbz correctly wraps error and maintains error information
//
// TestNewEbz 验证从 Kratos 错误创建基本的 Ebz
// 测试 NewEbz 正确包装错误并维护错误信息
func TestNewEbz(t *testing.T) {
	erk := errors.InternalServer("SERVER_ERROR", "database connection failed")
	ebz := ebzkratos.NewEbz(erk)
	require.NotNil(t, ebz)
	require.NotNil(t, ebz.Erk)

	t.Log(ebz.Erk.String())
}

// TestNew verifies Ebz creation using the short form
// Tests that New is an alias for NewEbz with identical behavior
//
// TestNew 验证使用短格式创建 Ebz
// 测试 New 是 NewEbz 的别名，具有相同行为
func TestNew(t *testing.T) {
	erk := errors.BadRequest("BAD_REQUEST", "invalid parameter: user_id")
	ebz := ebzkratos.New(erk)
	require.NotNil(t, ebz)
	require.NotNil(t, ebz.Erk)

	t.Log(ebz.Erk.String())
}

// TestNewEbz_NotImplementErrorInterface verifies that Ebz does not implement error interface
// Tests the key design decision that Ebz is not an error to avoid mental burden
// Validates that wrapped Erk still implements error interface
//
// TestNewEbz_NotImplementErrorInterface 验证 Ebz 不实现 error 接口
// 测试关键设计决策：Ebz 不是 error，以避免心智负担
// 验证包装的 Erk 仍然实现 error 接口
func TestNewEbz_NotImplementErrorInterface(t *testing.T) {
	t.Run("ebz not error", func(t *testing.T) {
		erk := errors.InternalServer("SERVER_ERROR", "database connection failed")
		ebz := ebzkratos.NewEbz(erk)
		require.NotNil(t, ebz)
		require.NotNil(t, ebz.Erk)

		var err interface{} = ebz
		res, ok := err.(error) // 不要实现 error 接口，而且注意一定不要实现，否则会加重开发者的心智负担
		require.False(t, ok)
		require.Nil(t, res)

		t.Log(ebz.Erk.String())
	})

	t.Run("erk is error", func(t *testing.T) {
		erk := errors.InternalServer("SERVER_ERROR", "database connection failed")
		ebz := ebzkratos.NewEbz(erk)
		require.NotNil(t, ebz)
		require.NotNil(t, ebz.Erk)

		var err interface{} = ebz.Erk
		res, ok := err.(error) // 已经实现 error 接口
		require.True(t, ok)
		require.Error(t, res)

		t.Log(ebz.Erk.String())
	})
}

// TestAs verifies type-safe conversion from generic error to Kratos error type
// Tests both non-nil error conversion and nil pointer interface handling
// Validates that conversion preserves error information and handles Go's nil interface trap
//
// TestAs 验证从通用 error 到 Kratos 错误类型的类型安全转换
// 测试非 nil 错误转换和 nil 指针接口处理
// 验证转换保留错误信息并处理 Go 的 nil 接口陷阱
func TestAs(t *testing.T) {
	t.Run("non nil", func(t *testing.T) {
		erk := errors.BadRequest("BAD_REQUEST", "invalid input")
		err := error(erk)
		// t.Log(erk != nil) // true
		// t.Log(err != nil) // true
		// 具体原因请看这里 https://go.dev/doc/faq#nil_error 因为类型和值都为nil的才是nil否则不是

		res, ok := ebzkratos.As(err)
		require.True(t, ok)
		t.Log(res)
		require.NotNil(t, res)
	})

	t.Run("nil value", func(t *testing.T) {
		erk := (*errors.Error)(nil)
		err := error(erk)
		// t.Log(erk != nil) // false
		// t.Log(err != nil) // true
		// 具体原因请看这里 https://go.dev/doc/faq#nil_error 因为类型和值都为nil的才是nil否则不是

		res, ok := ebzkratos.As(err)
		require.True(t, ok)
		t.Log(res)
		require.Nil(t, res)
	})
}

// TestIs validates error comparison logic with reason and code matching
// Verifies that errors with same type are considered equal regardless of message
// Tests compatibility with Kratos errors.Is function
//
// TestIs 验证基于 reason 和 code 匹配的错误比较逻辑
// 验证具有相同类型的错误无论消息如何都被视为相等
// 测试与 Kratos errors.Is 函数的兼容性
func TestIs(t *testing.T) {
	t.Run("compare ebz", func(t *testing.T) {
		ebz1 := ebzkratos.NewEbz(errors.BadRequest("BAD_REQUEST", "invalid input-1"))
		ebz2 := ebzkratos.NewEbz(errors.BadRequest("BAD_REQUEST", "invalid input-2"))
		require.True(t, ebzkratos.Is(ebz1, ebz2))
	})

	t.Run("compare erk", func(t *testing.T) {
		ebz := ebzkratos.NewEbz(errors.BadRequest("BAD_REQUEST", "invalid input"))
		require.True(t, errors.Is(ebz.Erk, ebz.Erk)) // 还是相等
	})
}

// TestFrom validates conversion from generic error to Ebz format
// Tests that converted errors maintain compatibility with original errors
// Verifies error equivalence after conversion through Is comparison
//
// TestFrom 验证从通用 error 到 Ebz 格式的转换
// 测试转换后的错误与原始错误保持兼容性
// 通过 Is 比较验证转换后的错误等价性
func TestFrom(t *testing.T) {
	ebz1 := ebzkratos.NewEbz(errors.InternalServer("SERVER_ERROR", "database error"))
	var err error = ebz1.Erk
	ebz2 := ebzkratos.From(err)
	require.True(t, ebzkratos.Is(ebz1, ebz2))
}
