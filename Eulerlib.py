#Import time library to time algorithms
import time

answer = 0

#This function returns the current time and is used to time algorithms
def getTime():
	return time.time()

#This function takes in a number and will return True if its prime and False if
#it is a consonants
def primeCheck(n):
	if n < 2:
		return False
	elif n <= 3:
		return True
	elif n % 2 == 0:
		return False

	#Onces we have ruled out numbers less than 3 and all even numbers we can just
	#check the odd numbers less than n
	for i in range(3, int(n**.5) + 1, 2):
		if n % i == 0:
			return False

	return True

#This funtion takes in a number and will generate a list of all primes less than
#that number using a Sieve of Eratosthenes
def primeList(n):
	#We make a list from 0 to n and declare them all to true we can then already
	#mark off 0 and 1 as primes for the sieve
	sieve = [True] * (n + 1)
	sieve[0] = sieve [1] = False

	#We can iterate through the list and when a prime is found(an entry marked true)
	#we can mark off all of its multiples eventually leaving only the prime nubmers
	for i in range(int(n**.5) + 1):
		if sieve[i]:
			for j in range(i * i, len(sieve), i):
				sieve[j] = False

	#We now can make a list that we will populate with the actual prime numbers
	#from our True, False table
	primes = []

	for (j, isprime) in enumerate(sieve):
		if isprime:
			primes.append(j)

	return primes

#This function takes in a number and will generate that many prime numbers
def primeGen(n):
	i = 0
	j = 2
	primes = []

	while i < n:
		if primeCheck(j):
			primes.append(j)
			i += 1
		j += 1

	return primes