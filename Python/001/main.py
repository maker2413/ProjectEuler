import time


limit = 1_000


def solve(to):
    answer = 0

    for x in range(0, to):
        if x % 3 == 0 or x % 5 == 0:
            answer += x

    return answer


start = time.time()
answer = solve(limit)
end = time.time()
print(answer)
print("--------------------------------------------------")
print(f"  {end - start:.6f} seconds")
