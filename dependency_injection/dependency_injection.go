package dependencyinjection

import (
	"io"
)

func Greet(w io.Writer, s string) {
	w.Write(([]byte)("Hello, " + s))
}
