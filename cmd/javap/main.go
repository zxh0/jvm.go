package main

import (
	"fmt"
	"strings"

	"github.com/zxh0/jvm.go/classfile"
	"github.com/zxh0/jvm.go/classpath"
	"github.com/zxh0/jvm.go/options"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func main() {
	opts, args := options.Parse()
	if opts.HelpFlag || len(args) == 0 {
		printUsage()
	}
	printClassInfo(opts, args[0])
}

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

func printUsage() {
	fmt.Println("usage: javap [-options] class [args...]")
}

func printClassInfo(opts options.Options, className string) {
	cp := classpath.Parse(opts)
	_, classData, err := cp.ReadClass(className)

	if err != nil {
		panic(err)
	}

	cf, err := classfile.Parse(classData)

	fmt.Printf("%s %s",
		accessFlagsForClass(cf.AccessFlags),
		strings.ReplaceAll(cf.GetThisClassName(), "/", "."))

	superClassName := cf.GetSuperClassName()
	interfaceNames := cf.GetInterfaceNames()

	if superClassName != "" {
		fmt.Printf(" extends %s",
			strings.ReplaceAll(superClassName, "/", "."))
	}

	if len(interfaceNames) > 0 {
		fmt.Printf(" implements %s",
			strings.ReplaceAll(strings.Join(interfaceNames, ","), "/", "."))
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

	accessFlags := heap.AccessFlags(af)
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

	accessFlags := heap.AccessFlags(af)
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

	accessFlags := heap.AccessFlags(af)
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
			return strings.ReplaceAll(strings.TrimSuffix(strings.TrimPrefix(descriptor, "[L"), ";"), "/", ".") + "[]"
		} else {
			return strings.ReplaceAll(strings.TrimSuffix(strings.TrimPrefix(descriptor, "L"), ";"), "/", ".")
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
