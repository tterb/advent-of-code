# This hash function simulates tying a knot in a circle of string with 256 marks on it. Based on the input to be hashed, the function repeatedly selects a span of string, brings the ends together, and gives the span a half-twist to reverse the order of the marks within it. After doing this many times, the order of the marks is used to build the resulting hash.

seq = [x for x in range(256)];
lengths = [int(x) for x in input().split(",")];
pos, skip = 0, 0;

for j in lengths:
  for i in range(j // 2):
    n = (pos + i) % len(seq);
    temp = (pos + j - 1 - i) % len(seq);
    seq[n], seq[temp] = seq[temp], seq[n];
  pos += j + skip;
  skip += 1;

print(seq[0] * seq[1])
