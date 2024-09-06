# jwt-decode

Simple tool to decode JSON Web token. No validation is included!

![Build](https://github.com/jadolg/jwt-decode/workflows/Go/badge.svg)
![Release](https://github.com/jadolg/jwt-decode/workflows/Release/badge.svg)

## How to use

Just pass it as an argument `jwt-decode -t "ABeautifulToken"`
or pipe it in `echo "ABeautifulToken" | jwt-decode`

### Running from Docker

You are the kind of person that would love to run the world using docker? This is for you: `docker run --rm -it guamulo/jwt-decode -c -t token`

## How to install

Download the latest release (from the GitHub releases) for your platform, place it in a directory within your `$PATH`, and you are good to go.

or

If you are using Ubuntu you can just `snap install jwt-decode-claims` 

or

If you use homebrew: 

```bash
brew tap jadolg/homebrew-tap
brew install jwt-decode
```

or

From my private apt repository
```
wget -O - https://deb.akiel.dev/gpg.pub.key | gpg --dearmor | sudo tee /etc/apt/trusted.gpg.d/akiel.gpg
sudo apt-add-repository deb "https://deb.akiel.dev/ all main"
sudo apt install jwt-decode
```
