package gooci

// #include "gooci.h"
import "C"

func ConnectionPoolCreate(
	env *Env,
	e *Error,
	cpool *CPool,
	connectionString string,
	min int,
	max int,
	incr int,
	username string,
	password string,
	mode Mode,
) (string, error) {
	var poolName *C.uchar
	var poolLen C.sb4

	connstr := goStringToCString(connectionString)
	uname := goStringToCString(username)
	pword := goStringToCString(password)

	result := C.OCIConnectionPoolCreate(
		env.handle,
		e.handle,
		cpool.handle,
		&poolName,
		&poolLen,
		connstr,
		C.sb4(len(connectionString)),
		C.ub4(min),
		C.ub4(max),
		C.ub4(incr),
		uname,
		C.sb4(len(username)),
		pword,
		C.sb4(len(password)),
		C.ub4(mode),
	)
	err := checkResult(result, e)

	if nil != err {
		return "", err
	}
	return cStringToGoString(poolName, int(poolLen)), nil
}
