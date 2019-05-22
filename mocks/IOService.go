// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import domain "github.com/nestybox/sysvisor/sysvisor-fs/domain"
import mock "github.com/stretchr/testify/mock"
import os "os"

// IOService is an autogenerated mock type for the IOService type
type IOService struct {
	mock.Mock
}

// CloseNode provides a mock function with given fields: i
func (_m *IOService) CloseNode(i domain.IOnode) error {
	ret := _m.Called(i)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.IOnode) error); ok {
		r0 = rf(i)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIOnode provides a mock function with given fields: n, p, attr
func (_m *IOService) NewIOnode(n string, p string, attr os.FileMode) domain.IOnode {
	ret := _m.Called(n, p, attr)

	var r0 domain.IOnode
	if rf, ok := ret.Get(0).(func(string, string, os.FileMode) domain.IOnode); ok {
		r0 = rf(n, p, attr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.IOnode)
		}
	}

	return r0
}

// OpenNode provides a mock function with given fields: i
func (_m *IOService) OpenNode(i domain.IOnode) error {
	ret := _m.Called(i)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.IOnode) error); ok {
		r0 = rf(i)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PathNode provides a mock function with given fields: i
func (_m *IOService) PathNode(i domain.IOnode) string {
	ret := _m.Called(i)

	var r0 string
	if rf, ok := ret.Get(0).(func(domain.IOnode) string); ok {
		r0 = rf(i)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// PidNsInode provides a mock function with given fields: i
func (_m *IOService) PidNsInode(i domain.IOnode) (uint64, error) {
	ret := _m.Called(i)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(domain.IOnode) uint64); ok {
		r0 = rf(i)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.IOnode) error); ok {
		r1 = rf(i)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadAtNode provides a mock function with given fields: i, p, off
func (_m *IOService) ReadAtNode(i domain.IOnode, p []byte, off int64) (int, error) {
	ret := _m.Called(i, p, off)

	var r0 int
	if rf, ok := ret.Get(0).(func(domain.IOnode, []byte, int64) int); ok {
		r0 = rf(i, p, off)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.IOnode, []byte, int64) error); ok {
		r1 = rf(i, p, off)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadDirAllNode provides a mock function with given fields: i
func (_m *IOService) ReadDirAllNode(i domain.IOnode) ([]os.FileInfo, error) {
	ret := _m.Called(i)

	var r0 []os.FileInfo
	if rf, ok := ret.Get(0).(func(domain.IOnode) []os.FileInfo); ok {
		r0 = rf(i)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]os.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.IOnode) error); ok {
		r1 = rf(i)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadLineNode provides a mock function with given fields: i
func (_m *IOService) ReadLineNode(i domain.IOnode) string {
	ret := _m.Called(i)

	var r0 string
	if rf, ok := ret.Get(0).(func(domain.IOnode) string); ok {
		r0 = rf(i)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ReadNode provides a mock function with given fields: i, p
func (_m *IOService) ReadNode(i domain.IOnode, p []byte) (int, error) {
	ret := _m.Called(i, p)

	var r0 int
	if rf, ok := ret.Get(0).(func(domain.IOnode, []byte) int); ok {
		r0 = rf(i, p)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.IOnode, []byte) error); ok {
		r1 = rf(i, p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SeekResetNode provides a mock function with given fields: i
func (_m *IOService) SeekResetNode(i domain.IOnode) (int64, error) {
	ret := _m.Called(i)

	var r0 int64
	if rf, ok := ret.Get(0).(func(domain.IOnode) int64); ok {
		r0 = rf(i)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.IOnode) error); ok {
		r1 = rf(i)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StatNode provides a mock function with given fields: i
func (_m *IOService) StatNode(i domain.IOnode) (os.FileInfo, error) {
	ret := _m.Called(i)

	var r0 os.FileInfo
	if rf, ok := ret.Get(0).(func(domain.IOnode) os.FileInfo); ok {
		r0 = rf(i)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(os.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.IOnode) error); ok {
		r1 = rf(i)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WriteNode provides a mock function with given fields: i, p
func (_m *IOService) WriteNode(i domain.IOnode, p []byte) (int, error) {
	ret := _m.Called(i, p)

	var r0 int
	if rf, ok := ret.Get(0).(func(domain.IOnode, []byte) int); ok {
		r0 = rf(i, p)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.IOnode, []byte) error); ok {
		r1 = rf(i, p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}