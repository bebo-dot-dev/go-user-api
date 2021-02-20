#!/usr/bin/env bash
git config --global user.email "joe@jjssoftware.co.uk"
git config --global user.name "Joe"
git fetch --all
git add .
git commit -m "automated codegen commit"
git pull --rebase
git push --force