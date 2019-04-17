package gooci

// #include "gooci.h"
import "C"
import (
	"unsafe"
)

const maxErrorMessageSize = 3024

func ErrorGet(handle Handle, code *int, message *string) Result {
	var result C.sword
	var state *C.uchar
	var cint C.sb4
	cstr := make([]C.uchar, maxErrorMessageSize)

	switch v := handle.(type) {
	case *Env:
		result = C.OCIErrorGet(
			unsafe.Pointer((*C.OCIEnv)(v)),
			C.ub4(1),
			state,
			&cint,
			&cstr[0],
			C.ub4(maxErrorMessageSize),
			C.OCI_HTYPE_ENV,
		)
		break
	case *Error:
		result = C.OCIErrorGet(
			unsafe.Pointer((*C.OCIError)(v)),
			C.ub4(1),
			state,
			&cint,
			&cstr[0],
			C.ub4(maxErrorMessageSize),
			C.OCI_HTYPE_ERROR,
		)
		break
	default:
		result = C.OCI_INVALID_HANDLE
		break
	}
	*code = int(cint)
	*message = cStringToGoString(&cstr[0], firstNullByteIndex(cstr)) //ignore null byte
	return Result(result)
}
