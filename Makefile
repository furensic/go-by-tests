test:
	go test -v -cover ./...


pkgsite:
	pkgsite -http "localhost:5555" .
