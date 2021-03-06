package main

/*
BIG: FAT: FUCKING: TODO:
FINISH ADDING SUPPORT FOR DA MOFUKING LIBRARIEEEEEEEEEES
and also make it so the main file is named after the project?
TODO: start showing the options when using etr
*/

import (
	"bufio"
	"embed"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var GREEN string = "\033[38;5;184m"
var ORANGE string = "\033[38;5;202m"
var GRAY string = "\033[38;5;109m"
var RED string = "\033[38;5;160m"
var RESET string = "\033[0m"

//go:embed templates
var basefiles embed.FS

type UserData struct {
	PJName    string
	Language  string
	License   string
	LicNum    string
	Author    string
	Year      string
	VCS       string
	Pwd       string
	Build     string
	Lib       bool
	LibString string
}

func isBetween(num, min, max int) bool {
	//exclusive on top
	if (num >= min) && (num < max) {
		return true
	} else {
		return false
	}
}

func ErrCheck(err error) {
	if nil != err {
		fmt.Printf("Error: %v\n", err)
		os.Exit(-1)
	}
}

func getTime() string {
	t := time.Now()
	currentyear := t.Year()

	return strconv.Itoa(currentyear)
}

func SanitizeStrings(dirtyString *string) {

	*dirtyString = strings.ReplaceAll(*dirtyString, "\n", "")
	*dirtyString = strings.TrimSpace(*dirtyString)
}

func getData(ud *UserData) {
	var err error

	ud.Pwd, err = os.Getwd()
	ErrCheck(err)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(GREEN + "Project" + RESET + " name: ")
	ud.PJName, err = reader.ReadString('\n')
	ErrCheck(err)
	SanitizeStrings(&ud.PJName)

	fmt.Printf(GREEN + "Language: " + RESET)
	fmt.Printf(ORANGE + "\n\tc | cpp | go | rs | zig\n" + RESET)
	fmt.Printf("Selection: " + GREEN)
	ud.Language, err = reader.ReadString('\n')
	fmt.Printf(RESET + "\n")
	ErrCheck(err)
	SanitizeStrings(&ud.Language)

	if ud.Language != "go" && ud.Language != "rs" && ud.Language != "zig" {
		fmt.Printf(GREEN + "Build system: " + RESET)
		fmt.Printf(ORANGE + "\n\tmake | cmake | meson | premake\n" + RESET)
		fmt.Printf("Selection: " + GREEN)
		ud.Build, err = reader.ReadString('\n')
		fmt.Printf(RESET + "\n")
		ErrCheck(err)
		SanitizeStrings(&ud.Build)
		if ud.Build == "cmake" {
			fmt.Printf(GREEN + "Is this a library?" + RESET)
			fmt.Printf("\n\t[Y]es | [N]o\n")
			fmt.Printf("Selection: " + GREEN)
			ud.LibString, err = reader.ReadString('\n')
			fmt.Printf(RESET + "\n")
			ErrCheck(err)
			SanitizeStrings(&ud.LibString)
			if ud.LibString == "yes" || ud.LibString == "y" || ud.LibString == "Y" {
				ud.Lib = true
			} else {
				ud.Lib = false
			}
		}
	}

	fmt.Printf(GREEN + "License:\n" + RESET)
	fmt.Printf("\t" + RED + "bsd" + RESET + " | " + ORANGE + "gpl" + RESET + " | " + "mpl | mit" + " | " + "[n]one\n")
	fmt.Printf("Selection: " + GREEN)
	ud.License, err = reader.ReadString('\n')
	fmt.Printf(RESET + "\n")
	ErrCheck(err)
	SanitizeStrings(&ud.License)
	if ud.License[0] == 'b' || ud.License[0] == 'B' {
		fmt.Printf("Which " + RED + "BSD" + RESET + " license?\n")
		fmt.Printf(RED + "\t 1 | 2 | 3 | 4\n" + RESET)
		fmt.Printf("Selection: " + RED)
		ud.LicNum, err = reader.ReadString('\n')
		fmt.Printf(RESET + "\n")
		ErrCheck(err)
		SanitizeStrings(&ud.LicNum)
		dummy, _ := strconv.Atoi(ud.LicNum)
		if !isBetween(dummy, 1, 5) {
			fmt.Printf("Defaulting to BSD 3\n")
		}
	} else if ud.License[0] == 'g' || ud.License[0] == 'G' {
		fmt.Printf("Which " + ORANGE + "GPL" + RESET + " license?")
		fmt.Printf(ORANGE + "\n\t2 | 2 or later | 3 | 3 or later\n" + RESET)
		fmt.Printf("Selection: " + ORANGE)
		ud.LicNum, err = reader.ReadString('\n')
		fmt.Printf(RESET + "\n")
		ErrCheck(err)
		SanitizeStrings(&ud.LicNum)
	}

	if ud.Language == "go" || ud.Language == "rus" {
		fmt.Print(GREEN + "Author: " + RESET)
		ud.Author, err = reader.ReadString('\n')
		ErrCheck(err)
		SanitizeStrings(&ud.Author)
	}

	fmt.Print(ORANGE + "Git" + RESET + " or " + GRAY + "HG? " + RESET)
	ud.VCS, err = reader.ReadString('\n')
	ErrCheck(err)
	SanitizeStrings(&ud.VCS)
}

/********************************
 * File and directories
 ********************************/
func DirectoryCreation(patheroo string) {
	_, err := os.Stat(patheroo)
	if true == os.IsNotExist(err) {
		os.Mkdir(patheroo, 0700)
	} else {
		fmt.Printf("Could not create %s already EXISTS\n", strings.ReplaceAll(patheroo, "./", ""))
	}
}

func FileExistence(patheroo string) bool {
	_, err := os.Stat(patheroo)
	if true == os.IsNotExist(err) {
		return true
	} else {
		return false
	}
}

/********************************
 * Build systems
 ********************************/
func BuildSystemC(ud UserData) {
	switch ud.Build {
	case "make":
		TemplateHandling("templates/MakeC.tmpl", "./Makefile", ud)
	case "cmake":
		if ud.Lib == true {
			TemplateHandling("templates/CMakeLibs.tmpl", "./CMakeLists.txt", ud)
		} else {
			TemplateHandling("templates/CMakeLists.tmpl", "./CMakeLists.txt", ud)
		}
	case "meson":
		TemplateHandling("templates/c.meson.tmpl", "./meson.build", ud)
	case "premake":
		TemplateHandling("templates/premake5.tmpl", "./premake5.lua", ud)
	default:
		fmt.Printf("\nDefaulting to CMake!\n")
		TemplateHandling("templates/CMakeLists.tmpl", "./CMakeLists.txt", ud)
	}
}

func BuildSystemCXX(ud UserData) {
	switch ud.Build {
	case "make":
		TemplateHandling("templates/MakeCxx.tmpl", "./Makefile", ud)
	case "cmake":
		if ud.Lib == true {
			TemplateHandling("templates/CMakeLibs.cxx.tmpl", "./CMakeLists.txt", ud)
		} else {
			TemplateHandling("templates/CMakeLists.cxx.tmpl", "./CMakeLists.txt", ud)
		}
	case "meson":
		TemplateHandling("templates/cxx.meson.tmpl", "./meson.build", ud)
	case "premake":
		TemplateHandling("templates/premake5.cxx.tmpl", "./premake5.lua", ud)
	default:
		fmt.Printf("\nDefaulting to CMake!\n")
		TemplateHandling("templates/CMakeLists.cxx.tmpl", "./CMakeLists.txt", ud)
	}
}

func BuildSystemZig(ud UserData) {
	TemplateHandling("templates/buildzig.tmpl", "./build.zig", ud)
}

/********************************
 * templates
 ********************************/

func TemplateHandling(tmplPath, filePath string, ud UserData) {
	if !FileExistence(filePath) {
		fmt.Printf("File %s already exists\n", strings.ReplaceAll(filePath, "./", ""))
		return
	}

	t, err := template.ParseFS(basefiles, tmplPath)
	ErrCheck(err)
	f, err := os.Create(filePath)
	defer f.Close()

	ErrCheck(err)
	err = t.Execute(f, ud)
	ErrCheck(err)
}

/********************************
 * version control system
 ********************************/
func vcsHandling(ud UserData) {
	switch ud.VCS {
	case "hg":
		hgFiles(ud)
	case "git":
		gitFiles(ud)
	case "both":
		fmt.Println("Using HG and Git!")
		gitFiles(ud)
		hgFiles(ud)
	default:
		fmt.Println("Defaulting to git!")
		gitFiles(ud)
	}
}
func gitFiles(ud UserData) {
	TemplateHandling("templates/gitignore.tmpl", "./.gitignore", ud)
	TemplateHandling("templates/gitattributes.tmpl", "./.gitattributes", ud)
}
func hgFiles(ud UserData) {
	TemplateHandling("templates/hgignore.tmpl", "./.hgignore", ud)
}

/********************************
 * Language files
 ********************************/
func cFiles(ud UserData) {
	BuildSystemC(ud)
	TemplateHandling("templates/main.c.tmpl", "./src/main.c", ud)
	TemplateHandling("templates/clang-format.tmpl", "./.clang-format", ud)
}
func cppFiles(ud UserData) {
	BuildSystemCXX(ud)
	TemplateHandling("templates/main.cxx.tmpl", "./src/main.cxx", ud)
	TemplateHandling("templates/clang-format.tmpl", "./.clang-format", ud)
}
func cppEmb(ud UserData) {
	TemplateHandling("templates/MakeCxxEmb.tmpl", "./Makefile", ud)
	TemplateHandling("templates/main.cxx.tmpl", "./src/main.cxx", ud)
	TemplateHandling("templates/clang-format.tmpl", "./.clang-format", ud)
}
func goFiles(ud UserData) {
	TemplateHandling("templates/taskfile.tmpl", "./Taskfile.yml", ud)
	TemplateHandling("templates/main.go.tmpl", "./src/main.go", ud)
}
func rsFiles(ud UserData) {
	TemplateHandling("templates/main.rs.tmpl", "./src/main.rs", ud)
	TemplateHandling("templates/rustfmt.toml.tmpl", "./rustfmt.toml", ud)
	TemplateHandling("templates/cargo.toml.tmpl", "./Cargo.toml", ud)
}
func zigFiles(ud UserData) {
	TemplateHandling("templates/mainzig.tmpl", "./src/main.zig", ud)
	TemplateHandling("templates/zlsjson.tmpl", "./zls.json", ud)
	BuildSystemZig(ud)
}

func FileCreation(ud UserData) {
	switch ud.Language {
	case "c":
		cFiles(ud)
	case "cpp":
		cppFiles(ud)
	case "go":
		goFiles(ud)
	case "rs":
		rsFiles(ud)
	case "zig":
		zigFiles(ud)
	default:
		fmt.Printf("\nDefaulting to C!\n")
		cFiles(ud)
	}
}

func LicenseCreation(ud UserData) {
	switch ud.License {
	case "gpl":
		switch ud.LicNum {
		case "2":
			TemplateHandling("templates/gpl2.tmpl", "./LICENSE", ud)
		case "2 or later":
			TemplateHandling("templates/gpl2orl.tmpl", "./LICENSE", ud)
		case "3":
			TemplateHandling("templates/gpl3.tmpl", "./LICENSE", ud)
		case "3 or later":
			TemplateHandling("templates/gpl3orl.tmpl", "./LICENSE", ud)
		default:
			TemplateHandling("templates/gpl3.tmpl", "./LICENSE", ud)
		}
	case "g":
		switch ud.LicNum {
		case "2":
			TemplateHandling("templates/gpl2.tmpl", "./LICENSE", ud)
		case "2 or later":
			TemplateHandling("templates/gpl2orl.tmpl", "./LICENSE", ud)
		case "3":
			TemplateHandling("templates/gpl3.tmpl", "./LICENSE", ud)
		case "3 or later":
			TemplateHandling("templates/gpl3orl.tmpl", "./LICENSE", ud)
		default:
			TemplateHandling("templates/gpl3.tmpl", "./LICENSE", ud)
		}

	case "bsd":
		switch ud.LicNum {
		case "1":
			TemplateHandling("templates/bsd1.tmpl", "./LICENSE", ud)
		case "2":
			TemplateHandling("templates/bsd1.tmpl", "./LICENSE", ud)
		case "3":
			TemplateHandling("templates/bsd1.tmpl", "./LICENSE", ud)
		case "4":
			TemplateHandling("templates/bsd1.tmpl", "./LICENSE", ud)
		default:
			TemplateHandling("templates/bsd3.tmpl", "./LICENSE", ud)
		}
	case "b":
		switch ud.LicNum {
		case "1":
			TemplateHandling("templates/bsd1.tmpl", "./LICENSE", ud)
		case "2":
			TemplateHandling("templates/bsd1.tmpl", "./LICENSE", ud)
		case "3":
			TemplateHandling("templates/bsd1.tmpl", "./LICENSE", ud)
		case "4":
			TemplateHandling("templates/bsd1.tmpl", "./LICENSE", ud)
		default:
			TemplateHandling("templates/bsd3.tmpl", "./LICENSE", ud)
		}
	case "mit":
		TemplateHandling("templates/mit.tmpl", "./LICENSE", ud)
	case "mpl":
		TemplateHandling("templates/mpl2.tmpl", "./LICENSE", ud)
	case "n":
		fmt.Println("No license!")
	default:
		fmt.Println("Defaulting to GPL3")
		TemplateHandling("templates/gpl3.tmpl", "./LICENSE", ud)
	}
}

func FileHandling(ud UserData) {
	if ud.Language == "c" || ud.Language == "cpp" {
		DirectoryCreation("./include")
	}
	DirectoryCreation("./lib")
	DirectoryCreation("./src")
	DirectoryCreation("./tests")
	DirectoryCreation("./.builds")
	TemplateHandling("templates/manifest.tmpl", "./.builds/manifest.yml", ud)
	FileCreation(ud)
	LicenseCreation(ud)
	vcsHandling(ud)
	TemplateHandling("templates/README.tmpl", "./README.md", ud)
}

func embeddedFileHandling(ud UserData) {
	DirectoryCreation("./include")
	DirectoryCreation("./lib")
	DirectoryCreation("./src")
	DirectoryCreation("./tests")
	DirectoryCreation("./.builds")
	TemplateHandling("./templates/manifest.tmpl", "./.builds/manifest.yml", ud)
	cppEmb(ud)
	LicenseCreation(ud)
	vcsHandling(ud)
	TemplateHandling("templates/READMEORG.tmpl", "./README.org", ud)
}

func argParse() {
	var empty UserData
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "git":
			fmt.Printf("Only creating" + GREEN + " gitignore" + RESET + " and" + GREEN + " gitattributes.\n" + RESET)
			gitFiles(empty)
			os.Exit(0)
		case "hg":
			fmt.Printf("Only creating" + GREEN + " hgignore.\n" + RESET)
			hgFiles(empty)
			os.Exit(0)
		case "both":
			fmt.Printf("Only creating" + GREEN + " gitignore, gitattributes" + RESET + " and" + GREEN + " hgignore.\n" + RESET)
			hgFiles(empty)
			gitFiles(empty)
			os.Exit(0)
		case "org":
			TemplateHandling("templates/READMEORG.tmpl", "./README.org", empty)
			os.Exit(0)
		case "emb":
			var ud UserData
			getData(&ud)
			ud.Year = getTime()
			embeddedFileHandling(ud)
			os.Exit(0)
		default:
			fmt.Printf("If you're going to pass arguments, the only usage is for\nadding gitignore (and gittributes), hgignore or both of them.\n\nDo so by typing 'etr " + GREEN + "hg" + RESET + "|" + GREEN + "git" + RESET + "|" + GREEN + "both'" + RESET + "\n")
			os.Exit(1)
		}
	}

}
func InitETR() {
	argParse()

	var ud UserData
	getData(&ud)
	ud.Year = getTime()

	FileHandling(ud)
}
