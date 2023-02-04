#!/bin/sh
# scripts/completions.sh
set -e
rm -rf completions
mkdir completions
for sh in bash zsh fish; do
	go run main.go completion "$sh" >"completions/jwt-decode.$sh"
	chmod +x completions/jwt-decode.$sh
done