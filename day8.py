
def parse(data):
    c, m = data.pop(0), data.pop(0)
    vals = list()
    total = 0
    for i in range(c):
        temp, meta, data = parse(data)
        total += temp
        vals.append(meta)
    total += sum(data[:m])
    if c == 0:
        meta = sum(data[:m])
    else:
        meta = sum(vals[k - 1] for k in data[:m] if k > 0 and k <= len(vals))
    return (total, meta, data[m:])

def main():
    data = [int(i) for i in input().split(' ')]
    total, value, remaining = parse(data)
    print(total)


if __name__ == '__main__':
    main()