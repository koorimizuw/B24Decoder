package b24decoder

type (
	CodeSetContoller struct {
		Buffer map[string]Code
		//
		SingleShift  string
		GraphicLeft  string
		GraphicRight string
		// escape sequence control
		EscapeSequence      []byte
		EscapeSequenceCount int
		EscapeBufferIndex   string
		EscapeDrcs          bool
	}
)

func (c *CodeSetContoller) invoke(bufferIndex string, area string, lockingShift bool) {
	switch area {
	case GL:
		if lockingShift {
			c.GraphicLeft = bufferIndex
		} else {
			c.SingleShift = bufferIndex
		}
	case GR:
		c.GraphicRight = bufferIndex
	}
	c.EscapeSequenceCount = 0
}

func (c *CodeSetContoller) setEscape(bufferIndex string, drcs bool) {
	if bufferIndex != "" {
		c.EscapeBufferIndex = bufferIndex
	}
	c.EscapeDrcs = drcs
	c.EscapeSequenceCount += 1
}

func (c *CodeSetContoller) degignate(code byte) {
	if c.EscapeDrcs {
		c.Buffer[c.EscapeBufferIndex] = CODE_SET_DRCS[code]
	} else {
		c.Buffer[c.EscapeBufferIndex] = CODE_SET_G[code]
	}
	c.EscapeSequenceCount = 0
}

func (c *CodeSetContoller) getCurrentCode(b byte) Code {
	if b >= 0x21 && b <= 0x7e {
		if c.SingleShift != "" {
			code := c.Buffer[c.SingleShift]
			c.SingleShift = ""
			return code
		}
		return c.Buffer[c.GraphicLeft]
	} else if b >= 0xa1 && b <= 0xfe {
		return c.Buffer[c.GraphicRight]
	}
	return Code{}
}
