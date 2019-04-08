package gooci

// #include "gooci.h"
import "C"
import "unsafe"

const (
	ociHtypeEnv     = C.OCI_HTYPE_ENV
	ociHtypeServer  = C.OCI_HTYPE_SERVER
	ociHtypeError   = C.OCI_HTYPE_ERROR
	ociHtypeSession = C.OCI_HTYPE_SESSION
	ociHtypeService = C.OCI_HTYPE_SVCCTX
)

/*
 * Environment Handle
 */

type Environment struct {
	handle *C.OCIEnv
}

func (env *Environment) Close() error {
	return free(unsafe.Pointer(env.handle), ociHtypeEnv)
}

func NewEnvironment(mode EnvironmentMode) (*Environment, error) {
	env := &Environment{handle: nil}
	m := C.ub4(mode)

	result := C.OCIEnvNlsCreate(&env.handle, m, nil, nil, nil, nil, C.size_t(0), nil, 0, 0)

	if ociSuccess != result {
		return nil, getError(unsafe.Pointer(env.handle), ociHtypeEnv)
	}
	return env, nil
}

/*
 * Server Handle
 */

type Server struct {
	handle *C.OCIServer
}

func (srv *Server) Close() error {
	return free(unsafe.Pointer(srv.handle), ociHtypeServer)
}

func NewServer(env *Environment) (*Server, error) {
	var handle unsafe.Pointer
	var buffer *unsafe.Pointer

	result := C.OCIHandleAlloc(
		unsafe.Pointer(env.handle),
		&handle,
		ociHtypeServer,
		C.size_t(0),
		buffer,
	)

	if ociSuccess != result {
		return nil, getError(unsafe.Pointer(env.handle), ociHtypeEnv)
	}
	return &Server{handle: (*C.OCIServer)(handle)}, nil
}

/*
 * Error Handle
 */

type Error struct {
	handle *C.OCIError
}

func (e *Error) Close() error {
	return free(unsafe.Pointer(e.handle), ociHtypeError)
}

func NewError(env *Environment) (*Error, error) {
	var handle unsafe.Pointer
	var buffer *unsafe.Pointer

	result := C.OCIHandleAlloc(
		unsafe.Pointer(env.handle),
		&handle,
		ociHtypeError,
		C.size_t(0),
		buffer,
	)

	if ociSuccess != result {
		return nil, getError(unsafe.Pointer(env.handle), ociHtypeEnv)
	}
	return &Error{handle: (*C.OCIError)(handle)}, nil
}

/*
 * Session Handle
 */

type Session struct {
	handle *C.OCISession
}

func (s *Session) Close() error {
	return free(unsafe.Pointer(s.handle), ociHtypeSession)
}

func NewSession(env *Environment) (*Session, error) {
	var handle unsafe.Pointer
	var buffer *unsafe.Pointer

	result := C.OCIHandleAlloc(
		unsafe.Pointer(env.handle),
		&handle,
		ociHtypeSession,
		C.size_t(0),
		buffer,
	)

	if ociSuccess != result {
		return nil, getError(unsafe.Pointer(env.handle), ociHtypeEnv)
	}
	return &Session{handle: (*C.OCISession)(handle)}, nil
}

/*
 * Service Handle
 */

type Service struct {
	handle *C.OCISvcCtx
}

func (s *Service) Close() error {
	return free(unsafe.Pointer(s.handle), ociHtypeService)
}

func NewService(env *Environment) (*Service, error) {
	var handle unsafe.Pointer
	var buffer *unsafe.Pointer

	result := C.OCIHandleAlloc(
		unsafe.Pointer(env.handle),
		&handle,
		ociHtypeService,
		C.size_t(0),
		buffer,
	)

	if ociSuccess != result {
		return nil, getError(unsafe.Pointer(env.handle), ociHtypeEnv)
	}
	return &Service{handle: (*C.OCISvcCtx)(handle)}, nil
}
