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

func printJson(x interface{}) {
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))
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

	var claims map[string]interface{}

	token, err := jwt.ParseSigned(tokenString)
	if err != nil {
		log.Fatal(err)
	}

	err = token.UnsafeClaimsWithoutVerification(&claims)
	if err != nil {
		log.Fatal(err)
	}

	printJson(claims)
}
