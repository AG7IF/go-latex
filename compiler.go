package latex

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/ag7if/go-files"
	"github.com/pkg/errors"
)

type Compiler struct {
	cmd      string
	assetDir *files.Directory
	buildDir *files.Directory
}

func NewCompiler(engine BuildEngine, assetDir, buildDir *files.Directory) Compiler {
	cmd := engine.CompileCommand()

	return Compiler{
		cmd:      cmd,
		assetDir: assetDir,
		buildDir: buildDir,
	}
}

func (lc Compiler) GenerateLaTeX(latex LaTeXer, outputFile files.File, assets []files.File) (files.File, error) {
	for _, asset := range assets {
		_, err := asset.ForceCopy(lc.buildDir.Path())
		if err != nil {
			return files.File{}, errors.WithMessagef(err, "failed to copy build asset: %s", asset.FullPath())
		}
	}

	inputFile, err := files.NewFile(filepath.Join(lc.buildDir.Path(), fmt.Sprintf("%s.tex", outputFile.Base())))
	if err != nil {
		return files.File{}, errors.WithStack(err)
	}

	err = inputFile.WriteString(latex.LaTeX())
	if err != nil {
		return files.File{}, errors.WithStack(err)
	}

	return inputFile, nil
}

func (lc Compiler) CompileLaTeX(inputFile, outputFile files.File) error {
	// First run
	jobname := fmt.Sprintf("-jobname=%s", outputFile.Base())
	cmd := exec.Command(lc.cmd, jobname, "-halt-on-error", inputFile.Name())
	cmd.Dir = lc.buildDir.Path()

	err := cmd.Run()
	if err != nil {
		return errors.WithStack(err)
	}

	// Second run (xelatex usually needs two runs to get formatting right)
	cmd = exec.Command(lc.cmd, jobname, "-halt-on-error", inputFile.Name())
	cmd.Dir = lc.buildDir.Path()

	cachedBuildFile, err := files.NewFile(filepath.Join(lc.buildDir.Path(), fmt.Sprintf("%s.tex", outputFile.Base())))
	if err != nil {
		return errors.WithStack(err)
	}

	outputDir := outputFile.Dir()
	_, err = outputDir.MoveFileToDir(cachedBuildFile)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
