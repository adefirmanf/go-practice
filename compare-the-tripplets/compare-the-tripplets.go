package comparetriplets

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the compareTriplets function below.
func compareTriplets(a []int32, b []int32) []int32 {
    if len(a) != len(b){
        return []int32{0}
    }
    score := []int32{0,0}

    for i, v := range a {
        if v > b[i]{ // A win
            score[0] = score[0] + 1
        } else if b[i] > v { // B win
            score[1] = score[1] + 1
        }
    }
    return score
}
