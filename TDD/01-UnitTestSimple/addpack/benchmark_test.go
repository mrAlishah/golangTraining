/*
goos, goarch, pkg, and cpu describe the operating system, architecture, package, and CPU specifications, respectively.
BenchmarkPrimeNumbers-4 denotes the name of the benchmark function that was run.
The -4 suffix denotes the number of CPUs used to run the benchmark,
as specified by GOMAXPROCS.

On the right side of the function name, you have two values, 14588 and 82798 ns/op.
The former indicates the total number of times the loop was executed,
while the latter is the average amount of time each iteration took to complete, expressed in nanoseconds per operation.
*/

package addpack

import (
	"testing"
)

var num = 1000

func BenchmarkPrimeNumbers(b *testing.B) {
	//b.N specifies the number of iterations; the value is not fixed, but dynamically allocated
    for i := 0; i < b.N; i++ {
        primeNumbers(num)
    }
}
