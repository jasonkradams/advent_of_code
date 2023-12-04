#!/usr/bin/env python3

import re

# only 12 red cubes, 13 green cubes, and 14 blue cubes
max_cubes = {
    "red": 12,
    "green": 13,
    "blue": 14
}

def task_1():
    total = 0
    # [1] = Game Number
    # [2] = Game Results
    pattern = '^Game\s+(\d+):\s+(.*)$'
    with open("2023/02/input", mode="r") as input:
        for line in input.readlines():
            impossible_game = False
            match = re.match(pattern, line)
            game_number = int(match[1])
            match_results = match[2].split(";")
            match_results = [s.strip() for s in match_results]
            for set in match_results:
                if impossible_game == True:
                    break
                for h in set.split(","):
                    count, color = h.strip().split(" ", 1)
                    count = int(count)
                    if count > max_cubes.get(color):
                        impossible_game = True
                        break
            if impossible_game == True:
                continue
            total += game_number
    return total

if __name__ == "__main__":
    print(f"task 1: {task_1()}")
