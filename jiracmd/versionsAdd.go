package jiracmd

import (
	"github.com/coryb/figtree"
	"github.com/coryb/oreo"
	"github.com/go-jira/jira"
	"github.com/go-jira/jira/jiracli"
	"github.com/go-jira/jira/jiradata"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type VersionAddOptions struct {
	jiracli.CommonOptions `yaml:",inline" json:",inline" figtree:",inline"`
	Project               string `json:"project,omitempty" yaml:"project,omitempty"`
	Name                  string `json:"name,omitempty" yaml:"name,omitempty"`
	Description           string `json:"description,omitempty" yaml:"description,omitempty"`
}

func CmdVersionAddRegistry() *jiracli.CommandRegistryEntry {
	opts := VersionAddOptions{
		CommonOptions: jiracli.CommonOptions{
			Template: figtree.NewStringOption("version"),
		},
	}

	return &jiracli.CommandRegistryEntry{
		"Adds new version to the project",
		func(fig *figtree.FigTree, cmd *kingpin.CmdClause) error {
			jiracli.LoadConfigs(cmd, fig, &opts)
			return CmdVersionAddUsage(cmd, &opts)
		},
		func(o *oreo.Client, globals *jiracli.GlobalOptions) error {
			return CmdVersionAdd(o, globals, &opts)
		},
	}
}

func CmdVersionAddUsage(cmd *kingpin.CmdClause, opts *VersionAddOptions) error {
	jiracli.BrowseUsage(cmd, &opts.CommonOptions)
	jiracli.TemplateUsage(cmd, &opts.CommonOptions)
	jiracli.GJsonQueryUsage(cmd, &opts.CommonOptions)
	cmd.Arg("PROJECT", "project key").Required().StringVar(&opts.Project)
	cmd.Arg("NAME", "version name").Required().StringVar(&opts.Name)
	return nil
}

// View will get issue data and send to "view" template
func CmdVersionAdd(o *oreo.Client, globals *jiracli.GlobalOptions, opts *VersionAddOptions) error {

	versionCreate := jiradata.Version{
		Name:    opts.Name,
		Project: opts.Project,
	}

	data, err := jira.CreateProjectVersion(o, globals.Endpoint.Value, &versionCreate)
	if err != nil {
		return err
	}
	if err := opts.PrintTemplate(data); err != nil {
		return err
	}
	return nil
}
