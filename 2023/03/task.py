#!/usr/bin/env python3

def task_1(x: int, y: int) -> None:
    if searched_grid[x][y]:
        return 0
    if is_int(x=x, y=y):
        num = char_grid[x][y]
        lnum = search_left(x=x, y=y)
        rnum =search_right(x=x, y=y)
        num = lnum + num + rnum

        return int(num)
    return 0

def is_symbol(char: str):
    return (char.isalnum() == False and char != '.')

def create_grids(input: str):
    char_grid = []
    symbol_grid = []
    searched_grid = []
    
    with open(input, mode="r") as input:
        for x, line in enumerate(input, start=0):
            """
            Create three grids:
              - One for mapping all characters
              - One for mapping only symbols
              - One to track which cells have already been searched as a number
            """
            char_grid.append([])
            symbol_grid.append([])
            searched_grid.append([])

            for char in line.strip():
                char_grid[x].append(char)
                symbol_grid[x].append(is_symbol(char))
                searched_grid[x].append(False)

    return char_grid, symbol_grid, searched_grid

def is_int(x: int, y: int) -> bool:
    return char_grid[x][y].isdigit()

def get_adjacent_cells(x: int, y: int) -> list:
    return [
        ((x - 1), (y - 1)),
        ((x - 1), (y)),
        ((x - 1), (y + 1)),
        ((x), (y - 1)),
        (x, (y + 1)),
        ((x + 1), (y - 1)),
        ((x + 1), (y)),
        ((x + 1), (y + 1))
    ]

def searched(x: int, y: int) -> None:
    """
    Mark a cell as having been searched.
    This prevents us from trying to incorporate the same number in multiple results.
    """
    try:
        searched_grid[x][y] = True
    except IndexError:
        return None

def get_char(x: int, y: int) -> str:
    try:
        char = char_grid[x][y]
    except IndexError:
        return ""
    return char

def search_left(x: int, y: int) -> str:
    num = ""
    y = y - 1
    char = get_char(x=x, y=y)
    searched(x=x, y=y)

    while char.isdigit():
        num = char + num
        y = y - 1
        char = get_char(x=x, y=y)
        searched(x=x, y=y)
    return num

def search_right(x: int, y: int):
    num = ""
    y = y + 1
    char = get_char(x=x, y=y)
    searched(x=x, y=y)

    while char.isdigit():
        num = num + char
        y = y + 1
        char = get_char(x=x, y=y)
        searched(x=x, y=y)
    return num

if __name__ == "__main__":
    total_task_1 = 0
    total_task_2 = 0
    input_file = "2023/03/input"

    char_grid, symbol_grid, searched_grid = create_grids(input_file)

    for x, line in enumerate(symbol_grid):
        for y, is_symbol in enumerate(line):
            if is_symbol:
                # Found a symbol, Search adjacent spaces for a number
                # ...
                # .*.
                # ...
                total_task_1 += sum([task_1(x=cx, y=cy) for cx, cy in get_adjacent_cells(x=x, y=y)])
                    
print(f"Task 1 Total: {total_task_1}")
print(f"Task 2 Total: {total_task_2}")
