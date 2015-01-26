package ansifmt

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/scanner"
	"unicode"
)

//AnsiSize return the number of  Lines and  Columns in the current "terminal" (if available)
func AnsiSize() (lines, cols int, err error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Warning cannot exec 'stty size': %v", err)
		return
	}
	scols := strings.TrimSuffix(string(out), "\n")

	_, err = fmt.Sscanf(scols, "%d %d", &lines, &cols)
	if err != nil {
		log.Printf("Warning invalid 'stty size' result %q, %v", scols, err)
		return
	}
	return
}

//LineWrap executes the linewrap algorithm escaping ansi escape code.
// it returns a slice of lines.
//
//'length' is the max numbers in a column. It is respected unless a single word is too big to fit in
func LineWrap(txt string, length int) (lines []string) {

	var s scanner.Scanner
	s.Init(strings.NewReader(txt))

	lines = make([]string, 0, 10)  //lines
	line := new(bytes.Buffer)      //line buffer
	word := new(bytes.Buffer)      //word buffer (between splitter position)
	var linelength, wordlength int //lenght of current line and current word
	//nb line length an len(line) are not the same: the former is in visible char, the last is in byte

	var inescape bool // if I'm in "escape sequence"

	for tok := s.Next(); true; tok = s.Next() { //loop over all the text
		//wordlength should be inc, only if this is a visible char
		// in particular, if I'm in an ansi escape sequence \x1b[0;0;m that's a lot of char "unvisible"
		if tok != scanner.EOF && tok != '\n' {
			word.WriteRune(tok)
		}
		switch {
		case tok == '\n':
			word.WriteTo(line)
			lines = append(lines, line.String())
			line.Reset()
			linelength = 0
			wordlength = 0
		case inescape: // we are in escape mod
			//just check if we should get out
			if tok == 'm' { // en of it
				inescape = false //exit inescape sequence
			}
		case tok == '\x1b': //we enter escape mode
			inescape = true
		case unicode.IsSpace(tok) || unicode.IsPunct(tok) || tok == scanner.EOF: //we reach a word boundary
			wordlength++                                              //we count it
			if linelength+wordlength < length && tok != scanner.EOF { // not enought to break line
				word.WriteTo(line)
				linelength += wordlength
				wordlength = 0
			} else { // we need to flush

				//FLUSHING

				if (tok != scanner.EOF && linelength > 0) || (tok == scanner.EOF && linelength+wordlength >= length) { //we can use the line alone, just do it
					//flush and add
					lines = append(lines, line.String())
					line.Reset()
					linelength = 0

					word.WriteTo(line)
					linelength += wordlength
					wordlength = 0

				} else { //unfortunately the current is empty so we can't use it
					//so we use the word despite beeing too big
					word.WriteTo(line)
					lines = append(lines, line.String())
					line.Reset()
					linelength = 0
					wordlength = 0
				}

				if tok == scanner.EOF && linelength > 0 {
					// there are stuff to be flushed
					lines = append(lines, line.String())
					line.Reset()
				}

			}

		default: //not escape in or out not word boundary
			wordlength++
		}
		if tok == scanner.EOF {
			return
		}
	}
	return
}

//ToTitle maps the strings.ToTitle but support text containing ansi escape code
//
//     strings.ToTitle("\x1b[1mHello")
//     > "\x1b[1MHELLO"
//
//This is not a valid ansi escaped code.
//
func ToTitle(txt string) (title string) {

	var s scanner.Scanner
	s.Init(strings.NewReader(txt))

	var inescape bool // if I'm in "escape sequence"
	out := new(bytes.Buffer)

	for tok := s.Next(); tok != scanner.EOF; tok = s.Next() { //loop over all the text
		//wordlength should be inc, only if this is a visible char
		// in particular, if I'm in an ansi escape sequence \x1b[0;0;m that's a lot of char "unvisible"
		switch {
		case inescape: // we are in escape mod
			//just check if we should get out
			if tok == 'm' { // en of it
				inescape = false //exit inescape sequence
			}
		case tok == '\x1b': //we enter escape mode
			inescape = true
		default: //not escape in or out not word boundary
			tok = unicode.ToUpper(tok)
		}
		out.WriteRune(tok)
	}
	return out.String()
}
