#https://projecteuler.net/problem=10

import Eulerlib

#Start timer to time algorithm
start = Eulerlib.getTime()

#We can just call our prime list function to list all primes less
#than 2000000 and then sum the list
primes = Eulerlib.primeList(1999999)

Eulerlib.answer = sum(primes)

#Stop timer to time algorithm
end = Eulerlib.getTime()

#The answer should be 142913828922
print ("The answer was found in %f seconds" % (end - start))
print ("And the answer is: %d" % Eulerlib.answer)