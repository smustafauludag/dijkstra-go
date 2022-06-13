package colourfulterminal

import "fmt"

var cEND string = "\033[0m"
var cRED2 string = "\033[91m"
var cGREEN2 string = "\033[92m"
var cYELLOW2 string = "\033[93m"
var cBLUE2 string = "\033[94m"


func Info(inf string){ fmt.Println(cBLUE2 +"[INFO]: "+ inf + cEND) }
  
func Success(scc string){ fmt.Println(cGREEN2 +"[SUCCESS]: "+ scc + cEND) }

func Warning(warn string){ fmt.Println(cYELLOW2 +"[WARNING]: "+ warn + cEND) }

func Error(err string){ fmt.Println(cRED2 +"[ERROR]: "+ err + cEND) }
