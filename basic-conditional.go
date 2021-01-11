package main

import(
    "fmt"
    "runtime"
)

func main(){

    // Maybe, just maybe, you cannot create a newline with else
    // Meaning after closing bracket `}` the word `else` must follow
    angka:=2

    if angka == 1{
        fmt.Println("Angkanya adalah 1")
    }else if angka == 2{
        fmt.Println("Angkanya adalah 2")
    }else{
        fmt.Println("Angkanya bukanlah 1 atau 2")
    }

    // This is switch case
    os:= runtime.GOOS
    switch os{
        case "darwin":
            fmt.Println("OS X")
        case "linux":
            fmt.Println("Linux")
        default:
            fmt.Println(os)
    }
}