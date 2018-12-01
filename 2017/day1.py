# The captcha requires you to review a sequence of digits (your puzzle input) and find the sum of all digits that match the next digit in the list. The list is circular, so the digit after the last digit is the first digit in the list.

seq = input();
length = len(seq);
result = 0;
temp = -1;
for i in range(length+1):
	if(temp >= 0 and temp == int(seq[i%length])):
		result += temp;
	temp = int(seq[i%length]);
print(result);