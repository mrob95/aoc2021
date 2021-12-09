import fileinput
import numpy as np
from scipy import ndimage

a = np.array([list(s.strip()) for s in fileinput.input()], dtype=int)
up    = np.array([[0, -1, 0], [0, 1, 0], [0, 0, 0]])
right = np.array([[0, 0, 0], [0, 1, -1], [0, 0, 0]])
down  = np.array([[0, 0, 0], [0, 1, 0], [0, -1, 0]])
left  = np.array([[0, 0, 0], [-1, 1, 0], [0, 0, 0]])
lowpoints = (ndimage.convolve(a, up, mode='constant', cval=10) < 0) & \
            (ndimage.convolve(a, right, mode='constant', cval=10) < 0) & \
            (ndimage.convolve(a, down, mode='constant', cval=10) < 0) & \
            (ndimage.convolve(a, left, mode='constant', cval=10) < 0)

print(lowpoints.sum() + a[lowpoints].sum())

highpoints = a == 9
labels, num_labels = ndimage.measurements.label(~highpoints)
area_sizes = sorted([(labels == label).sum() for label in range(1, num_labels + 1)])

print(area_sizes[-1] * area_sizes[-2] * area_sizes[-3])



# Using scipy felt like cheating, so here is a simple recursive floodfill...
def fill(x: int, y: int, val: int, a: np.ndarray) -> int:
    result = 1
    a[x, y] = val
    for next_x, next_y in [(x + 1, y), (x - 1, y), (x, y + 1), (x, y - 1)]:
        if next_x < 0 or next_y < 0 or next_x >= len(a) or next_y >= len(a[0]):
            continue # Out of bounds
        if a[next_x, next_y] != 0:
            continue # Already filled
        result += fill(next_x, next_y, val, a)
    return result

labels = np.zeros_like(a)
labels[highpoints] = -1
sizes = []

val = 1
for x, y in np.argwhere(lowpoints):
    sizes.append(fill(x, y, val, labels))
    val += 1

sorted_sizes = sorted(sizes)

print(sorted_sizes[-1] * sorted_sizes[-2] * sorted_sizes[-3])
