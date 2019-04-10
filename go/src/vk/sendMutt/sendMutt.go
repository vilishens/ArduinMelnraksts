package sendMutt

import (
	"fmt"
	"os/exec"

	vutils "vk/utils"
)

func SendMail(addr string, subject string, txtMsg string, htmlMsg string, file string) (res string, err error) {
	// cmdIP := "curl -s checkip.dyndns.org | sed -e 's/.*Current IP Address: //' -e 's/<.*$//'"
	// cmdMsg := "echo "MSG-TXT" | mutt -s "SUBJECT" -F <alternativs .muttrc> <e-mail>
	cmdMsg := fmt.Sprintf("echo \"%s\" | mutt -s \"%s\" -F %s %s", txtMsg, subject, file, addr)
	fmt.Printf("MUTT MAIL cmd: ###%s###", cmdMsg)

	tmpRes := []byte{}
	tmpRes, err = exec.Command("bash", "-c", cmdMsg).Output()
	if err != nil {
		err = vutils.ErrFuncLine(err)
	} else {
		res = string(tmpRes)
	}

	return
}
