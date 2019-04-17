package gooci

// #cgo pkg-config: oci8
// #include "gooci.h"
import "C"
import (
	"unsafe"
)

type Result int

const (
	ResultSuccess         = Result(C.OCI_SUCCESS)
	ResultSuccessWithInfo = Result(C.OCI_SUCCESS_WITH_INFO)
	ResultNoData          = Result(C.OCI_NO_DATA)
	ResultError           = Result(C.OCI_ERROR)
	ResultInvalidHandle   = Result(C.OCI_INVALID_HANDLE)
	ResultNeedData        = Result(C.OCI_NEED_DATA)
	ResultStillExecuting  = Result(C.OCI_STILL_EXECUTING)
	ResultContinue        = Result(C.OCI_CONTINUE)
	ResultRowCbkDone      = Result(C.OCI_ROWCBK_DONE)
)

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
	ModeLogon2CPool = Mode(C.OCI_LOGON2_CPOOL)
	ModeLogon2SPool = Mode(C.OCI_LOGON2_SPOOL)
	ModeLogon2StmtCache = Mode(C.OCI_LOGON2_STMTCACHE)
	ModeLogon2Proxy = Mode(C.OCI_LOGON2_STMTCACHE)
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
