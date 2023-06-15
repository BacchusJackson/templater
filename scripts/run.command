#!/bin/bash

cd $(dirname ${BASH_SOURCE[0]})

./templater | pbcopy

read -p "Press enter to continue"
