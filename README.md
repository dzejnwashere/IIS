# ISS project

## How to install GO

1. Download on [this site](https://go.dev/dl/) go version `1.21.1`
2. Untar it to `/usr/local`

>`sudo tar -xvzf go1.21.1.linux-amd64.tar.gz -C /usr/local/`

4. Add `export PATH=/usr/local/go/bin:$PATH` to `.bashrc`
4. `source .bashrc`
5. `go version` to verify installation

## Running our project
> go run main.go