from sys import stdin

vals = {int(i) for i in stdin.readlines()}
arr = list(vals)
for i in range(len(arr) - 1):
    for j in range(i, len(arr)):
        if (2020 - arr[i] - arr[j]) in vals:
            print(arr[i] * arr[j] * (2020 - arr[i] - arr[j]))
            exit(0)
