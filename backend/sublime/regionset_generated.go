// Copyright 2013 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

// This file was generated as part of a build step and shouldn't be manually modified
package sublime

import (
	"fmt"
	"github.com/limetext/gopy/lib"
	"github.com/limetext/lime/backend"
	"github.com/quarnster/util/text"
)

var (
	_ = backend.View{}
	_ = text.Region{}
	_ = fmt.Errorf
)

var _region_setClass = py.Class{
	Name:    "sublime.RegionSet",
	Pointer: (*RegionSet)(nil),
}

type RegionSet struct {
	py.BaseObject
	data *text.RegionSet
}

func (o *RegionSet) PyInit(args *py.Tuple, kwds *py.Dict) error {
	return fmt.Errorf("Can't initialize type RegionSet")
}
func (o *RegionSet) Py_add(tu *py.Tuple) (py.Object, error) {
	var (
		arg1 text.Region
	)
	if v, err := tu.GetItem(0); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(text.Region); !ok {
				return nil, fmt.Errorf("Expected type text.Region for text.RegionSet.Add() arg1, not %s", v.Type())
			} else {
				arg1 = v2
			}
		}
	}
	o.data.Add(arg1)
	return toPython(nil)
}

func (o *RegionSet) Py_add_all(tu *py.Tuple) (py.Object, error) {
	var (
		arg1 []text.Region
	)
	if v, err := tu.GetItem(0); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.([]text.Region); !ok {
				return nil, fmt.Errorf("Expected type []text.Region for text.RegionSet.AddAll() arg1, not %s", v.Type())
			} else {
				arg1 = v2
			}
		}
	}
	o.data.AddAll(arg1)
	return toPython(nil)
}

func (o *RegionSet) Py_clear() (py.Object, error) {
	o.data.Clear()
	return toPython(nil)
}

func (o *RegionSet) Py_contains(tu *py.Tuple) (py.Object, error) {
	var (
		arg1 text.Region
	)
	if v, err := tu.GetItem(0); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(text.Region); !ok {
				return nil, fmt.Errorf("Expected type text.Region for text.RegionSet.Contains() arg1, not %s", v.Type())
			} else {
				arg1 = v2
			}
		}
	}
	ret0 := o.data.Contains(arg1)
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *RegionSet) PySeqGet(arg0 int64) (py.Object, error) {
	var (
		pyret0 py.Object
		err    error
	)
	if l := o.data.Len(); int(arg0) >= l || arg0 < 0 {
		return nil, py.NewError(py.IndexError, "%d >= %d || %d < 0", arg0, l, arg0)
	}

	ret0 := o.data.Get(int(arg0))
	pyret0, err = toPython(ret0)
	if err != nil {
		// TODO: do the py objs need to be freed?
		return nil, err
	}
	return pyret0, err
}

func (o *RegionSet) PySeqLen() int64 {
	return int64(o.data.Len())
}

func (o *RegionSet) Py_regions() (py.Object, error) {
	ret0 := o.data.Regions()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *RegionSet) Py_substract(tu *py.Tuple) (py.Object, error) {
	var (
		arg1 text.Region
	)
	if v, err := tu.GetItem(0); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(text.Region); !ok {
				return nil, fmt.Errorf("Expected type text.Region for text.RegionSet.Substract() arg1, not %s", v.Type())
			} else {
				arg1 = v2
			}
		}
	}
	o.data.Substract(arg1)
	return toPython(nil)
}
