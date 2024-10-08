package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/go-jose/go-jose/v4"
	"github.com/spf13/cobra"
	"log"
	"os"

	"github.com/TylerBrock/colorjson"
	"github.com/go-jose/go-jose/v4/jwt"
)

var (
	// Version is the version of the application.
	Version = "dev"
	// Commit is the git commit of the build
	Commit = "none"
	// Date is the date of the build
	Date = "unknown"
	// BuiltBy is the user/tool that built the binary
	BuiltBy = "dirty hands"

	rootCmd = &cobra.Command{
		Use:              "jwt-decode",
		Short:            "Simple tool to decode JSON Web token.",
		Example:          "jwt-decode -t \"ABeautifulToken\"",
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			tokenString := ""

			if tokenArgument == "" {
				reader := bufio.NewReader(os.Stdin)
				line, err := reader.ReadString('\n')
				newToken := ""
				for err == nil {
					newToken = newToken + line
					line, err = reader.ReadString('\n')
				}
				tokenString = newToken
			} else {
				tokenString = tokenArgument
			}

			decodedToken, err := decodeToken(tokenString, colorizeArgument)

			if err != nil {
				log.Fatalf("could not decode token %v", err)
			}

			fmt.Print(decodedToken)
		},
	}
	tokenArgument    string
	colorizeArgument bool

	cmdVersion = &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version: %s, Commit: %s, Date: %s, BuiltBy: %s\n", Version, Commit, Date, BuiltBy)
		},
	}
)

func init() {
	rootCmd.AddCommand(cmdVersion)
	rootCmd.PersistentFlags().StringVarP(&tokenArgument, "token", "t", "", "Token to decode. If not specified will try to read it from stdin.")
	rootCmd.PersistentFlags().BoolVarP(&colorizeArgument, "colorize", "c", false, "Colorize the output")
}

func toJSONString(x interface{}, colorize bool) (string, error) {
	if colorize {
		f := colorjson.NewFormatter()
		f.Indent = 2

		s, err := f.Marshal(x)
		if err != nil {
			return "", err
		}
		return string(s), nil
	}
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func decodeToken(tokenString string, colorize bool) (string, error) {
	var claims map[string]interface{}

	token, err := jwt.ParseSigned(
		tokenString,
		[]jose.SignatureAlgorithm{
			jose.EdDSA,
			jose.HS256,
			jose.HS384,
			jose.HS512,
			jose.RS256,
			jose.RS384,
			jose.RS512,
			jose.ES256,
			jose.ES384,
			jose.ES512,
			jose.PS256,
			jose.PS384,
			jose.PS512,
		},
	)
	if err != nil {
		return "", err
	}

	err = token.UnsafeClaimsWithoutVerification(&claims)
	if err != nil {
		return "", err
	}

	return toJSONString(claims, colorize)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
