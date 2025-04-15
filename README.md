# Go 泛型教程示例

这个项目展示了 Go 语言泛型的基本用法，包含了以下示例：

## 功能特性

1. 泛型栈（Stack）实现
   - 支持任意类型的压栈和出栈操作
   - 包含空栈检查

2. 泛型 Map 函数
   - 支持对切片元素进行类型转换
   - 支持自定义转换函数

3. 泛型 Filter 函数
   - 支持根据条件筛选切片元素
   - 支持自定义筛选条件

## 代码示例

### 泛型栈

```go
// 创建整数栈
intStack := &Stack[int]{}
intStack.Push(1)
val, ok := intStack.Pop()

// 创建字符串栈
strStack := &Stack[string]{}
strStack.Push("hello")
val, ok := strStack.Pop()
```

### 泛型 Map

```go
// 数字翻倍
numbers := []int{1, 2, 3}
doubled := Map(numbers, func(x int) int { return x * 2 })

// 转换为字符串
strings := Map(numbers, func(x int) string { return fmt.Sprintf("num-%d", x) })
```

### 泛型 Filter

```go
// 筛选偶数
numbers := []int{1, 2, 3, 4, 5}
even := Filter(numbers, func(x int) bool { return x%2 == 0 })
```

## 运行测试

```bash
go test -v
```

## 要求

- Go 1.18 或更高版本（支持泛型特性）

## 学习要点

1. 泛型类型声明：使用方括号 `[T any]` 声明类型参数
2. 泛型约束：使用 `any` 表示接受任意类型
3. 泛型函数：如何定义和使用带有类型参数的函数
4. 零值处理：如何在泛型代码中正确处理零值