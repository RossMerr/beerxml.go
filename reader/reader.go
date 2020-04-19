package reader

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"

	"golang.org/x/text/encoding/charmap"
)

func MakeCharsetReader(charset string, input io.Reader) (io.Reader, error) {
	if charset == "ISO-8859-1" {
		return charmap.Windows1252.NewDecoder().Reader(input), nil
	}

	return nil, fmt.Errorf("Unknown charset: %s", charset)
}

type Trimmer struct {
	*xml.Decoder
}

func (tr Trimmer) Token() (xml.Token, error) {
	t, err := tr.Decoder.Token()
	if cd, ok := t.(xml.CharData); ok {
		t = xml.CharData(bytes.TrimSpace(cd))
	}
	return t, err
}
