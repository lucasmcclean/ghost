build:
    go build -o ghost ./cmd/ghost/

run CONTAINER="":
    sudo ./ghost {{CONTAINER}}

test CONTAINER="": build
    sudo ./ghost {{CONTAINER}}

lint:
    go vet ./...

clean:
    rm -f ghost

fmt:
    gofumpt -w .
