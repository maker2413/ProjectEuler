#Import time library to time algorithms
import time

answer = 0

#This function returns the current time and is used to time algorithms
def getTime():
	return time.time()

#This function takes in a number and will return True if its prime and False if it is a consonants
def primeCheck(n):
	if n < 2:
		return False
	elif n <= 3:
		return True
	elif n % 2 == 0:
		return False
	for i in range(3, int(n**.5) + 1, 2):
		if n % i == 0:
			return False

	return True

#This funtion takes in a number and will generate a list of all primes less than that number
def primeList(n):
	primes = []

	for i in range(2, n+1):
		if primeCheck(i):
			primes.append(i)

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