package commands


import (
	"github.com/spf13/cobra"
)


func init(){

}



// faasCmd is the FaaS CLI root command and mimics the legacy client behaviour
// Every other command attached to FaasCmd is a child command to it.
var plantCmd = &cobra.Command{
	Use:   "plant",
	Short: "Manage your farm",
	Long: `
Allows you to manager your farm and run commmands`,
	Run: runPlant,
}


func runPlant(cmd *cobra.Command, args []string) {
	printFiglet()
	cmd.Help()
}