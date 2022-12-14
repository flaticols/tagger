import * as core from '@actions/core'
import * as gh from '@actions/github'
import {Label, RepositoryTags, SemVerUpdate, getInputs} from './inputs'

const TAG_PATTERN = /v([0-9]+)\.([0-9]+)\.([0-9]+)/gi

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

  const tags = refTags.repository.refs.edges.map(x => x.node.name)

  if (tags.length === 0) {
    const {default_tag} = getInputs()
    tags.push(default_tag)
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
    return `${major}.0.0`
  }

  if (ver.Minor) {
    minor++
    return `${major}.${minor}.0`
  }

  patch++
  return `${major}.${minor}.${patch}`
}

export async function createRelease(tag: string): Promise<void> {
  try {
    await octokit.rest.repos.createRelease({
      owner: gh.context.repo.owner,
      repo: gh.context.repo.repo,
      tag_name: tag,
      name: tag,
      draft: false,
      prerelease: false,
      generate_release_notes: true
    })
  } catch (e) {
    core.error(`Create release error: ${(e as Error).message}`)
  }
}

export async function updateRelease(release_id: number, tag: string): Promise<void> {
  await octokit.rest.repos.updateRelease({
    owner: gh.context.repo.owner,
    repo: gh.context.repo.repo,
    tag_name: tag,
    target_cmmitish: 'master',
    name: tag,
    draft: false,
    prerelease: false,
    release_id,
    generate_release_notes: true
  })
}

export async function createLabels(): Promise<void> {
  for (const label of LABELS) {
    try {
      await createLabel(label.name, label.description, label.color)
    } catch (e) {
      await updateLabel(label.name, label.description, label.color)
    }
  }
}

async function createLabel(name: string, description: string, color: string): Promise<void> {
  await octokit.rest.issues.createLabel({
    owner: gh.context.repo.owner,
    repo: gh.context.repo.repo,
    name,
    description,
    color
  })
}

async function updateLabel(name: string, description: string, color: string): Promise<void> {
  await octokit.rest.issues.updateLabel({
    owner: gh.context.repo.owner,
    repo: gh.context.repo.repo,
    name,
    new_name: name,
    description,
    color
  })
}
