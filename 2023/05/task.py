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


def task_2(seeds: list) -> int:
    locations = []
    for seed, seed_range in pair_generator(seeds):
        location = process_seed_range(seed, seed_range)
        locations.append(location)

    return min(locations)


def pair_generator(lst: list):
    """Yield successive pairs from lst."""
    for i in range(0, len(lst), 2):
        # Yield a pair of elements
        yield lst[i], lst[i + 1]


def process_seed_range(seed, seed_range):
    # Process a range in chunks of 500 seeds at a time
    min_location = float("inf")
    for start in range(seed, (seed + seed_range), 50000):
        end = min(start + 50000, (seed + seed_range))
        local_min = min(get_location(seed=_) for _ in range(start, end))
        min_location = min(min_location, local_min)
    return min_location


def read_lines(f: str) -> list:
    return [_.strip().split(" ") for _ in open(f).readlines() if _.strip() != ""]

def parse_lines(maps: dict, input_lines: list):
    for line in input_lines:
        if line[0] == "seeds:":
            # ['seeds:', '79', '14', '55', '13']
            seeds = [int(num) for num in line[1:]]
        elif line[1] == "map:":
            # ['temperature-to-humidity', 'map:']
            conversion_map = line[0].replace("-", "_")
        elif len(line) == 3 and conversion_map is not None:
            # ['60', '56', '37']
            maps[conversion_map].append((line[0], line[1], line[2]))
        else:
            raise Exception(f"Unknown value: {l}")
    return seeds


def get_location(seed: int) -> int:
    soil = x_to_y(maps=efficient_maps["seed_to_soil"], x=seed)
    fertilizer = x_to_y(maps=efficient_maps["soil_to_fertilizer"], x=soil)
    water = x_to_y(maps=efficient_maps["fertilizer_to_water"], x=fertilizer)
    light = x_to_y(maps=efficient_maps["water_to_light"], x=water)
    temperature = x_to_y(maps=efficient_maps["light_to_temperature"], x=light)
    humidity = x_to_y(maps=efficient_maps["temperature_to_humidity"], x=temperature)
    location = x_to_y(maps=efficient_maps["humidity_to_location"], x=humidity)

    return location


def x_to_y(maps: dict, x: int) -> int:
    for source_range_start, (destination_range_start, range) in maps.items():
        if source_range_start <= x <= (source_range_start + (range - 1)):
            delta = x - source_range_start
            return destination_range_start + delta
    return x


if __name__ == "__main__":
    total_task_1 = 0
    total_task_2 = 0
    input_file = "2023/05/input_example"
    input_file = "2023/05/input"

    maps = {_: [] for _ in MAP_TYPES}
    input_lines = read_lines(input_file)
    seeds = parse_lines(maps, input_lines)
    # Convert maps to a more efficient structure if possible
    efficient_maps = {_: {int(source_range_start): (int(destination_range_start), int(range)) for destination_range_start, source_range_start, range in maps[_]} for _ in MAP_TYPES}

    total_task_1 = task_1(seeds=seeds)
    total_task_2 = task_2(seeds=seeds)

    print(f"Task 1 Total: {total_task_1}")
    print(f"Task 2 Total: {total_task_2}")
