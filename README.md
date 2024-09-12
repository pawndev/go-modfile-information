# Fetch go mod information

[![Build](https://github.com/pawndev/go-modfile-information/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/pawndev/go-modfile-information/actions/workflows/go.yml)
[![Integration Test](https://github.com/pawndev/go-modfile-information/actions/workflows/integration.yml/badge.svg?branch=main)](https://github.com/pawndev/go-modfile-information/actions/workflows/integration.yml)

## Summary

This action will let you retrieve some information in your `go.mod` file and export it in github action variables.
For now, only the go `version` and go `module` will be exported.

See the [examples](#examples) for how to use it

### Inputs

| Input     | Description                | Default value |
| --------- | -------------------------- | ------------- |
| `modfile` | An example mandatory input | go.mod        |

### Outputs

| Output       | Description              |
| ------------ | ------------------------ |
| `go_version` | Go version of the module |
| `go_module`  | Go module name           |

## Examples

Or see [this workflow](https://github.com/pawndev/go-modfile-information/blob/main/.github/workflows/go.yml)

```yaml
    ...
    - name: Get go mod info
      id: gomod
      uses: pawndev/go-modfile-information@v1.0.0
      with:
        modfile: go.mod # optional default to `go.mod`
    - name: Print go mod information
      run: |
        echo "${{ steps.gomod.outputs.go_version }}"
        echo "${{ steps.gomod.outputs.go_module }}"
```

### Using the optional input

```yaml
with:
  modfile: other-go.mod
```

## Testing locally

Be sure to have [act](https://github.com/nektos/act) locally and available il your `$PATH`.
And then you can `make test` to launch the project locally.

Folder `tests` contains another go.mod file with old version format using only major and minor without the patch number `1.21`.