package internal

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"github.com/spf13/cobra"
	"hash"
	"os"
)

type RootOptions struct {
	algorithm string
	key       []byte
	text      string
}

func (o RootOptions) hashFunction() (func() hash.Hash, error) {
	switch o.algorithm {
	case "sha1":
		return sha1.New, nil
	case "sha256":
		return sha256.New, nil
	case "sha512":
		return sha512.New, nil
	}

	return nil, fmt.Errorf("%s is not a supported signing algorithm", o.algorithm)
}

var rootOptions = RootOptions{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hmac-signer",
	Short: "Generate HMAC signatures from text or file",
	//Long: `A longer description that spans multiple lines and likely contains`
	RunE: func(cmd *cobra.Command, args []string) error {
		hashFunction, err := rootOptions.hashFunction()
		if err != nil {
			return err
		}

		hash := hmac.New(hashFunction, rootOptions.key)
		hash.Write([]byte(rootOptions.text))

		fmt.Printf("%x\n", hash.Sum(nil))
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&rootOptions.algorithm, "algorithm", "a", "sha256", "hash function used for signature (sha1, sha256 or sha512)")
	rootCmd.Flags().BytesBase64VarP(&rootOptions.key, "key", "k", nil, "key encoded as Base64 used to sign the data")
	_ = rootCmd.MarkFlagRequired("key")
	rootCmd.Flags().StringVarP(&rootOptions.text, "text", "t", "", "plain-text that will be signed")
	_ = rootCmd.MarkFlagRequired("text")
}
