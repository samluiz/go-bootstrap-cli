package commands

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func createProjectDir(dir string) (string, error) {
	err := os.Mkdir(dir, 0777)

	if err != nil {
		return "", err
	}

	return filepath.Abs(dir)
}

func GenerateGoProject(module string, packages []string) {
	var dir string

	if strings.Contains(module, "/") {
		splitModule := strings.Split(module, "/")
		dir = splitModule[len(splitModule)-1]
	} else {
		dir = module
	}

	fullPath, err := createProjectDir(dir)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			fmt.Printf("Directory %s already exists\n", fullPath)
			return
		}
		exitOnError(fullPath, err)
	}

	fmt.Printf("1. Created directory: %s\n", fullPath)

	err = os.Chdir(dir)
	if err != nil {
		exitOnError(fullPath, err)
	}

	cmd := exec.Command("go", "mod", "init", module)
	out, err := cmd.CombinedOutput()

	if err != nil {
		exitOnError(fullPath, err)
	}

	fmt.Println(string(out))
	fmt.Println("2. Initialized go module.")

	toolsContent := `
//go:build tools

package main

import (
	%s
)
`
	var pkgContent strings.Builder

	for _, pkg := range packages {
		if pkg == "" {
			continue
		}
		output, err := installPackage(pkg)

		if err != nil {
			fmt.Printf("Error installing package %s: %s\n", pkg, err)
			continue
		}

		fmt.Println(output)
		pkgContent.WriteString(fmt.Sprintf("\n\t_ \"%s\"", pkg))
	}

	toolsContent = fmt.Sprintf(toolsContent, pkgContent.String())

	toolsFile, err := os.Create("tools.go")
	if err != nil {
		exitOnError(fullPath, err)
	}
	err = os.Chmod("tools.go", 0777)
	if err != nil {
		exitOnError(fullPath, err)
	}

	defer toolsFile.Close()

	_, err = toolsFile.WriteString(toolsContent)
	if err != nil {
		exitOnError(fullPath, err)
	}

	fmt.Println("3. Packages installed successfully.")

	cmd = exec.Command("go", "mod", "tidy")
	out, err = cmd.CombinedOutput()

	if err != nil {
		exitOnError(fullPath, err)
	}

	fmt.Println(string(out))
	fmt.Println("4. Module is ready. Happy coding!")
}

func installPackage(pkg string) (string, error) {
	cmd := exec.Command("go", "get", pkg)
	out, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	return string(out), nil
}

func exitOnError(dir string, err error) {
	fmt.Println(err)
	errRemove := os.RemoveAll(dir)
	if errRemove != nil {
		fmt.Printf("Error removing directory: %s\n", errRemove)
	}
	panic(err)
}
