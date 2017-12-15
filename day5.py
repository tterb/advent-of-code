# The message includes a list of the offsets for each jump. Jumps are relative: -1 moves to the previous instruction, and 2 skips the next one. Start at the first instruction in the list. The goal is to follow the jumps until one leads outside the list.
# In addition, these instructions are a little strange; after each jump, the offset of that instruction increases by 1. So, if you come across an offset of 3, you would move three instructions forward, but change it to a 4 for the next time it is encountered.

i, steps, data = 0, 0, [];
while True:
	try:
		data.append(int(input()));
	except EOFError:
		break;
while(i < len(data)):
	temp = i;
	i += data[i]
	data[temp] += 1;
	steps += 1;
print(steps);