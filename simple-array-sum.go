package simpleArraySum

import (
    "fmt"
)

/*
 * Complete the simpleArraySum function below.
 */
func simpleArraySum(ar []int32) int32 {
    /*
     * Write your code here.
     */

     var results int32
     for _, v := range ar {
         results = results + v
     }
     return results

}