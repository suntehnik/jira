package jiracmd

import (
	"github.com/coryb/figtree"
	"github.com/coryb/oreo"
	"github.com/go-jira/jira"
	"github.com/go-jira/jira/jiracli"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type VersionsOptions struct {
	jiracli.CommonOptions `yaml:",inline" json:",inline" figtree:",inline"`
	Project               string `yaml:"project,omitempty" json:"project,omitempty"`
}

func CmdVersionsRegistry() *jiracli.CommandRegistryEntry {
	opts := VersionsOptions{
		CommonOptions: jiracli.CommonOptions{
			Template: figtree.NewStringOption("versions"),
		},
	}

	return &jiracli.CommandRegistryEntry{
		"Prints list of versions",
		func(fig *figtree.FigTree, cmd *kingpin.CmdClause) error {
			jiracli.LoadConfigs(cmd, fig, &opts)
			return CmdVersionsUsage(cmd, &opts)
		},
		func(o *oreo.Client, globals *jiracli.GlobalOptions) error {
			return CmdVersions(o, globals, &opts)
		},
	}
}

func CmdVersionsUsage(cmd *kingpin.CmdClause, opts *VersionsOptions) error {
	jiracli.BrowseUsage(cmd, &opts.CommonOptions)
	jiracli.TemplateUsage(cmd, &opts.CommonOptions)
	jiracli.GJsonQueryUsage(cmd, &opts.CommonOptions)
	cmd.Arg("PROJECT", "project").Required().StringVar(&opts.Project)
	return nil
}

// View will get issue data and send to "view" template
func CmdVersions(o *oreo.Client, globals *jiracli.GlobalOptions, opts *VersionsOptions) error {
	data, err := jira.GetProjectVersions(o, globals.Endpoint.Value, opts.Project)
	if err != nil {
		return err
	}
	versions := map[string]interface{}{
		"versions": data,
	}
	if err := opts.PrintTemplate(versions); err != nil {
		return err
	}
	return nil
}
