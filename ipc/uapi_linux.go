package ipc

import (
	"errors"
	"fmt"
	"net"
	"os"
	"path"

	"github.com/caleb072350/wireguard-go/rwcancel"
	"golang.org/x/sys/unix"
)

var socketDirectory = "/var/run/wireguard"

const (
	IpcErrorIO        = -int64(unix.EIO)
	IpcErrorProtocol  = -int64(unix.EPROTO)
	IpcErrorInvalid   = -int64(unix.EINVAL)
	IpcErrorPortInUse = -int64(unix.EADDRINUSE)
	socketName        = "%s.sock"
)

type UAPIListener struct {
	listener        net.Listener //unix socket listener
	connNew         chan net.Conn
	connErr         chan error
	inotifyFd       int
	inotifyRWCancel *rwcancel.RWCancel
}

func UAPIOpen(name string) (*os.File, error) {
	// check if path exist
	err := os.MkdirAll(socketDirectory, 0755)
	if err != nil && !os.IsExist(err) {
		return nil, err
	}

	// open UNIX socket
	socketPath := path.Join(
		socketDirectory,
		fmt.Sprintf(socketName, name),
	)

	addr, err := net.ResolveUnixAddr("unix", socketPath)
	if err != nil {
		return nil, err
	}

	oldUmask := unix.Umask(0077)
	listener, err := func() (*net.UnixListener, error) {
		// initial connection attempt
		listener, err := net.ListenUnix("unix", addr)
		if err == nil {
			return listener, nil
		}

		// check if socket already active

		_, err = net.Dial("unix", socketPath)
		if err == nil {
			return nil, errors.New("unix socket in use")
		}

		// cleanup & attempt again

		err = os.Remove(socketPath)
		if err != nil {
			return nil, err
		}
		return net.ListenUnix("unix", addr)
	}()
	unix.Umask(oldUmask)

	if err != nil {
		return nil, err
	}
	return listener.File()
}
