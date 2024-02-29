package commands

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/buger/goterm"
)

func CreateProjectDir(module string) (string, string, error) {
	var dir string

	if strings.Contains(module, "/") {
		splitModule := strings.Split(module, "/")
		dir = splitModule[len(splitModule)-1]
	} else {
		dir = module
	}

	err := os.Mkdir(dir, 0777)

	if err != nil {
		if os.IsExist(err) {
			log.Fatal(goterm.Color(fmt.Sprintf("\nDirectory %s already exists\n", dir), goterm.RED))
			Fatal("", err)
		}
		log.Fatal(goterm.Color(fmt.Sprintf("\nError while creating directory: %s\n", err), goterm.RED))
	}

	fullPath, err := filepath.Abs(dir)

	if err != nil {
		Fatal(dir, err)
	}

	return fullPath, dir, nil
}

func GenerateGoProject(module string, dir string, fullPath string, packages []string) {
	err := os.Chdir(fullPath)
	if err != nil {
		Fatal(fullPath, err)
	}

	cmd := exec.Command("go", "mod", "init", module)
	out, err := cmd.CombinedOutput()

	if err != nil {
		Fatal(fullPath, err)
	}

	fmt.Println(string(out))
	fmt.Println(goterm.Color("Initialized go module.\n\n", goterm.BLUE))

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
		fmt.Println(goterm.Color("\nInstalling packages...\n", goterm.BLUE))
		output, err := installPackage(pkg)

		if err != nil {
			fmt.Println(goterm.Color(fmt.Sprintf("\nError installing package %s: %s\n", pkg, err), goterm.RED))
			continue
		}

		fmt.Println(output)
		pkgContent.WriteString(fmt.Sprintf("\n\t_ \"%s\"", pkg))
	}

	toolsContent = fmt.Sprintf(toolsContent, pkgContent.String())

	toolsFile, err := os.Create("tools.go")
	if err != nil {
		Fatal(fullPath, err)
	}
	err = os.Chmod("tools.go", 0777)
	if err != nil {
		Fatal(fullPath, err)
	}

	defer toolsFile.Close()

	_, err = toolsFile.WriteString(toolsContent)
	if err != nil {
		Fatal(fullPath, err)
	}

	fmt.Println(goterm.Color("Packages installed successfully.\n", goterm.BLUE))

	cmd = exec.Command("go", "mod", "tidy")
	out, err = cmd.CombinedOutput()

	if err != nil {
		Fatal(fullPath, err)
	}

	fmt.Println(string(out))
	fmt.Println(goterm.Color("Project is ready.\n", goterm.BLUE))
	fmt.Println(goterm.Color(fmt.Sprintf("\nYou can now run `cd %s`\n", dir), goterm.BLUE))
	fmt.Println(goterm.Color("\nThe packages are imported in the tools.go file, you can delete this file after you import them in other files.\n", goterm.BLUE))
	fmt.Println(goterm.Color("\nHappy coding!\n", goterm.CYAN))
}

func installPackage(pkg string) (string, error) {
	cmd := exec.Command("go", "get", pkg)
	out, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	return string(out), nil
}

func Fatal(dir string, err error) {
	if dir != "" {
		errRemove := os.RemoveAll(dir)
		if errRemove != nil {
			fmt.Println(goterm.Color(fmt.Sprintf("Error removing directory: %s\n", errRemove), goterm.RED))
		}
	}
	log.Fatal(goterm.Color(fmt.Sprintf("Error: %s\n", err), goterm.RED))
}
