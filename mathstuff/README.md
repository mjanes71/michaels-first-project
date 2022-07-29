# mathstuff

Since you like math so much, I thought we could look at a go project that does some math!

## What are these files?

`main.go` 
This is the file that will have the main code/program in it. 

`main_test.go`
This is the file that will have unit tests for the function(s) in the `main.go` file.

`go.mod`
This is the file to help the project keep track of what modules(imports) we are using

`go.sum`
Keeps track of what you have installed locally so you don't always have to reinstall imports.

Let's take a look at the `main.go` first. Then we'll take a look in the `main_test.go`

## Git Help

1. To clone this repository
* open a terminal window
* run `ls` to list all the items in your current location
* run `cd Documents` to change directory into your Documents folder
* run `git clone ______`: the blank is for the url of the repo

2. To make a new branch to work on
* run `ls` to see your new repo
* run `cd michaels-first-project` to change directory into your new repo
* run `git checkout -b 'practice-branch'` to make a new branch called practice-branch

3. To push your local changes back up to github
* run `git status` to see what changes you've made
* run `git add .` to add all your changes to a list to be committed
* run `git commit -m 'my first changes'` to commit your changes
* run `git push -u origin practice-branch` to push your changes up to github