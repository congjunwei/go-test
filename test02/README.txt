测试不同赋值方式的差别
go test -bench=. -benchmem go-test/test02
testing: warning: no tests to run
PASS
Benchmark_01	200000000	         8.54 ns/op	       0 B/op	       0 allocs/op
Benchmark_02	1000000000	         2.76 ns/op	       0 B/op	       0 allocs/op
Benchmark_03	200000000	         7.53 ns/op	       0 B/op	       0 allocs/op
ok  	go-test/test02	7.861s