#!/bin/bash

go build main.go

sudo rm /kodybin/*
sudo cp main /kodybin/main
sudo ln -s /kodybin/main /kodybin/ls
sudo ln -s /kodybin/main /kodybin/cat

