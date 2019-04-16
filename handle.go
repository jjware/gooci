package gooci

// #include "gooci.h"
import "C"
import (
	"errors"
	"unsafe"
)

type Handle interface{}

type Env C.OCIEnv

func (env *Env) Close() error {
	result := HandleFree(env)

	if ResultInvalidHandle == result {
		return errors.New("invalid handle")
	}
	return nil
}

type Server C.OCIServer

func (srv *Server) Close() error {
	result := HandleFree(srv)

	if ResultInvalidHandle == result {
		return errors.New("invalid handle")
	}
	return nil
}

type Error C.OCIError

func (e *Error) Close() error {
	result := HandleFree(e)

	if ResultInvalidHandle == result {
		return errors.New("invalid handle")
	}
	return nil
}

type Session C.OCISession

func (s *Session) Close() error {
	result := HandleFree(s)

	if ResultInvalidHandle == result {
		return errors.New("invalid handle")
	}
	return nil
}

type SvcCtx C.OCISvcCtx

func (svc *SvcCtx) Close() error {
	result := HandleFree(svc)

	if ResultInvalidHandle == result {
		return errors.New("invalid handle")
	}
	return nil
}

type CPool C.OCICPool

func (cp *CPool) Close() error {
	result := HandleFree(cp)

	if ResultInvalidHandle == result {
		return errors.New("invalid handle")
	}
	return nil
}

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

func HandleFree(handle Handle) Result {
	switch v := handle.(type) {
	case *Env:
		return Result(C.OCIHandleFree(unsafe.Pointer((*C.OCIEnv)(v)), C.OCI_HTYPE_ENV))
	case *Server:
		return Result(C.OCIHandleFree(unsafe.Pointer((*C.OCIServer)(v)), C.OCI_HTYPE_SERVER))
	case *Error:
		return Result(C.OCIHandleFree(unsafe.Pointer((*C.OCIError)(v)), C.OCI_HTYPE_ERROR))
	case *Session:
		return Result(C.OCIHandleFree(unsafe.Pointer((*C.OCISession)(v)), C.OCI_HTYPE_SESSION))
	case *SvcCtx:
		return Result(C.OCIHandleFree(unsafe.Pointer((*C.OCISvcCtx)(v)), C.OCI_HTYPE_SVCCTX))
	case *CPool:
		return Result(C.OCIHandleFree(unsafe.Pointer((*C.OCICPool)(v)), C.OCI_HTYPE_CPOOL))
	default:
		return ResultInvalidHandle
	}
}

func HandleAlloc(env *Env, handlepp Handle) Result {
	var result C.sword
	var handle unsafe.Pointer
	var buffer *unsafe.Pointer

	switch v := handlepp.(type) {
	case **Server:
		result = C.OCIHandleAlloc(
			unsafe.Pointer((*C.OCIEnv)(env)),
			&handle,
			C.OCI_HTYPE_SERVER,
			C.size_t(0),
			buffer,
		)
		*v = (*Server)((*C.OCIServer)(handle))
	case **Error:
		result = C.OCIHandleAlloc(
			unsafe.Pointer((*C.OCIEnv)(env)),
			&handle,
			C.OCI_HTYPE_ERROR,
			C.size_t(0),
			buffer,
		)
		*v = (*Error)((*C.OCIError)(handle))
	case **Session:
		result = C.OCIHandleAlloc(
			unsafe.Pointer((*C.OCIEnv)(env)),
			&handle,
			C.OCI_HTYPE_SESSION,
			C.size_t(0),
			buffer,
		)
		*v = (*Session)((*C.OCISession)(handle))
	case **SvcCtx:
		result = C.OCIHandleAlloc(
			unsafe.Pointer((*C.OCIEnv)(env)),
			&handle,
			C.OCI_HTYPE_SVCCTX,
			C.size_t(0),
			buffer,
		)
		*v = (*SvcCtx)((*C.OCISvcCtx)(handle))
	case **CPool:
		result = C.OCIHandleAlloc(
			unsafe.Pointer((*C.OCIEnv)(env)),
			&handle,
			C.OCI_HTYPE_CPOOL,
			C.size_t(0),
			buffer,
		)
		*v = (*CPool)((*C.OCICPool)(handle))
	default:
		return ResultInvalidHandle
	}
	return Result(result)
}
