package main

import (
    "fmt"
)

func main() {
    
    userch := make(chan string, 2)
    
    userch <- "Aakash"
    userch <- "Jack"
    user := <- userch
    fmt.Println(user)
    userch <- "Aryan"
    users := <- userch
    fmt.Println(users)


}

