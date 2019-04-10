package message

import (
	"fmt"
	"path/filepath"
	"strings"
	vomni "vk/omnibus"
)

func splitMsg(msg string) (msgType string, subpath string, name string, txt string) {

	flds := strings.Split(msg, vomni.UDPMessageSeparator)

	fldLen := len(flds)

	// p1 ::: p2 ::: ... ::: pN ::: <point> ::: <txt>
	// <p1-Type><SUBPATH-::: p2 ::: ... ::: pN-SUBPATH>:::<point-NAME>::: <txt-MESSAGE>
	for k, v := range flds {
		str := strings.Trim(string(v), " \x00")

		fmt.Println("INDEX", k, "TXT", str)

		switch k + 1 {
		case fldLen:
			txt = str
		case fldLen - 1:
			name = str
		case 1:
			msgType = strings.ToUpper(str)
		default:
			subpath = filepath.Join(subpath, str)
		}
	}

	return
}
