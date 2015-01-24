package class

func calcArgCount(descriptor string) (uint) {
    count := 0
    refStarted := false
    for _, char := range descriptor {
        if char == ';' {
            refStarted = false
        } else if !refStarted {
            switch char {
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
            }
        }
    }
    // todo
    return 0
}
