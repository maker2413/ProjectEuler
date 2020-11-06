#https://projecteuler.net/problem=9

import Eulerlib

#Start timer to time algorithm
start = Eulerlib.getTime()

#Since a < b < c the biggest A can be for this sum of these number to be 1000 is 332
#and the biggest b can be is 498(a could be 1, b could be 499, and c could be 500). 
#We can also save CPU time by starting b's range from one number bigger than a.
for a in range(1, 332):
  for b in range(a+1, 499):
    c = (a**2 + b**2)**.5
    if a + b + c == 1000:
      Eulerlib.answer = a*b*c

#Stop timer to time algorithm
end = Eulerlib.getTime()

#The answer should be 31875000
print ("The answer was found in %f seconds" % (end - start))
print ("And the answer is: %d" % Eulerlib.answer)
