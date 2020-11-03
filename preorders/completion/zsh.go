package completion

import (
	"github.com/DamianSkrzypczak/order/internal/orderfile"
)

// ZSHCompletionOrder is order with ZSH completion adding script
var ZSHCompletionOrder = &orderfile.Order{
	Description: "Create and register ZSH completion script",
	Script: []orderfile.Cmd{
		// Ensure completion directory
		"mkdir -p ~/.zsh_completion.d/order",

		// Define script string.
		`script='#compdef order

		  __list_orders() {
			local -a orders

			orders_raw=$(order -l)

			if [[ ${orders_raw} == *"Could not load orderfile"* ]]; then
				return;
			fi;

			orders=( $(echo "$(order -l)" | tail -n +2 | cut -f1 | sed -r "/^\s*$/d" | tr "\n" " ") );
			_describe "order" orders
		  }

		  _arguments \
			  "(--debug)"--debug \
			  "(-l --list)"{-l,--list} \
			  "(--no-color)"--no-color \
			  "(--no-command)"--no-command \
			  "(--no-level)"--no-level \
			  "(-p --path)"{-p,--path}": :_files" \
			  "(--version)"--version \
			  "*: :__list_orders" \
		'`,

		// Create completion file
		"echo -e \"${script}\" > ~/.zsh_completion.d/order/_order",

		// If .zshrc doesn't contain sourcing command, add it
		`grep -qxF 'export fpath=(~/.zsh_completion.d/order $fpath)' ~/.zshrc || echo -e "\nexport fpath=(~/.zsh_completion.d/order \$fpath)" >> ~/.zshrc`,
		`grep -qxF 'autoload -Uz compinit && compinit' ~/.zshrc || echo -e "autoload -Uz compinit && compinit" >> ~/.zshrc`,

		// Show success message with sourcing reminder
		`echo "ZSH autocompletion support successfully added, run \"source ~/.zshrc\" to activate it"`,
	},
}

