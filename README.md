<p align="center">
  <!-- logo -->
  <img src="https://github.com/edoardottt/pwdsafety/blob/master/images/logo.jpg"><br>
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
        <img src="https://github.com/edoardottt/pwdsafety/blob/master/images/version.svg" alt="version" />
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
        <img src="https://github.com/edoardottt/pwdsafety/blob/master/images/ubuntu-build.svg" alt="ubuntu-build" />
      </a>
  <br>
    <!-- gobadge -->
      <a href="https://edoardoottavianelli.it">
        <img src="https://github.com/edoardottt/pwdsafety/blob/master/images/gobadge" alt="gobadge" />
      </a>
    <!-- license GPLv3.0 -->
      <a href="https://github.com/edoardottt/READMENATOR/blob/master/LICENSE">
        <img src="https://github.com/edoardottt/READMENATOR/blob/master/images/license-GPL3.svg" alt="license-GPL3" />
      </a>
</p>

**This tool doesn't store any information!!**

**Remember, never use personal(or related to you) info in your password!**

Get Started 🎉
----------

- **Linux:**

  - After downloaded the package and extracted inside go/ folder, for comfort, create the alias:

  - Edit the file /home/{REPLACE_USER}/.bashrc and append this row:

       `alias pwdsafety="OLD=$(pwd) && cd YOUR-pwdsafety-PATH-FOLDER && go run main.go; cd $OLD; unset OLD"`

  - Then just type **pwdsafety** in your terminal where you want and press enter.

- **Other OS(Windows,MacOS and so on...):**

  - testing...
  - If you want to test it on these OS, let me know :)

Example :bar_chart:
----------

![Example](https://github.com/edoardottt/pwdsafety/blob/master/images/screen.jpg)

Description 🔦 
----------

It reads from standard input the entered password.

First, it searches in known-pwd.txt file if there is the password or the password reversed.

Then, just do little calculations, checking if the basic rules are respected, like if there are UPPERCASE CHARS, lowercase chars, numb3rs and symbols.

It stores the length of the password and the ratio [ unique different chars / total chars].

It calculates then the entropy of a password.

Password entropy is a measurement of how unpredictable a password is.

The formula for entropy is:
              ![CodeCogsEqn](https://github.com/edoardottt/pwdsafety/blob/master/images/CodeCogsEqn.gif)
              
Where E = password entropy

R = pool of unique characters

L = number of characters in your password

Then R^L = the number of possible passwords and the log of it is the number of bits of entropy.

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

Download 📡
----------

- GIT command on  CLI: `git -clone https://github.com/edoardottt/pwdsafety.git`

- Download by Browser: `https://github.com/edoardottt/pwdsafety`

- WGET on Linux(Recommended on linux): `wget https://github.com/edoardottt/pwdsafety/archive/master.zip`

Contributing 🛠
-------

Just open an issue/pull request. 

See also [CONTRIBUTING.md](https://github.com/edoardottt/pwdsafety/blob/master/CONTRIBUTING.md) and [CODE OF CONDUCT.md](https://github.com/edoardottt/pwdsafety/blob/master/CODE_OF_CONDUCT.md)



If you liked it drop a :star:
-------

https://www.edoardoottavianelli.it for contact me.


                                                                            Edoardo Ottavianelli
