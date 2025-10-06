// Package ebzmust: Assert functions for production Ebz error handling
// Provides fail-fast assertions with panic and logging
//
// ebzmust: 生产环境 Ebz 错误处理的断言函数
// 提供带 panic 和日志的快速失败断言
package ebzmust

import (
	"github.com/orzkratos/ebzkratos"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// Done enforces zero-error convention with immediate panic on breach
// Logs "NO ERROR BUG" and panics if ebz is not nil
//
// Done 强制执行零错误约定，违反时立即 panic
// 如果 ebz 非空则记录 "NO ERROR BUG" 并 panic
func Done(ebz *ebzkratos.Ebz) {
	if ebz != nil {
		zaplog.ZAPS.Skip1.LOG.Panic("NO ERROR BUG", zap.Error(ebz.Erk))
	}
}

// Must demands perfect execution with fail-fast termination
// Logs "ERROR" and panics if ebz is not nil
//
// Must 要求完美执行，错误时快速失败终止
// 如果 ebz 非空则记录 "ERROR" 并 panic
func Must(ebz *ebzkratos.Ebz) {
	if ebz != nil {
		zaplog.ZAPS.Skip1.LOG.Panic("ERROR", zap.Error(ebz.Erk))
	}
}
