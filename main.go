package main

import (
	"fmt"
	"github.com/shogo82148/androidbinary/apk"
	"log"
	"os"
)

import (
	"github.com/lunny/axmlParser"
)

type apkFileInfo struct {
	File         string
	Label        string
	VersionName  string
	VersionCode  string
	PackageName  string
	ActivityName string
}

func (f *apkFileInfo) checkError(err error) {
	if err != nil {
		log.Println("Err -> ", err)
	}
}

func (f apkFileInfo) String() string {
	return fmt.Sprintf("File => %v\n"+
		"Label => %v\n"+
		"VersionName => %v\n"+
		"VersionCode => %v\n"+
		"PackageName => %v\n"+
		"ActivityName => %v", f.File, f.Label, f.VersionName, f.VersionCode, f.PackageName, f.ActivityName)
}

func main() {
	f := apkFileInfo{}
	listener := new(axmlParser.AppNameListener)
	f.File = os.Args[1]

	_, err := axmlParser.ParseApk(f.File, listener)
	f.checkError(err)
	if err == nil {
		f.VersionName = listener.VersionName
		f.VersionCode = listener.VersionCode
		f.PackageName = listener.PackageName
		f.ActivityName = listener.ActivityName
	}

	pkg, _ := apk.OpenFile(f.File)
	defer pkg.Close()

	f.Label, _ = pkg.Label(nil)

	fmt.Println(f)
}
