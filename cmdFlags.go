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
  var(
    name string
    boolean bool
  )
  flag.StringVar(&name,"name","zia","short desc for name")
  flag.BoolVar(&boolean,"boolean",false,"short desc for boolean")


  flag.Parse()
  fmt.Println("Word:",*flag1ptr)
  fmt.Println("Number:",*flag2ptr)
  fmt.Println("Name:",name)
  fmt.Println("Boolean:",boolean)
}
