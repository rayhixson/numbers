#!/bin/bash

cat ../numbers/1.txt | nc localhost 8888 &
cat ../numbers/2.txt | nc localhost 8888 &
cat ../numbers/3.txt | nc localhost 8888 &
cat ../numbers/4.txt | nc localhost 8888 &
cat ../numbers/5.txt | nc localhost 8888 &
