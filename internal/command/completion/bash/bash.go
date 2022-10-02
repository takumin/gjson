package bash

import (
	"html/template"
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/takumin/boilerplate-golang-cli/internal/config"
)

const bashCompletion = `
#!/bin/bash

_cli_bash_autocomplete() {
  if [[ "${COMP_WORDS[0]}" != "source" ]]; then
    local cur opts base
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    if [[ "$cur" == "-"* ]]; then
      opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} ${cur} --generate-bash-completion )
    else
      opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} --generate-bash-completion )
    fi
    COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
    return 0
  fi
}

complete -o bashdefault -o default -o nospace -F _cli_bash_autocomplete {{.}}
`

func NewCommands(cfg *config.Config, flags []cli.Flag) *cli.Command {
	return &cli.Command{
		Name:     "bash",
		Usage:    "bash completion",
		HideHelp: true,
		Action: func(ctx *cli.Context) error {
			t, err := template.New("bashCompletion").Parse(strings.TrimSpace(bashCompletion) + "\n")
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
