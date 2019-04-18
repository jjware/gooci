package gooci

// #include "gooci.h"
import "C"
import (
	"fmt"
	"unsafe"
)

func EnvNlsCreate(env **Env, mode Mode) Result {
	return Result(C.OCIEnvNlsCreate(
		(**C.OCIEnv)(unsafe.Pointer(env)),
		C.ub4(mode),
		nil,
		nil,
		nil,
		nil,
		C.size_t(0),
		nil,
		C.ub2(0),
		C.ub2(0),
	))
}

func ConnectionPoolCreate(
	envp *Env,
	errp *Error,
	cpp *CPool,
	namep *string,
	connstr string,
	min int,
	max int,
	incr int,
	username string,
	password string,
	mode Mode,
) Result {
	var cstrName *C.uchar
	var nameLen C.sb4

	cstrConnstr := goStringToCString(connstr)
	cstrUsername := goStringToCString(username)
	cstrPassword := goStringToCString(password)

	result := C.OCIConnectionPoolCreate(
		(*C.OCIEnv)(envp),
		(*C.OCIError)(errp),
		(*C.OCICPool)(cpp),
		&cstrName,
		&nameLen,
		cstrConnstr,
		C.sb4(len(connstr)),
		C.ub4(min),
		C.ub4(max),
		C.ub4(incr),
		cstrUsername,
		C.sb4(len(username)),
		cstrPassword,
		C.sb4(len(password)),
		C.ub4(mode),
	)

	if nil != namep {
		*namep = cStringToGoString(cstrName, int(nameLen))
	}
	return Result(result)
}

func ConnectionPoolDestroy(cpp *CPool, errp *Error) Result {
	return Result(C.OCIConnectionPoolDestroy(
		(*C.OCICPool)(cpp),
		(*C.OCIError)(errp),
		C.OCI_DEFAULT,
	))
}

func Logon2(
	envp *Env,
	errp *Error,
	svcpp **SvcCtx,
	username string,
	password fmt.Stringer,
	dbname string,
	mode Mode,
) Result {
	handle := (*C.OCISvcCtx)(*svcpp)

	var cstrPassword *C.uchar
	var passwordLen int

	cstrUsername := goStringToCString(username)

	if nil != password {
		strPassword := password.String()
		cstrPassword= goStringToCString(strPassword)
		passwordLen = len(strPassword)
	} else {
		cstrPassword = nil
		passwordLen = 0
	}
	cstrDBName := goStringToCString(dbname)

	result := C.OCILogon2(
		(*C.OCIEnv)(envp),
		(*C.OCIError)(errp),
		&handle,
		cstrUsername,
		C.ub4(len(username)),
		cstrPassword,
		C.ub4(passwordLen),
		cstrDBName,
		C.ub4(len(dbname)),
		C.ub4(mode),
	)
	*svcpp = (*SvcCtx)(handle)
	return Result(result)
}

func Logoff(svcp *SvcCtx, errp *Error) Result {
	return Result(C.OCILogoff((*C.OCISvcCtx)(svcp), (*C.OCIError)(errp)))
}
