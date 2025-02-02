# nilinterface

Extremely simple `go/analysis`-based linter for forbidding [passing literal `nil` as an interface argument](https://go.dev/tour/methods/13) to function calls. The function expects something that implements the interface; why are you passing something that'll panic if used?

## Examples

See the testdata directory for short examples of programs that violate the rule.

## Usage

### Standalone

1. Install the binary; either run `go install github.com/lukasschwab/nilinterface@latest` or clone this repo and install from local source:

    ```bash
    go install ./cmd/standalone/nilinterface.go
    ```

2. Invoke `nilinterface` directly:

    ```bash
    nilinterface ./...
    ```

<details><summary>Example from project root</summary>

```bash
# Standalone binary installation
go install ./cmd/standalone/nilinterface.go
cd ./pkg/analyzer/testdata
# Invocation
go vet -vettool=$(which nilinterface) ./...
```

</details>

### With `go vet`

1. Follow the [Standalone instructions](#Standalone) to install a `nilinterface` binary.

2. Invoke `nilinterface` through `vet`:

    ```bash
    go vet -vettool=$(which nilinterface)
    ```

<details><summary>Example from project root</summary>

```bash
# Standalone binary installation
go install ./cmd/standalone/nilinterface.go
cd ./pkg/analyzer/testdata
# Invocation
go vet -vettool=$(which nilinterface) ./...
```

</details>

### With `golangci-lint`

- [ ] Basic local usage: building a custom binary and using it.
- [ ] Suggest a makefile to do lazy rebuild when config changes?
- [ ] *Briefly* mention the `*.so` shared binary plugin system; this requires building `golangci-lint` for a particular system *and* your plugin for the same, which seems tedious.
- [ ] Cover GitHub Actions cacheing strategy.
- [ ] Can you point VS Code at that custom binary?

```make
# NOTE: assumes custom-gcl.yml doesn't specify a different name.
custom-gcl: .custom-gcl.yml
	golangci-lint custom -v

lint: custom-gcl
	custom-gcl run
```
