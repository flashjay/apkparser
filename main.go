package main

import (
	"fmt"
	"github.com/lunny/axmlParser"
	"log"
	"os"
)

type apkFileInfo struct {
	FilePath string
	Version  string
	Package  string
	Activity string
}

func (f *apkFileInfo) checkError(err error) {
	if err != nil {
		log.Println("Err -> ", err)
	}
}

// 获取 apk 版本
func (f *apkFileInfo) GetApkVersion() (err error) {
	listener := new(axmlParser.AppNameListener)
	_, err = axmlParser.ParseApk(f.FilePath, listener)
	f.checkError(err)
	if err == nil {
		f.Version = listener.VersionName
		f.Package = listener.PackageName
		f.Activity = listener.ActivityName
	}
	return err
}

func main() {
	f := apkFileInfo{}
	listener := new(axmlParser.AppNameListener)
	f.FilePath = os.Args[1]
	_, err := axmlParser.ParseApk(f.FilePath, listener)
	f.checkError(err)
	if err == nil {
		f.Version = listener.VersionName
		f.Package = listener.PackageName
		f.Activity = listener.ActivityName
	}
	fmt.Println(f)
}
