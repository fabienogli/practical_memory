package practicalmemory

import (
	"bytes"
	"fmt"
	"io"
)

var data = []struct {
	input  []byte
	output []byte
}{
	{[]byte("abc"), []byte("abc")},
	{[]byte("elvis"), []byte("Elvis")},
	{[]byte("aElvis"), []byte("aElvis")},
	{[]byte("abcelvis"), []byte("abcElvis")},
	{[]byte("aelvis"), []byte("aElvis")},
	{[]byte("eelvis"), []byte("eElvis")},
	{[]byte("aabbeeeeeeeeeelvis"), []byte("aabbeeeeeeeeeElvis")},
	{[]byte("e l v i s"), []byte("e l v i s")},
	{[]byte("elvielvis"), []byte("elviElvis")},
}

func assembleOutputStream() []byte {
	var out []byte
	for _, d := range data {
		out = append(out, d.output...)
	}
	return out
}

func assembleInputStream() []byte {
	var out []byte
	for _, d := range data {
		out = append(out, d.input...)
	}
	return out
}

func main() {
	var output bytes.Buffer
	in := assembleInputStream()
	out := assembleOutputStream()

	find := []byte("elvis")
	repl := []byte("Elvis")

	fmt.Println("=================================================\nRunning Alg One")
	output.Reset()
	algOne(in, find, repl, &output)
	matched := bytes.Compare(out, output.Bytes())
	fmt.Printf("Matched: %v\nInp: [%s]\nExp: [%s]\nGot: [%s]\n", matched == 0, in, out, output.Bytes())
	fmt.Println("=================================================\nRunning Alg One")
	output.Reset()
	algTwo(in, find, repl, &output)
	matched = bytes.Compare(out, output.Bytes())
	fmt.Printf("Matched: %v\nInp: [%s]\nExp: [%s]\nGot: [%s]\n", matched == 0, in, out, output.Bytes())

}

func algOne(data []byte, find []byte, repl []byte, output *bytes.Buffer) {
	input := bytes.NewBuffer(data)

	size := len(find)

	// Declare the buffers we need to process the stream
	buf := make([]byte, size)
	end := size - 1

	// Read in an initial number of bytes we need to get started
	if n, err := io.ReadFull(input, buf[:end]); err != nil {
		output.Write(buf[:n])
		return
	}

	for {
		//Read in one byte from the input stream
		if _, err := io.ReadFull(input, buf[end:]); err != nil {
			// Flush the reset of the bytes we have
			output.Write(buf[:end])
			return
		}

		// if we have a match , replace the bytes
		if bytes.Equal(buf, find) {
			output.Write(repl)

			//Read a new initial number of bytes
			if n, err := io.ReadFull(input, buf[:end]); err != nil {
				output.Write(buf[:n])
				return
			}
			continue
		}
		// Write the front byte since it has been compared
		output.WriteByte(buf[0])
		// Slice tha front byte out
		copy(buf, buf[1:])
	}
}

func algTwo(data []byte, find []byte, repl []byte, output *bytes.Buffer) {
	// Use the bytes Reader to provide a stream to process
	input := bytes.NewReader(data)

	// The number of bytes we are looking for
	size := len(find)

	// Create an index variable to match bytes
	idx := 0

	for {

		// Read a single byte from our input
		b, err := input.ReadByte()
		if err != nil {
			break
		}

		// Does this byte match the byte at this offset?
		if b == find[idx] {

			// it matches so increment the index position
			idx++

			//If every byte has bee, matched, write
			// out the replacement
			if idx == size {
				output.Write(repl)
				idx = 0
			}
			continue

		}

		// Did we have any sort of match of any given byte ?

		if idx != 0 {
			// Write what we(v matched up to this point
			output.Write(find[:idx])

			// Unread the unmatched byte so it can be processed again
			input.UnreadByte()

			// Reset the offset to start matching from the beginning
			idx = 0

			continue
		}

		// There was no previous match. Write byte and reset
		output.WriteByte(b)
		idx = 0
	}
}
