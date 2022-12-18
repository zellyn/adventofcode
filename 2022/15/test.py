import numpy as np
limit = 4_000_000
count = 9_000_000
xs = np.array(range(300_000,300_000+count))
print(len(xs))
ys = np.array(range(-2_000_000,-2_000_000+count))
print(len(ys))
all_good = np.logical_and(np.logical_and(xs>=0, xs<=limit), np.logical_and(xs>=0, xs<=limit))
print(len(all_good))
xs = xs[all_good]
ys = ys[all_good]
print(len(xs))
print(len(ys))
