package raindrops

import (
    "fmt"
    "strconv"
)

func Convert(number int) string {

soundofRain := ""

if number %3 == 0 {
    soundofRain += "Pling"
}

if number %5 == 0 {
    soundofRain += "Plang"
}

if number %7 == 0 {
    soundofRain += "Plong"
}

if soundofRain == "" {
    return strconv.Itoa(number)
}
return soundofRain
}

func main() {

    fmt.Println(Convert(34))
}

