package gooci

// #include "gooci.h"
import "C"

func SessionBegin(
	svc *SvcCtx,
	e *Error,
	usr *Session,
	cred Cred,
	mode Mode,
) error {
	return nil
}
