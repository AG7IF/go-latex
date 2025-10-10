package latex

import (
	"github.com/ag7if/go-files"
	"github.com/pkg/errors"
)

// Style is experimental and incomplete. You probably shouldn't use it.
type Style interface {
	AddCommand(command Command, args ...string)
	AddDef(command CustomCommand)
	StyleDef() string
}

// Document is experimental and incomplete. You probably shouldn't use it.
type Document interface {
	LaTeXer
	AddAsset(asset string)
	AddInclude(packageName, options string)
	AddStyle(style Style)
	AssetDir() string
	Assets() []files.File
	BuildDir() files.Directory
	SetClass(className, options string)
}

// BuildDocument is an incomplete document construction system. I don't recommend using it.
func BuildDocument(document Document, outputFile files.File) error {
	c := NewCompiler(XeLaTeX, Biber, document.BuildDir())

	err := c.GenerateLaTeX(document, outputFile, document.Assets())
	if err != nil {
		return errors.WithStack(err)
	}

	err = c.CompileLaTeX(outputFile)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
