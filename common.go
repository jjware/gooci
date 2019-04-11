package gooci

// #cgo pkg-config: oci8
// #include "gooci.h"
import "C"
import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"unsafe"
)

const maxErrorMessageSize = 3024

type Mode int

const (
	ModeDefault            = Mode(C.OCI_DEFAULT)
	ModeThreaded           = Mode(C.OCI_THREADED)
	ModeObject             = Mode(C.OCI_OBJECT)
	ModeEvents             = Mode(C.OCI_EVENTS)
	ModeNoUCB              = Mode(C.OCI_NO_UCB)
	ModeNoMutex            = Mode(C.OCI_NO_MUTEX)
	ModeNewLengthSemantics = Mode(C.OCI_NEW_LENGTH_SEMANTICS)
)

func firstNullByteIndex(s []C.uchar) int {
	for i := 0; i < len(s); i++ {
		if 0 == s[i] {
			return i
		}
	}
	return -1
}

type cString []C.uchar

func (m cString) String() string {
	return fmt.Sprintf("%v", m[0:firstNullByteIndex(m)])
}

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

func getError(handlep unsafe.Pointer, handleType C.ub4) error {
	var sqlState *C.OraText
	var eCode C.sb4
	eMessage := make(cString, maxErrorMessageSize)
	eRecord := make(errorRecord)

	for {
		r2 := C.OCIErrorGet(
			handlep,
			C.ub4(1),
			sqlState,
			&eCode,
			&eMessage[0],
			C.uint(maxErrorMessageSize),
			handleType,
		)

		if C.OCI_ERROR == r2 {
			return errors.New("message larger than buffer")
		} else if C.OCI_INVALID_HANDLE == r2 {
			return errors.New("invalid handle")
		} else if C.OCI_SUCCESS == r2 {
			eRecord[int(eCode)] = eMessage.String()
		} else {
			break
		}
	}
	return eRecord
}

func checkResult(result C.int, err *Error) error {
	if C.OCI_SUCCESS == result {
		return nil
	} else if C.OCI_SUCCESS_WITH_INFO == result {
		log.Printf("info: %s", getError(unsafe.Pointer(err.handle), C.OCI_HTYPE_ERROR).Error())
		return nil
	}
	return getError(unsafe.Pointer(err.handle), C.OCI_HTYPE_ERROR)
}
