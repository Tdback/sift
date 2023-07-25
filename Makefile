sift:
	go mod tidy
	go build .

install:
	go mod init sift

.PHONY: clean
clean:
	$(RM) ./sift
	$(RM) ./go.mod
