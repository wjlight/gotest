package benchmarks

import (
	"testing"
)

/*
http://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
基准测试的代码放在 _test打头的文件中，规则类似Test
如下是斐波那契的例子，
*/
func BenchmarkFib10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//note:这里的不能使用b.N作为函数Fib的参数传入
		Fib(10)
	}
}

/*
规则：
1、签名必须是Benchmarkxxx(*testing.B)
2、通过testing包，基准测试会运行多次，每次b.N的值会自增，直到benchmark认为结果已经稳定了。
3、基准测试的代码必须执行b.N次。
*/

/*
使用如下如下命令，运行基准：
$go test -bench=.

-bench后面必须接正则，

运行后出现如下结果：
PASS
BenchmarkFib10   1000000   1790 ns/op
ok   _/E_/goWork/GoTest/benchmarks  3.639s
testing: warning: no tests to run


第一行：PASS， 这是来自test驱动的 testing 部分， 会请求go test 来运行基准测试代码，
使test可用。如果想跳过test，可以通过传入正则 到 -run 标记。
第二行：运行了1000000次；在b.N最终的值的情况下，平均所用的时间，即如果b.N最后的值为5，则运行5次所用的平均时间
*/

/*
其他的一些：
1、每次最小的运行时间默认为1秒，如果运行完一次，没有超过1秒，
那么b.N的值就会增加 1,2,5,10,20,50然后再次运行
2、使用-benchtime 来修改最小的运行时间
*/
