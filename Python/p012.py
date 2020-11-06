#https://projecteuler.net/problem=12

import Eulerlib

#This function takes in a number and will return a total of how many divisors
#the input number has
def findDivisors(n):
  total = 0
  for i in range(1, int(n**.5)+1):
    if n % i == 0:
      total += 2

  return total

#Start timer to time algorithm
start = Eulerlib.getTime()

#We can generate triangle numbers by just incrementing y and adding it to the
#previous triangle number used
x, y = 1, 2

#Keep generating triangle numbers until one is found that has more than 500 divisors
while findDivisors(x) <= 500:
  x += y
  y += 1

Eulerlib.answer = x

#Stop timer to time algorithm
end = Eulerlib.getTime()

#The answer should be 76576500
print ("The answer was found in %f seconds" % (end - start))
print ("And the answer is: %d" % Eulerlib.answer)
