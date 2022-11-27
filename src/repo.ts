import * as core from '@actions/core'
import * as gh from '@actions/github'
import {Label, RepositoryTags, SemVerUpdate, getInputs} from './inputs'

const TAG_PATTERN = /v([0-9])\.([0-9])\.([0-9])/gi
const LABELS: {name: string; description: string; color: string}[] = [
  {name: 'Major', description: 'Major version update. Minor and Patch will be reset', color: 'FE9B2B'},
  {name: 'Minor', description: 'Minor version update. Patch will be reset', color: 'C4D174'},
  {name: 'Patch', description: 'Patch version update', color: '836C76'}
]

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

  core.notice(JSON.stringify(refTags))
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
    overwrite: true,
    prerelease: false,
    generate_release_notes: false
  })
}

export async function createLabels(): Promise<void> {
  try {
    for (const label of LABELS) {
      await createLabel(label.name, label.description, label.color)
    }
  } catch (e) {
    core.warning(`Failed to create labels: ${e}`)
  }
}

async function createLabel(name: string, description: string, color: string): Promise<void> {
  await octokit.rest.issues.createLabel({
    owner: gh.context.repo.owner,
    repo: gh.context.repo.repo,
    name,
    description,
    color,
    overwrite: true
  })
}
