// Basic Number Sorting Algorithm by omerakgoz34
// Numbers in any lenght and count

package main

import "fmt"

func main() {
    // Numbers to be sorted
    n := [...]int{41, 63, 96, 42, 47, 07, 41, 38, 73, 06, 14}
    y := 0

    fmt.Println("\nNumbers:", n)
    for i := 0;; {
        if i+1 >= len(n) {
           if y > 0 {
               i = 0
               y = 0
           } else {
               break
           }
        }

        if n[i] > n[i+1]{
            y++
            x := n[i]
            n[i] = n[i+1]
            n[i+1] = x
            fmt.Print("\n> ", n)
        }

        i++
    }
    fmt.Println(" done.\n")
}
