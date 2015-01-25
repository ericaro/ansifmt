package ansifmt

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	//caveat: we need to start with Color because iota value matters for "colors"
	BlackColor = Color(iota)
	RedColor
	GreenColor
	YellowColor
	BlueColor
	MagentaColor
	CyanColor
	WhiteColor
	//the reset code
	Reset = "\x1b[0m"

	NormalWeight = Weight(iota)
	BoldWeight
	FaintWeight
)

//Weight one of NormalWeight, BoldWeight or FaintWeight
type Weight int

//Color  one of the ansi color
type Color int

//Coder is a function that will decorate the input string with ansi escape code.
type Coder func(string) string

//Format describe a terminal Format. Limited capacity due to limited ansi escape capacity
// works nicely with the following pattern
//
//    defer f.SetWeight( f.SetWeight(BoldWeight) )
//
// Because, `f.SetWeight(BoldWeight)` is evaluated immediatly, the format is now "Bold"
// It will be reset to previous value when the deferred call will happen
//
//
type Format struct {
	w  Weight
	r  bool  //reversed color
	s  bool  // striketrough
	u  bool  // underline
	fg Color //foreground
	bg Color // background
}

// private direct switcher
// using switcher is a good way to hold a frame
// defer set( set(newvalue) )
// the newvalue is set before (because it is evaluated as an argument to the defered function)
// but the oldvalue (the result of set(newvalue)) will be called 'after' (that what's the defer is for).

//SetWeight and return previous value
func (f *Format) SetWeight(w Weight) (c Weight) { c, f.w = f.w, w; return }
func (f Format) Weight() Weight                 { return f.w }

//SetForeground and return previous value
func (f *Format) SetForeground(color Color) (c Color) { c, f.fg = f.fg, color; return }
func (f Format) Foreground() Color                    { return f.fg }

//SetBackground and return previous value
func (f *Format) SetBackground(color Color) (c Color) { c, f.bg = f.bg, color; return }
func (f Format) Background() Color                    { return f.bg }

//SetReverse and return previous value
func (f *Format) SetReverse(r bool) (c bool) { c, f.r = f.r, r; return }
func (f Format) Reverse() bool               { return f.r }

//SetStrike and return previous value
func (f *Format) SetStrike(s bool) (c bool) { c, f.s = f.s, s; return }
func (f Format) Strike() bool               { return f.s }

//SetUnder and return previous value
func (f *Format) SetUnder(u bool) (c bool) { c, f.u = f.u, u; return }
func (f Format) Under() bool               { return f.u }

func (f Format) Coder() Coder {

	var zero Format //the empty Format
	prefix := AnsiCode(zero, f)
	suffix := AnsiCode(f, zero)
	return func(str string) string {
		return prefix + str + suffix
	}
}

//AnsiCode compute the ansi escape string to change from one Format to another
func AnsiCode(from, to Format) (code string) {

	codes := make([]int, 0)

	if from.w != to.w { // the weight has changed
		switch to.w {
		case BoldWeight:
			codes = append(codes, 1)
		case FaintWeight:
			codes = append(codes, 2)
		default:
			codes = append(codes, 22) //neither bold nor faint
		}
	}

	if from.u != to.u {
		if to.u {
			codes = append(codes, 4)
		} else {
			codes = append(codes, 24)
		}
	}

	if from.r != to.r {
		if to.r {
			codes = append(codes, 7)
		} else {
			codes = append(codes, 27)
		}
	}

	if from.s != to.s {
		if to.s {
			codes = append(codes, 9)
		} else {
			codes = append(codes, 29)
		}
	}

	if from.fg != to.fg {
		codes = append(codes, 30+int(to.fg))
	}
	if from.bg != to.bg {
		codes = append(codes, 40+int(to.bg))
	}

	if len(codes) == 0 { //empty no Format differences
		return ""
	}
	//Format the code
	scodes := make([]string, len(codes))
	for i, c := range codes {
		scodes[i] = strconv.FormatInt(int64(c), 10)
	}
	return fmt.Sprintf("\033[%sm", strings.Join(scodes, ";"))
}
