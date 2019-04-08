package gooci

// #cgo pkg-config: oci8
// #include "gooci.h"
import "C"
import (
	"bytes"
	"log"
	"strconv"
	"unsafe"
)

const (
	ociSuccess         = C.int(0)
	ociSuccessWithInfo = C.int(1)
	ociError           = C.int(-1)
	ociNoData          = C.int(100)
	ociInvalidHandle   = C.int(-2)
)

type EnvironmentMode int

const (
	Default            = EnvironmentMode(C.OCI_DEFAULT)
	Threaded           = EnvironmentMode(C.OCI_THREADED)
	Object             = EnvironmentMode(C.OCI_OBJECT)
	Events             = EnvironmentMode(C.OCI_EVENTS)
	NoUCB              = EnvironmentMode(C.OCI_NO_UCB)
	NoMutex            = EnvironmentMode(C.OCI_NO_MUTEX)
	NewLengthSemantics = EnvironmentMode(C.OCI_NEW_LENGTH_SEMANTICS)
)

func checkResult(result C.int, err *Error) error {
	if ociSuccess == result {
		return nil
	} else if ociSuccessWithInfo == result {
		log.Printf("info: %s", getError(unsafe.Pointer(err.handle), ociHtypeError).Error())
		return nil
	}
	return getError(unsafe.Pointer(err.handle), ociHtypeError)
}

func free(handle unsafe.Pointer, handleType C.uint) error {
	result := C.OCIHandleFree(handle, C.uint(handleType))

	if ociSuccess != result {
		return getError(unsafe.Pointer(handle), handleType)
	}
	return nil
}

const maxErrorMessageSize = 3024

type errorRecord map[int]string

func (e errorRecord) Error() string {
	str := ""
	ctr := 0

	for k, v := range e {

		if ctr > 0 {
			str = str + ": "
		}
		str = str + strconv.Itoa(k) + " - " + v
		ctr++
	}
	return str
}

func getError(handle unsafe.Pointer, handleType C.uint) error {
	recordNMB := C.uint(1)
	var sqlState *C.uchar = nil
	eRecord := make(errorRecord)

	for {
		eCode := C.int(0)
		eText := make([]byte, maxErrorMessageSize)

		eResult := C.OCIErrorGet(
			handle,
			recordNMB,
			sqlState,
			&eCode,
			(*C.OraText)(&eText[0]),
			C.uint(maxErrorMessageSize),
			handleType,
		)

		if ociNoData == eResult {
			break
		} else if ociSuccess == eResult {
			ndx := bytes.IndexByte(eText, 0)
			eRecord[int(eCode)] = string(eText[:ndx])
		}
	}
	return eRecord
}
