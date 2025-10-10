package latex

// A LaTeXer is a self-contained, compilable document.
type LaTeXer interface {
	LaTeX() string
}

// LaTeXString is a simple wraper around the string type to make it a LaTeXer.
type LaTeXString string

func (ls LaTeXString) LaTeX() string {
	return string(ls)
}

// A BibTeXer can generate an entire BibTex file.
type BibTeXer interface {
	BibTex() string
}

// A LaTeXRenderer can be rendered as a LaTeX fragment, but not an entire document.
type LaTeXRenderer interface {
	RenderLaTeX() string
}

// A BibTexRenderer can be rendered as a BibTex fragment (e.g. a single bib resource), but not an entire .bib file.
type BibTeXRenderer interface {
	RenderBibTex() string
}

// Command is experimental and incomplete. You probably shouldn't use it.
type Command interface {
	LaTeXCmd(args ...string) string
}

// CustomCommand is experimental and incomplete. You probably shouldn't use it.
type CustomCommand interface {
	Command
	Define() string
}
