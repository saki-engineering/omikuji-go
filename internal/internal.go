/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (https://www.swig.org).
 * Version 4.1.1
 *
 * Do not make changes to this file unless you know what you are doing - modify
 * the SWIG interface file instead.
 * ----------------------------------------------------------------------------- */

// source: omikuji.i

package internal

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stddef.h>
#include <stdint.h>


typedef int intgo;
typedef unsigned int uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;



#cgo CPPFLAGS: -I${SRCDIR}/include
#cgo LDFLAGS: -L${SRCDIR}/lib -lomikuji

typedef _gostring_ swig_type_1;
typedef _gostring_ swig_type_2;
typedef _gostring_ swig_type_3;
typedef _gostring_ swig_type_4;
typedef _gostring_ swig_type_5;
typedef _gostring_ swig_type_6;
typedef _gostring_ swig_type_7;
typedef _gostring_ swig_type_8;
typedef _gostring_ swig_type_9;
extern void _wrap_Swig_free_internal_8d5c9551d044389f(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_internal_8d5c9551d044389f(swig_intgo arg1);
extern uintptr_t _wrap_new_PaperArray_internal_8d5c9551d044389f(swig_intgo arg1);
extern void _wrap_delete_PaperArray_internal_8d5c9551d044389f(uintptr_t arg1);
extern uintptr_t _wrap_PaperArray_getitem_internal_8d5c9551d044389f(uintptr_t arg1, swig_intgo arg2);
extern void _wrap_PaperArray_setitem_internal_8d5c9551d044389f(uintptr_t arg1, swig_intgo arg2, uintptr_t arg3);
extern swig_voidp _wrap_new_shortArray_internal_8d5c9551d044389f(swig_intgo arg1);
extern void _wrap_delete_shortArray_internal_8d5c9551d044389f(swig_voidp arg1);
extern short _wrap_shortArray_getitem_internal_8d5c9551d044389f(swig_voidp arg1, swig_intgo arg2);
extern void _wrap_shortArray_setitem_internal_8d5c9551d044389f(swig_voidp arg1, swig_intgo arg2, short arg3);
extern swig_voidp _wrap_new_strp_internal_8d5c9551d044389f(void);
extern swig_voidp _wrap_copy_strp_internal_8d5c9551d044389f(swig_type_1 arg1);
extern void _wrap_delete_strp_internal_8d5c9551d044389f(swig_voidp arg1);
extern void _wrap_strp_assign_internal_8d5c9551d044389f(swig_voidp arg1, swig_type_2 arg2);
extern swig_type_3 _wrap_strp_value_internal_8d5c9551d044389f(swig_voidp arg1);
extern swig_intgo _wrap_Error_NONE_internal_8d5c9551d044389f(void);
extern swig_intgo _wrap_Error_BREAK_BOX_internal_8d5c9551d044389f(void);
extern swig_intgo _wrap_Error_STUCKKED_internal_8d5c9551d044389f(void);
extern void _wrap_Paper_status_set_internal_8d5c9551d044389f(uintptr_t arg1, swig_type_4 arg2);
extern swig_type_5 _wrap_Paper_status_get_internal_8d5c9551d044389f(uintptr_t arg1);
extern void _wrap_Paper_description_set_internal_8d5c9551d044389f(uintptr_t arg1, swig_type_6 arg2);
extern swig_type_7 _wrap_Paper_description_get_internal_8d5c9551d044389f(uintptr_t arg1);
extern uintptr_t _wrap_new_Paper__SWIG_0_internal_8d5c9551d044389f(void);
extern uintptr_t _wrap_new_Paper__SWIG_1_internal_8d5c9551d044389f(swig_type_8 arg1, swig_type_9 arg2);
extern void _wrap_delete_Paper_internal_8d5c9551d044389f(uintptr_t arg1);
extern void _wrap_delete_OmikujiItf_internal_8d5c9551d044389f(uintptr_t arg1);
extern swig_intgo _wrap_OmikujiItf_DrawStatus_internal_8d5c9551d044389f(uintptr_t arg1, swig_voidp arg2);
extern swig_intgo _wrap_OmikujiItf_DrawPaper_internal_8d5c9551d044389f(uintptr_t arg1, uintptr_t arg2);
extern swig_intgo _wrap_OmikujiItf_MultiDrawPaper_internal_8d5c9551d044389f(uintptr_t arg1, uintptr_t arg2, swig_intgo arg3);
extern uintptr_t _wrap_OmikujiItf_CreateBox_internal_8d5c9551d044389f(void);
#undef intgo
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"


type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


func getSwigcptr(v interface { Swigcptr() uintptr }) uintptr {
	if v == nil {
		return 0
	}
	return v.Swigcptr()
}


type _ sync.Mutex

//export cgo_panic__internal_8d5c9551d044389f
func cgo_panic__internal_8d5c9551d044389f(p *byte) {
	s := (*[1024]byte)(unsafe.Pointer(p))[:]
	for i, b := range s {
		if b == 0 {
			panic(string(s[:i]))
		}
	}
	panic(string(s))
}


type swig_gostring struct { p uintptr; n int }
func swigCopyString(s string) string {
  p := *(*swig_gostring)(unsafe.Pointer(&s))
  r := string((*[0x7fffffff]byte)(unsafe.Pointer(p.p))[:p.n])
  Swig_free(p.p)
  return r
}

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_internal_8d5c9551d044389f(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_internal_8d5c9551d044389f(C.swig_intgo(_swig_i_0)))
	return swig_r
}

func New_PaperArray(arg1 int) (_swig_ret Paper) {
	var swig_r Paper
	_swig_i_0 := arg1
	swig_r = (Paper)(SwigcptrPaper(C._wrap_new_PaperArray_internal_8d5c9551d044389f(C.swig_intgo(_swig_i_0))))
	return swig_r
}

func Delete_PaperArray(arg1 Paper) {
	_swig_i_0 := getSwigcptr(arg1)
	C._wrap_delete_PaperArray_internal_8d5c9551d044389f(C.uintptr_t(_swig_i_0))
}

func PaperArray_getitem(arg1 Paper, arg2 int) (_swig_ret Paper) {
	var swig_r Paper
	_swig_i_0 := getSwigcptr(arg1)
	_swig_i_1 := arg2
	swig_r = (Paper)(SwigcptrPaper(C._wrap_PaperArray_getitem_internal_8d5c9551d044389f(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))))
	return swig_r
}

func PaperArray_setitem(arg1 Paper, arg2 int, arg3 Paper) {
	_swig_i_0 := getSwigcptr(arg1)
	_swig_i_1 := arg2
	_swig_i_2 := getSwigcptr(arg3)
	C._wrap_PaperArray_setitem_internal_8d5c9551d044389f(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.uintptr_t(_swig_i_2))
}

func New_shortArray(arg1 int) (_swig_ret *int16) {
	var swig_r *int16
	_swig_i_0 := arg1
	swig_r = (*int16)(C._wrap_new_shortArray_internal_8d5c9551d044389f(C.swig_intgo(_swig_i_0)))
	return swig_r
}

func Delete_shortArray(arg1 *int16) {
	_swig_i_0 := arg1
	C._wrap_delete_shortArray_internal_8d5c9551d044389f(C.swig_voidp(_swig_i_0))
}

func ShortArray_getitem(arg1 *int16, arg2 int) (_swig_ret int16) {
	var swig_r int16
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (int16)(C._wrap_shortArray_getitem_internal_8d5c9551d044389f(C.swig_voidp(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}

func ShortArray_setitem(arg1 *int16, arg2 int, arg3 int16) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	C._wrap_shortArray_setitem_internal_8d5c9551d044389f(C.swig_voidp(_swig_i_0), C.swig_intgo(_swig_i_1), C.short(_swig_i_2))
}

func New_strp() (_swig_ret *string) {
	var swig_r *string
	swig_r = (*string)(C._wrap_new_strp_internal_8d5c9551d044389f())
	return swig_r
}

func Copy_strp(arg1 string) (_swig_ret *string) {
	var swig_r *string
	_swig_i_0 := arg1
	swig_r = (*string)(C._wrap_copy_strp_internal_8d5c9551d044389f(*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_0))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

func Delete_strp(arg1 *string) {
	_swig_i_0 := arg1
	C._wrap_delete_strp_internal_8d5c9551d044389f(C.swig_voidp(_swig_i_0))
}

func Strp_assign(arg1 *string, arg2 string) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_strp_assign_internal_8d5c9551d044389f(C.swig_voidp(_swig_i_0), *(*C.swig_type_2)(unsafe.Pointer(&_swig_i_1)))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func Strp_value(arg1 *string) (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_strp_value_internal_8d5c9551d044389f(C.swig_voidp(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

type OmikujiError int
func _swig_getError_NONE() (_swig_ret OmikujiError) {
	var swig_r OmikujiError
	swig_r = (OmikujiError)(C._wrap_Error_NONE_internal_8d5c9551d044389f())
	return swig_r
}

var Error_NONE OmikujiError = _swig_getError_NONE()
func _swig_getError_BREAK_BOX() (_swig_ret OmikujiError) {
	var swig_r OmikujiError
	swig_r = (OmikujiError)(C._wrap_Error_BREAK_BOX_internal_8d5c9551d044389f())
	return swig_r
}

var Error_BREAK_BOX OmikujiError = _swig_getError_BREAK_BOX()
func _swig_getError_STUCKKED() (_swig_ret OmikujiError) {
	var swig_r OmikujiError
	swig_r = (OmikujiError)(C._wrap_Error_STUCKKED_internal_8d5c9551d044389f())
	return swig_r
}

var Error_STUCKKED OmikujiError = _swig_getError_STUCKKED()
type SwigcptrPaper uintptr

func (p SwigcptrPaper) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrPaper) SwigIsPaper() {
}

func (arg1 SwigcptrPaper) SetStatus(arg2 string) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_Paper_status_set_internal_8d5c9551d044389f(C.uintptr_t(_swig_i_0), *(*C.swig_type_4)(unsafe.Pointer(&_swig_i_1)))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (arg1 SwigcptrPaper) GetStatus() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_Paper_status_get_internal_8d5c9551d044389f(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrPaper) SetDescription(arg2 string) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_Paper_description_set_internal_8d5c9551d044389f(C.uintptr_t(_swig_i_0), *(*C.swig_type_6)(unsafe.Pointer(&_swig_i_1)))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (arg1 SwigcptrPaper) GetDescription() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_Paper_description_get_internal_8d5c9551d044389f(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func NewPaper__SWIG_0() (_swig_ret Paper) {
	var swig_r Paper
	swig_r = (Paper)(SwigcptrPaper(C._wrap_new_Paper__SWIG_0_internal_8d5c9551d044389f()))
	return swig_r
}

func NewPaper__SWIG_1(arg1 string, arg2 string) (_swig_ret Paper) {
	var swig_r Paper
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (Paper)(SwigcptrPaper(C._wrap_new_Paper__SWIG_1_internal_8d5c9551d044389f(*(*C.swig_type_8)(unsafe.Pointer(&_swig_i_0)), *(*C.swig_type_9)(unsafe.Pointer(&_swig_i_1)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func NewPaper(a ...interface{}) Paper {
	argc := len(a)
	if argc == 0 {
		return NewPaper__SWIG_0()
	}
	if argc == 2 {
		return NewPaper__SWIG_1(a[0].(string), a[1].(string))
	}
	panic("No match for overloaded function call")
}

func DeletePaper(arg1 Paper) {
	_swig_i_0 := getSwigcptr(arg1)
	C._wrap_delete_Paper_internal_8d5c9551d044389f(C.uintptr_t(_swig_i_0))
}

type Paper interface {
	Swigcptr() uintptr
	SwigIsPaper()
	SetStatus(arg2 string)
	GetStatus() (_swig_ret string)
	SetDescription(arg2 string)
	GetDescription() (_swig_ret string)
}

type SwigcptrOmikujiItf uintptr

func (p SwigcptrOmikujiItf) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrOmikujiItf) SwigIsOmikujiItf() {
}

func DeleteOmikujiItf(arg1 OmikujiItf) {
	_swig_i_0 := getSwigcptr(arg1)
	C._wrap_delete_OmikujiItf_internal_8d5c9551d044389f(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrOmikujiItf) DrawStatus(arg2 *string) (_swig_ret OmikujiError) {
	var swig_r OmikujiError
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (OmikujiError)(C._wrap_OmikujiItf_DrawStatus_internal_8d5c9551d044389f(C.uintptr_t(_swig_i_0), C.swig_voidp(_swig_i_1)))
	return swig_r
}

func (arg1 SwigcptrOmikujiItf) DrawPaper(arg2 Paper) (_swig_ret OmikujiError) {
	var swig_r OmikujiError
	_swig_i_0 := arg1
	_swig_i_1 := getSwigcptr(arg2)
	swig_r = (OmikujiError)(C._wrap_OmikujiItf_DrawPaper_internal_8d5c9551d044389f(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1)))
	return swig_r
}

func (arg1 SwigcptrOmikujiItf) MultiDrawPaper(arg2 Paper, arg3 int) (_swig_ret OmikujiError) {
	var swig_r OmikujiError
	_swig_i_0 := arg1
	_swig_i_1 := getSwigcptr(arg2)
	_swig_i_2 := arg3
	swig_r = (OmikujiError)(C._wrap_OmikujiItf_MultiDrawPaper_internal_8d5c9551d044389f(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1), C.swig_intgo(_swig_i_2)))
	return swig_r
}

func OmikujiItfCreateBox() (_swig_ret OmikujiItf) {
	var swig_r OmikujiItf
	swig_r = (OmikujiItf)(SwigcptrOmikujiItf(C._wrap_OmikujiItf_CreateBox_internal_8d5c9551d044389f()))
	return swig_r
}

type OmikujiItf interface {
	Swigcptr() uintptr
	SwigIsOmikujiItf()
	DrawStatus(arg2 *string) (_swig_ret OmikujiError)
	DrawPaper(arg2 Paper) (_swig_ret OmikujiError)
	MultiDrawPaper(arg2 Paper, arg3 int) (_swig_ret OmikujiError)
}


