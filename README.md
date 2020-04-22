# 🔒pwd-safety🔒

![gobadge](https://github.com/edoardottt/pwd-safety/blob/master/Images/gobadge)

This is a command line tool that checks how much a password is safe.

<p align="center">
  <img src="https://github.com/edoardottt/pwd-safety/blob/master/Images/logo.png">
</p>

**This tool doesn't store any information!!**

**If you want to use it anywhere, type on cli**

"alias pwdsafety='cd /home/<YOUR_USER>/go/pwd-safety && go run main.go && cd $OLDPWD'"

DESCRIPTION 🔦 
-------------------------------------------------

Just write **go run main.go** in your terminal inside ../pwd-safety$

So: ../pwd-safety$ go run main.go

It reads from standard input the entered password.

First, it searches in known-pwd1.txt and known-pwd2.txt files if there is the password or the reverse password.

The calculation isn't immediate, **It searches for over 10.5 million hacked passwords**

Then, just do little calculations, checking if the basic rules are respected, like if there are UPPERCASE CHARS, lowercase chars, numb3rs and symbols.

It stores the length of the password and the ratio [ unique different chars / total chars].

It calculates then the entropy of a password.

Password entropy is a measurement of how unpredictable a password is.

The formula for entropy is:
              ![CodeCogsEqn](https://github.com/edoardottt/pwd-safety/blob/master/Images/CodeCogsEqn.gif)
              
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

SOME EXAMPLES :bar_chart:
-------------------------------------------------
- **Very weak password:**

![veryWeak](https://github.com/edoardottt/pwd-safety/blob/master/Images/veryWeak.png)

- **Weak password:**

![weak](https://github.com/edoardottt/pwd-safety/blob/master/Images/weak.png)

- **Reasonable password**

![reasonable](https://github.com/edoardottt/pwd-safety/blob/master/Images/reasonable.png)

- **Strong password**

![strong](https://github.com/edoardottt/pwd-safety/blob/master/Images/strong.png)

- **Very strong password**

![veryStrong](https://github.com/edoardottt/pwd-safety/blob/master/Images/veryStrong.png)


DOWNLOAD 📡
-------------------------------------------------

- GIT command on prompt: 
            
            git -clone https://github.com/edoardottt/pwd-safety.git

- Download by Browser on: 

https://github.com/edoardottt/pwd-safety


VERSIONING :books:
--------------------------------------------

**[v0.1.1](https://github.com/edoardottt/pwd-safety/releases/tag/v0.1.1):**
  
          - Added:
                  - Generates random password when score <= REASONABLE
                  - Errors code table

**[v0.1](https://github.com/edoardottt/pwd-safety/releases/tag/v0.1):**
  
          - First pre-release

--------------------------
If you liked it drop a :star:
--------------------------

https://www.edoardoottavianelli.it for contact me.


                                      Edoardo Ottavianelli ©
