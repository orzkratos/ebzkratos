// Package ebzkratos: Error wrapper that addresses Go's nil interface challenge
// Does not implement error interface to avoid (*T)(nil) != nil issue
// CORE CONSTRAINT: ebz != nil ⇒ ebz.Erk must be non-nil (enforced by must.Full)
//
// ebzkratos: 处理 Go nil 接口挑战的错误包装器
// 不实现 error 接口以避免 (*T)(nil) != nil 问题
// 核心约定：ebz != nil ⇒ ebz.Erk 必须非空（由 must.Full 强制执行）
package ebzkratos

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/yyle88/must"
)

// Ebz wraps Kratos errors without implementing error interface
// STRUCTURAL INVARIANT: non-nil Ebz must have non-nil Erk component
// All constructors enforce this constraint through must.Full validation
//
// Ebz 包装 Kratos 错误但不实现 error 接口
// 结构不变量：非空 Ebz 必须有非空 Erk 组件
// 所有构造函数通过 must.Full 验证来强制执行此约定
type Ebz struct {
	Erk *errors.Error // Non-nil when Ebz is non-nil (enforced by constructors) // 当 Ebz 非空时必须非空（由构造函数强制）
}

// NewEbz creates error wrapper instance
// Ensures Erk is always non-nil through must.Full validation
//
// NewEbz 创建错误包装器实例
// 通过 must.Full 验证确保 Erk 始终非空
func NewEbz(erk *errors.Error) *Ebz {
	return &Ebz{
		Erk: must.Full(erk), // Enforces non-nil constraint: ebz != nil ⇒ erk must be non-nil // 强制非空约定：ebz != nil ⇒ erk 必须非空
	}
}

// New provides concise constructor alias
//
// New 提供简洁的构造函数别名
func New(erk *errors.Error) *Ebz {
	return NewEbz(must.Full(erk)) // Enforces non-nil constraint: ebz != nil ⇒ erk must be non-nil // 强制非空约定：ebz != nil ⇒ erk 必须非空
}

// As performs type conversion from generic error to Ebz
// Returns converted result or nil with safe nil handling
//
// As 执行从通用错误到 Ebz 的类型转换
// 返回转换结果或 nil，安全处理 nil
func As(err error) (ebz *Ebz, ok bool) {
	var erk *errors.Error
	if ok = errors.As(err, &erk); ok {
		if erk == nil {
			return nil, true
		}
		return NewEbz(erk), true
	}
	return nil, false
}

// Is checks if two Ebz values are equivalent
// Safely manages nil cases during comparison
//
// Is 检查两个 Ebz 值是否等价
// 比较时安全处理 nil 情况
func Is(ebz1 *Ebz, ebz2 *Ebz) bool {
	if ebz1 == nil || ebz2 == nil {
		return ebz1 == nil && ebz2 == nil
	}
	_ = must.Full(ebz1.Erk) // Enforces non-nil constraint: ebz != nil ⇒ erk must be non-nil // 强制非空约定：ebz != nil ⇒ erk 必须非空
	_ = must.Full(ebz2.Erk) // Enforces non-nil constraint: ebz != nil ⇒ erk must be non-nil // 强制非空约定：ebz != nil ⇒ erk 必须非空
	return ebz1.Erk.Is(ebz2.Erk)
}

// FromError transforms generic error into safe Ebz instance
// Returns nil-safe result for Kratos errors
//
// FromError 将通用错误转换为安全的 Ebz 实例
// 为 Kratos 错误返回 nil 安全结果
func FromError(err error) *Ebz {
	erk := errors.FromError(err)
	if erk != nil {
		return NewEbz(erk)
	}
	return nil
}

// From is an alias for FromError function
//
// From 是 FromError 函数的别名
func From(err error) *Ebz {
	ebz := FromError(err)
	if ebz != nil {
		_ = must.Full(ebz.Erk) // Enforces non-nil constraint: ebz != nil ⇒ erk must be non-nil // 强制非空约定：ebz != nil ⇒ erk 必须非空
	}
	return ebz
}
