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
    <!-- open-issues -->
      <a href="https://edoardoottavianelli.it">
        <img src="https://img.shields.io/github/issues/Naereen/StrapDown.js.svg" alt="open issues" />
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
      <a href="https://edoardoottavianelli.it">
        <img src="https://github.com/edoardottt/pwdsafety/workflows/Go/badge.svg?branch=master" alt="workflows" />
      </a>
    <!-- ubuntu-build -->
      <a href="https://edoardoottavianelli.it">
        <img src="https://github.com/edoardottt/images/blob/main/pwdsafety/ubuntu-build.svg" alt="ubuntu-build" />
      </a>
  <br>
    <!-- gobadge -->
      <a href="https://edoardoottavianelli.it">
        <img src="https://github.com/edoardottt/images/blob/main/pwdsafety/gobadge" alt="gobadge" />
      </a>
    <!-- license GPLv3.0 -->
      <a href="https://github.com/edoardottt/READMENATOR/blob/master/LICENSE">
        <img src="https://github.com/edoardottt/READMENATOR/blob/master/images/license-GPL3.svg" alt="license-GPL3" />
      </a>
</p>

**This tool doesn't store any information!!**  
**Remember, never use personal(or related to you) info in your password!**  
 - Use a password manager  
 - Don't use the same password for different services  
 - Enable 2FA when possible  
 
 Table of Contents 📽
 ------
 
 - [Example](https://github.com/edoardottt/pwdsafety#example-bar_chart)
 - [Get Started](https://github.com/edoardottt/pwdsafety#get-started-)
 - [Description](https://github.com/edoardottt/pwdsafety#description-)
 - [Scoring](https://github.com/edoardottt/pwdsafety#scoring-)
 - [Contributing](https://github.com/edoardottt/pwdsafety#contributing-)


Example :bar_chart:
----------

![Example](https://github.com/edoardottt/images/blob/main/pwdsafety/screen.gif)

Get Started 🎉
----------

- First of all, clone the repo locally

  - `git clone https://github.com/edoardottt/pwdsafety.git`

- Scilla has external dependencies, so they need to be pulled in:

  - `go get`

- Linux (Requires high perms, run with sudo)

  - `make linux`

  - `make unlinux`

- Windows (executable works only in pwdsafety folder. Alias?)

  - `make windows`

  - `make unwindows`

Description 🔦 
----------

It reads from standard input the entered password.

First, it searches in known-pwd.txt file if there is the password or the password reversed.

Then, just do little calculations, checking if the basic rules are respected, like if there are UPPERCASE CHARS, lowercase chars, numb3rs and symbols.

It stores the length of the password and the ratio [ unique different chars / total chars].

It calculates then the entropy of a password.

Password entropy is a measurement of how unpredictable a password is.

The formula for entropy is:
              ![formula](https://github.com/edoardottt/pwdsafety/blob/master/images/formula.png)
              
Where E = password entropy

R = pool of unique characters

L = number of characters in your password

Then R^L = the number of possible passwords

When the score <= 68(reasonable) it generates a random password using a list of all english words.

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

[![](https://sourcerer.io/fame/edoardottt/edoardottt/pwdsafety/images/0)](https://sourcerer.io/fame/edoardottt/edoardottt/pwdsafety/links/0)[![](https://sourcerer.io/fame/edoardottt/edoardottt/pwdsafety/images/1)](https://sourcerer.io/fame/edoardottt/edoardottt/pwdsafety/links/1)[![](https://sourcerer.io/fame/edoardottt/edoardottt/pwdsafety/images/2)](https://sourcerer.io/fame/edoardottt/edoardottt/pwdsafety/links/2)[![](https://sourcerer.io/fame/edoardottt/edoardottt/pwdsafety/images/3)](https://sourcerer.io/fame/edoardottt/edoardottt/pwdsafety/links/3)[![](https://sourcerer.io/fame/edoardottt/edoardottt/pwdsafety/images/4)](https://sourcerer.io/fame/edoardottt/edoardottt/pwdsafety/links/4)[![](https://sourcerer.io/fame/edoardottt/edoardottt/pwdsafety/images/5)](https://sourcerer.io/fame/edoardottt/edoardottt/pwdsafety/links/5)[![](https://sourcerer.io/fame/edoardottt/edoardottt/pwdsafety/images/6)](https://sourcerer.io/fame/edoardottt/edoardottt/pwdsafety/links/6)[![](https://sourcerer.io/fame/edoardottt/edoardottt/pwdsafety/images/7)](https://sourcerer.io/fame/edoardottt/edoardottt/pwdsafety/links/7)

Just open an issue/pull request. 

See also [CONTRIBUTING.md](https://github.com/edoardottt/pwdsafety/blob/master/CONTRIBUTING.md) and [CODE OF CONDUCT.md](https://github.com/edoardottt/pwdsafety/blob/master/CODE_OF_CONDUCT.md)



If you liked it drop a :star:
-------

https://www.edoardoottavianelli.it for contact me.


                                                                        Edoardo Ottavianelli
