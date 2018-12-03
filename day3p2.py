from sys import stdin
from collections import defaultdict

def checkOverlap(id,offset,dim,coords,unique):
    for i in range(dim[0]):
        for j in range(dim[1]):
            pos = (offset[0]+i, offset[1]+j)
            coords[pos].append(id)
            if len(coords[pos]) > 1:
                print(f'{coords[pos]}')
                [unique.remove(v) for v in coords[pos] if v in unique]
                    # if v in unique:
                    #     unique.remove(v)
                print(f'unique: {unique}')
                return unique
    unique.add(id)
    return unique
    
def removeOverlap(coords, count):
    unique = set(range(1, count+1))
    for i in coords.keys():
        if len(coords[i]) > 1:
            [unique.remove(j) for j in coords[i] if j in unique]
    return unique

count = 0
unique = set()
coords = defaultdict(list)
for claim in stdin.readlines():
    id = int(claim.split(' @ ')[0].strip()[1:])
    claim = claim.split(' @ ')[-1].strip()
    offset = [int(i) for i in claim.split(': ')[0].split(',')]
    dim = [int(i) for i in claim.split(': ')[1].split('x')]
    # unique = checkOverlap(id,offset,dim,coords,unique)
    for i in range(dim[0]):
        for j in range(dim[1]):
            coords[(offset[0]+i, offset[1]+j)].append(id)
    count = id
unique = removeOverlap(coords, count)
print(unique)
                
            
# print(sum(1 for i in coords.values() if i > 1))