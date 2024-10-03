//
// Copyright (c) 2021-2024 Markku Rossi
//
// All rights reserved.
//

package hexdump

import (
	"bytes"
	"encoding/hex"
	"testing"
)

var hexDumpTests = []string{
	"",
	"Hello, world!",
	`This is a longer input
with multiple lines
and still more lines
to test!`,
}

func TestHexDumpParse(t *testing.T) {
	for _, input := range hexDumpTests {
		ibuf := []byte(input)

		obuf, err := Parse([]byte(hex.Dump(ibuf)))
		if err != nil {
			t.Errorf("hexdump.Parse failed: %s", err)
			continue
		}
		if bytes.Compare(ibuf, obuf) != 0 {
			t.Errorf("Got invalid result: input:\n%soutput:\n%s",
				hex.Dump(ibuf), hex.Dump(obuf))
		}
	}
}

var hexdumpTests = []struct {
	input  string
	output []byte
}{
	{
		input: `0000000 6261 6463 6665 6867 6a69 000a
000000b
`,
		output: []byte{
			0x62, 0x61, 0x64, 0x63, 0x66, 0x65, 0x68, 0x67,
			0x6a, 0x69, 0x00, 0x0a,
		},
	},
}

func TestHexdumpParse(t *testing.T) {
	for _, test := range hexdumpTests {
		ibuf := []byte(test.input)
		obuf, err := Parse(ibuf)
		if err != nil {
			t.Errorf("hexdump.Parse failed: %s", err)
			continue
		}
		if bytes.Compare(obuf, test.output) != 0 {
			t.Errorf("Got invalid result:\n%sexpected:\n%s",
				hex.Dump(obuf), hex.Dump(test.output))
		}
	}
}
