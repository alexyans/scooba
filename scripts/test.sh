#!/bin/bash

cd fixtures
rm -rf testrepo
git clone --recurse-submodules origin testrepo
cd ..
go test
