package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/zxh0/jvm.go/classfile"
	"github.com/zxh0/jvm.go/classpath"
	"github.com/zxh0/jvm.go/vm"
	"github.com/zxh0/jvm.go/vmutils"
)

var (
	versionFlag bool
	helpFlag    bool
)

var (
	primitiveMap = map[string]string{
		"B":  "byte",
		"C":  "char",
		"D":  "double",
		"F":  "float",
		"I":  "int",
		"J":  "long",
		"S":  "short",
		"V":  "void",
		"Z":  "boolean",
		"[Z": "boolean[]",
		"[B": "byte[]",
		"[C": "char[]",
		"[S": "short[]",
		"[I": "int[]",
		"[L": "long[]",
		"[F": "float[]",
		"[D": "double[]",
	}
)

func main() {
	opts, args := parseOptions()
	if helpFlag || len(args) == 0 {
		printUsage()
	}
	printClassInfo(opts, args[0])
}

func parseOptions() (*vm.Options, []string) {
	options := &vm.Options{}
	flag.StringVar(&options.ClassPath, "classpath", "", "Specifies a list of directories, JAR files, and ZIP archives to search for class files.")
	flag.StringVar(&options.ClassPath, "cp", "", "Specifies a list of directories, JAR files, and ZIP archives to search for class files.")
	flag.BoolVar(&helpFlag, "help", false, "Displays usage information and exit.")
	flag.BoolVar(&helpFlag, "h", false, "Displays usage information and exit.")
	flag.BoolVar(&helpFlag, "?", false, "Displays usage information and exit.")
	flag.BoolVar(&versionFlag, "version", false, "Displays version information and exit.")
	flag.Parse()
	return options, flag.Args()
}

func printUsage() {
	fmt.Println("usage: javap [-options] class [args...]")
}

func printClassInfo(opts *vm.Options, className string) {
	cp := classpath.Parse(opts)
	_, classData := cp.ReadClass(className)
	if classData == nil {
		panic("class not found: " + className)
	}

	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s %s",
		accessFlagsForClass(cf.AccessFlags),
		vmutils.SlashToDot(cf.GetThisClassName()))

	superClassName := cf.GetSuperClassName()
	interfaceNames := cf.GetInterfaceNames()

	if superClassName != "" {
		fmt.Printf(" extends %s",
			vmutils.SlashToDot(superClassName))
	}

	if len(interfaceNames) > 0 {
		fmt.Printf(" implements %s",
			vmutils.SlashToDot(strings.Join(interfaceNames, ",")))
	}

	fmt.Println(" {")

	for _, f := range cf.Fields {
		fmt.Printf(" %s %s %s\n",
			accessFlagsForField(f.AccessFlags),
			descriptorToRealName(cf.GetUTF8(f.DescriptorIndex)),
			cf.GetUTF8(f.NameIndex))
	}

	for _, m := range cf.Methods {
		returnType := strings.Split(cf.GetUTF8(m.DescriptorIndex), ")")[1]
		inputTypes := strings.Split(cf.GetUTF8(m.DescriptorIndex), ")")[0][1:]
		fmt.Printf(" %s %s %s(%s)\n",
			accessFlagsForMethod(m.AccessFlags),
			descriptorToRealName(returnType),
			cf.GetUTF8(m.NameIndex),
			inputTypesToRealNames(inputTypes))
	}
	fmt.Println("}")
}

func accessFlagsForClass(af uint16) string {
	var result []string

	accessFlags := classfile.AccessFlags(af)
	if accessFlags.IsPublic() {
		result = append(result, "public")
	}
	if accessFlags.IsFinal() {
		result = append(result, "final")
	}
	if accessFlags.IsInterface() {
		result = append(result, "interface")
	}
	if accessFlags.IsAbstract() {
		result = append(result, "abstract")
	}
	if accessFlags.IsEnum() {
		result = append(result, "enum")
	} else if !accessFlags.IsInterface() {
		result = append(result, "class")
	}
	return strings.Join(result, " ")
}

func accessFlagsForField(af uint16) string {
	var result []string

	accessFlags := classfile.AccessFlags(af)
	if accessFlags.IsPublic() {
		result = append(result, "public")
	}
	if accessFlags.IsPrivate() {
		result = append(result, "private")
	}
	if accessFlags.IsProtected() {
		result = append(result, "protected")
	}
	if accessFlags.IsStatic() {
		result = append(result, "static")
	}
	if accessFlags.IsFinal() {
		result = append(result, "final")
	}
	if accessFlags.IsVolatile() {
		result = append(result, "volatile")
	}
	if accessFlags.IsTransient() {
		result = append(result, "transient")
	}
	if accessFlags.IsEnum() {
		result = append(result, "enum")
	}
	return strings.Join(result, " ")
}

func accessFlagsForMethod(af uint16) string {
	var result []string

	accessFlags := classfile.AccessFlags(af)
	if accessFlags.IsPublic() {
		result = append(result, "public")
	}
	if accessFlags.IsPrivate() {
		result = append(result, "private")
	}
	if accessFlags.IsProtected() {
		result = append(result, "protected")
	}
	if accessFlags.IsStatic() {
		result = append(result, "static")
	}
	if accessFlags.IsFinal() {
		result = append(result, "final")
	}
	if accessFlags.IsSynchronized() {
		result = append(result, "synchronized")
	}
	if accessFlags.IsAbstract() {
		result = append(result, "abstract")
	}
	return strings.Join(result, " ")
}

func descriptorToRealName(descriptor string) string {
	if primitiveMap[descriptor] != "" {
		return primitiveMap[descriptor]
	} else {
		if strings.Contains(descriptor, "[L") {
			return vmutils.SlashToDot(strings.TrimSuffix(strings.TrimPrefix(descriptor, "[L"), ";")) + "[]"
		} else {
			return vmutils.SlashToDot(strings.TrimSuffix(strings.TrimPrefix(descriptor, "L"), ";"))
		}
	}
}

func inputTypesToRealNames(inputTypes string) string {
	var result []string
	for _, inputType := range strings.Split(inputTypes, ";") {
		if inputType != "" {
			if strings.Contains(inputType, "/") {
				result = append(result, descriptorToRealName(inputType))
			} else {
				// need further split for case like II[CII
				for _, inputType := range furtherSplit(inputType) {
					result = append(result, descriptorToRealName(inputType))
				}
			}
		}
	}
	return strings.Join(result, ",")
}

func furtherSplit(inputTypes string) []string {
	var result []string
	for i := 0; i < len(inputTypes); i++ {
		if inputTypes[i] == '[' {
			result = append(result, inputTypes[i:i+2])
			i++
		} else {
			result = append(result, inputTypes[i:i+1])
		}
	}
	return result
}
