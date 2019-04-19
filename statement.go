package gooci

// #include "gooci.h"
import "C"
import "fmt"

type Syntax int

const (
	SyntaxV7      = Syntax(C.OCI_V7_SYNTAX)
	SyntaxNTV     = Syntax(C.OCI_NTV_SYNTAX)
	SyntaxForeign = Syntax(C.OCI_FOREIGN_SYNTAX)
)

type StmtType int

const (
	StmtTypeSelect  = StmtType(C.OCI_STMT_SELECT)
	StmtTypeUpdate  = StmtType(C.OCI_STMT_UPDATE)
	StmtTypeDelete  = StmtType(C.OCI_STMT_DELETE)
	StmtTypeInsert  = StmtType(C.OCI_STMT_INSERT)
	StmtTypeCreate  = StmtType(C.OCI_STMT_CREATE)
	StmtTypeDrop    = StmtType(C.OCI_STMT_DROP)
	StmtTypeAlter   = StmtType(C.OCI_STMT_ALTER)
	StmtTypeBegin   = StmtType(C.OCI_STMT_BEGIN)
	StmtTypeDeclare = StmtType(C.OCI_STMT_DECLARE)
	StmtTypeCall    = StmtType(C.OCI_STMT_CALL)
)

func StmtPrepare2(
	svcp *SvcCtx,
	stmtpp **Stmt,
	errp *Error,
	stmttext string,
	key fmt.Stringer,
	language Syntax,
	mode Mode,
) Result {
	handle := (*C.OCIStmt)(*stmtpp)

	var cstrKey *C.uchar
	var keyLen int

	if nil != key {
		strKey := key.String()
		cstrKey = goStringToCString(strKey)
		keyLen = len(strKey)
	} else {
		cstrKey = nil
		keyLen = 0
	}
	cstrStmtText := goStringToCString(stmttext)

	result := C.OCIStmtPrepare2(
		(*C.OCISvcCtx)(svcp),
		&handle,
		(*C.OCIError)(errp),
		cstrStmtText,
		C.ub4(len(stmttext)),
		cstrKey,
		C.ub4(keyLen),
		C.ub4(language),
		C.ub4(mode),
	)
	*stmtpp = (*Stmt)(handle)
	return Result(result)
}

func StmtRelease(
	stmtp *Stmt,
	errp *Error,
	key fmt.Stringer,
	mode Mode,
) Result {
	var cstrKey *C.uchar
	var keyLen int

	if nil != key {
		strKey := key.String()
		cstrKey = goStringToCString(strKey)
		keyLen = len(strKey)
	} else {
		cstrKey = nil
		keyLen = 0
	}

	return Result(C.OCIStmtRelease(
		(*C.OCIStmt)(stmtp),
		(*C.OCIError)(errp),
		cstrKey,
		C.ub4(keyLen),
		C.ub4(mode),
	))
}
