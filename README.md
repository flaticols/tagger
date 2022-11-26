# Tagger

Creating a new release using the labels from the Pul Requests

## How to use

### Create label

There are 3 labels (comply with [semver](https://semver.org)) waiting for this action:

  - Major
  - Minor
  - Patch

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
      - uses: flaticols/tagger
```