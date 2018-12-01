import numpy as np;
RNG = 256;

def partOne(seq, num_lst, pos, skip):
	lst_len = len(num_lst)
	curr_pos = pos
	skip_size = skip
	lengths = seq.split(',')

	for leng in lengths:
		leng = int(leng.strip())
		if leng > lst_len:
			continue
		rev_end = (curr_pos + leng) % lst_len
		if rev_end == 0:
			rev_end = lst_len
		inds = list(range(curr_pos, rev_end))
		if leng > 0 and rev_end <= curr_pos:
			inds = list(range(curr_pos, lst_len)) + list(range(rev_end)) 

		num_lst[inds] = np.flipud(num_lst[inds])                  
		curr_pos = (curr_pos + leng + skip_size) % lst_len
		skip_size += 1

	return num_lst[0] * num_lst[1], num_lst, curr_pos, skip_size


def partTwo(seq):
	sp = np.array(range(RNG))
	pos, skip, blockSize = 0, 0, 16;
	dense = []
	byt = ''.join([str(ord(char)) + ',' for char in seq.strip()]);
	byt += "17,31,73,47,23";

	for i in range(64):
		mult, sp, pos, skip = partOne(byt, sp, pos, skip);
	for block in range(0, RNG, blockSize):
		xored = 0;
		for i in range(block, block + blockSize): 
				xored ^= sp[i];
		dense.append(xored);
	return ''.join([('0' + format(num, 'x'))[-2:] for num in dense]);