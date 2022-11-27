# Tagger

Creating a new release using the labels from the Pul Requests

## 🚀 How to use

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
 - `pr_number` - GitHub Token. Default `${{ github.event.number }}`
 - `default_tag` - GitHub Token. Default `v0.0.0`

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
      - uses: flaticols/tagger@v0.5.0
```

### Overrite inputs:

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
          pr_number: 10
          default_tag: '0.1.0'
      - uses: flaticols/tagger
```
