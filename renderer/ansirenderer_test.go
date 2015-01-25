package renderer

import "fmt"

func ExampleIndentText() {
	txt := IndentText("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus scelerisque a purus et sagittis. Cras odio tellus, maximus non nunc.", 2, 20)
	//txt := IndentText("Lorem ipsum dolor sit amet, consectetur", 2, 20)
	fmt.Printf("%s", txt)
	//Output:Lorem ipsum
	//   dolor sit amet,
	//   consectetur
	//   adipiscing elit.
	//   Phasellus
	//   scelerisque a
	//   purus et
	//   sagittis. Cras
	//   odio tellus,
	//   maximus non
	//   nunc.
}

func ExampleIndentText_Singleline() {
	txt := IndentText("List Support", 1, 170)
	fmt.Printf("%q", txt)
	//Output: "List Support"
}