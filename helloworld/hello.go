package helloworld

/*
#include <stdlib.h>
#include <stddef.h>
#include "../libwaku.h"
#cgo LDFLAGS:-L${SRCDIR}/../ -lwaku -lrln -lm -lminiupnpc -lnatpmp -lbacktrace_wrapper -lbacktrace
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func WakuNew(configJSON *C.char, cb C.WakuCallBack, userData unsafe.Pointer) {
	result := C.waku_new(configJSON, cb, userData)
	fmt.Println(result)
}
