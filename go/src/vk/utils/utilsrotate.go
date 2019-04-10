package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

	vomni "vk/omnibus"
)

func SetRotateCfg(file2Rotate string, cfg2UseFile string, cfg2RunFile string, newRotation bool) (err error) {

	usr := new(user.User)

	// the current user data
	usr, err = user.Current()
	if nil != err {
		return ErrFuncLine(err)
	}

	name := usr.Username

	group := ""
	if usrGrp, err := user.LookupGroupId(usr.Gid); nil != err {
		return ErrFuncLine(err)
	} else {
		group = usrGrp.Name
	}

	var format []byte
	// read necessary data file rotation configuration
	if format, err = ioutil.ReadFile(cfg2UseFile); nil != err {
		return ErrFuncLine(err)
	}

	// put required data into configuration template
	str := fmt.Sprintf(string(format), file2Rotate, name, group)

	// delete the existing running configuration if is required
	// to start with the brand new configuration file
	if newRotation {
		has, err := PathExists(cfg2RunFile)
		if nil != err {
			return ErrFuncLine(err)
		}
		if has {
			if err = FileDelete(cfg2RunFile); nil != err {
				return ErrFuncLine(err)
			}
		}
	}

	// append the new configuration part to the running configuration file
	if err = FileAppend(cfg2RunFile, str); nil != err {
		return ErrFuncLine(err)
	}

	return
}

func LogReAssign(f *os.File, path string) (fNew *os.File, err error) { // (fNew *os.File, err error) {

	var perms os.FileMode
	flags := vomni.LogFileFlags

	if nil != f {
		var stat os.FileInfo

		stat, err = f.Stat()
		if nil != err {
			return nil, ErrFuncLine(fmt.Errorf("Could not get stat of the file %s", path))
		}

		perms = stat.Mode()

		if err = f.Close(); nil != err {
			return nil, ErrFuncLine(fmt.Errorf("Could not close the file %s", path))
		}
	} else {
		perms = vomni.LogUserPerms
	}

	return OpenFile(path, flags, perms)
}

func RunRotate(cfg2RunFile string) (err error) {
	//	find the local status file
	dirpath := filepath.Dir(cfg2RunFile)
	statusF := filepath.Join(dirpath, vomni.LogStatusFile)

	if has, _ := PathExists(statusF); has {
		FileDelete(statusF)
	}

	// logrotate <conf.file> -s <localstatus.file>
	cmd := exec.Command("logrotate", cfg2RunFile, "-s", statusF)

	if err = cmd.Run(); nil != err {
		return ErrFuncLine(err)
	}

	return
}

func PrepareRunRotate(file2Rotate string, cfg2UseFile string, cfg2RunFile string, newRotation bool) (err error) {

	if err = SetRotateCfg(file2Rotate, cfg2UseFile, cfg2RunFile, newRotation); nil != err {
		return ErrFuncLine(err)
	}

	if err = RunRotate(cfg2RunFile); nil != err {
		return ErrFuncLine(err)
	}

	return
}
