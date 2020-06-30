package util

import "testing"

//t.Log t.Logf     # 正常信息
//t.Error t.Errorf # 测试失败信息
//t.Fatal t.Fatalf # 致命错误，测试程序退出的信息
//t.Fail     # 当前测试标记为失败
//t.Failed   # 查看失败标记
//t.FailNow  # 标记失败，并终止当前测试函数的执行，需要注意的是，我们只能在运行测试函数的 Goroutine 中调用 t.FailNow 方法，而不能在我们在测试代码创建出的 Goroutine 中调用它
//t.Skip     # 调用 t.Skip 方法相当于先后对 t.Log 和 t.SkipNow 方法进行调用，而调用 t.Skipf 方法则相当于先后对 t.Logf 和 t.SkipNow 方法进行调用。方法 t.Skipped 的结果值会告知我们当前的测试是否已被忽略
//t.Parallel # 标记为可并行运算

func TestGenShortId(t *testing.T) {
	shortId, err := GenShortId()
	if shortId == "" || err != nil {
		t.Error("GenShortId failed!")
	}

	t.Log("GenShortId test pass")
}

// go test -v   功能测试
// go test -test.bench=".*"    进行性能测试
// go test -bench=".*" -cpuprofile=cpu.profile ./util   进行性能测试并生成性能测试文件
// go tool pprof util.test cpu.profile 查看性能
// go test -coverprofile=cover.out 测试覆盖率

func BenchmarkGenShortId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenShortId()
	}
}

func BenchmarkGenShortIdTimeConsuming(b *testing.B) {
	b.StopTimer() // 调用该函数停止压力测试的时间计数

	shortId, err := GenShortId()
	if shortId == "" || err != nil {
		b.Error(err)
	}

	b.StartTimer() // 重新开始时间

	for i := 0; i < b.N; i++ {
		GenShortId()
	}
}
