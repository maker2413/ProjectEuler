import time


limit = 1000


def palindromeCheck(n):
    pal = str(n)

    return pal == pal[::-1]


def solve(to):
    answer = 0

    for i in range(to - 1, 99, -1):
        if i * i <= answer:
            break

        for j in range(i - 1, 99, -1):
            product = i * j

            if product <= answer:
                break

            if palindromeCheck(i * j):
                answer = i * j

    return answer


# The answer should be 906609
start = time.time()
answer = solve(limit)
end = time.time()
print(answer)
print("--------------------------------------------------")
print(f"  {end - start:.6f} seconds")
