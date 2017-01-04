// cpuTicket.go
package goid

import (
	"runtime"
)

func GoId() int64 {
	return runtime.GoId()
}

func GoPc() int {
	return runtime.GoPc()
}
