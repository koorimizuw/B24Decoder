package b24decoder

import "reflect"

type (
	AribArray struct {
		ByteArray []byte
		JisArray  []byte
		String    string
		// control
		Control CodeSetContoller
	}
)

func NewAribArray(b []byte) AribArray {
	return AribArray{
		ByteArray: b,
		JisArray:  []byte{},
		String:    "",
		Control: CodeSetContoller{
			Buffer: map[string]Code{
				G0: CODE_SET_G[0x42],
				G1: CODE_SET_G[0x4a],
				G2: CODE_SET_G[0x30],
				G3: CODE_SET_G[0x31],
			},
			SingleShift:         "",
			GraphicLeft:         G0,
			GraphicRight:        G2,
			EscapeSequence:      []byte{},
			EscapeSequenceCount: 0,
			EscapeBufferIndex:   G0,
			EscapeDrcs:          false,
		},
	}
}

func (arr *AribArray) toString() {
	if len(arr.JisArray) == 0 {
		return
	}

	decoded, err := IsoDec(arr.JisArray)
	if err != nil {
		panic(err)
	}

	arr.String += string(decoded)
}

func (arr *AribArray) appendByte(escapeSequence []byte, char ...byte) {
	if reflect.DeepEqual(arr.Control.EscapeSequence, escapeSequence) {
		arr.JisArray = append(arr.JisArray, escapeSequence...)
		arr.Control.EscapeSequence = escapeSequence
	}
	arr.JisArray = append(arr.JisArray, char...)
}

// return skip
func (arr *AribArray) convert(idx int) bool {
	skip := false

	b1 := arr.ByteArray[idx]
	b2 := byte(0x0)
	code := arr.Control.getCurrentCode(b1)

	if code.Size == 2 {
		skip = true
		if idx+1 < len(arr.ByteArray) {
			b2 = arr.ByteArray[idx+1]
		}
	}
	if b1 >= 0xa1 && b1 <= 0xfe {
		b1 = b1 & 0x7f
		b2 = b2 & 0x7f
	}

	switch code.Code {
	// Kanji
	case KANJI, JIS_KANJI_PLANE_1, JIS_KANJI_PLANE_2:
		arr.JisArray = append(arr.JisArray, ESC_SEQ_ZENKAKU...)
		arr.JisArray = append(arr.JisArray, b1, b2)
	// Alphabet
	case ALPHANUMERIC, PROP_ALPHANUMERIC:
		arr.JisArray = append(arr.JisArray, ESC_SEQ_ASCII...)
		arr.JisArray = append(arr.JisArray, b1)
	// Hiragana
	case HIRAGANA, PROP_HIRAGANA:
		arr.JisArray = append(arr.JisArray, ESC_SEQ_ZENKAKU...)
		if b1 >= 0x77 {
			arr.JisArray = append(arr.JisArray, 0x21, ARIB_HIRAGANA_MAP[b1])
		} else {
			arr.JisArray = append(arr.JisArray, 0x24, b1)
		}
	// Katakana
	case KATAKANA, PROP_KATAKANA:
		arr.JisArray = append(arr.JisArray, ESC_SEQ_ZENKAKU...)
		if b1 >= 0x77 {
			arr.JisArray = append(arr.JisArray, 0x21, ARIB_KATAKANA_MAP[b1])
		} else {
			arr.JisArray = append(arr.JisArray, 0x25, b1)
		}
	// Hankaku katakana
	case JIS_X0201_KATAKANA:
		arr.JisArray = append(arr.JisArray, ESC_SEQ_HANKAKU...)
		arr.JisArray = append(arr.JisArray, b1)
	// Arib gaiji
	case ADDITIONAL_SYMBOLS:
		arr.toString()
		arr.JisArray = []byte{}
		arr.String += GAIJI_MAP[(int(b1)<<8)+int(b2)]
	}

	return skip
}

func (arr *AribArray) control(b byte) {
	switch b {
	case 0x0f:
		arr.Control.invoke(G0, GL, true) // LS0
	case 0x0e:
		arr.Control.invoke(G1, GL, true) // LS1
	case 0x019:
		arr.Control.invoke(G2, GL, false) // SS2
	case 0x1d:
		arr.Control.invoke(G3, GL, false) // SS3
	case 0x1b:
		arr.Control.EscapeSequenceCount = 1
	}
}

func (arr *AribArray) escape(b byte) {
	switch arr.Control.EscapeSequenceCount {
	case 1:
		switch b {
		case 0x6e:
			arr.Control.invoke(G2, GL, true) // LS2
		case 0x6f:
			arr.Control.invoke(G3, GL, true) // LS3
		case 0x7e:
			arr.Control.invoke(G1, GR, true) // LS1R
		case 0x7d:
			arr.Control.invoke(G2, GR, true) // LS2R
		case 0x7c:
			arr.Control.invoke(G3, GR, true) // LS3R
		case 0x24, 0x28:
			arr.Control.setEscape(G0, false)
		case 0x29:
			arr.Control.setEscape(G1, false)
		case 0x2a:
			arr.Control.setEscape(G2, false)
		case 0x2b:
			arr.Control.setEscape(G3, false)
		}
	case 2:
		switch b {
		case 0x20:
			arr.Control.setEscape("", true)
		case 0x28:
			arr.Control.setEscape(G0, false)
		case 0x29:
			arr.Control.setEscape(G1, false)
		case 0x2a:
			arr.Control.setEscape(G2, false)
		case 0x2b:
			arr.Control.setEscape(G3, false)
		default:
			arr.Control.degignate(b)
		}
	case 3:
		switch b {
		case 0x20:
			arr.Control.setEscape("", true)
		default:
			arr.Control.degignate(b)
		}
	case 4:
		arr.Control.degignate(b)
	}
}
