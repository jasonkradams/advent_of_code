#!/usr/bin/env python3

import re

PATTERN = '^Card\s+(\d+):\s+([^|]+)\|\s+(.*)$'

def task_1(line: str) -> None:
    card_value = 0
    match = re.match(pattern=PATTERN, string=line.strip())
    winning_numbers = [_.strip() for _ in match[2].split(' ') if _.strip() != '']
    your_numbers = [_.strip() for _ in match[3].split(' ') if _.strip() != '']

    has_winning_number = False
    for n in your_numbers:
        if n in winning_numbers:
            if has_winning_number:
                card_value = card_value * 2
            else:
                has_winning_number = True
                card_value = 1
    return card_value

if __name__ == "__main__":
    total_task_1 = 0
    total_task_2 = 0
    input_file = "2023/04/input"

    with open(input_file) as f:
        for line in f.readlines():
            total_task_1 += task_1(line=line)

print(f"Task 1 Total: {total_task_1}")
print(f"Task 2 Total: {total_task_2}")
