package ansifmt

import "fmt"

func ExampleFormat_Coder() {
	var f Format
	f.SetStrike(true)
	striker := f.Coder()

	fmt.Printf("%q", striker("toto"))
	//Output: "\x1b[9mtoto\x1b[29m"
}
