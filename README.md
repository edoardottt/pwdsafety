# 🔒pwdsafety🔒

![gobadge](https://github.com/edoardottt/pwdsafety/blob/master/images/gobadge)
![version](https://github.com/edoardottt/pwdsafety/blob/master/images/version.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/edoardottt/pwdsafety)](https://goreportcard.com/report/github.com/edoardottt/pwdsafety)
![Go](https://github.com/edoardottt/pwdsafety/workflows/Go/badge.svg?branch=master)
![ubuntu-build](https://github.com/edoardottt/pwdsafety/blob/master/images/ubuntu-build.svg)


**Command line tool that checks how much a password is safe.**

<p align="center">
  <img src="https://github.com/edoardottt/pwdsafety/blob/master/images/logo.png">
</p>

**This tool doesn't store any information!!**

**Remember, never use personal(or related to you) info in your password!**

USAGE 💡
-------------------------------------------------

- **Linux:**

  - After downloaded the package and extracted inside go/ folder, for comfort, create the alias:

  - Edit the file /home/{REPLACE_USER}/.bashrc and append this row:

        alias pwdsafety='cd YOUR-pwdsafety-PATH-FOLDER && go run main.go && cd $OLDPWD'

  - Then just type **pwdsafety** in your terminal where you want and press enter.

- **Other OS(Windows,MacOS and so on...):**

  - testing...

SOME EXAMPLES :bar_chart:
-------------------------------------------------
- **Very weak password:**

![veryWeak](https://github.com/edoardottt/pwdsafety/blob/master/images/veryWeak.png)

- **Weak password:**

![weak](https://github.com/edoardottt/pwdsafety/blob/master/images/weak.png)

- **Reasonable password**

![reasonable](https://github.com/edoardottt/pwdsafety/blob/master/images/reasonable.png)

- **Strong password**

![strong](https://github.com/edoardottt/pwdsafety/blob/master/images/strong.png)

- **Very strong password**

![veryStrong](https://github.com/edoardottt/pwdsafety/blob/master/images/veryStrong.png)

DESCRIPTION 🔦 
-------------------------------------------------

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

SCORING 💯
-------------------------------------------------

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

DOWNLOAD 📡
-------------------------------------------------

- GIT command on  CLI: `git -clone https://github.com/edoardottt/pwdsafety.git`

- Download by Browser: `https://github.com/edoardottt/pwdsafety`

- WGET on Linux(Recommended on linux): `wget https://github.com/edoardottt/pwdsafety/archive/master.zip`

--------------------------
If you liked it drop a :star:
--------------------------

https://www.edoardoottavianelli.it for contact me.


                                                                            Edoardo Ottavianelli ©
