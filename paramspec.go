package glib

/*
#include <glib-object.h>
*/
import "C"


type ParamSpec struct {
	p C.gpointer
}

func (p *ParamSpec) g() *C.GParamSpec {
	return (*C.GParamSpec)(p.p)
}

func (p *ParamSpec) SetPtr (ptr Pointer) {
	p.p = C.gpointer (ptr)
}

func (p *ParamSpec) ValueType () Type {
	return Type (p.g().value_type)
}

func (p *ParamSpec) String() string {
	return C.GoString ((*C.char) (p.g().name))
}

func (p *ParamSpec) Name() string {
	return p.String()
}
