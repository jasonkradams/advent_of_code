#!/usr/bin/env python3

import re

# only 12 red cubes, 13 green cubes, and 14 blue cubes
MAX_CUBES = {
    "red": 12,
    "green": 13,
    "blue": 14
}

PATTERN = '^Game\s+(\d+):\s+(.*)$'

def split_line(line: str):
    # match[1] = Game Number
    # match[2] = Game Results
    match = re.match(PATTERN, line)
    game_number = int(match[1])
    match_results = match[2].split(";")
    match_results = [s.strip() for s in match_results]

    return game_number, match_results

def task_1(line: str):
    game_number, match_results = split_line(line)
    for set in match_results:
        for h in set.split(","):
            count, color = h.strip().split(" ", 1)
            count = int(count)
            if count > MAX_CUBES.get(color):
                return 0
    return game_number

def task_2(line):
    min_required = {
        "red": 0,
        "green": 0,
        "blue": 0,
    }
    game_number, match_results = split_line(line)
    for set in match_results:
        for h in set.split(","):
            count, color = h.strip().split(" ", 1)
            count = int(count)
            if count > min_required.get(color):
                min_required[color] = count

    return min_required.get("red") * min_required.get("green") * min_required.get("blue")

if __name__ == "__main__":
    total_task_1 = 0
    total_task_2 = 0

    with open("2023/02/input", mode="r") as input:
        for line in input.readlines():
            total_task_1 += task_1(line)
            total_task_2 += task_2(line)
    

    print(f"Task 1: {total_task_1}")
    print(f"Task 2: {total_task_2}")
