// cpuTicket_test.go
package goid

import (
	"runtime"
	"testing"
	"time"
)

func TestGoid(t *testing.T) {
	t.Logf("GoId: %v", runtime.GoId())
	t.Logf("GoId: %v", runtime.GoId())
}

func TestBenchGoid(t *testing.T) {
	N := 10000 * 1000
	start := time.Now()
	for i := 0; i < N; i++ {
		_ = runtime.GoId()
	}
	end := time.Now()
	use := end.Sub(start)
	op := use / time.Duration(N)
	t.Logf("Times: %10v, Use: %14v %10v/op\n", N, use, op)
}

//func TestPcid(t *testing.T) {
//	t.Logf("GoId: %v", runtime.ProcId())
//	t.Logf("GoId: %v", runtime.ProcId())
//}

//func TestBenchPcid(t *testing.T) {
//	N := 10000 * 1000
//	start := time.Now()
//	for i := 0; i < N; i++ {
//		_ = runtime.ProcId()
//	}
//	end := time.Now()
//	use := end.Sub(start)
//	op := use / time.Duration(N)
//	t.Logf("Times: %10v, Use: %14v %10v/op\n", N, use, op)
//}
