#!/bin/python3

a = list(map(int, input().rstrip().split()))
b = list(map(int, input().rstrip().split()))
x = y = 0
for i in range(len(a)):
    if a[i]>b[i]:
        x += 1
    elif b[i] > a[i]:
        y += 1
print([x,y])
