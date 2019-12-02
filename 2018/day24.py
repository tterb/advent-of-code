from sys import stdin
import re

reg = re.compile(r"(\d+) units each with (\d+) hit points (\([^)]*\) )?with an attack that does (\d+) (\w+) damage at initiative (\d+)")

inp = [i for i in stdin.readlines()]
s = inp.index('\n')
immunes, infections = inp[1:s], inp[s+2:]
teams = list()
for lines in [immunes, infections]:
  team = list()
  for i in lines:
    i = i.strip()
    units, hp, attr, dmg, type, init =  reg.findall(i)
    immune, weak = list(), list()
    if attr:
      attr = attr.rstrip(' )').lstrip('(')
      for a in attr.split("; "):
        if a.startswith("weak to "):
          weak = a[len("weak to "):].split(", ")
        elif s.startswith("immune to "):
          immune = a[len("immune to "):].split(", ")
        else:
          assert False
    unit = [int(units), int(hp), int(dmg), type, int(init), set(immune), set(weak)]
    team.append(unit)
  teams.append(team)