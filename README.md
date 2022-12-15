# Tagger 0.8

Creating a new release using the labels from the Pull Requests

## 🚀 How to use

Add a one of `major`, `minor` or `path` label to your PR and merge it.

-  **Major** - reset `minor` and `patch`
-  **Minor** - reset `patch`

**Examples**

##### Just add `patch` label to PR:
  With default settings you will get release: v0.0.1

##### Add `patch` and I have latest tag `v0.4.10`
  Result: release with tag and name `v0.4.11`

##### Add `Minor` and I have latest tag `v0.4.10`
  Result: release with tag and name `v0.5.0`

##### Add `Major` and I have latest tag `v0.4.10`
  Result: release with tag and name `v1.0.0`


### First run

#### Labels

There are 3 labels will be created or updated (comply with [semver](https://semver.org)):

  - Major
  - Minor
  - Patch

#### Tags

If your repository does not contain any tags, the next one after the `default_tag` will be created:

Default: `v0.0.0` -> first tag `v0.0.1`

### inputs

 - `github-token` - GitHub Token. Default `${{ github.token }}`
 - `pr-number` - GitHub Token. Default `${{ github.event.number }}`
 - `default-tag` - GitHub Token. Default `0.0.0`
 - `prefix` - Tag prefix. Default `v`

### Add new GitHub Action Workflow file:

```yaml
name: 'Release'

on:
  pull_request:
    types: [closed]

jobs:
  release:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: flaticols/tagger@v1.0.0
```

### Override inputs:

```yaml
name: 'Release'

on:
  pull_request:
    types: [closed]

jobs:
  release:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
        with:
          github-token: 'My token'
          pr-number: 10
          default-tag: '0.1.0'
          prefix: "ver"
      - uses: flaticols/tagger@v1.0.0
```

### Run locally

```bash
go run . create -o flaticols -r tagger -t $(gh auth token)
```

### Use as package

```go
package main

import C "github.com/flaticols/tagger/commands"

func main() {
	root := cobra.Command{}

	root.AddCommand(C.CreateCommand())

	_, err := root.ExecuteC()
	if err != nil {
		return
	}
}

```
