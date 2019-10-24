package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/docopt/docopt-go"
	"github.com/zxh0/jvm.go/classfile"
	"github.com/zxh0/jvm.go/module"
	"github.com/zxh0/jvm.go/vmutils"
)

const (
	version = "jmod.go 0.0.1"
	usage   = `jmod.

Usage:
  jmod list <file>
  jmod describe <file>
  jmod -h | --help
  jmod --version

Commands:
  list      Prints the names of all the entries.
  describe  Prints the module details.

Options:
  -h --help     Print this help message
  --version     Print version information`
)

func main() {
	args := os.Args[1:]
	if opts, err := docopt.ParseArgs(usage, args, version); err != nil {
		fmt.Println(usage)
	} else if opts["list"].(bool) {
		list(opts["<file>"].(string))
	} else if opts["describe"].(bool) {
		describe(opts["<file>"].(string))
	}
}

func list(filename string) {
	f, err := vmutils.OpenJModFile(filename)
	if err != nil {
		panic(err)
	}
	for _, f := range f.ListFiles() {
		fmt.Println(f)
	}
}

func describe(filename string) {
	jmodFile, err := vmutils.OpenJModFile(filename)
	if err != nil {
		panic(err)
	}

	classData, err := jmodFile.ReadFile("classes/module-info.class")
	if err != nil {
		panic(err)
	}

	modInfo := module.ParseModuleInfo(classData)
	describeModule(jmodFile, modInfo)
}

func describeModule(jmodFile *vmutils.JModFile, modInfo *module.Info) {
	fmt.Printf("%s@%s\n", modInfo.Name, modInfo.Version)
	// unqualified exports (sorted by package)
	for _, export := range modInfo.Exports {
		if len(export.To) == 0 {
			fmt.Printf("exports %s\n", vmutils.SlashToDot(export.Package))
			// TODO: flags
		}
	}
	// dependencies
	for _, require := range modInfo.Requires {
		fmt.Printf("requires %s", require.Name)
		flags := classfile.AccessFlags(require.Flags)
		if flags.IsMandated() {
			fmt.Print(" mandated")
		} else if flags.IsTransitive() {
			fmt.Print(" transitive")
		}
		fmt.Println()
	}
	// service uses
	for _, use := range modInfo.Uses {
		fmt.Printf("uses %s\n", vmutils.SlashToDot(use))
	}
	// service provides
	for _, provide := range modInfo.Provides {
		fmt.Printf("provides %s with", vmutils.SlashToDot(provide.Service))
		for _, impl := range provide.Impls {
			fmt.Printf(" %s", vmutils.SlashToDot(impl))
		}
		fmt.Println()
	}
	// qualified exports
	for _, export := range modInfo.Exports {
		if len(export.To) > 0 {
			fmt.Printf("qualified exports %s to %s\n",
				vmutils.SlashToDot(export.Package), strings.Join(export.To, " "))
		}
	}
	// open packages
	for _, open := range modInfo.Opens {
		if len(open.To) == 0 {
			fmt.Printf("opens %s\n", vmutils.SlashToDot(open.Package))
			// TODO: flags
		}
	}
	for _, open := range modInfo.Opens {
		if len(open.To) > 0 {
			fmt.Printf("qualified opens %s to  %s\n",
				vmutils.SlashToDot(open.Package), strings.Join(open.To, " "))
			// TODO: flags
		}
	}
	// non-exported/non-open packages
	for _, pkg := range getPrivatePackages(jmodFile, modInfo) {
		fmt.Printf("contains %s\n", vmutils.SlashToDot(pkg))
	}
	// TODO: platform & hashes
}

func getPrivatePackages(jmodFile *vmutils.JModFile, modInfo *module.Info) []string {
	nonPrivatePkgMap := map[string]bool{}
	for _, export := range modInfo.Exports {
		nonPrivatePkgMap[export.Package] = true
	}
	for _, open := range modInfo.Opens {
		nonPrivatePkgMap[open.Package] = true
	}

	privatePkgMap := map[string]bool{}
	for _, f := range jmodFile.ListFiles() {
		pkg := getPackage(f)
		if pkg != "" && !nonPrivatePkgMap[pkg] {
			privatePkgMap[pkg] = true
		}
	}

	privatePackages := make([]string, 0, len(privatePkgMap))
	for pkg, _ := range privatePkgMap {
		privatePackages = append(privatePackages, pkg)
	}
	sort.Strings(privatePackages)
	return privatePackages
}

func getPackage(pathInJModFile string) string {
	if strings.Index(pathInJModFile, ".class") < 0 {
		return ""
	}

	slashIdx := strings.IndexByte(pathInJModFile, '/')
	lastSlashIdx := strings.LastIndexByte(pathInJModFile, '/')
	if lastSlashIdx > slashIdx {
		return pathInJModFile[slashIdx+1 : lastSlashIdx]
	} else {
		return ""
	}
}
