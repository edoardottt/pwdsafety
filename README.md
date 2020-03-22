# 🔒pwd-safety🔒

This is a light command line tool that checks how much a password is safe.

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

##############################################à
##################################################SCORINGGGGG

-------------------------------------------------
SOME EXAMPLES :bar_chart:
-------------------------------------------------
**Take a look**

//IMAGES


-------------------------------------------------
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
