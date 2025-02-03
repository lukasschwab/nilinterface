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
    - [ ] Pointing VS Code at your custom binary.
- [ ] Suggest a makefile to do lazy rebuild when config changes?
- [ ] *Briefly* mention the `*.so` shared binary plugin system; this requires building `golangci-lint` for a particular system *and* your plugin for the same, which seems tedious.
    + Any instructions about `*.so` files or the `CGO_ENABLED=1` flag are probably about the Go plugin system, *not* the module plugin system;  disregard them. Old references don't distinguish, because the Go plugin system came first. [Example](https://github.com/golangci/golangci-lint-action-plugin-example)
- [ ] Cover GitHub Actions cacheing strategy.
- [ ] Consider demoing in this repository; too confusing?

```yaml
version: v1.63.4
plugins:
  - module: 'github.com/lukasschwab/nilinterface'
    import: 'github.com/lukasschwab/nilinterface/pkg/golangci-linter'
    version: v0.0.6
```

```make
.PHONY: lint
lint: custom-gcl
	./custom-gcl run

custom-gcl: .custom-gcl.yml
	golangci-lint custom -v
```
