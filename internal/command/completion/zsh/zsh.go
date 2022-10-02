package zsh

import (
	"html/template"
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/takumin/boilerplate-golang-cli/internal/config"
)

const zshCompletion = `
#compdef {{.}}

_cli_zsh_autocomplete() {
  local -a opts
  local cur

  cur=${words[-1]}
  if [[ "$cur" == "-"* ]]; then
    opts=("${(@f)$(_CLI_ZSH_AUTOCOMPLETE_HACK=1 ${words[@]:0:#words[@]-1} ${cur} --generate-bash-completion)}")
  else
    opts=("${(@f)$(_CLI_ZSH_AUTOCOMPLETE_HACK=1 ${words[@]:0:#words[@]-1} --generate-bash-completion)}")
  fi

  if [[ "${opts[1]}" != "" ]]; then
    _describe 'values' opts
  else
    _files
  fi

  return
}

compdef _cli_zsh_autocomplete {{.}}
`

func NewCommands(cfg *config.Config, flags []cli.Flag) *cli.Command {
	return &cli.Command{
		Name:     "zsh",
		Usage:    "zsh completion",
		HideHelp: true,
		Action: func(ctx *cli.Context) error {
			t, err := template.New(ctx.App.Name).Parse(strings.TrimSpace(zshCompletion) + "\n")
			if err != nil {
				return err
			}
			if err = t.Execute(os.Stdout, ctx.App.Name); err != nil {
				return err
			}
			return nil
		},
	}
}
