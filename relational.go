package gooci

// #include "gooci.h"
import "C"
import "unsafe"

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
	*namep = cStringToGoString(cstrName, int(nameLen))
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
	password string,
	dbname string,
	mode Mode,
) Result {
	handle := (*C.OCISvcCtx)(*svcpp)

	var cstrUsername *C.uchar
	usernameLen := len(username)

	var cstrPassword *C.uchar
	passwordLen := len(password)

	var cstrDBName *C.uchar
	dbnameLen := len(dbname)

	if usernameLen > 0 {
		cstrUsername = goStringToCString(username)
	} else {
		cstrUsername = nil
	}

	if passwordLen > 0 {
		cstrPassword = goStringToCString(password)
	} else {
		cstrPassword = nil
	}

	if dbnameLen > 0 {
		cstrDBName = goStringToCString(dbname)
	} else {
		cstrDBName = nil
	}

	result := C.OCILogon2(
		(*C.OCIEnv)(envp),
		(*C.OCIError)(errp),
		&handle,
		cstrUsername,
		C.ub4(usernameLen),
		cstrPassword,
		C.ub4(passwordLen),
		cstrDBName,
		C.ub4(dbnameLen),
		C.ub4(mode),
	)
	*svcpp = (*SvcCtx)(handle)
	return Result(result)
}
