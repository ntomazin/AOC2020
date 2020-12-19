r = open('day19_in').read()
input = r.split('\n\n')

#Part 1
ruletext = input[0].splitlines()
rules = {}
for text in ruletext:
    data = text.split(': ')
    id = data[0]
    content = [seq.split(' ') for seq in data[1].split(' | ')]
    rules[id] = content

received = input[1].splitlines()

def check(rules, id, sample, start):
    rule = rules[id]
    if rule[0][0][0] == '"':
        return {start + 1} if start < len(sample) and rule[0][0][1] == sample[start] else set()
    else:
        endings = set()
        for subrule in rule:
            buffer = {start}
            for part in subrule:
                temp = set()
                for loc in buffer:
                    temp = temp | check(rules, part, sample, loc)
                buffer = temp
            endings = endings | buffer
        return endings

results = [len(sample) in check(rules, '0', sample, 0) for sample in received]
print(results.count(True))

#Part 2
rules['8'] = [['42'],['42','8']]
rules['11'] = [['42','31'],['42','11','31']]
results = [len(sample) in check(rules, '0', sample, 0) for sample in received]
print(results.count(True))