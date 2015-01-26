//renderer for russross's Blackfriday markdown processor.
package ansiblackfriday

import (
	"bytes"
	"fmt"
	"html"
	"strings"

	"github.com/ericaro/ansifmt"
	"github.com/russross/blackfriday"
)

type ansirenderer struct {
	doc                                 *bytes.Buffer // the final formatted doc
	listdepth, listcounter              int           // for nested list
	maxcols                             int           // for text wrapping
	striker, em1, em2, em3, under, code ansifmt.Coder
}

//NewAnsiRenderer creates a mardown renderer using ansi escape code as only formatting rules.
func NewAnsiRenderer() blackfriday.Renderer {
	_, cols, err := ansifmt.AnsiSize()
	if err != nil {
		cols = 80
	}

	r := &ansirenderer{
		//txtRenderer: txtRenderer{},
		maxcols: cols, //compute from tty the current number of col, or 80
	}

	//build the Coder (strike, and the three emphasis)
	var b ansifmt.Format
	b.SetStrike(true)
	r.striker = b.Coder()
	b.SetStrike(false)

	b.SetWeight(ansifmt.BoldWeight)
	r.em1 = b.Coder()

	b.SetWeight(ansifmt.FaintWeight)
	r.code = b.Coder()

	b.SetWeight(ansifmt.NormalWeight)
	b.SetReverse(true)
	r.em2 = b.Coder()

	b.SetWeight(ansifmt.BoldWeight)
	r.em3 = b.Coder()

	b.SetWeight(ansifmt.NormalWeight)
	b.SetReverse(false)
	b.SetUnder(true)
	r.under = b.Coder()

	return r

}

/* strategy: format is done on the "out" buffer for terminals (normal text, etc.)
on the other hand, top level (highest one) are flush the out buffer into the printer one, with some "rendering capability"
*/

func (d *ansirenderer) Header(out *bytes.Buffer, text func() bool, level int, id string) {
	//log.Printf("Header level=%v, id=%s", level, id)
	text()

	title := out.String() //by convention, I don't write useless line break in the title.
	title = ansifmt.ToTitle(title)
	out.Reset()
	block := indentText(title, level-1, d.maxcols)
	h := strings.Repeat(" ", level-1)
	// log.Printf("indent(%q,%v,%v)-> %q", title, level-1, d.maxcols, block)
	fmt.Fprintf(d.doc, "\n%s%s\n\n", h, d.em1(block))
}

func (d *ansirenderer) Paragraph(out *bytes.Buffer, text func() bool) {
	//log.Printf("Paragraph ")
	text()

	par := out.String() //by convention, I don't write useless line breaks
	out.Reset()

	h := strings.Repeat(" ", 4)

	fmt.Fprintf(d.doc, "\n%s%s\n", h, indentText(par, 4, d.maxcols))

}

func (d *ansirenderer) HRule(out *bytes.Buffer) {
	h := strings.Repeat(" ", d.maxcols/2-2)
	fmt.Fprintf(d.doc, "\n%s\u2500\u2500*\u2500\u2500\n", h)
}

func (d *ansirenderer) List(out *bytes.Buffer, text func() bool, flags int) {
	//log.Printf("List flags %v", flags)
	// I can do some nesting here
	d.listdepth++
	prev := d.listcounter
	d.listcounter = 0
	text()
	d.listdepth--
	d.listcounter = prev

	if d.listdepth == 0 { //top level list
		itemtext := out.String()
		out.Reset()
		fmt.Fprintln(d.doc, itemtext)
	}

}

func (d *ansirenderer) ListItem(out *bytes.Buffer, text []byte, flags int) {
	//log.Printf("List Item %q : %v", text, flags)
	bullet := "  \u2022 "
	d.listcounter++
	if flags&blackfriday.LIST_TYPE_ORDERED != 0 {
		switch {
		case d.listcounter < 10:
			bullet = fmt.Sprintf("  %v. ", d.listcounter)
		case d.listcounter < 100:
			bullet = fmt.Sprintf(" %v. ", d.listcounter)
		}
	}

	i := 4 //+ d.listdepth*4
	block := indentText(string(text), i, d.maxcols)
	// now I need to put the bullet in the right place
	h := strings.Repeat(" ", 4) //d.listdepth*4)
	fmt.Fprintf(out, "%s%s%s\n", h, bullet, block)
}

func (d *ansirenderer) BlockCode(out *bytes.Buffer, text []byte, lang string) {
	txt := string(text)
	//indent it

	fmt.Fprintln(d.doc)
	for _, line := range strings.Split(txt, "\n") {
		fmt.Fprintf(d.doc, "        %s\n", d.code(line))
	}
}

func (d *ansirenderer) NormalText(out *bytes.Buffer, text []byte) {
	out.Write(text)
}

func (d *ansirenderer) Entity(out *bytes.Buffer, entity []byte) {
	out.WriteString(html.UnescapeString(string(entity)))
}
func (d *ansirenderer) CodeSpan(out *bytes.Buffer, text []byte) {
	txt := string(text)
	out.WriteString(d.code(txt))
}
func (d *ansirenderer) StrikeThrough(out *bytes.Buffer, text []byte) {
	//log.Printf("StrikeThrough %q", text)
	txt := string(text)
	out.WriteString(d.striker(txt))

}

func (d *ansirenderer) AutoLink(out *bytes.Buffer, link []byte, kind int) {
	//log.Printf("AutoLink %q: %v", link, kind)
	txt := string(link)
	out.WriteString(d.under(txt))
}

func (d *ansirenderer) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {
	//log.Printf("Link [%q](%q)", content, link)
	txt := fmt.Sprintf("%s ( %s )", content, link)
	out.WriteString(d.under(txt))

}

//text is simply copied
func (d *ansirenderer) Emphasis(out *bytes.Buffer, text []byte) {
	txt := string(text)
	out.WriteString(d.em1(txt))
}
func (d *ansirenderer) DoubleEmphasis(out *bytes.Buffer, text []byte) {
	txt := string(text)
	out.WriteString(d.em2(txt))
}
func (d *ansirenderer) TripleEmphasis(out *bytes.Buffer, text []byte) {
	txt := string(text)
	out.WriteString(d.em3(txt))
}

func (d *ansirenderer) DocumentHeader(out *bytes.Buffer) {
	d.doc = new(bytes.Buffer)
}
func (d *ansirenderer) DocumentFooter(out *bytes.Buffer) {
	out.Reset()
	d.doc.WriteTo(out)
}
func (d *ansirenderer) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {
	out.Write(alt) //that what's alt is for, right ?
}
func (d *ansirenderer) LineBreak(out *bytes.Buffer) { out.WriteString("\n\n") }
func (d *ansirenderer) GetFlags() int               { return 0 }

func (d *ansirenderer) Footnotes(out *bytes.Buffer, text func() bool)                 { /*ignored */ }
func (d *ansirenderer) TableRow(out *bytes.Buffer, text []byte)                       { /*ignored */ }
func (d *ansirenderer) TableHeaderCell(out *bytes.Buffer, text []byte, flags int)     { /*ignored */ }
func (d *ansirenderer) TableCell(out *bytes.Buffer, text []byte, flags int)           { /*ignored */ }
func (d *ansirenderer) FootnoteItem(out *bytes.Buffer, name, text []byte, flags int)  { /*ignored */ }
func (d *ansirenderer) TitleBlock(out *bytes.Buffer, text []byte)                     { /*ignored */ }
func (d *ansirenderer) FootnoteRef(out *bytes.Buffer, ref []byte, id int)             { /*ignored */ }
func (d *ansirenderer) RawHtmlTag(out *bytes.Buffer, tag []byte)                      { /*ignored */ }
func (d *ansirenderer) Table(out *bytes.Buffer, header []byte, body []byte, cd []int) { /*ignored */ }
func (d *ansirenderer) BlockQuote(out *bytes.Buffer, text []byte)                     { /*ignored */ }
func (d *ansirenderer) BlockHtml(out *bytes.Buffer, text []byte)                      { /*ignored */ }

//indentText will break the txt into sm
func indentText(txt string, indent, length int) string {

	out := new(bytes.Buffer)
	for i, line := range ansifmt.LineWrap(txt, length-indent-4) {
		h := "\n" + strings.Repeat(" ", indent)
		if i > 0 {
			out.WriteString(h)
		}
		out.WriteString(line)
	}
	return out.String()
}
