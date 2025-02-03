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

> [!WARNING]
> This guide describes using `golangci-lint`'s [Module Plugin System](https://golangci-lint.run/plugins/module-plugins), *not* the [Go plugin system](https://golangci-lint.run/plugins/go-plugins/). These are easily confused, and old documentation on `golangci-lint` extensibility doesn't explicitly distinguish the two.
>
> If you're troubleshooting, advice regarding `.so` files and the `CGO_ENABLED` build flag is likely specific to the *Go plugin system.*

This guide assumes you're already using `golangci-lint`. See also `golangci-lint`'s (currently crummy) docs: [Module Plugin System / The Automatic Way](https://golangci-lint.run/plugins/module-plugins/#the-automatic-way).

1. In your project, create a new `.custom-gcl.yml` file pinning a version of this linter:

    ```yaml
    version: v1.63.4
    plugins:
      - module: 'github.com/lukasschwab/nilinterface'
        import: 'github.com/lukasschwab/nilinterface/pkg/golangci-linter'
        version: v0.0.6
    ```

    For an example, see [.custom-gcl.yml](.custom-gcl.yml).

2. Build `custom-gcl`, a version of `golangci-lint` with `nilinterface` (and any others pinned in `.custom-gcl.yml`) available to it.

    ```bash
    golangci-lint custom -v
    ```

3. If you have a `.golangci.yml` configuration file, update it to enable the custom linter.

    ```yml
    linters-settings:
      custom:
        nilinterface:
          type: 'module'
          description: 'forbids passing `nil` as an interface argument to function calls'

    # If you have `disable-all: true`, add nilinterface to your enabled linters.
    linters:
      disable-all: true
      enable:
        - nilinterface
    ```

When you update `.custom-gcl.yml` (e.g. to pin a new linter version), you will need to rebuild `custom-gcl` to reflect those changes. If you use `make`, you can integrate the linter build step into your Makefile targets; see [Makefile](Makefile).

#### In VS Code

`custom-gcl` is a drop-in replacement for `golangci-lint`. You can specify the substitution in VS Code with the `go.alternateTools` setting. If `custom-gcl` is built in your VS Code workspace root:

```json
{
  "go.lintTool": "golangci-lint",
  "go.alternateTools": {
    "golangci-lint": "${workspaceFolder}/custom-gcl"
  },
}
```

#### In GitHub Actions

At time of writing, `golangci-lint`'s official GitHub Action doesn't support building and running custom linters.

+ Consider using [`golangci-lint-custom-plugins-action`](https://github.com/marketplace/actions/golangci-lint-custom).
+ To build your own, see @fr12k's [comment](https://github.com/golangci/golangci-lint-action/issues/1076#issuecomment-2624479984) suggesting a workaround custom action.
