package cmdline

const (
    CLASSPATH = "-classpath"
)

// optionName -> hasValue
var optionMap = map[string]bool{CLASSPATH: true}
// shortName -> fullName
var nameMap = map[string]string{"-cp": CLASSPATH}

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
