package gooci

// #include "gooci.h"
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type Env struct {
	handle *C.OCIEnv
}

func (env *Env) Close() error {
	result := C.OCIHandleFree(unsafe.Pointer(env.handle), C.OCI_HTYPE_ENV)

	if C.OCI_INVALID_HANDLE == result {
		return fmt.Errorf("invalid handle")
	}
	return nil
}

func NewEnv(mode Mode) (*Env, error) {
	var env Env

	r1 := C.OCIEnvNlsCreate(
		&env.handle,
		C.ub4(mode),
		nil,
		nil,
		nil,
		nil,
		C.size_t(0),
		nil,
		C.ub2(0),
		C.ub2(0),
	)

	if C.OCI_ERROR == r1 {
		return nil, getError(unsafe.Pointer(env.handle), C.OCI_HTYPE_ENV)
	}
	return &env, nil
}

type Server struct {
	handle *C.OCIServer
}

func (srv *Server) Close() error {
	result := C.OCIHandleFree(unsafe.Pointer(srv.handle), C.OCI_HTYPE_SERVER)

	if C.OCI_INVALID_HANDLE == result {
		return fmt.Errorf("invalid handle")
	}
	return nil
}

func NewServer(env *Env) (*Server, error) {
	var handle unsafe.Pointer
	var buffer *unsafe.Pointer

	result := C.OCIHandleAlloc(
		unsafe.Pointer(env.handle),
		&handle,
		C.OCI_HTYPE_SERVER,
		C.size_t(0),
		buffer,
	)

	if C.OCI_INVALID_HANDLE == result {
		return nil, errors.New("invalid handle")
	} else if C.OCI_ERROR == result {
		return nil, getError(unsafe.Pointer(env.handle), C.OCI_HTYPE_ENV)
	}
	return &Server{handle: (*C.OCIServer)(handle)}, nil
}

type Error struct {
	handle *C.OCIError
}

func (e *Error) Close() error {
	result := C.OCIHandleFree(unsafe.Pointer(e.handle), C.OCI_HTYPE_ERROR)

	if C.OCI_INVALID_HANDLE == result {
		return fmt.Errorf("invalid handle")
	}
	return nil
}

func NewError(env *Env) (*Error, error) {
	var handle unsafe.Pointer
	var buffer *unsafe.Pointer

	result := C.OCIHandleAlloc(
		unsafe.Pointer(env.handle),
		&handle,
		C.OCI_HTYPE_ERROR,
		C.size_t(0),
		buffer,
	)

	if C.OCI_INVALID_HANDLE == result {
		return nil, errors.New("invalid handle")
	} else if C.OCI_ERROR == result {
		return nil, getError(unsafe.Pointer(env.handle), C.OCI_HTYPE_ENV)
	}
	return &Error{handle: (*C.OCIError)(handle)}, nil
}

type Session struct {
	handle *C.OCISession
}

func (s *Session) Close() error {
	result := C.OCIHandleFree(unsafe.Pointer(s.handle), C.OCI_HTYPE_SESSION)

	if C.OCI_INVALID_HANDLE == result {
		return fmt.Errorf("invalid handle")
	}
	return nil
}

func NewSession(env *Env) (*Session, error) {
	var handle unsafe.Pointer
	var buffer *unsafe.Pointer

	result := C.OCIHandleAlloc(
		unsafe.Pointer(env.handle),
		&handle,
		C.OCI_HTYPE_SESSION,
		C.size_t(0),
		buffer,
	)

	if C.OCI_INVALID_HANDLE == result {
		return nil, errors.New("invalid handle")
	} else if C.OCI_ERROR == result {
		return nil, getError(unsafe.Pointer(env.handle), C.OCI_HTYPE_ENV)
	}
	return &Session{handle: (*C.OCISession)(handle)}, nil
}

type SvcCtx struct {
	handle *C.OCISvcCtx
}

func (s *SvcCtx) Close() error {
	result := C.OCIHandleFree(unsafe.Pointer(s.handle), C.OCI_HTYPE_SVCCTX)

	if C.OCI_INVALID_HANDLE == result {
		return fmt.Errorf("invalid handle")
	}
	return nil
}

func NewSvcCtx(env *Env) (*SvcCtx, error) {
	var handle unsafe.Pointer
	var buffer *unsafe.Pointer

	result := C.OCIHandleAlloc(
		unsafe.Pointer(env.handle),
		&handle,
		C.OCI_HTYPE_SVCCTX,
		C.size_t(0),
		buffer,
	)

	if C.OCI_INVALID_HANDLE == result {
		return nil, errors.New("invalid handle")
	} else if C.OCI_ERROR == result {
		return nil, getError(unsafe.Pointer(env.handle), C.OCI_HTYPE_ENV)
	}
	return &SvcCtx{handle: (*C.OCISvcCtx)(handle)}, nil
}

type CPool struct {
	handle *C.OCICPool
}

func (cp *CPool) Close() error {
	result := C.OCIHandleFree(unsafe.Pointer(cp.handle), C.OCI_HTYPE_CPOOL)

	if C.OCI_INVALID_HANDLE == result {
		return fmt.Errorf("invalid handle")
	}
	return nil
}

func NewConnectionPool(env *Env) (*CPool, error) {
	var handle unsafe.Pointer
	var buffer *unsafe.Pointer

	result := C.OCIHandleAlloc(
		unsafe.Pointer(env.handle),
		&handle,
		C.OCI_HTYPE_CPOOL,
		C.size_t(0),
		buffer,
	)

	if C.OCI_INVALID_HANDLE == result {
		return nil, errors.New("invalid handle")
	} else if C.OCI_ERROR == result {
		return nil, getError(unsafe.Pointer(env.handle), C.OCI_HTYPE_ENV)
	}
	return &CPool{handle: (*C.OCICPool)(handle)}, nil
}