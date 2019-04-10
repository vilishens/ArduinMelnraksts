package utils

import (
	"fmt"
	"path/filepath"
	"strings"

	vomni "vk/omnibus"
)

func SplitMsg(msg string) (msgType string, subpath string, name string, point string, txt string, err error) {
	flds := strings.Split(msg, vomni.UDPMessageSeparator)

	errLoc := fmt.Errorf("Incorrect data format of \"%s\"", msg)

	fldCount := len(flds)
	if 0 == fldCount {
		err = ErrFuncLine(errLoc)
		return
	}

	msgType = strings.Trim(strings.ToUpper(flds[0]), " ")

	limit := vomni.MessageTypeLimits[vomni.MessageOmnibus]
	if fldCount < limit {
		err = ErrFuncLine(errLoc)
		return
	}

	subpath, name, txt = msgParts(flds[1:])

	point = filepath.Join(subpath, name)

	if ("" == msgType) || ("" == point) {
		err = ErrFuncLine(errLoc)
		return
	}

	return
}

func TypeOfMsg(msg string) (msgType string, err error) {
	flds := strings.Split(msg, vomni.UDPMessageSeparator)

	fldCount := len(flds)
	if 0 == fldCount {
		err = ErrFuncLine(fmt.Errorf("Incorrect data format of \"%s\"", msg))
		return
	}

	msgType = strings.Trim(strings.ToUpper(flds[0]), " ")

	if "" == msgType {
		err = ErrFuncLine(fmt.Errorf("Incorrect data format of \"%s\"", msg))
		return
	}

	return
}

func msgParts(flds []string) (subpath string, name string, txt string) {

	fldCount := len(flds)

	// p1 ::: p2 ::: ... ::: pN ::: <point> ::: <txt>
	// <p1-Type><SUBPATH-::: p2 ::: ... ::: pN-SUBPATH>:::<point-NAME>::: <txt-MESSAGE>
	for k, v := range flds {
		str := strings.Trim(string(v), " \x00")

		switch k + 1 {
		case fldCount:
			txt = str
		case fldCount - 1:
			name = str
		default:
			subpath = filepath.Join(subpath, str)
		}
	}

	return
}
