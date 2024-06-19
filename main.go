GNU nano 7.2         main.go
package main

import (
    "fmt"
    "strconv"
)

func main() {

    vector31, err1 := strconv.Atoi("45")
    vector32, err2 := strconv.Atoi("25")


    if err1 != nil || err2 != nil {
        fmt.Println("Error converting strings to >
        return
}
    fmt.Println(vector31 * vector32)
}
