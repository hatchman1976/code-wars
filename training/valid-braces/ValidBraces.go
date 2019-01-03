package kata

import "strings"

func ValidBraces(str string) bool {

	for strings.Index(str, "{}") != -1 || strings.Index(str, "[]") != -1 || strings.Index(str, "()") != -1 {
		str = strings.Replace(strings.Replace(strings.Replace(str, "{}", "", -1), "[]", "", -1), "()", "", -1)

	}

	return len(str) == 0
	/* 	braces :=  str
	   	while(braces.indexOf("{}") != -1 || braces.indexOf("()") != -1 || braces.indexOf("[]") != -1) {
	   		braces = braces.replace("{}", "").replace("()", "").replace("[]", "");
	   	  }
	   	  return len(braces) == 0
	*/

}
