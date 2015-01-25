package classfile

type DescriptorReader struct {
    d   string
    r   *ClassReader
}
func (self DescriptorReader) readParamStart() {
    b := self.r.readUint8()
    if b != '(' {
        self.causePanic()
    }
}
func (self DescriptorReader) causePanic() {
    panic("BAD descriptor: " + self.d)
}

// descriptor looks like: (IDLjava/lang/Thread;)Ljava/lang/Object;
func calcArgCount(descriptor string) (uint) {
    cr := newClassReader([]byte(descriptor))
    dr := &DescriptorReader{descriptor, cr}
    dr.readParamStart()

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
