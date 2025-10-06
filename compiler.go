package latex

import (
	"fmt"
	"os/exec"

	"github.com/ag7if/go-files"
	"github.com/pkg/errors"
)

type Compiler struct {
	buildCommand BuildCommand
	bibCommand   BibCommand
	buildDir     files.Directory
}

func NewCompiler(buildCmd BuildCommand, bibCmd BibCommand, buildDir files.Directory) Compiler {
	return Compiler{
		buildCommand: buildCmd,
		bibCommand:   bibCmd,
		buildDir:     buildDir,
	}
}

func (lc Compiler) GenerateLaTeX(latex LaTeXer, outputFile files.File, assets []files.File) error {
	for _, asset := range assets {
		_, err := lc.buildDir.CopyFileToDir(asset)
		if err != nil {
			// TODO: React to whether this build asset has already been copied
			return errors.WithMessagef(err, "failed to copy build asset: %s", asset)
		}
	}

	texSourceFileName := fmt.Sprintf("%s.tex", outputFile.Base())
	texSource, err := lc.buildDir.NewFile(texSourceFileName)
	if err != nil {
		return errors.WithStack(err)
	}
	err = texSource.WriteString(latex.LaTeX())
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (lc Compiler) CompileLaTeX(outputFile files.File) error {
	texSourceFile, err := lc.buildDir.NewFile(fmt.Sprintf("%s.tex", outputFile.Base()))
	if err != nil {
		return errors.WithStack(err)
	}

	// First run
	cmd := exec.Command(lc.buildCommand.String(), "-halt-on-error", texSourceFile.Name())
	cmd.Dir = lc.buildDir.Path()
	err = cmd.Run()
	if err != nil {
		return errors.WithStack(err)
	}

	// Build Bibliography
	if lc.bibCommand != NoBib {
		cmd := exec.Command(lc.bibCommand.String(), texSourceFile.Base())
		cmd.Dir = lc.buildDir.Path()
		err = cmd.Run()
		if err != nil {
			return errors.WithStack(err)
		}
	}

	// Second run
	cmd = exec.Command(lc.buildCommand.String(), "-halt-on-error", texSourceFile.Name())
	cmd.Dir = lc.buildDir.Path()
	err = cmd.Run()
	if err != nil {
		return errors.WithStack(err)
	}

	// Third run
	cmd = exec.Command(lc.buildCommand.String(), "-halt-on-error", texSourceFile.Name())
	cmd.Dir = lc.buildDir.Path()
	err = cmd.Run()
	if err != nil {
		return errors.WithStack(err)
	}

	cachedBuildfile, err := lc.buildDir.NewFile(fmt.Sprintf("%s.pdf", outputFile.Base()))
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = outputFile.Dir().MoveFileToDir(cachedBuildfile)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
