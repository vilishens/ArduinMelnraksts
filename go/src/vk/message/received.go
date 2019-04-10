package message

import (
	"fmt"
	"net"
	"path/filepath"
	vdatafiles "vk/dataFiles"
	vomni "vk/omnibus"
	vpoint "vk/points"
	vutils "vk/utils"
)

func ReceivedMsg(conn *net.UDPConn, addr *net.UDPAddr, msg string) (point string, back string, err error) {
	msgType, subpath, name, txt := splitMsg(msg)
	if "" == msgType || "" == name {
		fmt.Println("TIpe ", msgType, "NAME ", name, "DUB ", subpath, "tzt", txt)

		err = vutils.ErrFuncLine(fmt.Errorf("Incorrect data format of \"%s\"", msg))
		return
	}

	point = name

	if nil == err {
		switch msgType {
		case vomni.MessageTypeCmd:
			back, err = vpoint.Cmd(msg, *addr)
		case vomni.MessageTypeStart:
			back, err = vpoint.StartPoint(msg, *addr)
		case vomni.MessageTypeEvent, vomni.MessageTypeError:
			if msgType != vomni.MessageTypeError {
				//if back, err = Get(msg); nil != err {
				//	break
				//}
			}
			if msgType, _, subpath, point, txt, err1 := vutils.SplitMsg(msg); nil != err1 {
				err = err1
			} else {
				err = vdatafiles.AddRecord(msgType, subpath, point, txt)
			}

		default:
			err = vutils.ErrFuncLine(fmt.Errorf("Unsupported message type \"%s\"", msgType))
		}
	}

	point = filepath.Join(subpath, name)

	return
}
