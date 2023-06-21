#!/bin/bash

# read from template file
template_file="template.go"

echo $template

day=$1

mkdir $1
cd $1

# initialize go module
go mod init $1

# create files
touch demo.txt
touch input.txt
cp ../$template_file ./
mv $template_file main.go



