#!/usr/bin/env python3

from dataclasses import dataclass


def task_1(lines: str) -> int:
    for line in lines:
        if line[0].lower() == "time:":
            times = [int(_) for _ in line[1:]]
        if line[0].lower() == "distance:":
            records = [int(_) for _ in line[1:]]

    opportunities_list = []
    for i in range(len(times)):
        time = times[i]
        record = records[i]
        opportunities_list.append(winning_opportunities(time=time, record=record))

    total = opportunities_list.pop()
    for record in opportunities_list:
        total *= record

    return total


def task_2(seeds: list) -> int:
    pass


def read_lines(f: str) -> list[str]:
    return [_.strip().split() for _ in open(f).readlines() if _.strip() != ""]


def winning_opportunities(time: int, record: int) -> int:
    winning_rounds = 0
    for i in range(time):
        ms_remaining = time - i
        distance_per_ms = i
        total_distance = distance_per_ms * ms_remaining
        if total_distance > record:
            winning_rounds += 1

    return winning_rounds


if __name__ == "__main__":
    total_task_1 = 0
    total_task_2 = 0
    # input_file = "2023/06/input_example"
    input_file = "2023/06/input"

    lines = read_lines(input_file)
    total_task_1 = task_1(lines=lines)

    # total_task_2 = task_2()

    print(f"Task 1 Total: {total_task_1}")
    print(f"Task 2 Total: {total_task_2}")
