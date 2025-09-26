package iface

import (
	"fmt"
	"unsafe"
)

type Eface struct {
	_type unsafe.Pointer
	data  unsafe.Pointer
}

type Iface struct {
	tab  unsafe.Pointer
	data unsafe.Pointer
}

func ShowEface(i interface{}) {
	ef := (*Eface)(unsafe.Pointer(&i))
	fmt.Printf("Eface: type=%p data=%p\n", ef._type, ef.data)
}
