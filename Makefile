test:
	go test -v
build:
	go build -o grub main.go api.go env.go cart.go order.go random.go countdown.go
run:
	go run main.go api.go env.go cart.go order.go random.go countdown.go
container:
	docker run --publish 3333:3333 --name grub --rm grub