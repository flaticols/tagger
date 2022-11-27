import * as core from '@actions/core'
import * as repo from './repo'

async function run(): Promise<void> {
  try {
    await repo.createLabels()

    const pr = await repo.getPR()
    const ver = repo.getRequiredLabels(pr.data.labels)
    const tags = await repo.getLatestTag()

    const newVersion = repo.getNewTag(ver, tags)
    core.notice(`Next release: v${newVersion}`)

    repo.createRelease(`v${newVersion}`)
  } catch (error) {
    if (error instanceof Error) core.setFailed(JSON.stringify(error))
  }
}

run()
