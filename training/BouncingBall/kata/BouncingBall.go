package kata


func BouncingBall(h, bounce, window float64) int {
  
  if h < 0 || window >= h {
    return -1
  } else if (bounce <=0 || bounce >= 1) {
    return -1
  } else if (h * bounce) < window {
    return -1
  }

  var returnValue = -1
 for ; h > window; h *= bounce {
    returnValue+= 2
  }

  return returnValue
}