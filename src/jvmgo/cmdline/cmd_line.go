package cmdline 

// java [ options ] class [ arguments ]
type Command struct {
    options []*Option
    class   string
    args    []string
}

type Option struct {
    
}

func ParseCommand() {
    // todo
}