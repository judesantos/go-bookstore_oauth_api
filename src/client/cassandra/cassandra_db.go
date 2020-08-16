package cassandra

// cassandra_db.go

import (
	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	// initialize cassandra session on startup
	InitSession()
}

func InitSession() {
	/*
	 * Initialize cassandra session
	 */

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	var err error

	// get session now!
	session, err = cluster.CreateSession()

	if err != nil {
		panic(err)
	}
}

// GetSession - Get Cassandra session
//
// If connection is closed after session is established,
// try and create new session. Discard old session.
//
func GetSession() *gocql.Session {
	// check if current db session is still valid
	if session != nil && true == session.Closed() {
		InitSession()
	}

	return session
}

func DestroySession() {
	if session != nil {
		session.Close()
	}
}
