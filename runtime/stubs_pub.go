// stubs_pub.go
package runtime

func GoId() int64 {
	return getg().goid
}

func GoPc() int {
	return int(getg().gopc)
}
