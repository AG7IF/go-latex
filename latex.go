package latex

type LaTeXer interface {
	LaTeX() string
}

type Command interface {
	LaTeXCmd(args ...string) string
}

type CustomCommand interface {
	Command
	Define() string
}
