#https://projecteuler.net/problem=4 

import Eulerlib

#This function takes in a number and turns it into a string and compares the characters
#to see if it si a palindrome(number can be read the same forward and backward).
def palindromeCheck(n):
  n = str(n)
  for i in range(0, len(n)/2):
    if n[i] != n[len(n)-i-1]:
      return False

  return True

#Start timer to time algorithm
start = Eulerlib.getTime()

#We can save CPU time by starting the second loop from i since we don't need to recheck
#Those same multiples again.
for i in range(100, 1000):
  for j in range(i, 1000):
    if palindromeCheck(i*j):
      if i*j > Eulerlib.answer:
        Eulerlib.answer = i*j

#Stop timer to time algorithm
end = Eulerlib.getTime()

#The answer should be 906609
print ("The answer was found in %f seconds" % (end - start))
print ("And the answer is: %d" % Eulerlib.answer)
