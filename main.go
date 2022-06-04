package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"

	"github.com/TylerBrock/colorjson"
	"gopkg.in/square/go-jose.v2/jwt"
)

var (
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

	// Cobra will only enable the completion subcommand if you have at least one subcommand
	cmdDummy = &cobra.Command{
		Use:    "dummy",
		Hidden: true,
	}
)

func init() {
	rootCmd.AddCommand(cmdDummy)
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

	token, err := jwt.ParseSigned(tokenString)
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
