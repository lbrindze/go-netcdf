// Copyright 2014 The Go-NetCDF Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// These files are autogenerated from nc_double.go using generate.sh

package netcdf

import (
	"fmt"
	"unsafe"
)

// #include <stdlib.h>
// #include <netcdf.h>
import "C"

// WriteFloat64s writes data as the entire data for variable v.
func (v Var) WriteFloat64s(data []float64) error {
	if err := okData(v, DOUBLE, len(data)); err != nil {
		return err
	}
	return newError(C.nc_put_var_double(C.int(v.ds), C.int(v.id), (*C.double)(unsafe.Pointer(&data[0]))))
}

// ReadFloat64s reads the entire variable v into data, which must have enough
// space for all the values (i.e. len(data) must be at least v.Len()).
func (v Var) ReadFloat64s(data []float64) error {
	if err := okData(v, DOUBLE, len(data)); err != nil {
		return err
	}
	return newError(C.nc_get_var_double(C.int(v.ds), C.int(v.id), (*C.double)(unsafe.Pointer(&data[0]))))
}

// WriteFloat64s sets the value of attribute a to val.
func (a Attr) WriteFloat64s(val []float64) error {
	// We don't need okData here because netcdf library doesn't know
	// the length or type of the attribute yet.
	cname := C.CString(a.name)
	defer C.free(unsafe.Pointer(cname))
	return newError(C.nc_put_att_double(C.int(a.v.ds), C.int(a.v.id), cname,
		C.nc_type(DOUBLE), C.size_t(len(val)), (*C.double)(unsafe.Pointer(&val[0]))))
}

// ReadFloat64s reads the entire attribute value into val.
func (a Attr) ReadFloat64s(val []float64) (err error) {
	if err := okData(a, DOUBLE, len(val)); err != nil {
		return err
	}
	cname := C.CString(a.name)
	defer C.free(unsafe.Pointer(cname))
	err = newError(C.nc_get_att_double(C.int(a.v.ds), C.int(a.v.id), cname,
		(*C.double)(unsafe.Pointer(&val[0]))))
	return
}

// Float64sReader is a interface that allows reading a sequence of values of fixed length.
type Float64sReader interface {
	Len() (n uint64, err error)
	ReadFloat64s(val []float64) (err error)
}

// GetFloat64s reads the entire data in r and returns it.
func GetFloat64s(r Float64sReader) (data []float64, err error) {
	n, err := r.Len()
	if err != nil {
		return
	}
	data = make([]float64, n)
	err = r.ReadFloat64s(data)
	return
}

// TestReadFloat64s writes somes data to v. N is v.Len().
// This function is only used for testing.
func testWriteFloat64s(v Var, n uint64) error {
	data := make([]float64, n)
	for i := 0; i < int(n); i++ {
		data[i] = float64(i + 10)
	}
	return v.WriteFloat64s(data)
}

// TestReadFloat64s reads data from v and checks that it's the same as what
// was written by testWriteDouble. N is v.Len().
// This function is only used for testing.
func testReadFloat64s(v Var, n uint64) error {
	data := make([]float64, n)
	if err := v.ReadFloat64s(data); err != nil {
		return err
	}
	for i := 0; i < int(n); i++ {
		if val := float64(i + 10); data[i] != val {
			return fmt.Errorf("data at position %d is %v; expected %v\n", i, data[i], val)
		}
	}
	return nil
}
