package main
import(
  "fmt"
  "flag"
)
func main(){
  //one way of declaring flags
  flag1ptr := flag.String("word","defaultValueOfWord","short description for word")
  flag2ptr := flag.Int("num",0,"short description for number")


  //another way
  var name string
  flag.StringVar(&name,"name","zia","short desc for name")


  flag.Parse()
  fmt.Println("Word:",*flag1ptr)
  fmt.Println("Number:",*flag2ptr)
  fmt.Println("Name:",name)
}
