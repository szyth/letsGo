//vim code formatting gg=G
package main
import(
  "fmt"
  "flag"

  "io/ioutil"

  "os"
  "bufio"
  "log"
)
var(
  passfile string
  dictfile string
)
func init(){
  flag.StringVar(&passfile,"p","","Path to Shadow File")
  flag.StringVar(&dictfile,"d","","Path to dictionary File")
}


func main(){
  flag.Parse()

  //check if no file is provided
  if passfile =="" || dictfile == ""{
    fmt.Println("./unixPasswordCrack -h")
    os.Exit(0)
  }

  //READ file using ioutil method; opens whole file at once; one-liner code(ReadFile)(ioutil)
  dictfileData,err := ioutil.ReadFile(dictfile)
  if err!=nil {
    fmt.Println("Error opening file: ",err)
    os.Exit(0)
  }
  fmt.Println(string(dictfileData))
  //END of READ ioutil


  //READ file using scanner method; opens file line by line; 6-liner code(Open,Close,NewScanner,Scan,Text,Err)(os,log,bufio)
  passfileData, err := os.Open(passfile)
  if err!=nil {
    fmt.Println("Error opening file: ",err)
  }
  //CLOSE file; defer executes at function return
  defer func(){
    if err = passfileData.Close(); err!=nil{
      log.Fatal(err)
    }
  }()
  scan := bufio.NewScanner(passfileData)
  for scan.Scan(){
    fmt.Println(scan.Text(),":")
  }
  err = scan.Err()
  if err!=nil {
    log.Fatal(err)
  }
  //END of READ scanner
}

