package demo1

import (
	"fmt"
	"os"
	"testing"
)

/*
进入demo1目录
go test
go test -v
执行指定正则函数
go test -v -run=dd

只运行TestMyAdd/dddd子用例
go test -v -run="TestMyAdd/dddd"

测试覆盖率
go test -cover

go test -cover -coverprofile=c.out
go tool cover -html=c.out
*/

/*
基准测试
go test -bench=*
go test -bench=MyAdd
go test -run=TestMyAddAllSuccess -bench=MyAdd
go test -run=^$ -bench=*
go test -bench=MyAdd -benchmem
go test -bench=. -benchtime=20s
go test -run=^$ -bench=^BenchmarkMyAddParallel$ -benchtime=5s -benc

go test -run=^$ -bench=. -test.cpuprofile cpu.out -test.memprofile mem.out -benchmem -benchtime=20s
go tool pprof -http :9099 cpu.out

*/
func TestMyAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1+1-a", args: struct {
			a int
			b int
		}{a: 1, b: 1}, want: 2},
		{name: "1+1-ab", args: struct {
			a int
			b int
		}{a: 1, b: 1}, want: 2},
		{name: "1+2-c", args: struct {
			a int
			b int
		}{a: 1, b: 2}, want: 3},
		{name: "dddd", args: struct {
			a int
			b int
		}{a: 1, b: 2}, want: 4},
		{name: "1+2-e", args: struct {
			a int
			b int
		}{a: 1, b: 2}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MyAdd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MyAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyAddAllSuccess(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1+1-a", args: struct {
			a int
			b int
		}{a: 1, b: 1}, want: 2},
		{name: "1+1-b", args: struct {
			a int
			b int
		}{a: 1, b: 1}, want: 2},
		{name: "1+2-c", args: struct {
			a int
			b int
		}{a: 1, b: 2}, want: 3},
		{name: "1+2-d", args: struct {
			a int
			b int
		}{a: 1, b: 2}, want: 3},
		{name: "1+2-e", args: struct {
			a int
			b int
		}{a: 1, b: 2}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MyAdd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MyAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMyAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {

		//time.Sleep(time.Millisecond*1)
		//b.ResetTimer之前的处理不会放到执行时间里，也不会输出到报告中，所以可以在之前做一些不计划作为测试报告的操作。
		//b.ResetTimer()

		MyAdd(1, 2)
	}
}

func BenchmarkMyAddTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MyAddTwo(1, 2)
	}
}

//go test -run=^$ -bench=^BenchmarkMyAddParallel$ -benchtime=5s -benchmem
func BenchmarkMyAddParallel(b *testing.B) {
	b.SetParallelism(4) // 设置使用的CPU数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			MyAddTwo(1, 2)
		}
	})
}

func TestMain(m *testing.M) {
	fmt.Println("write setup code here...") // 测试之前的做一些设置
	// 如果 TestMain 使用了 flags，这里应该加上flag.Parse()
	retCode := m.Run()                          // 执行测试
	fmt.Println("write teardown code here...&") // 测试之后做一些拆卸工作
	os.Exit(retCode)                            // 退出测试
}

//为你的代码编写示例代码有如下三个用处：
//示例函数能够作为文档直接使用，例如基于web的godoc中能把示例函数与对应的函数或包相关联。
//示例函数只要包含了// Output:也是可以通过go test运行的可执行测试。
//go test -run Example
func ExampleSplit() {
	fmt.Println(MyAdd(1, 2))
	fmt.Println(MyAdd(2, 3))
	// Output:
	// 3
	// 5

	//9 //这个9不能加 加上就报错了
}
