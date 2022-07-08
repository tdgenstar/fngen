go test -json -coverprofile cover.out ./... -cover
go tool cover -html=cover.out -o cover-count.html

open http://localhost:63342/fn/cover-count.html#file1