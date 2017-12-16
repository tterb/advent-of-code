# There are sixteen memory banks; each memory bank can hold any number of blocks. The goal of the reallocation routine is to balance the blocks between the memory banks.
# The reallocation routine operates in cycles. In each cycle, it finds the memory bank with the most blocks (ties won by the lowest-numbered memory bank) and redistributes those blocks among the banks. To do this, it removes all of the blocks from the selected bank, then moves to the next (by index) memory bank and inserts one of the blocks. It continues doing this until it runs out of blocks; if it reaches the last memory bank, it wraps around to the first one.

# The debugger would like to know how many redistributions can be done before a blocks-in-banks configuration is produced that has been seen before.

while True:
	try:
		banks = [int(s) for s in input().split()]; 
	except EOFError:
		break;
s = set();
cycles = 0;
while(len(s) == cycles):
	cycles += 1;
	m = max(banks);
	i = banks.index(m);
	banks[i] = 0;
	while(m > 0):
		i += 1;
		banks[i % len(banks)] += 1;
		m -= 1;
	s.add(str(banks));
print(cycles);
