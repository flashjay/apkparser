package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/lunny/axmlParser"
	"github.com/shogo82148/androidbinary/apk"
)

type apkFileInfo struct {
	File         string
	Icon         string
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
		"ActivityName => %v\n"+
		"Icon => %v",
		f.File, f.Label, f.VersionName, f.VersionCode, f.PackageName, f.ActivityName, f.Icon)
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
	dat, err := pkg.Icon(nil)
	if err == nil {
		ico, _ := os.Create(f.File[0:len(f.File)-4] + ".ico")
		png.Encode(ico, dat)
		f.Icon = ico.Name()
	} else {
		f.Icon = fmt.Sprintf("## Error {%v}", err)
	}
	fmt.Println(f)
}
