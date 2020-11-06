#https://projecteuler.net/problem=5

import Eulerlib

x, y = 1, 1

#Start timer to time algorithm
start = Eulerlib.getTime()

#We can save a huge amount of CPU time by finding the least common multiple of 20 and 19
#and then finding multples of that number until we find a number divisible by 18 and
#multiple that until we find a multiple divisible by 17 etc.
for i in range(20 , 0, -1):
  while y % i != 0:
    y += x
  x = y

Eulerlib.answer = y

#Stop timer to time algorithm
end = Eulerlib.getTime()

#The answer should be 232792560
print ("The answer was found in %f seconds" % (end - start))
print ("And the answer is: %d" % Eulerlib.answer)
