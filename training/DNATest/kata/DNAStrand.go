package kata

import "strings"
import "bytes"


func DNAStrand(dna string) string {
  // your code here

  s := strings.Split(dna, "")
  
  var replica bytes.Buffer

  for _, element := range s {

		switch element {
    case "A":
      replica.WriteString("T")
    case "T":
      replica.WriteString("A")
    case "C":
      replica.WriteString("G")
    case "G":
      replica.WriteString("C")
    }

  }
  return replica.String()
  
}