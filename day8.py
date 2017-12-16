# Each instruction consists of several parts: the register to modify, whether to increase or decrease that register's value, the amount by which to increase or decrease it, and a condition. If the condition fails, skip the instruction without modifying the register.
# What is the largest value in any register after completing the instructions in your puzzle input?

import operator

d = {};
oper = { 
	"<": operator.lt, 
	"<=": operator.le, 
	">": operator.gt, 
	">=": operator.ge, 
	"!=": operator.ne, 
	"==": operator.eq,
	"inc": operator.add,
	"dec": operator.sub
}

while True:
	try:
		inst, cond = [i.split() for i in input().split(' if ')];
		if d.get(cond[0]) == None:
			d[cond[0]] = 0;
		if d.get(inst[0]) == None:
			d[inst[0]] = 0;
		if oper[cond[1]](d[cond[0]], int(cond[2])):
			d[inst[0]] = oper[inst[1]](d[inst[0]], int(inst[2]));
	except EOFError:
		break;
print(max(d.values()));