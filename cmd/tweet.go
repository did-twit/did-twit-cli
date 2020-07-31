package cmd

import (
	"errors"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/did-twit/did-twit-cli/pkg/did"
	"github.com/did-twit/did-twit-cli/pkg/storage"
	"github.com/did-twit/did-twit-cli/pkg/tweet"
)

var (
	tweetFlag  = "tweet"
	createFlag = "create"
)

func init() {
	addPersistentBoolFlag(TweetCmd, createFlag, "c", "Set to true if initial Tweet.", false, NotRequired)
	addPersistentStringFlag(TweetCmd, tweetFlag, "t", "Text of a tweet.", NotRequired)
	TweetCmd.AddCommand(VerifyCmd)
	RootCmd.AddCommand(TweetCmd)
}

var TweetCmd = &cobra.Command{
	Use:   "tweet",
	Short: "create a tweet for a given did",
	Long:  `create a tweet for a given did`,
	RunE: func(cmd *cobra.Command, args []string) error {
		didTwit := viper.GetString(didFlag)
		if didTwit == "" {
			return errors.New("must supply a user to author the tweet")
		}

		create := viper.GetBool(createFlag)
		var text string
		if create {
			text = createTweet(didTwit)
		} else {
			text = viper.GetString(tweetFlag)
		}
		if text == "" {
			return errors.New("tweet text cannot be empty")
		}

		db, err := storage.NewConnection()
		if err != nil {
			fmt.Println("Problem initiating storage")
			return err
		}
		didPrivKey, err := db.Read(didTwit)
		if err != nil || *didPrivKey == "" {
			fmt.Printf("Key could not be found for did: %s", didTwit)
			return err
		}

		privKey := base58.Decode(*didPrivKey)
		t, err := tweet.SignTweet(privKey, didTwit, text)
		if err != nil {
			fmt.Println("Problem signing tweet")
			return err
		}
		tweetTxt, err := tweet.GenerateTweetText(*t)
		if err != nil {
			fmt.Println("Problem generating tweet text")
			return err
		}
		fmt.Printf("Copy & Paste the following text into Twitter: \n"+
			"------------------------------------\n"+
			"%s"+
			"\n------------------------------------", *tweetTxt)
		return nil
	},
}

func createTweet(did string) string {
	return fmt.Sprintf("I've created a did:twit method with the the identifier %s! I will be using this identitiy"+
		" to verify Tweets, and other information originating from this account. Find out more at didtwit.io!", did)
}

var VerifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "verify a tweet for a given did",
	Long:  `verify a tweet for a given did`,
	RunE: func(cmd *cobra.Command, args []string) error {
		didTwit := viper.GetString(didFlag)
		if didTwit == "" {
			return errors.New("must supply the author of the tweet")
		}

		text := viper.GetString(tweetFlag)
		if text == "" {
			return errors.New("tweet text cannot be empty")
		}

		t, err := tweet.ReconstructTweet(text)
		if err != nil {
			fmt.Println("Could not reconstruct tweet")
			return err
		}

		pubKey, err := did.ExpandDID(didTwit)
		if err != nil {
			fmt.Printf("Author DID could not be expanded: %s\n", didTwit)
			return err
		}
		if err := tweet.VerifyTweet(*t, pubKey); err != nil {
			fmt.Printf("Tweet not valid for the given author: %s\n", didTwit)
			return err
		}

		fmt.Println("Tweet valid.")
		return nil
	},
}
