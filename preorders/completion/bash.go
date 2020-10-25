package completion

import (
	"github.com/DamianSkrzypczak/order/internal/orderfile"
)

// BashCompletionOrder is order with bash completion adding script
var BashCompletionOrder = &orderfile.Order{
	Description: "Create and register bash completion script",
	Script: []orderfile.Cmd{
		// Ensure completion directory
		"mkdir -p ~/.bash_completion.d",

		// Define script string.
		`script='_order_completion() {
			local orders;
			local orders_raw;
			local curr_arg;

			orders_raw=$(order -l);

			if [[ ${orders_raw} == *"Could not load orderfile"* ]]; then
				return;
			fi;

			orders=$(echo "${orders_raw}" | tail -n +2 | cut -f1 | sed -r "/^\s*$/d");
			COMPREPLY=($(compgen -c | echo "${orders}" | grep "^${COMP_WORDS[COMP_CWORD]}"));
		  };

		  complete -F _order_completion order
		  '`,

		// Create completion file
		"echo ${script} > ~/.bash_completion.d/order.bash",

		// If .bashrc doesn't contain sourcing command, add it
		`grep -qxF 'source ~/.bash_completion.d/order.bash' ~/.bashrc || echo "source ~/.bash_completion.d/order.bash" >> ~/.bashrc`,

		// Show success message with sourcing reminder
		`echo "Bash autocompletion support successfully added, run \"source ~/.bashrc\" to activate it"`,
	},
}
