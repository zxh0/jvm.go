package class

func createVTable(class *Class) {
	// todo
}

func countVirtualMethod(class *Class) int {
	count := 0
	for _, m := range class.methods {
		if isVirtualMethod(m) {
			count++
		}
	}
	return count
}

func isVirtualMethod(method *Method) bool {
	return !method.IsStatic() &&
		!method.IsFinal() &&
		!method.IsPrivate() &&
		method.Name() != constructorName
}
