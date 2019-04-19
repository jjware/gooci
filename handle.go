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

type Error C.OCIError

func (e *Error) Close() error {
	result := HandleFree(e)

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

type Stmt C.OCIStmt

func (s *Stmt) Close() error {
	result := HandleFree(s)

	if ResultInvalidHandle == result {
		return errors.New("invalid handle")
	}
	return nil
}

type Bind C.OCIBind

func (b *Bind) Close() error {
	result := HandleFree(b)

	if ResultInvalidHandle == result {
		return errors.New("invalid handle")
	}
	return nil
}

type Define C.OCIDefine

func (d *Define) Close() error {
	result := HandleFree(d)

	if ResultInvalidHandle == result {
		return errors.New("invalid handle")
	}
	return nil
}

type Describe C.OCIDescribe

func (d *Describe) Close() error {
	result := HandleFree(d)

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

type Session C.OCISession

func (s *Session) Close() error {
	result := HandleFree(s)

	if ResultInvalidHandle == result {
		return errors.New("invalid handle")
	}
	return nil
}

type AuthInfo C.OCIAuthInfo

func (a *AuthInfo) Close() error {
	result := HandleFree(a)

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

type SPool C.OCISPool

func (sp *SPool) Close() error {
	result := HandleFree(sp)

	if ResultInvalidHandle == result {
		return errors.New("invalid handle")
	}
	return nil
}

type Trans C.OCITrans

func (t *Trans) Close() error {
	result := HandleFree(t)

	if ResultInvalidHandle == result {
		return errors.New("invalid handle")
	}
	return nil
}

type ComplexObject C.OCIComplexObject

func (co *ComplexObject) Close() error {
	result := HandleFree(co)

	if ResultInvalidHandle == result {
		return errors.New("invalid handle")
	}
	return nil
}

func HandleFree(handle Handle) Result {
	switch v := handle.(type) {
	case *Env:
		return Result(C.OCIHandleFree(unsafe.Pointer((*C.OCIEnv)(v)), C.OCI_HTYPE_ENV))
	case *Error:
		return Result(C.OCIHandleFree(unsafe.Pointer((*C.OCIError)(v)), C.OCI_HTYPE_ERROR))
	case *SvcCtx:
		return Result(C.OCIHandleFree(unsafe.Pointer((*C.OCISvcCtx)(v)), C.OCI_HTYPE_SVCCTX))
	case *Stmt:
		return Result(C.OCIHandleFree(unsafe.Pointer((*C.OCIStmt)(v)), C.OCI_HTYPE_STMT))
	case *Bind:
		return Result(C.OCIHandleFree(unsafe.Pointer((*C.OCIBind)(v)), C.OCI_HTYPE_BIND))
	case *Define:
		return Result(C.OCIHandleFree(unsafe.Pointer((*C.OCIDefine)(v)), C.OCI_HTYPE_DEFINE))
	case *Describe:
		return Result(C.OCIHandleFree(unsafe.Pointer((*C.OCIDescribe)(v)), C.OCI_HTYPE_DESCRIBE))
	case *Server:
		return Result(C.OCIHandleFree(unsafe.Pointer((*C.OCIServer)(v)), C.OCI_HTYPE_SERVER))
	case *Session:
		return Result(C.OCIHandleFree(unsafe.Pointer((*C.OCISession)(v)), C.OCI_HTYPE_SESSION))
	case *AuthInfo:
		return Result(C.OCIHandleFree(unsafe.Pointer((*C.OCIAuthInfo)(v)), C.OCI_HTYPE_AUTHINFO))
	case *CPool:
		return Result(C.OCIHandleFree(unsafe.Pointer((*C.OCICPool)(v)), C.OCI_HTYPE_CPOOL))
	case *SPool:
		return Result(C.OCIHandleFree(unsafe.Pointer((*C.OCISPool)(v)), C.OCI_HTYPE_SPOOL))
	case *Trans:
		return Result(C.OCIHandleFree(unsafe.Pointer((*C.OCITrans)(v)), C.OCI_HTYPE_TRANS))
	case *ComplexObject:
		return Result(C.OCIHandleFree(unsafe.Pointer((*C.OCIComplexObject)(v)), C.OCI_HTYPE_COMPLEXOBJECT))
	default:
		return ResultInvalidHandle
	}
}

func HandleAlloc(env *Env, handlepp Handle) Result {
	var result C.sword
	var handle unsafe.Pointer
	var buffer *unsafe.Pointer

	switch v := handlepp.(type) {
	case **Env:
		result = C.OCIHandleAlloc(
			unsafe.Pointer((*C.OCIEnv)(env)),
			&handle,
			C.OCI_HTYPE_ENV,
			C.size_t(0),
			buffer,
		)
		*v = (*Env)((*C.OCIEnv)(handle))
	case **Error:
		result = C.OCIHandleAlloc(
			unsafe.Pointer((*C.OCIEnv)(env)),
			&handle,
			C.OCI_HTYPE_ERROR,
			C.size_t(0),
			buffer,
		)
		*v = (*Error)((*C.OCIError)(handle))
	case **SvcCtx:
		result = C.OCIHandleAlloc(
			unsafe.Pointer((*C.OCIEnv)(env)),
			&handle,
			C.OCI_HTYPE_SVCCTX,
			C.size_t(0),
			buffer,
		)
		*v = (*SvcCtx)((*C.OCISvcCtx)(handle))
	case **Stmt:
		result = C.OCIHandleAlloc(
			unsafe.Pointer((*C.OCIEnv)(env)),
			&handle,
			C.OCI_HTYPE_STMT,
			C.size_t(0),
			buffer,
		)
		*v = (*Stmt)((*C.OCIStmt)(handle))
	case **Bind:
		result = C.OCIHandleAlloc(
			unsafe.Pointer((*C.OCIEnv)(env)),
			&handle,
			C.OCI_HTYPE_BIND,
			C.size_t(0),
			buffer,
		)
		*v = (*Bind)((*C.OCIBind)(handle))
	case **Define:
		result = C.OCIHandleAlloc(
			unsafe.Pointer((*C.OCIEnv)(env)),
			&handle,
			C.OCI_HTYPE_DEFINE,
			C.size_t(0),
			buffer,
		)
		*v = (*Define)((*C.OCIDefine)(handle))
	case **Describe:
		result = C.OCIHandleAlloc(
			unsafe.Pointer((*C.OCIEnv)(env)),
			&handle,
			C.OCI_HTYPE_DESCRIBE,
			C.size_t(0),
			buffer,
		)
		*v = (*Describe)((*C.OCIDescribe)(handle))
	case **Server:
		result = C.OCIHandleAlloc(
			unsafe.Pointer((*C.OCIEnv)(env)),
			&handle,
			C.OCI_HTYPE_SERVER,
			C.size_t(0),
			buffer,
		)
		*v = (*Server)((*C.OCIServer)(handle))
	case **Session:
		result = C.OCIHandleAlloc(
			unsafe.Pointer((*C.OCIEnv)(env)),
			&handle,
			C.OCI_HTYPE_SESSION,
			C.size_t(0),
			buffer,
		)
		*v = (*Session)((*C.OCISession)(handle))
	case **AuthInfo:
		result = C.OCIHandleAlloc(
			unsafe.Pointer((*C.OCIEnv)(env)),
			&handle,
			C.OCI_HTYPE_AUTHINFO,
			C.size_t(0),
			buffer,
		)
		*v = (*AuthInfo)((*C.OCIAuthInfo)(handle))
	case **CPool:
		result = C.OCIHandleAlloc(
			unsafe.Pointer((*C.OCIEnv)(env)),
			&handle,
			C.OCI_HTYPE_CPOOL,
			C.size_t(0),
			buffer,
		)
		*v = (*CPool)((*C.OCICPool)(handle))
	case **SPool:
		result = C.OCIHandleAlloc(
			unsafe.Pointer((*C.OCIEnv)(env)),
			&handle,
			C.OCI_HTYPE_SPOOL,
			C.size_t(0),
			buffer,
		)
		*v = (*SPool)((*C.OCISPool)(handle))
	case **Trans:
		result = C.OCIHandleAlloc(
			unsafe.Pointer((*C.OCIEnv)(env)),
			&handle,
			C.OCI_HTYPE_TRANS,
			C.size_t(0),
			buffer,
		)
		*v = (*Trans)((*C.OCITrans)(handle))
	case **ComplexObject:
		result = C.OCIHandleAlloc(
			unsafe.Pointer((*C.OCIEnv)(env)),
			&handle,
			C.OCI_HTYPE_COMPLEXOBJECT,
			C.size_t(0),
			buffer,
		)
		*v = (*ComplexObject)((*C.OCIComplexObject)(handle))
	default:
		return ResultInvalidHandle
	}
	return Result(result)
}
