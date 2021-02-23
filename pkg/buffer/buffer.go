package buffer

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

type Buffer struct {
	buffer bytes.Buffer
}

func (b *Buffer) Print(s string, values ...interface{}) {
	if _, err := fmt.Fprintf(&b.buffer, s, values...); err != nil {
		log.Println("err:", err)
	}
}

func (b *Buffer) Println(s string, values ...interface{}) {
	b.Print(s, values...)
	b.Print("\n")
}

func (b *Buffer) Flush(writer io.Writer) {
	if _, err := writer.Write(b.buffer.Bytes()); err != nil {
		log.Println(err)
	}
	b.buffer.Reset()
}
