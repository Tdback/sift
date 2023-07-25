target_file := sift.go
exec_file := ./sift

sift:
	go mod tidy
	gofmt -w $(target_file)
	go build .

install:
	go mod init sift

.PHONY: clean
clean:
	$(RM) $(exec_file)
	$(RM) ./go.mod
