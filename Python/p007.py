#https://projecteuler.net/problem=7

import Eulerlib

#Start timer to time algorithm
start = Eulerlib.getTime()

#We can just generate the first 10001 and primes and then read the last prime in the list
primes = Eulerlib.primeGen(10001)

Eulerlib.answer = primes[len(primes)-1]

#Stop timer to time algorithm
end = Eulerlib.getTime()

#The answer should be 104743
print ("The answer was found in %f seconds" % (end - start))
print ("And the answer is: %d" % Eulerlib.answer)
