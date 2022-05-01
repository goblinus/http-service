package mysql

import (
	"fmt"
)

func NewDSN(user User, passwd Passwd, host Host, port Port, dbname DbName) DSN {
	return DSN(fmt.Sprintf("%s:%s@(%s:%s)/%s",
		user.String(),
		passwd.String(),
		host.String(),
		port.String(),
		dbname.String(),
	))
}

func NewUser(c *cli.Context) User {
	return User(c.String("dbase_user"))
}

func NewPasswd(c *cli.Context) Passwd {
	return Passwd(c.String("dbase_passwd"))
}

func NewHost(c *cli.Context) Host {
	return Host(c.String("dbase_host"))
}

func NewPort(c *cli.Context) Port {
	return Port(c.Int64("dbase_port"))
}

func NewDbName(c *cli.Context) DbName {
	return DbName(c.String("dbase_dbname"))
}

type DSN string
type User string
type Passwd string
type Host string
type Port int64
type DbName string

func (d DSN) String() string {
	return string(d)
}

func (u User) String() string {
	return string(u)
}

func (p Passwd) String() string {
	return string(p)
}

func (h Host) String() string {
	return string(h)
}

func (p Port) String() string {
	return string(p)
}

func (n DbName) String() string {
	return string(n)
}
