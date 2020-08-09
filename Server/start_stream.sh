#!/bin/bash
raspivid -ih -a 4 -a 'OSS+ROV %d/%m/%Y %X' -vf -hf -n -w 960 -h 480 -fps 25 -b 17000000 -pf high -lev 4 -drc high -t 0 -o -| nc 192.168.0.29 5001
