package main

import (
 "bufio"
 "fmt"
 "os"
)

func main() {
 var in *bufio.Reader
 var out *bufio.Writer
 in = bufio.NewReader(os.Stdin)
 out = bufio.NewWriter(os.Stdout)
 defer out.Flush()

 var t int
 fmt.Fscan(in, &t)

 for i := 0; i < t; i++ {
  var number string
  fmt.Fscan(in, &number)

  n := len(number)
  if n == 1 {
   fmt.Fprintln(out, 0)
   continue
  }

  maxNumber := ""
  found := false

  for j := 0; j < n-1; j++ {
   if number[j] < number[j+1] {
    maxNumber = number[:j] + number[j+1:]
    found = true
    break
   }
  }
  
  if !found {
   maxNumber = number[:n-1]
  }
  
  fmt.Fprintln(out, maxNumber)
 }
}
