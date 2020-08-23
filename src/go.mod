module github.com/judesantos/go-bookstore_oauth_api/src

go 1.14

require (
	github.com/federicoleon/golang-restclient v0.0.0-20191104170228-162ed620df66
	github.com/gin-gonic/gin v1.6.3
	github.com/gocql/gocql v0.0.0-20200815110948-5378c8f664e9
	github.com/golang/snappy v0.0.1 // indirect
	github.com/judesantos/go-bookstore_users_api v0.0.0-20200821011309-64ac7d73130a
	github.com/judesantos/go-bookstore_utils v0.0.0-20200821013116-f3a26ad513fc
	github.com/stretchr/testify v1.4.0
	golang.org/x/sys v0.0.0-20200821140526-fda516888d29 // indirect
)

//replace github.com/judesantos/go-bookstore_utils => /home/judesantos/dev/go/projects/bookstore/go-bookstore_utils
