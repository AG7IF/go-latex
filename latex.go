package latex

type LaTeXer interface {
	LaTeX() string
}

type Commander interface {
	LaTeXCmd(args ...string) string
}

type CustomCommander interface {
	Commander
	Define() string
}
