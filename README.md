# Fetch go mod information

[![Action Template](https://img.shields.io/badge/Action%20Template-Go%20Container%20Action-blue.svg?colorA=24292e&colorB=0366d6&style=flat&longCache=true&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAA4AAAAOCAYAAAAfSC3RAAAABHNCSVQICAgIfAhkiAAAAAlwSFlzAAAM6wAADOsB5dZE0gAAABl0RVh0U29mdHdhcmUAd3d3Lmlua3NjYXBlLm9yZ5vuPBoAAAERSURBVCiRhZG/SsMxFEZPfsVJ61jbxaF0cRQRcRJ9hlYn30IHN/+9iquDCOIsblIrOjqKgy5aKoJQj4O3EEtbPwhJbr6Te28CmdSKeqzeqr0YbfVIrTBKakvtOl5dtTkK+v4HfA9PEyBFCY9AGVgCBLaBp1jPAyfAJ/AAdIEG0dNAiyP7+K1qIfMdonZic6+WJoBJvQlvuwDqcXadUuqPA1NKAlexbRTAIMvMOCjTbMwl1LtI/6KWJ5Q6rT6Ht1MA58AX8Apcqqt5r2qhrgAXQC3CZ6i1+KMd9TRu3MvA3aH/fFPnBodb6oe6HM8+lYHrGdRXW8M9bMZtPXUji69lmf5Cmamq7quNLFZXD9Rq7v0Bpc1o/tp0fisAAAAASUVORK5CYII=)](https://github.com/pawndev/go-modfile-information)
[![Actions Status](https://github.com/pawndev/go-modfile-information/workflows/Build/badge.svg)](https://github.com/pawndev/go-modfile-information/actions)
[![Actions Status](https://github.com/pawndev/go-modfile-information/workflows/Integration%20Test/badge.svg)](https://github.com/pawndev/go-modfile-information/actions)

## Summary

This action will let you retrieve some information in your `go.mod` file and export it in github action variables.
For now, only the go `version` and go `module` will be exported.

See the [examples](#examples) for how to use it

### Inputs

| Input                       | Description                     | Default value |
|-----------------------------|---------------------------------|---------------|
| `modfile`                   | An example mandatory input      | go.mod        |

### Outputs

| Output       | Description              |
|--------------|--------------------------|
| `go_version` | Go version of the module |
| `go_module`  | Go module name           |

## Examples

```yaml
    ...
    - name: Get go mod info
      id: gomod
      uses: pawndev/go-modfile-information@v1
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