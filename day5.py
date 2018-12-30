# --- Day 5: Alchemical Reduction ---

def react(poly):
  i = len(poly)-1
  while(i > 0):
    if abs(ord(poly[i])-ord(poly[i-1])) == 32:
      poly.pop(i)
      poly.pop(i-1)
      if i < len(poly)-2:
        i += 2
      else:
        i = len(poly)-1
    i -= 1
  if abs(ord(poly[len(poly)-2])-ord(poly[len(poly)-1])) == 32:
    poly = poly[:-2]
  return poly

# poly = list(input().strip())
poly = react(list(input().strip()))
print(f'Part 1: {len(poly)}')
minVal = len(poly)
for i in range(65, 91):
  mod = list(filter((chr(i+32)).__ne__, list(filter((chr(i)).__ne__, poly))))
  val = len(react(mod))
  if val < minVal:
    minVal = val
print(f'Part 2: {minVal}')
    
    
  
  
