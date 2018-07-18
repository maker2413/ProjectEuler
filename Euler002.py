#https://projecteuler.net/problem=2

import Eulerlib

#We have to start the Fibonacci sequence with two 1's
x, y = 1, 1

#Start timer to time algorithm
start = Eulerlib.getTime()

#This will just add up the even numbers of the Fibonacci sequence less than 4000000
while y < 4000000:
	y += x
	x = y - x
	if y % 2 == 0:
		Eulerlib.answer += y

#Stop timer to time algorithm
end = Eulerlib.getTime()

#The answer should be 4613732
print ("The answer was found in %f seconds" % (end - start))
print ("And the answer is: %d" % Eulerlib.answer)