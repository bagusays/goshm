run:
	go build -o goshm

mockery:
	mockery --all

test:
	go test ./... -cover