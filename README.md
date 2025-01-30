# nilinterface

Extremely simple `go/analysis`-based linter for forbidding passing `nil` as an
interface argument to function calls. The function expects something that
implements the interface; why are you passing something that'll panic if used?

> [!WARNING]
> This package structure is currently overfit to serving as a golangci-lint module plugin, *not* as a standalone binary or with `go vet --vettool`. Watch this space.

<!-- ## Usage

To install the linter, run

```
go install github.com/lukasschwab/nilinterface/cmd/nilinterface@latest
```

You can invoke it as `nilinterface`; in your Go project, run `nilinterface ./...`.

### From source

1. Clone this repository
2. Build the binary: `go build cmd/nilinterface.go`
3. Do with it what you will -->

## Examples

See the testdata directory for short examples of programs that violate the rule.
