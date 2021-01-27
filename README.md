# panictest

`go test ./... -v`

```
=== RUN   Test1
    Test1: main_test.go:6: Test1 START
=== RUN   Test1/xs
--- FAIL: Test1 (0.00s)
    --- FAIL: Test1/xs (0.00s)
panic: runtime error: index out of range [0] with length 0 [recovered]
        panic: runtime error: index out of range [0] with length 0

goroutine 7 [running]:
testing.tRunner.func1.1(0x1130a60, 0xc000018200)
        /Users/main/.anyenv/envs/goenv/versions/1.14.2/src/testing/testing.go:940 +0x2f5
testing.tRunner.func1(0xc000120240)
        /Users/main/.anyenv/envs/goenv/versions/1.14.2/src/testing/testing.go:943 +0x3f9
panic(0x1130a60, 0xc000018200)
        /Users/main/.anyenv/envs/goenv/versions/1.14.2/src/runtime/panic.go:969 +0x166
github.com/mkql/panictest.Test1.func1(0xc000120240)
        /Users/main/go/src/github.com/mkql/panictest/main_test.go:9 +0x95
testing.tRunner(0xc000120240, 0xc00000c0c0)
        /Users/main/.anyenv/envs/goenv/versions/1.14.2/src/testing/testing.go:991 +0xdc
created by testing.(*T).Run
        /Users/main/.anyenv/envs/goenv/versions/1.14.2/src/testing/testing.go:1042 +0x357
FAIL    github.com/mkql/panictest       0.054s
=== RUN   TestSubPackage
    TestSubPackage: sub_test.go:6: TestSubPackage START
    TestSubPackage: sub_test.go:7: TestSubPackage END
--- PASS: TestSubPackage (0.00s)
PASS
ok      github.com/mkql/panictest/sub   (cached)
FAIL
```