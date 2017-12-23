data = [];
while True:
	try:
		line = input();
		data.append(line);
	except EOFError:
		break;

i, j = 0, data[0].index('|');
steps, letters = 0, [];
direction = 'down';
current = '|';
while current != ' ':
	steps += 1;
	if direction == 'down':
		i += 1;
	elif direction == 'up':
		i -= 1;
	elif direction == 'left':
		j -= 1;
	elif direction == 'right':
		j += 1;
	current = data[i][j];
	if current == '+':
		if direction in ('down', 'up'):
			if data[i][j-1] != ' ':
				direction = 'left';
			else:
				direction = 'right';
		else:
			if data[i-1][j] != ' ':
				direction = 'up';
			else:
				direction = 'down';
	elif current.isalpha():
		letters.append(current);

print(''.join(letters));
print(steps);