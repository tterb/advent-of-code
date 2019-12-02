from sys import stdin
import re

oper = {
    'addr': lambda reg, a, b: reg[a] + reg[b],
    'addi': lambda reg, a, b: reg[a] + b,
    'mulr': lambda reg, a, b: reg[a] * reg[b],
    'muli': lambda reg, a, b: reg[a] * b,
    'banr': lambda reg, a, b: reg[a] & reg[b],
    'bani': lambda reg, a, b: reg[a] & b,
    'borr': lambda reg, a, b: reg[a] | reg[b],
    'bori': lambda reg, a, b: reg[a] | b,
    'setr': lambda reg, a, b: reg[a],
    'seti': lambda reg, a, b: a,
    'gtir': lambda reg, a, b: 1 if a > reg[b] else 0,
    'gtri': lambda reg, a, b: 1 if reg[a] > b else 0,
    'gtrr': lambda reg, a, b: 1 if reg[a] > reg[b] else 0,
    'eqir': lambda reg, a, b: 1 if a == reg[b] else 0,
    'eqri': lambda reg, a, b: 1 if reg[a] == b else 0,
    'eqrr': lambda reg, a, b: 1 if reg[a] == reg[b] else 0,
}

def decipherCodes(monitor, opcodes):
    opts = dict()
    found = set()
    for i in monitor:
        before, after = i[0], i[2]
        op = i[1][0]
        a,b,c = i[1][1], i[1][2], i[1][3]
        if op not in opts.keys():
            opts[op] = set(oper.keys())
        possible = list(opts[op])
        for j in possible:
            if j in opts[op] and oper[j](before,a,b) != after[c]:
                opts[op].remove(j)
        if len(opts[op]) == 1:
            opcodes[list(opts[op])[0]] = op
            found.add(list(opts[op])[0])
    # Narrow down possible operations
    incomplete = True
    while incomplete:
        incomplete = False
        for i in opts.keys():
            opts[i] = opts[i].difference(found)
            if len(opts[i]) == 1:
                opcodes[list(opts[i])[0]] = i
                found.add(list(opts[i])[0])
                incomplete = True
    return dict((v,k) for k,v in opcodes.items())
        

inp = [i.strip() for i in stdin.readlines()]
opts = list(oper.keys())
monitor, ops = list(), list()
i = 0
while 'Before:' in inp[i]:
    temp = list()
    temp.append([int(j) for j in re.findall('\d+', inp[i])])
    temp.append([int(j) for j in inp[i+1].split()])
    temp.append([int(j) for j in re.findall('\d+', inp[i+2])])
    monitor.append(temp)
    i += 4
opcodes = decipherCodes(monitor, {i:None for i in opts})
reg = [0, 0, 0, 0]
for i in range(i+2, len(inp)):
    code,a,b,c = [int(j) for j in inp[i].split()]
    print(f'{opcodes[code]}, {a}, {b}, {c}')
    reg[c] = oper[opcodes[code]](reg, a, b)
print(reg[0])
