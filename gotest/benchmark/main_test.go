package benchmark

import (
	"testing"
)

// 名称以 Benchmark 为名称前缀的函数，只能接受 *testing.B 的参数，这种测试函数是性能测试函数
func Benchmark(t *testing.B) {
	TestB()
}
