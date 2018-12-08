# --- Day 5: Alchemical Reduction ---

inp = list(input().strip())
i = len(inp)-1
while(i > 0):
  if abs(ord(inp[i])-ord(inp[i-1])) == 32:
    inp.pop(i)
    inp.pop(i-1)
    if i < len(inp)-2:
      i += 2
    else:
      i = len(inp)-1
  i -= 1
if abs(ord(inp[len(inp)-2])-ord(inp[len(inp)-1])) == 32:
  inp = inp[:-2]
print(''.join(inp))
print(len(inp))
