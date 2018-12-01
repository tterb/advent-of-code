# You need to understand the structure of these towers. You ask each program to yell out their name, their weight, and (if they're holding a disc) the names of the programs immediately above them balancing on that disc. In their panic, they don't do this in an orderly fashion; by the time you're done, you're not sure which program gave which information.
# Before you're ready to help them, you need to make sure your information is correct. What is the name of the bottom program?

names, children = set(), set();
while True:
	try:
		data = [s.replace(',', '') for s in input().split()];
		n = data[0];
		c = data[3:len(data)];
		names.add(n);
		children = set(c)|children;
	except EOFError:
		break;
print(set(names).difference(children).pop());
