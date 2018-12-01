# Packet Scanners

rng = {};
while True:
	try:
		(i, j) = [int(k) for k in input().split(': ')];
		rng[i] = j; 
	except EOFError:
		break;

pos, pico, sev = -1, 0, 0;
while pos <= max(rng.keys())+1:
	pos += 1;
	if pos in rng.keys(): 
		if pos == 0 or pos%((rng[pos]-1)*2) == 0:
			sev += pico*rng[pos];
	pico += 1;
print(sev);


