package ebzmust_test

import (
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/ebzkratos"
	"github.com/orzkratos/ebzkratos/must/ebzmust"
	"github.com/stretchr/testify/require"
)

// TestDone verifies Done function panics on error presence with structured logging
// Tests that nil Ebz passes without panic and non-nil Ebz triggers panic
//
// TestDone 验证 Done 函数在错误存在时 panic 并记录结构化日志
// 测试 nil Ebz 通过而非 nil Ebz 触发 panic
func TestDone(t *testing.T) {
	t.Run("no panic", func(t *testing.T) {
		var ebz *ebzkratos.Ebz
		ebzmust.Done(ebz)
	})

	t.Run("panic on error", func(t *testing.T) {
		require.Panics(t, func() {
			erk := errors.InternalServer("SERVER_ERROR", "database connection failed")
			ebzmust.Done(ebzkratos.New(erk))
		})
	})
}

// TestMust verifies Must function enforces error absence with panic
// Tests that nil Ebz passes without panic and non-nil Ebz triggers panic
//
// TestMust 验证 Must 函数强制错误缺失并使用 panic
// 测试 nil Ebz 通过而非 nil Ebz 触发 panic
func TestMust(t *testing.T) {
	t.Run("no panic", func(t *testing.T) {
		var ebz *ebzkratos.Ebz
		ebzmust.Must(ebz)
	})

	t.Run("panic on error", func(t *testing.T) {
		require.Panics(t, func() {
			erk := errors.BadRequest("BAD_REQUEST", "invalid transaction")
			ebzmust.Must(ebzkratos.New(erk))
		})
	})
}
