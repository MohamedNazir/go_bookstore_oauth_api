package cassendra

import (
	"time"

	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	// connect to the cassendra cluster

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 10
	//cluster.Authenticator = gocql.PasswordAuthenticator{Username: "Username", Password: "Password"} //replace the username and password fields with their real settings.

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}

}

func GetSession() *gocql.Session {

	return session
}
