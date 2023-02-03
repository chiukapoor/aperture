package cmd

import (
	"errors"
	"os"

	"github.com/spf13/cobra"

	"github.com/fluxninja/aperture/cmd/aperturectl/cmd/utils"
	"github.com/fluxninja/aperture/pkg/log"
)

var (
	policy  string
	cr      string
	dot     string
	mermaid string
)

func init() {
	compileCmd.Flags().StringVar(&policy, "policy", "", "Path to policy file")
	compileCmd.Flags().StringVar(&cr, "cr", "", "Path to policy custom resource file")
	compileCmd.Flags().StringVar(&dot, "dot", "", "Path to store the dot file")
	compileCmd.Flags().StringVar(&mermaid, "mermaid", "", "Path to store the mermaid file")
}

// compileCmd is the command to compile a circuit from a policy file or CR.
var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compile circuit from policy file",
	Long: `
Use this command to compile the Aperture policy circuit from a policy file to validate the circuit.
You can also generate the DOT and Mermaid graphs of the compiled policy circuit to visualize it.`,
	SilenceErrors: true,
	Example: `aperturectl compile --cr=policy-cr.yaml --mermaid --dot

aperturectl compile --policy=policy.yaml --mermaid --dot`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// check if policy or cr is provided
		if policy == "" && cr == "" || policy != "" && cr != "" {
			errStr := "either --policy or --cr must be provided"
			return errors.New(errStr)
		}

		var policyFile string
		var err error

		// check if cr is provided
		if cr != "" {
			policyFile, err = utils.FetchPolicyFromCR(cr)
			if err != nil {
				return err
			}
			defer os.Remove(policyFile)
		} else {
			policyFile = policy
		}

		circuit, err := utils.CompilePolicy(policyFile)
		if err != nil {
			log.Error().Err(err).Msg("error reading policy spec")
			return err
		}

		log.Info().Msg("Compilation successful")

		// if --dot flag is set, write dotfile
		// check if the dot flag is set
		if dot != "" {
			if err := utils.GenerateDotFile(circuit, dot); err != nil {
				return err
			}
		}
		// if --mermaid flag is set, write mermaid file
		if mermaid != "" {
			if err := utils.GenerateMermaidFile(circuit, mermaid); err != nil {
				return err
			}
		}
		return nil
	},
}