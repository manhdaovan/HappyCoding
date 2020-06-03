package somepackage

//ExportedFunc is a sample of exported function of a package.
func ExportedFunc(a int, b uint64) error {
	return nil
}

func unexportedFunc(a string, b float32) (float32, int) {
	f := float32(0.123)
	n := 2
	return f, n
}

func namedReturnFnc(a, b int) (c, d int, err error) {
	c = a + b
	d = a * b
	return
}

func exampleUsage() {
	ExportedFunc(1, 2)
	unexportedFunc("s", 2.123)
}
