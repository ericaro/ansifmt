package ansifmt

import "io"

//Writer implements Writer but it can change the output Format using ansi escape codes.
//
// It also implements Formatter interface, it is easy to specify a Format.
//
// Changes to the Format are buffered until the next Write. So it is not inefficient to toggle Format without writing anyting.
//
// Only changes in Format are written down.
//
// It has a method Reset that writedown immediatly a reset ansi code.
type Writer struct {
	Format
	current Format
	out     io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{out: w}
}

//Reset send the ansi code \033[0m immediatly to the stream.
//
//It does not change the current in memory Format. So the next call to Write() the ansi code will be recomputed.
//
// For instance
//
//    w:= NewWriter(os.Stdout)
//    w.SetWeight(BoldWeight)
//    fmt.Fprintf("Bolded Text")
//    fmt.Fprintf("Bolded Text")
//
//
func (p *Writer) Reset() error {
	_, err := p.out.Write(([]byte)(Reset))
	p.current = Format{} //make "current" match the actual Format (since it has been reseted, it is the defautl Format value)
	return err
}

func (p *Writer) Write(b []byte) (n int, err error) {
	// compute the possible change in Format
	pfix := AnsiCode(p.current, p.Format)
	if pfix != "" { // there are changes
		n, err = p.out.Write(([]byte)(pfix))
		p.current = p.Format //acknowledge that 'Format' as been set on the output, and therefore it is the new current
		if err != nil {
			return //on pfix error
		}
	}
	//n might not be zero, I need to append the second write to the first one
	secondN, err := p.out.Write(b)
	return n + secondN, err
}
