import time


limit = 600851475143

def solve(to):
  answer = 1

  i = 3
  while i <= int(to**.5):
    if to % i == 0:
      answer = i
      while to % i == 0:
        to = int(to / i)

    i += 2

  if to > answer:
    return to

  return answer

start = time.time()
answer = solve(limit)
end = time.time()
print(solve(limit))
print("--------------------------------------------------")
print(f'  {end - start:.6f} seconds')
