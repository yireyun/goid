// stubs_pub.go
package runtime

func GoId() int64 {
	return getg().goid
}
