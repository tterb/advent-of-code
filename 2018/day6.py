from sys import stdin
from scipy.spatial.distance import cdist
import numpy as np

coords = [tuple([int(j) for j in i.split(', ')]) for i in stdin.readlines()]
axes_ranges = np.stack([np.min(coords, axis=0), np.max(coords, axis=0)])  # [min, max], not [min, max)
axes = [np.arange(axis[0], axis[1] + 1) for axis in axes_ranges.T]
grid = np.array(np.meshgrid(*axes, indexing='ij')).reshape(len(axes), -1).T  # cartesian product
border_idx = np.any(axes_ranges[:, np.newaxis] == grid, axis=(0, -1))  # indices of border locs

dists = cdist(grid, coords, metric='cityblock')
min_dists = np.min(dists, axis=1)
idx_arr = (min_dists[..., np.newaxis] == dists)

not_shared_idx = (np.sum(idx_arr, axis=1) == 1)
idx_arr = idx_arr[not_shared_idx]  # remove non-unique distances
border_idx = border_idx[not_shared_idx]

infinite = np.any(idx_arr[border_idx], axis=0)
area = np.sum(idx_arr, axis=0)
area[infinite] = -1

print(np.max(area))