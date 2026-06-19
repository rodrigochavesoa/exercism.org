package hamming

import (
    "fmt"
    "errors"
)

func Distance(a, b string)(int, error) {
    
    distance := 0

    if len(a) != len(b){
        return 0, errors.New("Lenght Different")
    }

    for i := range a {
        if a[i] != b[i] {
            distance++
        }
    }

    return distance, nil
    
}

func main() {
    tapeA := "GAGCCTACTAACGGGAT"
    tapeB := "CATCGTAATGACGGCCT"
    
result, error := (Distance(tapeA, tapeB))    
   
    fmt.Println(result, error)
}
