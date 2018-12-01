# The spreadsheet consists of rows of apparently-random numbers. To make sure the recovery process is on the right track, they need you to calculate the spreadsheet's checksum. For each row, determine the difference between the largest value and the smallest value; the checksum is the sum of all of these differences.

result = 0;
while True:
	try:
		line = [int(x) for x in input().split()];
		result += max(line) - min(line);
	except EOFError:
		break;
print(result);