#!/bin/bash

cd fixtures
git clone --recurse-submodules origin testrepo
cd ..
go test
