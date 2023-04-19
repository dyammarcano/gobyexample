package main

import (
	"fmt"
	"net"
	"sync/atomic"
)

type ServerConn struct {
	Connection net.Conn
	ID string
	Open bool
}

func main() {
	p := atomic.Pointer[ServerConn]{}
	s := ServerConn{ ID : "first_conn"}
	p.Store( &s )
	fmt.Println(p.Load()) // Will display value stored.
}
