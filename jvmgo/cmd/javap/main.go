package main

import (
	"fmt"
	"github.com/zxh0/jvm.go/jvmgo/classfile"
	"github.com/zxh0/jvm.go/jvmgo/classpath"
	"github.com/zxh0/jvm.go/jvmgo/cmd"
	"github.com/zxh0/jvm.go/jvmgo/options"
	"os"
)

func main() {
	cmd, err := cmd.ParseCommand(os.Args)
	if err != nil {
		fmt.Println(err)
		printUsage()
	}
	printClassInfo(cmd)
}

func printUsage() {
	fmt.Println("usage: javap [-options] class [args...]")
}

func printClassInfo(cmd cmd.Command) {
	options.InitOptions(cmd.Options.VerboseClass, cmd.Options.Xss, cmd.Options.XuseJavaHome)

	cp := classpath.Parse(cmd.Options.Classpath)
	_, classData, err := cp.ReadClass(cmd.Class)

	if err != nil {
		panic(err)
	}

	cf, err := classfile.Parse(classData)

	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags)
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields))
	for _, f := range cf.Fields {
		fmt.Printf("  %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods))
	for _, m := range cf.Methods {
		fmt.Printf("  %s\n", m.Name())
	}

}