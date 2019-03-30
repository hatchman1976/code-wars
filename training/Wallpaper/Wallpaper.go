package kata

import "math"

func WallPaper(l, w, h float64) string {
  numbers := [21]string{
    "zero", "one", "two", "three", "four", 
    "five", "six", "seven", "eight", "nine", 
    "ten", "eleven", "twelve", "thirteen", "fourteen", 
    "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", 
    "twenty",
  }
  if l > 0 && h > 0 && w > 0 {
    // 1.15 for reserve, 5.2 for area per roll
    return numbers[int(math.Ceil(2.3 * (l * h + w * h) / 5.2))] 
  } else {
    return numbers[0]
  }
}