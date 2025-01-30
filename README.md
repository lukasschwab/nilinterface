# nilinterface

Extremely simple `go/analysis`-based linter for forbidding passing `nil` as an
interface argument to function calls. The function expects something that
implements the interface; why are you passing something that'll panic if used?

## Usage

To install the linter, run

```
go install github.com/lukasschwab/nilinterface@latest
```

You can invoke it as `nilinterface`; in your Go project, run `nilinterface ./...`.

### From source

1. Clone this repository
2. Build the binary: `go build -o nilinterface`
3. Do with it what you will

## Examples

See the testdata directory for short examples of programs that violate the rule.
