#!/usr/bin/env python3

import re

PATTERN = "^Card\s+(\d+):\s+([^|]+)\|\s+(.*)$"


def task_1(winning_numbers: list, selected_numbers: list) -> None:
    return sum_matching_numbers(winning_numbers=winning_numbers, selected_numbers=selected_numbers)


def update_grid(
    grid: dict, winning_numbers: list, selected_numbers: list, card_number: int
) -> list:
    grid_data = {
        "card_number": card_number,
        "remaining_cards": 1,
        "winning_numbers": winning_numbers,
        "selected_numbers": selected_numbers,
    }

    grid.update({f"card_{card_number}": grid_data})

    return grid


def match_line(line: str):
    match = re.match(pattern=PATTERN, string=line.strip())
    card_number = match[1]
    winning_numbers = [_.strip() for _ in match[2].split(" ") if _.strip() != ""]
    selected_numbers = [_.strip() for _ in match[3].split(" ") if _.strip() != ""]

    return card_number, winning_numbers, selected_numbers


def sum_matching_numbers(winning_numbers: list, selected_numbers: list) -> int:
    card_value = 0
    has_winning_number = False

    for n in selected_numbers:
        if n in winning_numbers:
            if has_winning_number:
                card_value = card_value * 2
            else:
                has_winning_number = True
                card_value = 1
    return card_value


def count_matching_numbers(winning_numbers: list, selected_numbers: list) -> int:
    match_count = 0
    for n in selected_numbers:
        if n in winning_numbers:
            match_count += 1
    return match_count


if __name__ == "__main__":
    total_task_1 = 0
    total_task_2 = 0
    input_file = "2023/04/input"
    card_grid = {}

    with open(input_file) as f:
        for line in f.readlines():
            card_number, winning_numbers, selected_numbers = match_line(line=line)

            total_task_1 += task_1(
                winning_numbers=winning_numbers, selected_numbers=selected_numbers
            )
            card_grid = update_grid(
                grid=card_grid,
                winning_numbers=winning_numbers,
                selected_numbers=selected_numbers,
                card_number=card_number,
            )

    for card_number in card_grid.keys():
        game = card_grid.get(card_number)
        for remaining_cards in range(game.get("remaining_cards"), 0, -1):
            selected_numbers = game.get("selected_numbers")
            winning_numbers = game.get("winning_numbers")
            match_count = count_matching_numbers(
                winning_numbers=winning_numbers, selected_numbers=selected_numbers
            )

            current_card_number = int(game.get("card_number"))
            for n in range(match_count):
                current_card_number += 1

                current_card = card_grid.get(f"card_{current_card_number}")
                remaining_cards = current_card.get("remaining_cards") + 1
                current_card["remaining_cards"] = remaining_cards
                card_grid[f"card_{current_card_number}"] = current_card

    for card in card_grid.keys():
        total_task_2 += card_grid[card]["remaining_cards"]

    print(f"Task 1 Total: {total_task_1}")
    print(f"Task 2 Total: {total_task_2}")
