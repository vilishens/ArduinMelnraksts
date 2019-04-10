// utils_file.go
package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	vomni "vk/omnibus"
)

func FileAbsPath(fPath string, file string) (full string) {

	abs := ""

	if !filepath.IsAbs(fPath) {
		abs = vomni.RootPath
	}

	abs = filepath.Join(abs, fPath, file)
	full = filepath.Clean(abs)

	return
}

func FileDir(full string) (err error) {

	permDir := os.FileMode(0700)

	dirpath := filepath.Dir(full)

	if err = os.MkdirAll(dirpath, permDir); nil != err {
		return
	}

	return
}

func FileAppend(fullPath string, strAdd string) (err error) {

	permDir := os.FileMode(0700)
	permFile := os.FileMode(0600)

	dirpath := filepath.Dir(fullPath)

	if err = os.MkdirAll(dirpath, permDir); nil != err {
		return
	}

	f, err := os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, permFile)
	if err != nil {
		return
	}

	defer f.Close()

	_, err = f.WriteString(strAdd)

	return
}

func FileDelete(full string) (err error) {
	return os.Remove(full)
}

/*
func FileOverwrite(f_path string, content []byte) (err error) {

	dir_only := filepath.Dir(f_path)

	err = os.MkdirAll(dir_only, vomni.DIR_PERMISSIONS)
	return ioutil.WriteFile(f_path, content, vomni.FILE_PERMISSIONS)
}
*/

func PathExists(full string) (exists bool, err error) {
	if _, err = os.Stat(full); os.IsNotExist(err) {
		return false, nil
	}

	return true, nil
}

func FileCopy(src string, dst string) (err error) {

	var srcF, dstF *os.File

	if srcF, err = os.Open(src); nil != err {
		return
	}
	defer srcF.Close()

	dir_only := filepath.Dir(dst)
	if err = os.MkdirAll(dir_only, vomni.DIR_PERMISSIONS); nil != err {
		return
	}

	if dstF, err = os.Create(dst); nil != err { // creates if file doesn't exist
		return
	}
	defer dstF.Close()

	if _, err = io.Copy(dstF, srcF); nil != err { // check first var for number of bytes copied
		return
	}

	return dstF.Sync()
}

func OpenFile(path string, fileFlags int, userPerms os.FileMode) (f *os.File, err error) {

	if err = FileDir(path); nil != err {
		return
	}

	f, err = os.OpenFile(path, fileFlags, userPerms)
	if err != nil {
		err = fmt.Errorf("error opening file: %v", err)
		return
	}

	return
}

func FileReadString(path string) (str string, err error) {

	if ok, err := PathExists(path); !ok {
		return "", ErrFuncLine(err)
	}

	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return "", ErrFuncLine(err)
	}

	return string(raw), err
}
