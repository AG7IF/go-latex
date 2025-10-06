package latex

type BuildCommand int

const (
	PdfLaTeX BuildCommand = iota
	XeLaTeX
	LuaLaTeX
)

func (bc BuildCommand) String() string {
	switch bc {
	case PdfLaTeX:
		return "pdflatex"
	case XeLaTeX:
		return "xelatex"
	case LuaLaTeX:
		return "lualatex"
	default:
		panic("invalid build command")
	}
}

type BibCommand int

const (
	NoBib BibCommand = iota
	BibTeX
	BibLaTeX
	Biber
)

func (bc BibCommand) String() string {
	switch bc {
	case BibTeX:
		return "bibtex"
	case BibLaTeX:
		return "biblatex"
	case Biber:
		return "biber"
	default:
		panic("invalid bib command")
	}
}
