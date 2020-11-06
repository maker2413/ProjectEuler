#https://projecteuler.net/problem=6

import Eulerlib

#Start timer to time algorithm
start = Eulerlib.getTime()

#First we can just sum all number less than 100 and then square it wince this will
#a bigger number than the sum of all squares of the numbers less than 100
for i in range(1, 101):
  Eulerlib.answer += i

Eulerlib.answer = Eulerlib.answer ** 2

#We can then just subtract the squares of all numbers less than 100 form the 
#square of the sum of all numbers less than 100
for j in range(1, 101):
  Eulerlib.answer -= (j ** 2)

#Stop timer to time algorithm
end = Eulerlib.getTime()

#The answer should be 25164150
print ("The answer was found in %f seconds" % (end - start))
print ("And the answer is: %d" % Eulerlib.answer)
