## Testing tools
[![CI Tests](https://github.com/xSolrac87/testing/actions/workflows/linter.yml/badge.svg)]
(https://github.com/xSolrac87/testing/actions/workflows/linter.yml)

### Execution Modes
#### parallel
* It enables to run test in parallel.
* Parallel test will be pause til sequential one are completed, then resumed.
* Parallel can be limited with `go test -parallel {threads} -v` by default it use `GOMAXPROCS`.

#### shuffle
* A best practice while writing tests is to make them isolated.
* We can use the `-shuffle` flag to randomize tests.
* It generates a seed value to run the test in the same order `go test -shuffle={seed} -v .`

### Table Driven
* Table-driven tests are an efficient technique for writing condensed tests and thus
 reducing boilerplate code.

### Sleeping
* Try to not use `time.Sleep()` to avoid Flaky tests.

### Race Flag
* Race detector it’s a tool to find data races that occur at runtime.
* It’s generally recommended to enable the race detector only during local testing or CI.
* We can use the `-race` flag to run test in data race mode.
* To run race for all test `go test -race -cover`.
* To run test with race for a specificity function `go test -race -run TestDataRace`.

### Benchmark
* API `b *testing.B`.
* `b.N` represents a variable number of iterations, if the benchmark completes in under 1 second, b.N is increased,
  and the benchmark runs again until b.N roughly matches benchtime ( default by 1 )
* Benchmark function must start with `Benchmark`.
* We can use the `-bench` flag to benchmark test.
* To run benchmark `go test -bench=.`.
* The `-{cores}` suffix denotes the number of CPUs used to run the benchmark, as specified by `GOMAXPROCS`.
* We can use `-count` flag to determine how many benchmarks will run `go test -bench=. -count 5`.
* We can use `-benchtime` flag to adjusting the minimum time `go test -bench=. -benchtime=10s`.
* Alternative way to control the amount of time a benchmark should run is by specifying the desired number of iterations
  `go test -bench=. -benchtime=100x`

#### timer
* In some cases, we need to perform operations before the benchmark loop, 
to avoid impact the benchmark results, we can use the `b.ResetTimer()` method before entering the loop.
* if we have to perform an expensive setup not just once but within each loop
we can surround the call with `b.StopTimer()` and `b.StartTimer()`.

#### compiler-optimizations
* https://github.com/golang/go/issues/14813
* `inlining` : an optimization that replaces a function call with the body of the called
  function and lets us prevent a function call. 
  Once the function is inlined, the compiler notices that the call has no side effects and replaces
  it with the following benchmark:
  ```
  func BenchmarkPopcnt1(b *testing.B) {
      for i := 0; i < b.N; i++ {
          // Empty
      }
  }
  ```

### Fuzzing
* Fuzzing or fuzz testing is a method of giving random unexpected input to your programs to test for possible crashes or edge cases.
* Go now supports fuzz testing natively as of Go 1.18.
* API `f *testing.F`.
* We can use the `-fuzz` flag to fuzzing test.
* Fuzzing function must tart with `Fuzz`.
* To run fuzzy for all test `go go test -fuzz .`.
* It generates a report in case of failure.
* It consumes a lot of CPU resources.
* More info you can check GopherCon2022 video about it -> https://www.youtube.com/watch?v=7KWPiRq3ZYI&t=2s

### Coverage
* Generates a report to check coverage.
* We can use the `-coverprofile` flag to create coverage file.
* To generate a coverage file `go test -coverprofile=coverage.out ./...`.
* We can then open using go tool cover `go tool cover -html=coverage.out`.
