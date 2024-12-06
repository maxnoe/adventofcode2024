import requests
import os


test_input = """47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
"""

cookie = os.environ["AOC_SESSION"]
r = requests.get(
    "https://adventofcode.com/2024/day/5/input",
    headers=dict(Cookie=f"session={cookie}")
)
real_input = r.text



def build_lookup(rules):
    greater = {}
    for a, b in rules:
        if a not in greater:
            greater[a] = set()
        if b not in greater:
            greater[b] = set()
        greater[a].add(b)
    return greater


def correct_order(job, rules):

    applicable_rules = [
        (a, b) 
        for a, b in rules
        if a in job and b in job
    ]

    greater = build_lookup(applicable_rules)
    correct_order = [k for k, _ in sorted(greater.items(), key=lambda kv: len(kv[1]), reverse=True)]

    mid = len(correct_order) // 2
    if job == correct_order:
        return correct_order[mid], 0
    else:
        return 0, correct_order[mid]


def solve(input_string):
    rules, jobs = input_string.split("\n\n")

    rules = [(int(a), int(b)) for a, b in map(lambda s: s.split("|"), rules.splitlines())]
    jobs = [list(map(int, line.split(","))) for line in jobs.splitlines()]

    part1 = 0
    part2 = 0
    for job in jobs:
        p1, p2 = correct_order(job, rules)
        part1 += p1
        part2 += p2

    print(part1, part2)


if __name__ == "__main__":
    print("Test:", end=" ")
    solve(test_input)
    print("Real:", end=" ")
    solve(real_input)
