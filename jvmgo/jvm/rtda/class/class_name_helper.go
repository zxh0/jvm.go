package class

func DotToSlash(name string) string {
	return util.ReplaceAll(name, ".", "/")
}
func SlashToDot(name string) string {
	return util.ReplaceAll(name, "/", ".")
}
