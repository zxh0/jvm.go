package cmdline

// optionName -> hasValue
var optionMap = map[string]bool{"-classpath": true}
// shortName -> fullName
var nameMap = map[string]string{"-cp": "-classpath"}

type Option struct {
    name    string
    value   string
}

func parseOption(args *CmdLineArgs) (*Option) {
    if !args.empty() {
        arg := fullName(args.first())
        hasVal, isOption := optionMap[arg]
        if isOption {
            args.removeFirst()
            option := &Option{name: arg}
            if hasVal {
                option.value = args.removeFirst()
            }
            return option
        }
    }
    return nil
}

func fullName(name string) (string) {
    fullName, isShortName := nameMap[name]
    if isShortName {
        return fullName
    } else {
        return name
    }
}
