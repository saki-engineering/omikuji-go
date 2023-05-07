%module internal
%{
#include "errorCode.h"
#include "omikuji.h"
%}

%insert(cgo_comment_typedefs) %{
#cgo CPPFLAGS: -I${SRCDIR}/include
#cgo LDFLAGS: -L${SRCDIR}/lib -lomikuji
%}

%include "carrays.i"
%array_functions(omikuji::Paper, PaperArray);
%array_functions(short, shortArray);

%include "cpointer.i"
%pointer_functions(char * , strp)

%include "errorCode.h"
%include "omikuji.h"