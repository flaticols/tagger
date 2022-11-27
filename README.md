# Tagger

Creating a new release using the labels from the Pul Requests

## 🚀 How to use

To create a new release and tag, just add the labels `major`, `minor`, or `patch` to your Pull Request to increase the version and merge!

For example, if you have the latest tag  `v4.56.4` and you added the label `minor` the next version will be `v4.57.0`

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


