import * as core from '@actions/core'
import type {components} from '@octokit/openapi-types'

export type ActionInput = {
  token: string
  pr_number: number
}

export type SemVerUpdate = {
  Major: boolean
  Minor: boolean
  Patch: boolean
}

export type RepositoryTags = {
  repository: {
    refs: {
      edges: {
        node: {name: string}
      }[]
    }
  }
}

export type PR = components['schemas']['pull-request-simple']
export type Label = components['schemas']['label']

export function getInputs(): ActionInput {
  const token = core.getInput('github-token', {required: true})
  const pr_number = parseInt(core.getInput('pr_number', {required: true}))
  return {token, pr_number}
}
