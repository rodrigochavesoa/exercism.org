package twofer

import "fmt"

func ShareWith (name string) string {

    if name == "" {
        return "One for you, one for me."
    } else {
        return "One for " + name + ", one for me."
    }
}

func main() {
    fmt.Println(ShareWith("Bob"))
}


