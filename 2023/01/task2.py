#!/usr/bin/env python3

numbers = {
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9",
}


class Calibration:
    def __init__(self, text: str = ""):
        self.text = text
        self.low_int = ""
        self.high_int = ""
        self.low_str = ""
        self.high_str = ""

        self.low_int_idx = None
        self.high_int_idx = None
        self.low_str_idx = None
        self.high_str_idx = None

        self._find_low_int()
        self._find_high_int()
        self._find_low_str()
        self._find_high_str()

    def _find_low_int(self):
        line = self.text
        for c in line:
            if c.isdigit():
                self.low_int_idx = line.find(c)
                self.low_int = c
                return

    def _find_high_int(self):
        line = self.text[::-1]
        for c in line:
            if c.isdigit():
                self.high_int = c
                self.high_int_idx = self.text.rfind(c)
                return

    def _find_low_str(self):
        for num in numbers.keys():
            if line.find(num) != -1:
                i = line.find(num)

                if self.low_str_idx is None:
                    self.low_str = num
                    self.low_str_idx = line.find(num)
                    continue

                if i < self.low_str_idx:
                    self.low_str = num
                    self.low_str_idx = line.find(num)
                    continue

    def _find_high_str(self, sub_line=None):
        for num in numbers.keys():
            if line.rfind(num) != -1:
                i = line.rfind(num)
                if self.high_str_idx is None:
                    self.high_str = num
                    self.high_str_idx = line.rfind(num)
                    continue

                if i > self.high_str_idx:
                    self.high_str = num
                    self.high_str_idx = line.rfind(num)
                    continue

    def _low_digit(self):
        if self.low_str_idx is None:
            return self.low_int

        if self.low_int_idx is None:
            return numbers.get(self.low_str)

        if self.low_int_idx < self.low_str_idx:
            return self.low_int

        return numbers.get(self.low_str)

    def _high_digit(self):
        if self.high_str_idx is None:
            return self.high_int

        if self.high_int_idx is None:
            return numbers.get(self.high_str)

        if self.high_int_idx > self.high_str_idx:
            return self.high_int

        return numbers.get(self.high_str)

    def result(self):
        return int(self._low_digit() + self._high_digit())


if __name__ == "__main__":
    sum = 0
    with open("2023/01/input", mode="r") as input:
        for line in input.readlines():
            calibration = Calibration(text=line)
            sum += calibration.result()

    print(sum)
