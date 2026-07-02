import time


limit = 4_000_000


def solve(to):
    answer = 0
    x, y = 1, 1

    while y < to:
        y += x
        x = y - x

        if y % 2 == 0:
            answer += y

    return answer


start = time.time()
answer = solve(limit)
end = time.time()
print(answer)
print("--------------------------------------------------")
print(f"  {end - start:.6f} seconds")
