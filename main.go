package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/TylerBrock/colorjson"
	"gopkg.in/square/go-jose.v2/jwt"
)

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
	tokenArgument := flag.String("t", "", "Token to decode. If not specified will try to read it from stdin.")
	colorizeArgument := flag.Bool("c", false, "Colorize the output")
	flag.Parse()

	tokenString := ""

	if *tokenArgument == "" {
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		newToken := ""
		for err == nil {
			newToken = newToken + line
			line, err = reader.ReadString('\n')
		}
		tokenString = newToken
	} else {
		tokenString = *tokenArgument
	}

	decodedToken, err := decodeToken(tokenString, *colorizeArgument)

	if err != nil {
		log.Fatalf("could not decode token %v", err)
	}

	fmt.Print(decodedToken)
}
