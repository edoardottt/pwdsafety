# 🔒pwd-safety🔒

This is a light command line tool that checks how much a password is safe.

**This tool doesn't store any information!**

![gobadge](https://github.com/edoardottt/pwd-safety/blob/devel/Images/gobadge)


DESCRIPTION 🔦 
-------------------------------------------------

Just write go run man.go in your terminal inside ../pwd-safety$

It reads from standard input the entered password.

First, it searches in the known-pwd.txt file if there is the password or the reverse password.

Then, just do little calculations, checking if the basic rules are respected. 

Like if there are UPPERCASE CHARS, lowercase chars, numb3rs and symbols.

It stores the length of the password and the ratio [ unique different chars / total chars].

It calculates then the entropy of a password.

Password entropy is a measurement of how unpredictable a password is.

The formula for entropy is:
              ![CodeCogsEqn](https://github.com/edoardottt/pwd-safety/blob/devel/Images/CodeCogsEqn.gif)
              
Where E = password entropy

R = pool of unique characters

L = number of characters in your password

Then R^L = the number of possible passwords and the log of it is the number of bits of entropy.


SCORING 💯
-------------------------------------------------

Max score: 100

Scores:
  - Very weak: 0 - 35
  - Weak: 36 - 59
  - Reasonable: 60 - 68
  - Strong: 69 - 80
  - Very strong: 81 -100
  
Scoring parameters:
  - Found in known password
  - Found in known password reverse
  - Composition:
      - numbers
      - symbols
      - uppercase
      - lowercase
  - Unique different characters
  - Length
  - Entropy

SOME EXAMPLES :bar_chart:
-------------------------------------------------

![veryWeak](https://github.com/edoardottt/pwd-safety/blob/devel/Images/veryWeak.png)

![weak](https://github.com/edoardottt/pwd-safety/blob/devel/Images/weak.png)

![reasonable](https://github.com/edoardottt/pwd-safety/blob/devel/Images/reasonable.png)

![strong](https://github.com/edoardottt/pwd-safety/blob/devel/Images/strong.png)

![veryStrong](https://github.com/edoardottt/pwd-safety/blob/devel/Images/veryStrong.png)


DOWNLOAD 📡
-------------------------------------------------

- GIT command on prompt: 
            
            git -clone https://github.com/edoardottt/pwd-safety.git

- Download by Browser on: 

https://github.com/edoardottt/pwd-safety

--------------------------
If you liked it drop a :star:
--------------------------

https://www.edoardoottavianelli.it for contact me.


                                      Edoardo Ottavianelli ©
