package gooci

// #cgo pkg-config: oci8
// #include "gooci.h"
import "C"
import (
	"errors"
	"log"
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
	ModeCPoolReinitialize  = Mode(C.OCI_CPOOL_REINITIALIZE)
)

type Cred int

const (
	CredRDBMS = Cred(C.OCI_CRED_RDBMS)
	CredExt   = Cred(C.OCI_CRED_EXT)
)

func firstNullByteIndex(s []C.uchar) int {
	for i := 0; i < len(s); i++ {
		if 0 == s[i] {
			return i
		}
	}
	return -1
}

func cStringToGoString(str *C.uchar, length int) (result string) {
	size := int(unsafe.Sizeof(*str))
	byt := C.GoBytes(unsafe.Pointer(str), (C.int)(size*length))
	return string(byt)
}

func goStringToCString(s string) *C.uchar {
	arr := make([]C.uchar, len(s)+1)
	i := 0

	for ; i < len(s); i++ {
		arr[i] = C.uchar(s[i])
	}
	arr[i] = 0
	return &arr[0]
}

type errorRecord map[int]string

func (e errorRecord) Error() string {
	str := ""
	ctr := 0

	for _, v := range e {

		if ctr > 0 {
			str = str + ": "
		}
		str += v
		ctr++
	}
	return str
}

func getError(handlep unsafe.Pointer, handleType C.ub4) error {
	var sqlState *C.uchar
	var eCode C.sb4
	eMessage := make([]C.uchar, maxErrorMessageSize)
	eRecord := make(errorRecord)

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
		eRecord[int(eCode)] = cStringToGoString(&eMessage[0], firstNullByteIndex(eMessage))
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
