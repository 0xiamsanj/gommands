package main

import (
    "fmt"
    "os"
    "log"
)
func ls( path string){
    entries, err := os.ReadDir(path)
    if err != nil {
        log.Fatal(err)
    }

    for _, e := range entries {
        fmt.Println(e.Name())
    }
}
func main() {
    var path string
    path = "/home/sanj"
    ls(path)
}
