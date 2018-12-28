package kata

import "strings"
import "fmt"

func DNAStrand(dna string) string {
  // your code here

  s := strings.Split(dna, "")
  
  for _, element := range s {

		fmt.Println(element)

  }
  return dna
  
}