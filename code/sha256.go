package sha256

import (
  "crypto/sha256"
  "encoding/hex"
  "io"
  "os"
)

func Sha256String(s string) (string,error){
  var hashValue string
  hash256 := sha256.New()
  _,err:=hash256.Write([]byte(s))
  if err !=nil{
    return hashValue,err
  }
  
  hashInBytes := hash256.Sum(nil)
  hashValue = hex.EncodeToString(hashInBytes)
  
  return hashValue,nil
}

func Sha256File(s string) (string,error){
  var hashValue string
  fileofopen,err:=os.Open(s)
  if err!=nil{
    return hashValue,err
  }
  
  defer fileofopen.Close()
  
  hash256:= sha256.New()
  if _,err :=io.Copy(hash256,fileofopen);err!=nil{
    return hashValue,err
  }
  
  hashInBytes:=hash256.Sum(nil)
  hashValue=hex.EncodeToString(hashInBytes)
  
  return hashValue,nil
}
