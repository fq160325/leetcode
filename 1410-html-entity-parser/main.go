package HTMLEntityParser

import (
	"bytes"
	"fmt"
)

var dict = map[string]byte{
	"&quot;":  '"',
	"&apos;":  '\'',
	"&amp;":   '&',
	"&gt;":    '>',
	"&lt;":    '<',
	"&frasl;": '/',
}

func Main() {
	fmt.Println(entityParser("&&gt;&&&"))
}

func entityParser(text string) string {
	var buff, tempstr bytes.Buffer
	var flag bool
	for i, l := 0, len(text); i < l; i++ {
		if text[i] == '&' {
			if flag {
				buff.WriteString(tempstr.String())
				tempstr.Reset()
			}
			flag = true
			tempstr.WriteByte(text[i])
			continue
		}
		if text[i] == ';' {
			flag = false
			tempstr.WriteByte(text[i])
			fmt.Println(tempstr.String())
			b, ok := dict[tempstr.String()]
			if ok {
				buff.WriteByte(b)
				tempstr.Reset()
				continue
			}
			buff.WriteString(tempstr.String())
			tempstr.Reset()
			continue
		}
		if flag {
			tempstr.WriteByte(text[i])
			continue
		}
		buff.WriteByte(text[i])
	}
	if flag {
		buff.WriteString(tempstr.String())
	}
	return buff.String()
}
