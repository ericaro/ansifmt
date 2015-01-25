package ansifmt

import (
	"bytes"
	"fmt"
)

func ExampleAnsiWriter_Write() {
	out := new(bytes.Buffer)
	w := NewWriter(out)

	func() { // a bit artifical here, but the idea is to show how to use Set, Unset
		defer w.SetWeight(w.SetWeight(BoldWeight))
		fmt.Fprintf(w, "Bolded")
	}()

	func() { // a bit artifical here, but the idea is to show how to use Set, Unset
		defer w.SetForeground(w.SetForeground(RedColor))
		fmt.Fprintf(w, "Red")
	}()

	//just print the output escaped so we can clearly see the result
	fmt.Printf("%q", out.String())
	// Output: "\x1b[1mBolded\x1b[22;31mRed"
}

func ExampleAnsiWriter_Reset() {
	out := new(bytes.Buffer) //somehow you get one

	w := NewWriter(out) // build the ansiwriter

	w.SetWeight(BoldWeight) // set any
	fmt.Fprintf(w, "Bolded Text")

	w.Reset()
	fmt.Fprintf(w, "Bolded Text Again") // the format has been reset so the ansicode reset the bold format
	fmt.Printf("%q", out.String())
	// Output: "\x1b[1mBolded Text\x1b[0m\x1b[1mBolded Text Again"

}
