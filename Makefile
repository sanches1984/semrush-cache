APP=semrush-cache
GO?=go

test:
	$(GO) test -v -race ./...

coverage:
	$(GO) test -v -covermode=atomic -coverpkg=./... -coverprofile=cover.out ./...

benchmark:
	$(GO) test -bench=. -benchmem ./...

modules:
	$(GO) mod tidy -v