# This spinlock's algorithm is simple but efficient, it starts with a circular buffer containing only the value 0, which it marks as the current position. It then steps forward through the circular buffer some number of steps (the puzzle input) before inserting the first new value, 1, after the value it stopped on. The inserted value becomes the current position. 

buf = [0];
pos, interval, count = 0, 316, 2018;
for i in range(1, count):
	pos = (pos + interval) % len(buf);
	buf.insert(pos+1, i);
	pos += 1;
print(buf[buf.index(2017)+1])
