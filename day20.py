from collections import defaultdict

class Particle(object):
	def __init__(self, p, v, a):
		self.pos = p;
		self.vel = v;
		self.acc = a;

	def next(self):
		for i in range(len(self.pos)):
			self.vel[i] += self.acc[i];
			self.pos[i] += self.vel[i];

	def dist(self):
		return sum([abs(x) for x in self.pos]);

def main():
	particles, i = {}, 0;
	while True:
		try:
			data = input().split(', ');
			p = [int(x) for x in data[0].split("=")[1][1:-1].split(",")];
			v = [int(x) for x in data[1].split("=")[1][1:-1].split(",")];
			a = [int(x) for x in data[2].split("=")[1][1:-1].split(",")];
			particles[i] = Particle(p,v,a); 
			i += 1;
		except EOFError:
			break;

	# part2 = True;
	while True:
		minDist = None;
		part = None;
		for i, p in particles.items():
			p.next();
			if minDist is None or p.dist() < minDist:
				part = i;
				minDist = p.dist();

		# if part2:
		#   dd = defaultdict(list);
		#   for i, part in particles.items():
		#     k = tuple(part.pos);
		#     dd[k].append(i);

		#   for k, v in dd.items():
		#     if len(v) > 1:
		#       for i in v:
		#         del particles[i];

		#   # print(len(particles));
		#   # break;
		# else:c
		print(part);

if __name__ == "__main__":
		main()
