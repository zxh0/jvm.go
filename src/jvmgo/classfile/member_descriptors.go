package classfile

func calcArgCount(descriptor string) (uint) {
    count := 0
    refStarted := false
    for pos, char := range descriptor {
        if pos > 0 {
            if refStarted {
                if char == ';' {
                    refStarted = false
                }
            } else {
                switch char {
                case ')': break
                case 'B': fallthrough // byte
                case 'C': fallthrough // char
                case 'D': fallthrough // double
                case 'F': fallthrough // float
                case 'I': fallthrough // int
                case 'J': fallthrough // long
                case 'S': fallthrough // short
                case 'z': count++     // boolean
                case 'L': fallthrough // reference
                case '[':             // array
                    count++
                    refStarted = true
                default: panic("calcArgCount()") // todo
                }
            }
        }
    }
    return uint(count - 1)
}
