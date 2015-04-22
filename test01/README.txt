测试for,range循环，在结构体，指针结构体循环的差别

go test -bench=. go-test/test01
testing: warning: no tests to run
PASS
Benchmark_loop1	   50000	     26015 ns/op
Benchmark_loop2	 2000000	       941 ns/op
Benchmark_loop3	 1000000	      1076 ns/op
Benchmark_loop4	 1000000	      1163 ns/op
ok  	go-test/test01	6.714s

对结构体的range循环是最耗时的，因为要做数据复制