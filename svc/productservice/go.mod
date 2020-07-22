module github.com/ampazdev/simple-go-project/svc/productservice

go 1.13

require (
	github.com/ampazdev/simple-go-project/gopkg v0.0.0-20200516032755-9181205842d7
	github.com/ampazdev/simple-go-project/myproto v0.0.0-00010101000000-000000000000
	github.com/labstack/echo/v4 v4.1.16
	github.com/lib/pq v1.5.2
	go.uber.org/config v1.4.0
	go.uber.org/multierr v1.5.0 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/tools v0.0.0-20200520220537-cf2d1e09c845 // indirect
	google.golang.org/grpc v1.29.1
	gopkg.in/yaml.v2 v2.3.0 // indirect
	honnef.co/go/tools v0.0.1-2020.1.4 // indirect
)

replace github.com/ampazdev/simple-go-project/myproto => github.com/ampazdev/simple-go-project/svc/proto v0.0.0-20200611071026-0fe827aacb20
