import * as gh from '@actions/github'
import {Label, RepositoryTags, SemVerUpdate, getInputs} from './inputs'

const TAG_PATTERN = /v([0-9])\.([0-9])\.([0-9])/gi

const {token, pr_number} = getInputs()
const octokit = gh.getOctokit(token)

// eslint-disable-next-line @typescript-eslint/explicit-function-return-type
export async function getPR() {
  return octokit.rest.pulls.get({
    owner: gh.context.repo.owner,
    repo: gh.context.repo.repo,
    pull_number: pr_number
  })
}

export function getRequiredLabels(labels: Label[]): SemVerUpdate {
  const semver: SemVerUpdate = {
    Major: false,
    Minor: false,
    Patch: false
  }

  for (const label of labels) {
    if (label.name.toLowerCase() === 'major') {
      semver.Major = true
    } else if (label.name.toLowerCase() === 'minor') {
      semver.Minor = true
    } else if (label.name.toLowerCase() === 'patch') {
      semver.Patch = true
    }
  }

  return semver
}

export async function getLatestTag(): Promise<string[]> {
  const refTags = await octokit.graphql<RepositoryTags>(`
    query {
      repository(owner: "${gh.context.repo.owner}", name: "${gh.context.repo.repo}") {
        refs(refPrefix: "refs/tags/", last: 1) {
          edges {
            node {
              name
            }
          }
        }
      }
    }`)

  const tags = refTags.repository.refs.edges.map(x => x.node.name)

  if (tags.length === 0) {
    tags.push('0.0.0')
  }

  return tags
}

export function getNewTag(ver: SemVerUpdate, tags: string[]): string {
  // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
  const parts = TAG_PATTERN.exec(tags[0])!

  let major = parseInt(parts[1])
  let minor = parseInt(parts[2])
  let patch = parseInt(parts[3])

  if (ver.Major) {
    major++
  }

  if (ver.Minor) {
    minor++
  }

  if (ver.Patch) {
    patch++
  }

  return `${major}.${minor}.${patch}`
}

export async function createRelease(tag: string): Promise<void> {
  await octokit.rest.repos.createRelease({
    owner: gh.context.repo.owner,
    repo: gh.context.repo.repo,
    tag_name: tag,
    name: tag,
    draft: false,
    prerelease: false,
    generate_release_notes: false
  })
}
