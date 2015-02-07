package util

func Int8ToUint8(jBytes []int8) (goBytes []byte) {
    goBytes = make([]byte, len(jBytes))
    for i, jByte := range jBytes {
        goBytes[i] = byte(jByte)
    }
    return
}
