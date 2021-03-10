package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func createGenericFiles(fpath string) error {
	sep := string(filepath.Separator)
	return createFile(fpath + sep + "README.md")
}

func initGit(fpath string) (outp string, err error) {
	sep := string(filepath.Separator)
	fmt.Println("Initializing git repo...")
	_ = createFile(fpath + sep + ".gitignore")
	return execute("git init")
}


func createProject(projectType string, projectName string, projectPath string, git bool) (result string, err error) {
	sep := string(filepath.Separator)
	fpath := projectPath + sep  + projectName

	switch projectType {
	case "python":
		result = fmt.Sprintf("Created a python project named '%s' at directory '%s'\n", projectName, projectPath)
	case "go":
		var errors []error
		errors = append(errors, createGenericFiles(fpath))                    // Also creates base directory
		errors = append(errors, os.Chdir(fpath))                              // CD into the base directory
		errors = append(errors, createDir( "vendor"))                   // create the vendor dir
		errors = append(errors, createDir("build"))                     // create the build directory for binaries
		_, ee := execute("go mod init " + projectName)                        // use mod init to generate the go.mod
		errors = append(errors, ee)
		errors = append(errors, createFile(fpath + sep + "go.sum"))     // manually create the go.sum
		for _, erline := range errors {
			if erline != nil {
				fmt.Println(erline)
			}
		}
		for _, line := range dirlist() {
			fmt.Println(line)
		}
		result = fmt.Sprintf("Created a go project named '%s' at directory '%s'\n", projectName, projectPath)
	default:
		return "", errors.New(fmt.Sprintf("Project type '%s' is not supported.\n", projectType))
	}

	if git {
		res, errr := initGit(fpath)
		if errr != nil {
			fmt.Println(res)
		}
	}

	return result, err
}

func main() {
	argCount := len(os.Args[1:])

	var projectName string
	var projectType string
	var projectPath string
	var git bool

	flag.StringVar(&projectName, "n", "", "Name of project.")
	flag.StringVar(&projectType, "t", "go", "Type of project.")
	flag.StringVar(&projectPath, "p", ".", "Directory path for project.")
	flag.BoolVar(&git, "g", false, "Initialize git repo.")
	flag.Parse()

	if argCount == 0 {
		flag.Usage()
	}

	if argCount > 0 {
		if projectName == "" {
			fmt.Println("Oops! No project name is provided. What do you want to call your project?")
			os.Exit(1)
		}
		result, err := createProject(projectType, projectName, projectPath, git)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	}

}
