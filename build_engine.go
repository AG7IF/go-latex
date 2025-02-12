package latex

import (
	"github.com/pkg/errors"
)

type BuildEngine int

const (
	LuaLaTex BuildEngine = iota
	PDFLaTeX
	PDFCSLaTeX
	XeLaTeX
)

func (be BuildEngine) CompileCommand() string {
	switch be {
	case LuaLaTex:
		return "lualatex"
	case PDFLaTeX:
		return "pdflatex"
	case PDFCSLaTeX:
		return "pdfcslatex"
	case XeLaTeX:
		return "xelatex"
	default:
		panic(errors.Errorf("unknown build engine: %d", be))
	}
}
