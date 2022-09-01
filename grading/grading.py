#!/bin/python3

from turtle import mode


grades = sum(list(map(int, input().rstrip().split())))
moderation = int(input().rstrip())
min_val = int(input().rstrip())
print([i if i<min_val or i%moderation <= 2 else i+(moderation-i%moderation)  for i in grades])
