package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/square/go-jose.v2/jwt"
	"log"
	"os"
)

func toJsonString(x interface{}) (string, error) {
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func decodeToken(tokenString string) (string, error) {
	var claims map[string]interface{}

	token, err := jwt.ParseSigned(tokenString)
	if err != nil {
		return "", err
	}

	err = token.UnsafeClaimsWithoutVerification(&claims)
	if err != nil {
		return "", err
	}

	return toJsonString(claims)
}

func main() {
	tokenArgument := flag.String("t", "", "Token to decode. If not specified will try to read it from stdin.")
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

	decodedToken, err := decodeToken(tokenString)

	if err != nil {
		log.Fatalf("could not decode token %v", err)
	}

	fmt.Print(decodedToken)
}
