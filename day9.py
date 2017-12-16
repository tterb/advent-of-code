# You sit for a while and record part of the stream (your puzzle input). The characters represent groups - sequences that begin with { and end with }. Within a group, there are zero or more other things, separated by commas: either another group or garbage. Since groups can contain other groups, a } only closes the most-recently-opened unclosed group - that is, they are nestable. Your puzzle input represents a single, large group which itself contains many smaller ones.

# Sometimes, instead of a group, you will find garbage. Garbage begins with < and ends with >. Between those angle brackets, almost any character can appear, including { and }. Within garbage, < has no special meaning.
# In a futile attempt to clean up the garbage, some program has canceled some of the characters within it using !: inside garbage, any character that comes after ! should be ignored, including <, >, and even another !.

# Your goal is to find the total score for all groups in your input. Each group is assigned a score which is one more than the score of the group that immediately contains it. (The outermost group gets a score of 1.)

import re

stream = input();
depth, score, garbageCount, garbage = 0, 0, 0, False;

while(stream.count('!') > 0):
	index = stream.index('!');
	stream = stream[:index] + stream[(index+2):];

for i in stream: 
	if garbage and i != ">":
		garbageCount += 1;
	elif i == "<":
		garbage = True;
	elif i == ">":
		garbage = False;
	elif i == "{": 
		depth += 1;
	elif i == "}":
		score += depth;
		depth -= 1;
print(str(score));
