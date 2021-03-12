package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func createGenericFiles(fpath string) error {
	sep := string(filepath.Separator)
	return createFile(fpath + sep + "README.md")
}

func initGit(fpath string) (outp string, err error) {
	sep := string(filepath.Separator)
	_ = createFile(fpath + sep + ".gitignore")
	result, err := execute("git init")
	return result, err
}

func createProject(projectType string, projectName string, projectPath string, git bool) (resultlist []string, errorlist []error) {
	sep := string(filepath.Separator)
	fpath := projectPath + sep + projectName

	switch projectType {
	case "python":
		if err := createGenericFiles(fpath); err != nil {
			errorlist = append(errorlist, err)
		}
		if err := os.Chdir(fpath); err != nil {
			errorlist = append(errorlist, err)
		}
		if err := createDir(projectName); err != nil { //the application goes in a folder by the same name
			errorlist = append(errorlist, err)
		}
		if err := createDir(projectName + "/tests"); err != nil { //the app tests go with the app
			errorlist = append(errorlist, err)
		}
		if err := createFile(fpath + sep + "requirements.txt"); err != nil {
			errorlist = append(errorlist, err)
		}
		if ! isInstalled("pipenv") {
			_, err := execute("python3 -m pip install pipenv")
			if err != nil {
				fmt.Println("Cannot create virtual environment: ", err.Error())
			}
		}
		result, err := execute("pipenv install")     // use pipenv to generate dependency files
		if err != nil {
			fmt.Println(err.Error(), result)
			errorlist = append(errorlist, err)
		}
		for _, record := range strings.Split(result, "\n") {
			if strings.Contains(record, "Virtualenv location"){
				resultlist = append(resultlist, record)
			}
		}

	case "go":
		if err := createGenericFiles(fpath); err != nil {
			errorlist = append(errorlist, err)
		}
		if err := os.Chdir(fpath); err != nil {
			errorlist = append(errorlist, err)
		}
		if err := createDir("vendor"); err != nil {
			errorlist = append(errorlist, err)
		}
		if err := createDir("build"); err != nil {
			errorlist = append(errorlist, err)
		}

		result, err := execute("go mod init " + projectName)     // use mod init to generate the go.mod
		if err != nil {
			errorlist = append(errorlist, err)
		}
		resultlist = append(resultlist, strings.Split(result, "\n")[0]) //just the first line

		if err := createFile(fpath + sep + "go.sum"); err != nil {
			errorlist = append(errorlist, err)
		}

	default:
		err := fmt.Errorf("Project type '%s' is not supported.\n", projectType)
		errorlist = append(errorlist, err)
	}

	if git {
		result, err := initGit(fpath)
		result = strings.Split(result, "\n")[0] //strip the carriage return
		if err != nil {
			errorlist = append(errorlist, err)
		}
		resultlist = append(resultlist, result)
	}

	return resultlist, errorlist
}

func main() {
	argCount := len(os.Args[1:])

	var (
		projectName string
		projectType string
		projectPath string
		git         bool
	)

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
		results, errs := createProject(projectType, projectName, projectPath, git)
			for _, err := range errs {
				fmt.Println("ERR: ", err)
		}
		result := fmt.Sprintf(
			"Created a '%s' project named '%s' at directory '%s'\n", projectType, projectName, projectPath)
		results = append(results, result)
		for _, result := range results {
				fmt.Println(result)
		}
	}

}
