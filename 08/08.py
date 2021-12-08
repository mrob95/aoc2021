import fileinput
from dataclasses import dataclass
from typing import Counter, Dict, List, Set
from collections import Counter

@dataclass
class Entry:
    patterns: List[str]
    outputs: List[str]

def parse_entries():
    result = []
    for line in fileinput.input():
        pats, outs = line.split(" | ")
        result.append(Entry(pats.split(" "), outs.split(" ")))
    return result

def part_one(es: List[Entry]) -> int:
    result = 0
    for e in es:
        result += sum([len(o) in {2, 4, 3, 7} for o in e.outputs])
    return result

# Reference:
#
#  dddd
# e    a
# e    a
#  ffff
# g    b
# g    b
#  cccc
#

# Map of "number of segments lit" to "segments used" for the four cases where a number requires a unique number of segments
DISTINCT_LENGTHS = {
    2: set("ab"),
    4: set("efab"),
    3: set("dab"),
    7: set("abcdefg"),
}

# Map of "number of times a letter appears in the input patterns" to "segment it maps to" for the ones where this is 1->1
DISTINCT_COUNTS = {
    9: "b",
    6: "e",
    4: "g",
}

# Map of alphabetically sorted segments to the values they represent
DISPLAY_MAP = {
    "abcdeg": "0",
    "ab": "1",
    "acdfg": "2",
    "abcdf": "3",
    "abef": "4",
    "bcdef": "5",
    "bcdefg": "6",
    "abd": "7",
    "abcdefg": "8",
    "abcdef": "9",
}

def get_mapping(patterns: List[str]) -> Dict[str, str]:
    options = {l: set("abcdefg") for l in "abcdefg"}
    counts = Counter("".join(patterns))
    # Some can be figured out by looking at the counts of values in the patterns, e.g.
    # only segment "b" in the diagram above appears in 9 digits
    for l in "abcdefg":
        if counts[l] in DISTINCT_COUNTS:
            mapped = DISTINCT_COUNTS[counts[l]]
            # "l" maps to "mapped" - remove it from all other options
            for k in options:
                options[k] -= set(l)
            options[mapped] = set(l)

    # The rest can be figured out by following the logic in part one - some numbers require a unique number of segments
    for pattern in patterns:
        if len(pattern) in DISTINCT_LENGTHS:
            for l in options:
                if l in DISTINCT_LENGTHS[len(pattern)]:
                    options[l] &= set(pattern)
                else:
                    options[l] -= set(pattern)

    # Return a mapping from the letters in the input patterns to the canonical ones above
    return {next(iter(s)): l for l, s in options.items()}


def part_two(es: List[Entry]) -> int:
    result = 0
    for e in es:
        mapping = get_mapping(e.patterns)
        digits = []
        for o in e.outputs:
            s = "".join(sorted([mapping[l] for l in o.strip()]))
            digits.append(DISPLAY_MAP[s])
        result += int("".join(digits))
    return result

entries = parse_entries()

print(part_one(entries))
print(part_two(entries))
