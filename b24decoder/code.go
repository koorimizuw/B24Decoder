package b24decoder

type Code struct {
	Code string
	Size int
}

var CODE_SET_G = map[byte]Code{
	0x42: {KANJI, 2},
	0x4A: {ALPHANUMERIC, 1},
	0x30: {HIRAGANA, 1},
	0x31: {KATAKANA, 1},
	0x32: {MOSAIC_A, 1},
	0x33: {MOSAIC_B, 1},
	0x34: {MOSAIC_C, 1},
	0x35: {MOSAIC_D, 1},
	0x36: {PROP_ALPHANUMERIC, 1},
	0x37: {PROP_HIRAGANA, 1},
	0x38: {PROP_KATAKANA, 1},
	0x49: {JIS_X0201_KATAKANA, 1},
	0x39: {JIS_KANJI_PLANE_1, 2},
	0x3A: {JIS_KANJI_PLANE_2, 2},
	0x3B: {ADDITIONAL_SYMBOLS, 2},
}

var CODE_SET_DRCS = map[byte]Code{
	0x40: {"", 2}, // DRCS-0
	0x41: {"", 1}, // DRCS-1
	0x42: {"", 1}, // DRCS-2
	0x43: {"", 1}, // DRCS-3
	0x44: {"", 1}, // DRCS-4
	0x45: {"", 1}, // DRCS-5
	0x46: {"", 1}, // DRCS-6
	0x47: {"", 1}, // DRCS-7
	0x48: {"", 1}, // DRCS-8
	0x49: {"", 1}, // DRCS-9
	0x4A: {"", 1}, // DRCS-10
	0x4B: {"", 1}, // DRCS-11
	0x4C: {"", 1}, // DRCS-12
	0x4D: {"", 1}, // DRCS-13
	0x4E: {"", 1}, // DRCS-14
	0x4F: {"", 1}, // DRCS-15
	0x70: {"", 1}, // MACRO
}

var ARIB_HIRAGANA_MAP = map[byte]byte{
	0x79: 0x3C,
	0x7A: 0x23,
	0x7B: 0x56,
	0x7C: 0x57,
	0x7D: 0x22,
	0x7E: 0x26,
	0x77: 0x35,
	0x78: 0x36,
}

var ARIB_KATAKANA_MAP = map[byte]byte{
	0x79: 0x3C,
	0x7A: 0x23,
	0x7B: 0x56,
	0x7C: 0x57,
	0x7D: 0x22,
	0x7E: 0x26,
	0x77: 0x33,
	0x78: 0x34,
}

var ESC_SEQ_ASCII = []byte{0x1B, 0x28, 0x42}
var ESC_SEQ_ZENKAKU = []byte{0x1B, 0x24, 0x42}
var ESC_SEQ_HANKAKU = []byte{0x1B, 0x28, 0x49}

const (
	G0 = "G0"
	G1 = "G1"
	G2 = "G2"
	G3 = "G3"
	//
	GL = "GL"
	GR = "GR"
	//
	KANJI              = "KANJI"
	ALPHANUMERIC       = "ALPHANUMERIC"
	HIRAGANA           = "HIRAGANA"
	KATAKANA           = "KATAKANA"
	MOSAIC_A           = "MOSAIC_A"
	MOSAIC_B           = "MOSAIC_B"
	MOSAIC_C           = "MOSAIC_C"
	MOSAIC_D           = "MOSAIC_D"
	PROP_ALPHANUMERIC  = "PROP_ALPHANUMERIC"
	PROP_HIRAGANA      = "PROP_HIRAGANA"
	PROP_KATAKANA      = "PROP_KATAKANA"
	JIS_X0201_KATAKANA = "JIS_X0201_KATAKANA"
	JIS_KANJI_PLANE_1  = "JIS_KANJI_PLANE_1"
	JIS_KANJI_PLANE_2  = "JIS_KANJI_PLANE_2"
	ADDITIONAL_SYMBOLS = "ADDITIONAL_SYMBOLS"
	UNSUPPORTED        = "UNSUPPORTED"
)
