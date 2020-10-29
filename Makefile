test:
	go test -v
build:
	go build -o grubAPI main.go api.go env.go cart.go order.go random.go countdown.go slack.go
run:
	go run main.go api.go env.go cart.go order.go random.go countdown.go slack.go