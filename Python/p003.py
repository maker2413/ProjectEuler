#https://projecteuler.net/problem=3

import Eulerlib

#This is the number we are trying to find the greatest common prime factor of
z = 600851475143

#Start timer to time algorithm
start = Eulerlib.getTime()

#We can save CPU time by just checking if the numbers less than the square root of z
#are factors of z. We can then check if each factor is prime to find the greatest 
#common prime factor
for i in range (3, int(z**.5)):
  if z % i == 0:
    if Eulerlib.primeCheck(i):
      Eulerlib.answer = i

#Stop timer to time algorithm
end = Eulerlib.getTime()

#The answer should be 6857
print ("The answer was found in %f seconds" % (end - start))
print ("And the answer is: %d" % Eulerlib.answer)
