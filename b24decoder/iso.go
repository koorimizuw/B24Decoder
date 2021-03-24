package b24decoder

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func IsoDec(b []byte) ([]byte, error) {
	r := bytes.NewBuffer(b)
	decoded, err := ioutil.ReadAll(transform.NewReader(r, japanese.ISO2022JP.NewDecoder()))
	if err != nil {
		return nil, err
	}
	return decoded, nil
}
