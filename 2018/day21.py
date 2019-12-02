from sys import stdin

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

inp = [i.strip() for i in stdin.readlines()]
ip = int(inp.pop(0).split()[1])
inst = [[i.split()[0],int(i.split()[1]),int(i.split()[2]),int(i.split()[3])] for i in inp]
reg = [0,0,0,0,0,0]
n = reg[ip]
values = set()
prev = None
while n < len(inst):
    op, a, b, c = inst[n]
    reg[c] = oper[op](reg, a, b)
    # print(f'{n}: ({op} {a} {b} {c}): {reg}')
    reg[ip] += 1
    n = reg[ip]
    if n == 29:
        if reg[5] not in values:
            # print(reg[5])
            values.add(reg[5])
            prev = reg[5]
        else:
            print(prev)
            break
print(len(values))