package simplenet

import (
	"net"
	"os"
	"syscall"
)

type netSocket struct {
	fd int
}

func NewNetSocket(ip net.IP, port int) (*netSocket, error) {
	// ForkLock docs state that socket syscall requires the lock.
	syscall.ForkLock.Lock()
	// AF_INET = Address Family for IPv4
	// SOCK_STREAM = virtual circuit service
	// 0: the protocol for SOCK_STREAM, there's only 1.
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		return nil, os.NewSyscallError("socket", err)
	}
	syscall.ForkLock.Unlock()

	// Allow reuse of recently-used addresses.
	err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
	if err != nil {
		syscall.Close(fd)
		return nil, os.NewSyscallError("setsockopt", err)
	}

	// Step 2. Bind the socket to a port
	sa := &syscall.SockaddrInet4{Port: port}
	copy(sa.Addr[:], ip)
	err = syscall.Bind(fd, sa)
	if err != nil {
		return nil, os.NewSyscallError("bind", err)
	}

	// Step 3. Listen for incoming connections.
	err = syscall.Listen(fd, syscall.SOMAXCONN)
	if err != nil {
		return nil, os.NewSyscallError("listen", err)
	}
	return &netSocket{fd: fd}, nil
}

func (ns netSocket) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}

	n, err := syscall.Read(ns.fd, p) // read from socket file descriptor
	if err != nil {
		n = 0
	}

	return n, err
}

func (ns netSocket) Write(p []byte) (int, error) {
	n, err := syscall.Write(ns.fd, p) // write to socket file descriptor
	if err != nil {
		n = 0
	}

	return n, err
}

// Creates a new netSocket for the next pending connection request.
func (ns *netSocket) Accept() (*netSocket, error) {
	// syscall.ForkLock doc states lock not needed for blocking accept.
	nfd, _, err := syscall.Accept(ns.fd)
	if err == nil {
		syscall.CloseOnExec(nfd)
	}

	if err != nil {
		return nil, err
	}

	return &netSocket{nfd}, nil
}

func (ns *netSocket) Close() error {
	return syscall.Close(ns.fd)
}

//type request struct {
//	method string // GET, POST, etc.
//	header simpleTextProto.MIMEHeader
//	body   []byte
//	uri    string // The raw URI from the request
//	proto  string // "HTTP/1.1"
//}
//
//func ParseRequest(c *netSocket) (*request, error) {
//	b := bufio.NewReader(*c)
//	tp := simpleTextProto.NewReader(b) // need replace
//	req := new(request)
//
//	// Parse request line: parse "GET /index.html HTTP/1.0"
//	var s string
//	s, _ = tp.ReadLine() // need replace
//	sp := strings.Split(s, " ")
//	req.method, req.uri, req.proto = sp[0], sp[1], sp[2]
//
//	// Parse request headers
//	mimeHeader, _ := tp.ReadMIMEHeader() // need replace
//	req.header = mimeHeader
//
//	// Parse request body
//	if req.method == "GET" || req.method == "HEAD" {
//		return req, nil
//	}
//	if len(req.header["Content-Length"]) == 0 {
//		return nil, errors.New("no content length")
//	}
//	length, err := strconv.Atoi(req.header["Content-Length"][0])
//	if err != nil {
//		return nil, err
//	}
//	body := make([]byte, length)
//	if _, err = io.ReadFull(b, body); err != nil {
//		return nil, err
//	}
//	req.body = body
//	return req, nil
//}
