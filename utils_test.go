package ansifmt

import "fmt"

func ExampleLineWrap_Simple() {

	lines := LineWrap("toto is a weird person", 6)
	for i, line := range lines {
		fmt.Printf("%v: %q\n", i, line)
	}
	//Output:
	//0: "toto "
	//1: "is a "
	//2: "weird "
	//3: "person"

}
func ExampleLineWrap() {

	lines := LineWrap("toto \x1b[1mis\x1b[0m a weird person", 6)
	for i, line := range lines {
		fmt.Printf("%v: %q\n", i, line)
	}
	//Output:
	//0: "toto "
	//1: "\x1b[1mis\x1b[0m a "
	//2: "weird "
	//3: "person"

}

func ExampleToTitle() {
	title := "this \x1b[1mis\x1b[0m a title"
	Title := ToTitle(title)
	fmt.Printf("%q", Title)
	//Output:"THIS \x1b[1mIS\x1b[0m A TITLE"
}
