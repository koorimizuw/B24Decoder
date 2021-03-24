package b24decoder

func Decode(b []byte) string {
	aribArray := NewAribArray(b)

	skip := false
	for i, v := range aribArray.ByteArray {
		if skip {
			skip = false
			continue
		}

		if aribArray.Control.EscapeSequenceCount > 0 {
			aribArray.escape(v)
		} else {
			if (v >= 0x21 && v <= 0x7E) || (v >= 0xA1 && v <= 0xFE) {
				skip = aribArray.convert(i)
			} else if v == 0x20 || v == 0xa0 || v == 0x09 {
				aribArray.appendByte(ESC_SEQ_ASCII, 0x20)
			} else if v == 0x0d || v == 0x0a {
				aribArray.appendByte(ESC_SEQ_ASCII, 0x0a)
			} else {
				aribArray.control(v)
			}
		}
	}

	aribArray.toString()
	return aribArray.String
}
