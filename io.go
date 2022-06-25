package fn

import (
	"bytes"
)

// Wrap bytes to Buffer(readable or writeable)
func Wrap(bys []byte) *bytes.Buffer {
	return bytes.NewBuffer(bys)
}

// WrapStr wrap string to Buffer(readable or writeable)
func WrapStr(s string) *bytes.Buffer {
	return bytes.NewBufferString(s)
}
