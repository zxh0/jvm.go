package cmdline 

import "fmt"

// java [ options ] class [ arguments ]
type Command struct {
    options []*Option
    class   string
    args    []string
}

type Option struct {
    
}

func ParseCommand(cmdLineArgs []string) {
    // todo
    for idx, arg := range cmdLineArgs {
        fmt.Printf("idx: %v arg:%v \n", idx, arg)
    }
}