package main

import "fmt"

func main(){
  
    var radiusOfCircle = 5
    var messageToDisplay = "The area of Circle for a given radius" + radiusOfCircle + "is "
    const PI = 3.14
    var result = PI * radiusOfCircle * radiusOfCircle
    fmt.Println(messageToDisplay,result)
  
}
