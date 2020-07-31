package cmd

import (
	"errors"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/did-twit/did-twit-cli/pkg/did"
	"github.com/did-twit/did-twit-cli/pkg/storage"
)

func init() {
	MgmtCmd.AddCommand(createDIDCmd, viewDIDCmd)
	RootCmd.AddCommand(MgmtCmd)
}

var MgmtCmd = &cobra.Command{
	Use:   "manage",
	Short: "administrativa for did:twit docs",
	Long:  `create and view did:twit documents`,
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := storage.NewConnection()
		if err != nil {
			fmt.Println("Problem initiating storage")
			return err
		}
		dids, err := db.ListDIDs()
		if err != nil {
			return err
		}
		printDIDs(dids)
		return nil
	},
}

func printDIDs(dids []string) {
	if len(dids) == 0 {
		fmt.Println("No DIDs available.")
		return
	}
	for _, d := range dids {
		fmt.Println(d)
	}
}

var createDIDCmd = &cobra.Command{
	Use:   "create",
	Short: "create a did:twit doc and store the private key",
	Long:  `create a did:twit doc and store the private key`,
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := storage.NewConnection()
		if err != nil {
			fmt.Println("Problem initiating storage")
			return err
		}
		user := viper.GetString(usernameFlag)
		if user == "" {
			return errors.New("username cannot be empty. supply a username with the \"--user\" flag")
		}
		did, privKey, err := did.CreateDID(user)
		if err != nil {
			fmt.Printf("Problem creating did: %s\n", user)
			return err
		}
		if err := db.Write(*did, base58.Encode(privKey)); err != nil {
			fmt.Printf("Problem writing did user: %s\n", user)
			return err
		}
		fmt.Printf("DID created successfully:\n%s\n", *did)
		return nil
	},
}

var viewDIDCmd = &cobra.Command{
	Use:   "view",
	Short: "view a did:twit DID",
	Long:  `view a did:twit base 58 encoded private key"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := storage.NewConnection()
		if err != nil {
			fmt.Println("Problem initiating storage")
			return err
		}
		id := viper.GetString(didFlag)
		if id == "" {
			return errors.New("DID cannot be empty. Supply a DID with the \"--did\" flag ")
		}
		res, err := db.Read(id)
		if err != nil {
			return err
		}
		if *res == "" {
			fmt.Printf("DID <%s> was not found be", id)
		} else {
			fmt.Printf("DID: %s\nPrivate Key (b58): %s\n", id, *res)
		}
		return nil
	},
}
