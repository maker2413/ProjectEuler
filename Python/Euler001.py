#https://projecteuler.net/problem=1

import Eulerlib

#Start timer to time algorithm
start = Eulerlib.getTime()

#We can just check all the numbers less than 1000 if they are divisible by 3 or 5
#CPU's are quick enough that the answer is found instantly
for x in range(0, 1000):
	if x % 3 == 0 or x % 5 == 0:
		Eulerlib.answer += x

#Stop timer to time algorithm
end = Eulerlib.getTime()

#The answer should be 233168
print ("The answer was found in %f seconds" % (end - start))
print ("And the answer is: %d" % Eulerlib.answer)