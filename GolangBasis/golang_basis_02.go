package main

func main() {
	// loop
	// for
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	// for without init state
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	var i := 0
	for ; sum < 1000; i ++{
		sum += sum
	}
	
	for sum < 1000 {
		sum += sum
	}
	// infinite loop
	for {
		break
	}

	// if
	func sqrt(x float64) string {
		if x < 0 {
			return sqrt(-x) + "i"
		}
		return fmt.Sprint(math.Sqrt(x))
	}
	
	func pow(x, n, lim float64) float64 {
		if v := math.Pow(x, n); v < lim {
			return v
		}
		return lim
	}
	// if and else
	func pow(x, n, lim float64) float64 {
		if v := math.Pow(x, n); v < lim {
			return v
		}else if v == 2{

		} else {
			fmt.Printf("%g >= %g\n", v, lim)
		}
		// can't use v here, though
		return lim
	}
	// switch
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	// switch with calculated-value
	// Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
	v := f()
	switch v {
	case g():
		fmt.Println("in g()")
	case h():
		fmt.Println("in h()")
	default:
		fmt.Println("not in h() or g()")
	}
	// switch with no condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// defer
	// A defer statement defers the execution of a function until the surrounding function returns.
	// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
	func CopyFile(dstName, srcName string) (written int64, err error) {
		src, err := os.Open(srcName)
		if err != nil {
			return
		}
	
		dst, err := os.Create(dstName)
		if err != nil {
			src.Close()
			return
		}
	
		written, err = io.Copy(dst, src)
		dst.Close()
		src.Close()
		return
	}
	func CopyFile(dstName, srcName string) (written int64, err error) {
		src, err := os.Open(srcName)
		if err != nil {
			return
		}
		defer src.Close()
		defer func(x int) int{
			// do something with x
			x
		}(x)
	
		dst, err := os.Create(dstName)
		if err != nil {
			return
		}
		defer dst.Close()
	
		return io.Copy(dst, src)
	}

	n, err := CopyFile("a.txt", "b.txt")

	// Ref: https://blog.golang.org/defer-panic-and-recover
}
