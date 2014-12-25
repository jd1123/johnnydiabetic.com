# Johnnydiabetic.com source code

This is the source code for my personal website. Originally, I had a website written in Python with Django
but I decided to rewrite the website in Go. It has been a learning experience and is a lot of fun.
The only libraries used besides the Golang standard lib is gorilla mux (http://gorillatoolkit.com/pkg/mux) and 
mgo (http://gopkg.in/mgo). On the frontend, it's Twitter Bootstrap and JQuery.

## Why rewrite this in Go?
Well, mainly because Go rocks. But another reason is to get a little bit deeper into web programming. You
need to do more without the help of the framework, and eventually, I think I will probably use some other 
libraries when the going gets really tough. But for now, it's fun to write it all from scratch.

## What does it do?
Not much. Right now it has a pretty bland blog hooked into a Mongo database. Further down the line, I plan to
rewrite some of the other apps in Golang.
