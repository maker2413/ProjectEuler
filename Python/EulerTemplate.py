#https://projecteuler.net/problem=

import Eulerlib

#Start timer to time algorithm
start = Eulerlib.getTime()



#Stop timer to time algorithm
end = Eulerlib.getTime()

#The answer should be
print ("The answer was found in %f seconds" % (end - start))
print ("And the answer is: %d" % Eulerlib.answer)
