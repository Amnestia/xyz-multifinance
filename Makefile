build:
	go build -o bin/main cmd/xyz-multifinance/main.go

runapp:
	go run cmd/xyz-multifinance/main.go

clean:
	rm bin/*

run: test runapp

check: test

test:
	go test ./...

inithooks:
	cp ./script/hooks/pre-push .git/hooks/
	chmod +x .git/hooks/*
