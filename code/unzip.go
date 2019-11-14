package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func Unzip(src_zip string,dest string) error {
	unzipFile,err:=zip.OpenReader(src_zip)
	if err!=nil{return err}
	defer unzipFile.Close()

	os.MkdirAll(dest,0755)
	for _,f:=range unzipFile.File{
		rc,err:=f.Open()
		if err!=nil{return err}
		defer rc.Close()

		path:=filepath.Join(dest,f.Name)
		if f.FileInfo().IsDir(){os.MkdirAll(path,f.Mode())
		} else {
			f,err:=os.OpenFile(path,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,f.Mode())
			if err!=nil{return err}
			defer f.Close()

			_,err=io.Copy(f,rc)
			if err!=nil {
				if err!=io.EOF{return err}
			}
		}
	}
	return nil
}
