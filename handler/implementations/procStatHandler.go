//
// Copyright: (C) 2019 Nestybox Inc.  All rights reserved.
//

package implementations

import (
	"io"
	"os"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/nestybox/sysbox-fs/domain"
	"github.com/nestybox/sysbox-fs/fuse"
)

//
// /proc/stat Handler
//
type ProcStatHandler struct {
	Name      string
	Path      string
	Type      domain.HandlerType
	Enabled   bool
	Cacheable bool
	Service   domain.HandlerService
}

func (h *ProcStatHandler) Lookup(
	n domain.IOnode,
	req *domain.HandlerRequest) (os.FileInfo, error) {

	logrus.Debugf("Executing Lookup() method on %v handler", h.Name)

	return n.Stat()
}

func (h *ProcStatHandler) Getattr(
	n domain.IOnode,
	req *domain.HandlerRequest) (*syscall.Stat_t, error) {

	logrus.Debugf("Executing Getattr() method on %v handler", h.Name)

	return nil, nil
}

func (h *ProcStatHandler) Open(
	n domain.IOnode,
	req *domain.HandlerRequest) error {

	logrus.Debugf("Executing %v Open() method", h.Name)

	flags := n.OpenFlags()
	if flags != syscall.O_RDONLY {
		return fuse.IOerror{Code: syscall.EACCES}
	}

	if err := n.Open(); err != nil {
		logrus.Debugf("Error opening file %v", h.Path)
		return fuse.IOerror{Code: syscall.EIO}
	}

	return nil
}

func (h *ProcStatHandler) Close(n domain.IOnode) error {

	logrus.Debugf("Executing Close() method on %v handler", h.Name)

	if err := n.Close(); err != nil {
		logrus.Debugf("Error closing file %v", h.Path)
		return fuse.IOerror{Code: syscall.EIO}
	}

	return nil
}

func (h *ProcStatHandler) Read(
	n domain.IOnode,
	req *domain.HandlerRequest) (int, error) {

	logrus.Debugf("Executing %v Read() method", h.Name)

	// Bypass emulation logic for now by going straight to host fs.
	ios := h.Service.IOService()
	len, err := ios.ReadNode(n, req.Data)
	if err != nil && err != io.EOF {
		return 0, err
	}

	req.Data = req.Data[:len]

	return len, nil
}

func (h *ProcStatHandler) Write(
	n domain.IOnode,
	req *domain.HandlerRequest) (int, error) {

	return 0, nil
}

func (h *ProcStatHandler) ReadDirAll(
	n domain.IOnode,
	req *domain.HandlerRequest) ([]os.FileInfo, error) {

	return nil, nil
}

func (h *ProcStatHandler) GetName() string {
	return h.Name
}

func (h *ProcStatHandler) GetPath() string {
	return h.Path
}

func (h *ProcStatHandler) GetEnabled() bool {
	return h.Enabled
}

func (h *ProcStatHandler) GetType() domain.HandlerType {
	return h.Type
}

func (h *ProcStatHandler) GetService() domain.HandlerService {
	return h.Service
}

func (h *ProcStatHandler) SetEnabled(val bool) {
	h.Enabled = val
}

func (h *ProcStatHandler) SetService(hs domain.HandlerService) {
	h.Service = hs
}
