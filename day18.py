# - snd X plays a sound with a frequency equal to the value of X.
# - set X Y sets register X to the value of Y.
# - add X Y increases register X by the value of Y.
# - mul X Y sets register X to the result of multiplying the value contained in register X by the value of Y.
# - mod X Y sets register X to the remainder of dividing the value contained in register X by the value of Y (that is, it sets X to the result of X modulo Y).
# - rcv X recovers the frequency of the last sound played, but only when the value of X is not zero. (If it is zero, the command does nothing.)
# - jgz X Y jumps with an offset of the value of Y, but only if the value of X is greater than zero. (An offset of 2 skips the next instruction, an offset of -1 jumps to the previous instruction, and so on.)

def getVal(n):
	try:
		val = int(n);
		return val;
	except ValueError:
		pass;
	return d[n];

d, inp = {}, [];
while True:
	try:
		data = input().split();
		d[data[1]] = 0;
		inp.append(data);
	except EOFError:
		break;

i, played = 0, 0;
while i < len(inp):
	op, reg, val = inp[i][0], inp[i][1], 0;
	if len(inp[i]) > 2:
		val = getVal(inp[i][2]);
	if op == 'snd':
		played = d[reg];
	elif op == 'set':
		d[reg] = val;
	elif op == 'add':
		d[reg] += val;
	elif op == 'mul':
		d[reg] = d[reg]*val;
	elif op == 'mod':
		d[reg] = d[reg]%val;
	elif op == 'rcv' and d[reg] != 0:
		print(played);
		i = len(inp);
	if op == 'jgz' and d[reg] > 0:
		i += val;
	else:
		i += 1;
