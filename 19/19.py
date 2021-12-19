import numpy as np
import sys

X_ROTATIONS = np.array([
    [ # x0
        [1, 0, 0],
        [0, 1, 0],
        [0, 0, 1],
    ],
    [ # x90
        [1, 0, 0],
        [0, 0, 1],
        [0, -1, 0],
    ],
    [ # x180
        [1, 0, 0],
        [0, -1, 0],
        [0, 0, -1],
    ],
    [ # x270
        [1, 0, 0],
        [0, 0, -1],
        [0, 1, 0],
    ],
])

Y_ROTATIONS = np.array([
    [ # y0
        [1, 0, 0],
        [0, 1, 0],
        [0, 0, 1],
    ],
    [ # y90
        [0, 0, -1],
        [0, 1, 0],
        [1, 0, 0],
    ],
    [ # y180
        [-1, 0, 0],
        [0, 1, 0],
        [0, 0, -1],
    ],
    [ # y270
        [0, 0, 1],
        [0, 1, 0],
        [-1, 0, 0],
    ],
])

Z_ROTATIONS = np.array([
    [ # z0
        [1, 0, 0],
        [0, 1, 0],
        [0, 0, 1],
    ],
    [ # z90
        [0, 1, 0],
        [-1, 0, 0],
        [0, 0, 1],
    ],
    [ # z180
        [-1, 0, 0],
        [0, -1, 0],
        [0, 0, 1],
    ],
    [ # z270
        [0, -1, 0],
        [1, 0, 0],
        [0, 0, 1],
    ],
])

ALL_ROTS = []
for xr in X_ROTATIONS:
    for yr in Y_ROTATIONS:
        for zr in Z_ROTATIONS:
            ALL_ROTS.append(xr @ yr @ zr)

ALL_ROTS = np.unique(np.array(ALL_ROTS), axis=0)

def parse(s: str):
    coords = [l.split(",") for l in s.strip().split("\n") if not l.startswith("--")]
    return np.array(coords, dtype=int)

scanners = []
with open(sys.argv[1]) as f:
    for block in f.read().split("\n\n"):
        scanners.append(parse(block))

def permutations(beacons):
    # all possible rotations of beacons
    for rotation in ALL_ROTS:
        yield np.array([rotation @ pos for pos in beacons])

def canonical_if_can_be_matched(known, unknown):
    # Is there a rotation we can apply to unknown so that it lines up with known?
    for perm in permutations(unknown):
        match_shift = matches(known, perm)
        if match_shift is not None:
            return perm + match_shift, match_shift
    return None, None

def matches(a, b):
    # Do these two sets of beacons match?
    for ba in a:
        for bb in b:
            candidate_shift = ba - bb
            hits = 0
            for bb2 in b:
                shifted = (bb2 + candidate_shift)
                contains = (a[:, None] == shifted).all(-1).any(-1).any()
                if contains:
                    hits += 1
            if hits >= 12:
                print("MATCH", ba, bb, candidate_shift)
                return candidate_shift
    return None

def manhattan(a, b):
    return abs((b-a).sum())

known_scanners = [scanners[0]]
unknown_scanners = scanners[1:]
known_is = {0}
tried = set()
positions = [np.array([0, 0, 0])]

while len(known_scanners) < len(scanners):
    for unknown_i in range(len(scanners)):
        if unknown_i in known_is:
            continue
        unknown = scanners[unknown_i]
        for known_i in range(len(known_scanners)):
            if (unknown_i, known_i) in tried:
                continue
            known = known_scanners[known_i]
            matched, scanner_pos = canonical_if_can_be_matched(known, unknown)
            if matched is not None:
                known_scanners.append(matched)
                known_is.add(unknown_i)
                positions.append(scanner_pos)
                break
            tried.add((unknown_i, known_i))


beacons = np.unique(np.concatenate(known_scanners), axis=0)
print(len(beacons))


max_dist = 0
for i in range(len(positions)):
    for j in range(len(positions)):
        if i == j:
            continue
        m = manhattan(positions[i], positions[j])
        if m > max_dist:
            max_dist = m

print(max_dist)
