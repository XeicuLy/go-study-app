.PHONY: setup fmt test fix

setup:
	mise install
	mise exec -- pnpm install --frozen-lockfile
	mise exec -- go mod download

fmt:
	mise exec -- pnpm run fmt && mise exec -- go fmt ./...

test:
	mise exec -- go test ./...

fix: fmt
	mise exec -- go vet ./...
