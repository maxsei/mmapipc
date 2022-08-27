default: run
fmt:
	go run golang.org/x/tools/cmd/goimports -w .
	go fmt
run: fmt
	go run main.go
clean:
	ipcs -m | awk '{if($$6 == 0){print $$2}}' | tee /dev/stderr | xargs -L 1 ipcrm -m
