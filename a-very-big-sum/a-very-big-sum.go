package averybigsum

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the aVeryBigSum function below.
func aVeryBigSum(ar []int64) int64 {
    var results int64
    for _, v := range ar {
        results = results + v
    }
    return results 
}

