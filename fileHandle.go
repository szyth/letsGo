package main
import(
  "fmt"
  "flag"
  "io/ioutil"
)
func main(){
  fileFlag := flag.String("file","No file selected","a file to read its content")
  flag.Parse()

  fileData, err := ioutil.ReadFile(*fileFlag)
  if err!=nil {
    fmt.Println("Error reading file:", err)
    return
  }
  fmt.Println("File Contents: \n", string(fileData))


}
