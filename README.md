<p align="center">
  <!-- logo -->
  <img src="https://github.com/edoardottt/images/blob/main/pwdsafety/logo.jpg"><br>
  <b>Command line tool that checks how much a password is safe</b><br>
  <sub>
    Coded with 💙 by edoardottt.
  </sub>
</p>

<!-- badges -->
<p align="center">
    <!-- mainteinance -->
      <a href="https://edoardoottavianelli.it">
        <img src="https://img.shields.io/badge/Maintained%3F-yes-green.svg" alt="Mainteinance yes" />
      </a>
    <!-- version -->
      <a href="https://edoardoottavianelli.it">
        <img src="https://github.com/edoardottt/images/blob/main/pwdsafety/version.svg" alt="version" />
      </a>
    <!-- pr-welcome -->
      <a href="https://edoardoottavianelli.it">
        <img src="https://github.com/edoardottt/READMENATOR/blob/master/images/pr-welcome.svg" alt="pr-welcome" />
      </a>
    <!-- ask-me-anything -->
      <a href="https://edoardoottavianelli.it">
        <img src="https://github.com/edoardottt/READMENATOR/blob/master/images/ask-me-anything.svg" alt="ask me anything" />
      </a>
  <br>
    <!-- go-report-card -->
      <a href="https://goreportcard.com/report/github.com/edoardottt/pwdsafety">
        <img src="https://goreportcard.com/badge/github.com/edoardottt/pwdsafety" alt="go-report-card" />
      </a>
    <!-- workflows -->
      <a href="https://github.com/edoardottt/pwdsafety/actions">
        <img src="https://github.com/edoardottt/pwdsafety/actions/workflows/go.yml/badge.svg" alt="workflows" />
      </a>
    <!-- ubuntu-build -->
      <a href="https://edoardoottavianelli.it">
        <img src="https://github.com/edoardottt/images/blob/main/pwdsafety/ubuntu-build.svg" alt="ubuntu-build" />
      </a>
  <br>
    <!-- license GPLv3.0 -->
      <a href="https://github.com/edoardottt/READMENATOR/blob/master/LICENSE">
        <img src="https://github.com/edoardottt/READMENATOR/blob/master/images/license-GPL3.svg" alt="license-GPL3" />
      </a>
</p>
<p align="center">
  <a href="#example-bar_chart">Example</a> •
  <a href="#get-started-">Get Started</a> •
  <a href="#description-">Description</a> •
  <a href="#scoring-">Scoring</a> •
  <a href="#contributing-">Contributing</a> •
  <a href="#changelog-%EF%B8%8F">Changelog</a>
</p>

**This tool doesn't store any information!!**  
**Remember, never use personal information in your password!**  
 - Use a password manager (I recommend [BitWarden](https://bitwarden.com/))  
 - Don't use the same password for different services  
 - Enable 2FA when possible  

Example :bar_chart:
----------

[![asciicast](https://asciinema.org/a/406710.svg)](https://asciinema.org/a/406710)

Get Started 🎉
----------

### Snap
```bash
sudo snap install pwdsafety
```

### Go1.17+
```bash
go install -v github.com/edoardottt/pwdsafety@latest
```

### From source

- First of all, clone the repo locally

  - `git clone https://github.com/edoardottt/pwdsafety.git`

- pwdsafety has external dependencies, so they need to be pulled in:

  - `cd pwdsafety/cmd && go get && cd ..`

- Linux (Requires high perms, run with sudo)

  - `make linux` (to install)

  - `make unlinux` (to uninstall)

- Windows (executable works only in pwdsafety folder. Alias?)

  - `make windows` (to install)

  - `make unwindows` (to uninstall)

Description 🔦 
----------

It reads from standard input the entered password.  
First, it searches if the password or the password reversed is a well known pwd.  
Then, just do little calculations, checking if the basic rules are respected, like if there are UPPERCASE CHARS, lowercase chars, numb3rs and symbols.  
It stores the length of the password and the ratio [ unique different chars / total chars].  
It calculates then the entropy of a password.  
Password entropy is a measurement of how unpredictable a password is.  
The formula for entropy is:  
              ![formula](https://github.com/edoardottt/images/blob/main/pwdsafety/formula.png)  
              
Where:
- E = password entropy  
- R = pool of unique characters  
- L = number of characters in your password  
- Then R^L = the number of possible passwords  

When the score <= 68(reasonable) it generates a random password.  

Scoring 💯
----------

**Max score: 100**

**Scores:**
  - Very weak: 0 - 35
  - Weak: 36 - 59
  - Reasonable: 60 - 68
  - Strong: 69 - 80
  - Very strong: 81 -100
  
**Scoring parameters:**
  - Found in known password
  - Found in known password reversed
  - Password composition:
      - numbers
      - symbols
      - uppercase
      - lowercase
  - Unique different characters
  - Length
  - Entropy

Contributing 🛠
-------

Just open an [issue](https://github.com/edoardottt/pwdsafety/issues) / [pull request](https://github.com/edoardottt/pwdsafety/pulls). 

Before opening a pull request, download [golangci-lint](https://golangci-lint.run/usage/install/) and run
```bash
golangci-lint run
```
If there aren't errors, go ahead :)

See also [CONTRIBUTING.md](https://github.com/edoardottt/pwdsafety/blob/master/CONTRIBUTING.md) and [CODE OF CONDUCT.md](https://github.com/edoardottt/pwdsafety/blob/master/CODE_OF_CONDUCT.md)

Thanks to [fabaff](https://github.com/fabaff) and [ecnepsnai](https://github.com/ecnepsnai/pwnedpassword/blob/master/pwned.go).

Changelog 📌
-------
Detailed changes for each release are documented in the [release notes](https://github.com/edoardottt/pwdsafety/releases).

License 📝
-------

This repository is under [GNU General Public License v3.0](https://github.com/edoardottt/pwdsafety/blob/master/LICENSE).  
[edoardoottavianelli.it](https://www.edoardoottavianelli.it) to contact me.
