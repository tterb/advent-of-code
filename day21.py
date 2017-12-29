# Part 1: 208

import numpy as np;
import itertools;

def printer(arr):
	for i in arr:
		print(''.join(i));
	print('')


def enhance(p, e, arr):
	# [print(len(i.split('/')),end=', ') for i in p.keys()]
	pat = [i for i in p.keys() if len(arr) == len(i.split('/')) ] 
	for i in pat:
		if arr in p[i]:
			arr = e[i];
	return arr;

def joinArr(b,index):
	n = np.concatenate([np.concatenate([b[j+i] for j in range(index)], 1).tolist() for i in range(0,index*2,2)]);
	return n;
	# n2 = np.concatenate([b[j+index] for j in range(index)], 1);
	# return np.concatenate((n, n2), 0).tolist();

def main():
	img = [list(i) for i in '.#./..#/###'.split('/')];
	pattern, rule, count = {}, {}, 0;
	while True:
		try:
			data = input().split(' => ');
			v = [];
			p = [list(i) for i in data[0].split('/')];
			rule[data[0]] = [list(i) for i in data[1].split('/')];
			for _ in range(4):
				v.append(p);
				v.append([list(reversed(i)) for i in p]);
				p = list(zip(*reversed(p)));
			pattern[data[0]] = v;
		except EOFError:
			break;

	while count < 5:
		size = len(img);
		seg = [];
		if size % 2 == 0 and size > 2:
			index = int(size/2);
			temp = [tuple(2*k+2*d for k in range(2)) for d in range(index)];
			for i, j in itertools.product(temp, temp):
				seg.append(np.array(img)[j[0]:j[1], i[0]:i[1]].tolist());
			blocks = [enhance(pattern, rule, s) for s in seg];
			img = np.concatenate([np.concatenate([blocks[j+i] for j in range(index)], 1).tolist() for i in range(0,size,2)]);

		elif size % 3 == 0 and size > 3:
			index = int(size/3);
			temp = [tuple(3*k+3*d for k in range(2)) for d in range(index)];
			print('temp: '+str(temp))
			for i, j in itertools.product(temp, temp):
				seg.append(np.array(img)[j[0]:j[1], i[0]:i[1]].tolist());
				# print(np.array(img)[j[0]:j[1], i[0]:i[1]].tolist())
			# print('seg: '+str(len(seg)))
			blocks = [enhance(pattern, rule, s) for s in seg];
			# print('blocks: '+str(len(blocks)))
			# [printer(b) for b in blocks]
			img = np.concatenate([np.concatenate([blocks[j] for j in range(index)], 1).tolist() for i in range(0,size,2)]);
			
		else:
				img = enhance(pattern, rule, img);
		count += 1;
		print('#'+str(count)+': ('+str(len(img))+')');
		printer(img);
	result = 0;
	# result += [str(i).count('#') for i in img];
	for i in img:
		result += str(i).count('#');
	print(result)


if __name__ == "__main__":
    main()
