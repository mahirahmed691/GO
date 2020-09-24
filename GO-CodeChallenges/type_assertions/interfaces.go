package main

import (
    "fmt"
)

type Developer struct {
    Name string
    Age int
}

func GetDeveloper(name interface{}, age interface{}) Developer {
    var dev Developer
    dev.Name = name.(string)
    dev.Age = age.(int)
    return dev
}

func main(){
    
    fmt.Println("Interfaces & Structs implementation")

    var name interface{} = "Mahir"
    var age interface{} = 27

    dynamicDev := GetDeveloper(name, age)
    fmt.Println("---Name---:")
    fmt.Println(dynamicDev.Name)
    fmt.Println("---Age---:")
    fmt.Println(dynamicDev.Age)
}
