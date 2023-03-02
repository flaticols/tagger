package commands

import (
	"fmt"

	"github.com/flaticols/tagger/gh"
	"github.com/flaticols/tagger/inputs"
	gha "github.com/sethvargo/go-githubactions"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func CreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a release",
		PreRun: func(cmd *cobra.Command, args []string) {
			isActions, _ := cmd.Flags().GetBool("actions")

			if !isActions {
				_ = cmd.MarkFlagRequired("owner")
				_ = cmd.MarkFlagRequired("token")
				_ = cmd.MarkFlagRequired("repo")
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			flags := cmd.Flags()
			isActions, _ := flags.GetBool("actions")

			params := inputs.Inputs{}

			major, _ := flags.GetBool("major")
			minor, _ := flags.GetBool("minor")
			patch, _ := flags.GetBool("patch")

			verUpd := inputs.PullRequestLabels{
				Major: major,
				Minor: minor,
				Patch: patch,
			}

			if isActions {
				fmt.Println("Running in GitHub Actions")
				r, err := inputs.GetInputs()
				if err != nil {
					return err
				}
				params = r
			} else {
				r, err := getRepoInfo(flags)
				if err != nil {
					return err
				}
				params = r
			}

			client, err := gh.NewGitHubClient(params.GitHubToken)
			if err != nil {
				return fmt.Errorf("failed to create GitHub client, %w", err)
			}

			if isActions {
				labels, err := client.GetPullRequestLabels(params.PullRequestNumber, params.Owner, params.Repository)
				if err != nil {
					return fmt.Errorf("failed to get pull request labels, %w", err)
				}
				verUpd = inputs.GetPRLabels(labels)
			}

			latestTag, err := client.GetLatestTag(params.Owner, params.Repository, params.GetDefaultTag())
			if err != nil {
				return fmt.Errorf("failed to get latest tag, %w", err)
			}

			newTag, err := inputs.GetNewTag(latestTag, verUpd)
			if err != nil {
				return fmt.Errorf("failed to get new tag, %w", err)
			}

			newTagString := newTag.StringWithPrefix(params.TagPrefix)
			_, err = client.CreateRelease(newTagString, params.Owner, params.Repository, false, true)
			if err != nil {
				return fmt.Errorf("failed to create release, %w", err)
			}

			gha.SetOutput("tag", newTagString)

			return nil
		},
	}

	cmd.Flags().BoolP("actions", "", false, "The run in GitHub Actions")
	cmd.Flags().StringP("owner", "o", "", "The owner of the repository")
	cmd.Flags().StringP("token", "t", "", "The GitHub token")
	cmd.Flags().StringP("repo", "r", "", "The repository name")
	cmd.Flags().StringP("tag", "", "0.0.0", "The default tag to create")
	cmd.Flags().StringP("prefix", "p", "v", "The tag prefix")

	cmd.Flags().BoolP("major", "", false, "The major version update")
	cmd.Flags().BoolP("minor", "", false, "The minor version update")
	cmd.Flags().BoolP("patch", "", false, "The patch version update")

	return cmd
}

func getRepoInfo(flags *pflag.FlagSet) (inputs.Inputs, error) {
	params := inputs.Inputs{}

	owner, err := flags.GetString("owner")
	if err != nil {
		return params, fmt.Errorf("get owner failed, %w", err)
	}
	params.Owner = owner

	repo, err := flags.GetString("repo")
	if err != nil {
		return params, fmt.Errorf("get repo failed, %w", err)
	}
	params.Repository = repo

	token, err := flags.GetString("token")
	if err != nil {
		return params, fmt.Errorf("get token failed, %w", err)
	}
	params.GitHubToken = token

	defaultTag, err := flags.GetString("tag")
	if err != nil {
		return params, fmt.Errorf("get tag failed, %w", err)
	}
	params.DefaultTag = defaultTag

	tagPrefix, err := flags.GetString("prefix")
	if err != nil {
		return params, fmt.Errorf("get tag prefix, %w", err)
	}
	params.TagPrefix = tagPrefix

	return params, nil
}
