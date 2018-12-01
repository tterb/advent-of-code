# The programs' dance consists of a sequence of dance moves:
#  - Spin, written sX, makes X programs move from the end to the front, but maintain their order otherwise. (For example, s3 on abcde produces cdeab).
#  - Exchange, written xA/B, makes the programs at positions A and B swap places.
#  - Partner, written pA/B, makes the programs named A and B swap places.

def exchange(x, y):
	temp = seq[y];
	seq[y] = seq[x];
	seq[x] = temp;

seq = [chr(i+97) for i in range(16)];
moves = input().split(','); 
for m in moves:
	if '/' not in m:
		x = int(m[1:]);
		seq = seq[-x:]+seq[:-x];
	else:
		i, j = m[1:].split('/');
		if m[0] == 'x':
			exchange(int(i), int(j));
		elif m[0] == 'p':
			exchange(seq.index(i), seq.index(j));
print(''.join(seq));