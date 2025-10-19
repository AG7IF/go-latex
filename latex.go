package latex

// A LaTeXer is a self-contained, compilable document.
type LaTeXer interface {
	LaTeX() []byte
}

// LaTeXString is a simple wraper around the string type to make it a LaTeXer.
type LaTeXString string

func (ls LaTeXString) LaTeX() []byte {
	return []byte(ls)
}

// A BibTeXer can generate an entire BibTex file.
type BibTeXer interface {
	BibTex() []byte
}

// A LaTeXRenderer can be rendered as a LaTeX fragment, but not an entire document.
type LaTeXRenderer interface {
	RenderLaTeX() []byte
}

// A BibTexRenderer can be rendered as a BibTex fragment (e.g. a single bib resource), but not an entire .bib file.
type BibTeXRenderer interface {
	RenderBibTex() []byte
}
