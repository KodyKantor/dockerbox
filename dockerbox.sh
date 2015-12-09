#!/bin/bash

binfile=/kodybin/dockerbox

mkdir /kodybin
cp /dockerbox $binfile
ln -s $binfile /kodybin/ls
ln -s $binfile /kodybin/cat
ln -s $binfile /kodybin/cc
