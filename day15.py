# The generators both work on the same principle. To create its next value, a generator will take the previous value it produced, multiply it by a factor (generator A uses 16807; generator B uses 48271), and then keep the remainder of dividing that resulting product by 2147483647. That final remainder is the value it produces next.

valA, valB = 16807, 48271;
mod = 2147483647;
num = 40000000;

def generator(factor, val, M=1):
	while True:
		factor = factor * val % mod;
		if factor % M == 0:
			yield factor & 0xffff;

A, B = input().split(' ').pop(), input().split(' ').pop();
A, B = int(A), int(B);
genA, genB = generator(A, valA), generator(B, valB);
print(sum(next(genA) == next(genB) for _ in range(40000000)))
genA, genB = generator(A, valA, 4), generator(B, valB, 8);
print(sum(next(genA) == next(genB) for _ in range(5000000)));