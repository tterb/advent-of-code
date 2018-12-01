# While this is very space-efficient (no squares are skipped), requested data must be carried back to square 1 (the location of the only access port for this memory system) by programs that can only move up, down, left, or right. They always take the shortest path: the Manhattan Distance between the location of the data and square 1.

import math;

n = int(input());
dim = math.ceil(math.sqrt(n));
center = (math.ceil(dim/2), math.ceil(dim/2));
pos = (n % dim, dim);
steps = sum(tuple(x-y for x, y in zip((n % dim, dim), center)));
print(steps);
