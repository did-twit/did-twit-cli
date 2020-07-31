package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	usernameFlag = "username"
	didFlag      = "did"
)

func init() {
	addPersistentStringFlag(RootCmd, usernameFlag, "u", "Twitter username.", NotRequired)
	addPersistentStringFlag(RootCmd, didFlag, "d", "did:twit DID.", NotRequired)
}

var RootCmd = &cobra.Command{
	Use:   "did-twit-cli",
	Short: "The did:twit cli",
	Long:  `A utility for using the did:twit DID method for your Twitter username. Helps you create a did:twit document and generate Tweets.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.UsageString())
		// uncomment me to generate docs
		//err := doc.GenMarkdownTree(cmd, "/tmp")
		//if err != nil {
		//	log.Fatal(err)
		//}
	},
}
