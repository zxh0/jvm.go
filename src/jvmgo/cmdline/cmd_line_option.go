package cmdline

var options = map[string]bool{"-cp": true, "-classpath": true}

type Option struct {
    name    string
    value   string
}
