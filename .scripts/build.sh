# Download our release binary builder
go get -u github.com/mitchellh/gox

# Specify platforms and release version
PLATFORMS="darwin/386 linux/amd64 linux/386 linux/arm"

# Build Inertia Go binaries for specified platforms
mkdir -p bin
gox \
  -output="bin/tezos-watcher.{{.OS}}.{{.Arch}}" \
  -osarch="$PLATFORMS" \
  github.com/bobheadxi/tezos-watcher/cmd/tezos-watcher
