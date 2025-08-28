# ebzkratos

Kratos 框架的 Go 错误包装器，解决 nil 接口问题。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 核心特性

🎯 **安全错误包装**: 通过不实现 error 接口来避免 Go 的 (*T)(nil) != nil 问题  
🔄 **类型转换**: 在通用错误和 Ebz 实例之间进行安全转换，正确处理 nil  
📋 **完整 API**: NewEbz、New、As、Is、FromError、From 函数提供完整错误处理

## 安装

```bash
go install github.com/orzkratos/ebzkratos@latest
```

## 使用方法

### 基本错误包装

```go
package main

import (
    "github.com/go-kratos/kratos/v2/errors"
    "github.com/orzkratos/ebzkratos"
)

func main() {
    // 创建 Kratos 错误
    erk := errors.BadRequest("INVALID_PARAM", "invalid parameter")
    
    // 用 Ebz 包装 - 避免 nil 接口问题
    ebz := ebzkratos.New(erk)
    
    // 安全的 nil 检查 - 没有 (*T)(nil) != nil 问题
    if ebz != nil {
        // 安全处理错误
        println("Error:", ebz.Erk.Message)
    }
}
```

### 类型转换

```go
func processError(err error) {
    // 将通用错误转换为 Ebz
    if ebz, ok := ebzkratos.As(err); ok {
        // 通过 Ebz 包装器处理 Kratos 错误
        println("Kratos error:", ebz.Erk.Reason)
    }
    
    // 替代转换方式
    ebz := ebzkratos.From(err)
    if ebz != nil {
        println("Converted:", ebz.Erk.Message)
    }
}
```

### 错误比较

```go
func compareErrors(ebz1, ebz2 *ebzkratos.Ebz) {
    // 带 nil 处理的安全相等性检查
    if ebzkratos.Is(ebz1, ebz2) {
        println("Errors are equivalent")
    }
}
```

### 生产环境断言

```go
import "github.com/orzkratos/ebzkratos/must/ebzmust"

func criticalOperation() {
    result, ebz := someOperation()
    
    // 断言无错误 - 如果 ebz 不为 nil 则 panic
    ebzmust.Done(ebz)
    
    // 安全继续处理结果
    processResult(result)
}
```

## 核心约定

**结构不变量**: `ebz != nil ⇒ ebz.Erk 必须非空`

此约定通过所有构造函数中的 `must.Full` 验证来强制执行，确保:

- 没有模糊的中间状态
- 简化错误处理逻辑  
- 快速失败错误检测
- 避免 nil 接口复杂性

## API 参考

### 构造函数

- `NewEbz(erk *errors.Error) *Ebz` - 创建带验证的错误包装器实例
- `New(erk *errors.Error) *Ebz` - 简洁的构造函数别名

### 转换函数  

- `As(err error) (*Ebz, bool)` - 从通用错误到 Ebz 的类型转换
- `FromError(err error) *Ebz` - 将通用错误转换为安全的 Ebz 实例
- `From(err error) *Ebz` - FromError 函数的别名

### 比较函数

- `Is(ebz1, ebz2 *Ebz) bool` - 检查两个 Ebz 值是否等价

### 断言函数

- `ebzmust.Done(ebz *Ebz)` - 断言零错误约定，违反时 panic
- `ebzmust.Must(ebz *Ebz)` - 要求完美执行，快速失败终止

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-08-28 08:33:43.829511 +0000 UTC -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **意见反馈？** 欢迎所有建议和宝贵意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Pull Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Pull Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**使用这个包快乐编程！** 🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub 标星点赞

[![Stargazers](https://starchart.cc/orzkratos/ebzkratos.svg?variant=adaptive)](https://starchart.cc/orzkratos/ebzkratos)