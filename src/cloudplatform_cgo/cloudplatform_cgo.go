// cloudplatform_cgo project cloudplatform_cgo.go
package cloudplatform_cgo

import (
	"fmt"
)

/*
#include <stdio.h>
void sayHi() {
   printf("Hi");
}
char say(){
  return 'a';
}
*/
import "C"

func Cgo() {
	C.sayHi()
	fmt.Println(string(C.say()))
}
