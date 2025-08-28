// Package tests: Testing utilities for panic and error validation
// Provides assertion mechanisms for fail-fast error handling verification
//
// tests: 用于 panic 和错误验证的测试工具
// 为快速失败错误处理验证提供断言机制
package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// ExpectPanic validates that provided function triggers panic
// Uses defer-recover to detect panic and fail test if no panic occurs
//
// ExpectPanic 验证提供的函数触发 panic
// 使用 defer-recover 检测 panic，如果没有 panic 则测试失败
func ExpectPanic(t *testing.T, run func()) {
	defer func() {
		if cause := recover(); cause != nil {
			t.Logf("expect panic then catch panic [%v] -> [SUCCESS]", cause)
			return
		}
		require.Fail(t, "expect panic while not panic -> [FAILURE]")
	}()

	run()
}
