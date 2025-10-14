package stddocs

import (
	"github.com/ag7if/go-files"

	"github.com/ag7if/go-latex"
)

type Article struct {
}

func NewArticle() *Article {
	panic("implement me")
}

func (a *Article) LaTeX() string {
	latex := `\documentclass{article}`
}

func (a *Article) AssetDir() *files.Directory {
	//TODO implement me
	panic("implement me")
}

func (a *Article) Assets() []files.File {
	//TODO implement me
	panic("implement me")
}

func (a *Article) BuildEngine() latex.BuildEngine {
	//TODO implement me
	panic("implement me")
}

func (a *Article) IncludePackage(pkg latex.Package, args ...string) {
	//TODO implement me
	panic("implement me")
}

func (a *Article) WriteClass() error {
	//TODO implement me
	panic("implement me")
}

func (a *Article) WritePackages() error {
	//TODO implement me
	panic("implement me")
}
