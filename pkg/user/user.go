package user

import (
	"fmt"
	"io"
	"math/rand"
	"net"
)

// User represent user's conn and name
type User struct {
	io.Writer
	io.Reader
	conn net.Conn // connection
	name string   // username
}

// Conn returns user's connection.
func (u *User) Conn() net.Conn {
	return u.conn
}

// Name returns user's name.
func (u *User) Name() string {
	return u.name
}

// SetName change user's name.
func (u *User) SetName(name string) {
	if name == "" {
		name = fmt.Sprintf("user%d", rand.Intn(1000000))
	}
	u.name = name
}

func (u *User) Write(p []byte) (n int, err error) {
	return u.Conn().(io.Writer).Write(p)
}

func (u *User) Read(p []byte) (n int, err error) {
	return u.Conn().(io.Reader).Read(p)
}

// NewUser creates a new user with specified connection and name.
func NewUser(conn net.Conn, name string) *User {
	if "" == name {
		name = fmt.Sprintf("user%d", rand.Intn(1000000))
	}

	return &User{conn: conn, name: name}
}
