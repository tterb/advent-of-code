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
while n < len(inst):
    op, a, b, c = inst[n]
    prev.append(n)
    if op in oper.keys():
        reg[c] = oper[op](reg, a, b)
    else:
        break
    print(f'({op} {a} {b} {c}): {reg}')
    reg[ip] += 1
    n = reg[ip]
print(reg[0])
