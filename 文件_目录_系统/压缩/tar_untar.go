//https://www.cnblogs.com/golove/p/3454630.html
package main
import(
    "fmt"
    "archive/tar"
    "errors"
    "io"
    "io/ioutil"
    "os"
    "path"
)
func main(){
    TarFile:="test.tar"
    src:="/code/script"
    dstDir:="/mnt"

    if err:=Tar(src,TarFile,false);err!=nil{
        fmt.Println(err)
    }

    if err:=UnTar(TarFile,dstDir);err!=nil{
        fmt.Println(err)
    }
}
func Tar(src, dstTar string,failIfExist bool) (err error){
    src=path.Clean(src)
    fmt.Println(src)
    if !Exists(src){
        return errors.New("the file path is not exist:"+src)
    }
    if FileExists(dstTar){
        if failIfExist{
            return errors.New("the file is exist:"+dstTar)
        } else {
            if er:=os.Remove(dstTar);er!=nil{
                return
            }
        }
    }
    fw,er:=os.Create(dstTar)
    if er!=nil{
        return er
    }
    defer fw.Close()

    tw:=tar.NewWriter(fw)
    defer func(){
        if er:=tw.Close();er!=nil{ err=er }
    }()

    fi,er:=os.Stat(src)
    if er!=nil{ return er }

    srcBase,srcRelative:=path.Split(path.Clean(src))

    if fi.IsDir(){
        tarDir(srcBase,srcRelative,tw,fi)
    } else {
        tarFile(srcBase,srcRelative,tw,fi)
    }
    return nil
}

func Exists(name string) bool{
    _,err:=os.Stat(name)
    return err==nil||os.IsExist(err)
}

func FileExists(filename string) bool{
    fi,err:=os.Stat(filename)
    return (err==nil||os.IsExist(err))&&!fi.IsDir()
}

func DirExists(dirname string) bool{
    fi,err:=os.Stat(dirname)
    return (err==nil||os.IsExist(err))&& fi.IsDir()
}


func tarFile(srcBase,srcRelative string,tw *tar.Writer,fi os.FileInfo)(err error){
    srcFull:=srcBase+srcRelative
    hdr,er:=tar.FileInfoHeader(fi,"")
    if er!=nil{
        return er
    }
    hdr.Name=srcRelative

    if er=tw.WriteHeader(hdr);er!=nil{
        return er
    }
    fr,er:=os.Open(srcFull)
    if er!=nil{
        return er
    }
    defer fr.Close()

    if _,er=io.Copy(tw,fr);er!=nil{
        return er
    }
    return nil
}

func tarDir(srcBase,srcRelative string,tw *tar.Writer,fi os.FileInfo)(err error){
    srcFull:=srcBase+srcRelative

    last:=len(srcRelative)-1
    if srcRelative[last]!=os.PathSeparator{
        srcRelative+=string(os.PathSeparator)
    }

    fis,er:=ioutil.ReadDir(srcFull)
    if er!=nil{ return er }

    for _,fi:=range fis{
        if fi.IsDir(){
            tarDir(srcBase,srcRelative+fi.Name(),tw,fi)
        } else{
            tarFile(srcBase,srcRelative+fi.Name(),tw,fi)
        }
    }

    if len(srcRelative)>0{
        hdr,er:=tar.FileInfoHeader(fi,"")
        if er!=nil{ return er    }
        hdr.Name=srcRelative
        if er=tw.WriteHeader(hdr);er!=nil{ return er }
    }
    return nil
}

func unTarFile(dstFile string,tr *tar.Reader) error{
    fw,er:=os.Create(dstFile)
    if er!=nil{ return er }
    defer fw.Close()

    _,er=io.Copy(fw,tr)
    if er!=nil{ return er }

    return nil
}

func UnTar(srcTar,dstDir string) (err error){
    dstDir=path.Clean(dstDir)+string(os.PathSeparator)

    fr,er:=os.Open(srcTar)
    if er!=nil{ return er   }
    defer fr.Close()

    tr:=tar.NewReader(fr)

    for hdr,er:=tr.Next();er!=io.EOF;hdr,er=tr.Next(){
        if er!=nil{ return er }
        fi:=hdr.FileInfo()

        dstFullPath:=dstDir+hdr.Name
        if hdr.Typeflag==tar.TypeDir{
            os.MkdirAll(dstFullPath,fi.Mode().Perm())
            os.Chmod(dstFullPath,fi.Mode().Perm())
        }else {
            os.MkdirAll(path.Dir(dstFullPath),os.ModePerm)
            if er:=unTarFile(dstFullPath,tr);er!=nil{ return er }
            os.Chmod(dstFullPath,fi.Mode().Perm())
        }
    }
    return nil
}