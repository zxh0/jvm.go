package classfile

// descriptor looks like: (IDLjava/lang/Thread;)Ljava/lang/Object;
func calcArgCount(descriptor string) (uint) {
    count := 0
    refStarted := false
    for pos, char := range descriptor {
        if pos == 0 {
            if char != '(' {
                panic("BAD method descriptor: " + descriptor)
            }
        } else if refStarted {
            if char == ';' {
                refStarted = false
            }
        } else {
            switch char {
            case 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z':
                count++
            case 'L', '[':
                count++
                refStarted = true
            case ')': return uint(count) // break not works
            default: panic("BAD method descriptor: " + descriptor)
            }
        }
    }
    return uint(count)
}
