#
# sysbox-fs Makefile
#
# Note: targets must execute from the $SYSFS_DIR

.PHONY: clean sysbox-fs-debug sysbox-fs-static lint list-packages

GO := go
ARCH := amd64

SYSFS_BUILDROOT := build
SYSFS_BUILDDIR := $(SYSFS_BUILDROOT)/$(ARCH)
SYSFS_TARGET := sysbox-fs
SYSFS_DEBUG_TARGET := sysbox-fs-debug
SYSFS_STATIC_TARGET := sysbox-fs-static
SYSFS_DIR := $(CURDIR)
SYSFS_SRC := $(shell find . 2>&1 | grep -E '.*\.(c|h|go)$$')

SYSIPC_DIR := ../sysbox-ipc
SYSIPC_SRC := $(shell find $(SYSIPC_DIR) 2>&1 | grep -E '.*\.(c|h|go|proto)$$')

LIBSECCOMP_DIR := ../lib/seccomp-
LIBSECCOMP_SRC := $(shell find $(LIBSECCOMP_DIR) 2>&1 | grep -E '.*\.(go)')

LIBPIDMON_DIR := ../sysbox-libs/pidmonitor
LIBSPIDMON_SRC := $(shell find $(LIBPIDMON_DIR) 2>&1 | grep -E '.*\.(go)')

NSENTER_DIR := ../sysbox-runc/libcontainer/nsenter
NSENTER_SRC := $(shell find $(NSENTER_DIR) 2>&1 | grep -E '.*\.(c|h|go)')

COMMIT_NO := $(shell git rev-parse HEAD 2> /dev/null || true)
COMMIT ?= $(if $(shell git status --porcelain --untracked-files=no),$(COMMIT_NO)-dirty,$(COMMIT_NO))
BUILT_AT := $(shell date)
BUILT_BY := $(shell git config user.name)

LDFLAGS := '-X "main.edition=${EDITION}" -X main.version=${VERSION} \
		-X main.commitId=$(COMMIT) -X "main.builtAt=$(BUILT_AT)" \
		-X "main.builtBy=$(BUILT_BY)"'

ifeq ($(ARCH),armel)
	GO_XCOMPILE := CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=6 CC=arm-linux-gnueabi-gcc
else ifeq ($(ARCH),armhf)
	GO_XCOMPILE := CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=7 CC=arm-linux-gnueabihf-gcc
else ifeq ($(ARCH),arm64)
	GO_XCOMPILE = CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC=aarch64-linux-gnu-gcc
else
	GO_XCOMPILE = GOARCH=amd64
endif

.DEFAULT: sysbox-fs

sysbox-fs: $(SYSFS_BUILDDIR)/$(SYSFS_TARGET)

$(SYSFS_BUILDDIR)/$(SYSFS_TARGET): $(SYSFS_SRC) $(SYSIPC_SRC) $(LIBSECCOMP_SRC) $(LIBPIDMON_SRC) $(NSENTER_SRC)
	$(GO_XCOMPILE) $(GO) build -ldflags ${LDFLAGS}	-o $(SYSFS_BUILDDIR)/sysbox-fs ./cmd/sysbox-fs

sysbox-fs-debug: $(SYSFS_BUILDDIR)/$(SYSFS_DEBUG_TARGET)

$(SYSFS_BUILDDIR)/$(SYSFS_DEBUG_TARGET): $(SYSFS_SRC) $(SYSIPC_SRC) $(LIBSECCOMP_SRC) $(LIBPIDMON_SRC) $(NSENTER_SRC)
	$(GO_XCOMPILE) $(GO) build -gcflags="all=-N -l" -ldflags ${LDFLAGS} -o $(SYSFS_BUILDDIR)/sysbox-fs ./cmd/sysbox-fs

sysbox-fs-static: $(SYSFS_BUILDDIR)/$(SYSFS_STATIC_TARGET)

$(SYSFS_BUILDDIR)/$(SYSFS_STATIC_TARGET): $(SYSFS_SRC) $(SYSIPC_SRC) $(LIBSECCOMP_SRC) $(LIBPIDMON_SRC) $(NSENTER_SRC)
	$(GO_XCOMPILE) CGO_ENABLED=1 $(GO) build -tags "netgo osusergo static_build" \
		-installsuffix netgo -ldflags "-w -extldflags -static" -ldflags ${LDFLAGS} \
		-o $(SYSFS_BUILDDIR)/sysbox-fs ./cmd/sysbox-fs

lint:
	$(GO) vet $(allpackages)
	$(GO) fmt $(allpackages)

listpackages:
	@echo $(allpackages)

clean:
	rm -f $(SYSFS_BUILDROOT)/sysbox-fs

distclean: clean
	rm -rf $(SYSFS_BUILDROOT)

# memoize allpackages, so that it's executed only once and only if used
_allpackages = $(shell $(GO) list ./... | grep -v vendor)
allpackages = $(if $(__allpackages),,$(eval __allpackages := $$(_allpackages)))$(__allpackages)
