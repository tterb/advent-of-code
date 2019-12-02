import sys

def diff_by_one(a,b):
    diff = None
    for i in range(len(a)):
        if a[i] != b[i]:
            if diff:
                return None
            diff = i
    return diff
        

two, three = 0, 0
inp = [i.strip() for i in sys.stdin.readlines()]
for i in range(len(inp)-1):
    for j in range(i, len(inp)):
        index = diff_by_one(inp[i], inp[j])
        if index != None:
            print(inp[i]+'\n'+inp[j])
            print(inp[j][:index]+inp[j][index+1:])
            sys.exit()
