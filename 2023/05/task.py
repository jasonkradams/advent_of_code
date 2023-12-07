#!/usr/bin/env python3

MAP_TYPES = [
    "seed_to_soil",
    "soil_to_fertilizer",
    "fertilizer_to_water",
    "water_to_light",
    "light_to_temperature",
    "temperature_to_humidity",
    "humidity_to_location",
]


def task_1(seeds: list) -> int:
    return min([get_location(seed=seed) for seed in seeds])


def task_2(x: int, y: int) -> None:
    return 0


def read_lines(f: str) -> list:
    return [_.strip().split(" ") for _ in open(f).readlines() if _.strip() != ""]


def parse_lines(maps: dict, input_lines: list):
    for l in input_lines:
        if l[0] == "seeds:":
            # ['seeds:', '79', '14', '55', '13']
            seeds = [int(_) for _ in l[1:]]
        elif l[1] == "map:":
            # ['temperature-to-humidity', 'map:']
            conversion_map = l[0].replace("-", "_")
            maps.update({conversion_map: []})
        elif len(l) == 3:
            # ['60', '56', '37']
            maps[conversion_map].append((l[0], l[1], l[2]))
        else:
            raise Exception(f"Unknown value: {l}")
    return seeds


def get_location(seed: int) -> int:
    soil = x_to_y(maps=maps["seed_to_soil"], x=seed)
    fertilizer = x_to_y(maps=maps["soil_to_fertilizer"], x=soil)
    water = x_to_y(maps=maps["fertilizer_to_water"], x=fertilizer)
    light = x_to_y(maps=maps["water_to_light"], x=water)
    temperature = x_to_y(maps=maps["light_to_temperature"], x=light)
    humidity = x_to_y(maps=maps["temperature_to_humidity"], x=temperature)
    location = x_to_y(maps=maps["humidity_to_location"], x=humidity)

    return location


def x_to_y(maps: list, x: int) -> int:
    y = None
    for i, map in enumerate(maps):
        destination_range_start = int(map[0])
        source_range_start = int(map[1])
        range = int(map[2])

        if source_range_start <= x <= (source_range_start + (range - 1)):
            delta = int(x) - int(source_range_start)
            y = destination_range_start + delta
    return x if y is None else y


if __name__ == "__main__":
    total_task_1 = 0
    total_task_2 = 0
    input_file = "2023/05/input"
    # input_file = "2023/05/input_example"

    maps = {_: [] for _ in MAP_TYPES}
    input_lines = read_lines(input_file)
    seeds = parse_lines(maps, input_lines)
    total_task_1 = task_1(seeds=seeds)

    print(f"Task 1 Total: {total_task_1}")
    print(f"Task 2 Total: {total_task_2}")
