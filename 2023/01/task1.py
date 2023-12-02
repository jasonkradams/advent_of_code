#!/usr/bin/env python3


def first_num(line: str):
    for c in line:
        if c.isdigit():
            return c


def last_num(line: str):
    line = line[::-1]
    for c in line:
        if c.isdigit():
            return c


if __name__ == "__main__":
    sum = 0
    with open("2023/01/input", mode="r") as input:
        for line in input.readlines():
            first = first_num(line)
            last = last_num(line)
            sum += int((first + last))

    print(sum)
