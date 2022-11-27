import * as core from '@actions/core'
import * as repo from './repo'

async function run(): Promise<void> {
  try {
    await repo.createLabels()
  } catch (error) {
    if (error instanceof Error) core.setFailed(JSON.stringify(error))
  }
}

run()
