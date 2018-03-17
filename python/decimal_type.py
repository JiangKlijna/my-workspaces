from sys import exit
from math import gcd, sqrt
from collections import Counter

# 只包含2和5的集合
two_five = set([2, 5])

# 获得分数
def get_fraction():
	try:
		numerator, denominator = input('Enter two strictly positive integers: ').split(',')
		numerator, denominator = int(numerator), int(denominator)
		if numerator <= 0 or denominator <= 0: raise ValueError
		return numerator, denominator
	except ValueError:
		print('Incorrect input, giving up.')
		exit()

# 获得最大公约数，并简化分数
def simplify(a, b):
	n = gcd(a, b)
	return a // n, b // n

# 质因数分解
def get_primes(n):
	result = []
	for i in range(2, n+1):
		s = 0
		while n % i == 0:
			n = n / i
			s += 1
		if s > 0:
			for k in range(s):
				result.append(i)
			if n == 1:
				return result

# 获得循环部分和循环前的部分
# a分子 b分母 n循环节从第几位开始
def get_loop_part(a, b, n):
	arr = [] #存放循环节中的数字
	rem = [] #存放余数

	if a >= b:
		a = a % b
	j = 0 #j是下标要从0开始
	while True:
		if a < b:
			arr.append(a*10//b)
			rem_num = (a*10)%b
			rem.append(rem_num)
			if len(rem) > n+1:
				if rem[n] == rem_num:
					return arr[0:n], arr[n:j]
			a = (a*10)%b
			j += 1
			
def main():
	result = {}
	# 获得分数
	numerator, denominator = get_fraction()
	# 简化分数
	numerator, denominator = simplify(numerator, denominator)
	# 直接计算得到整数部分
	integral_part = numerator // denominator
	result['整数部分'] = integral_part
	# 获得分母的质因数集合
	primes = get_primes(denominator)
	# 统计每个质因数出现的次数
	ele_count = Counter(primes)
	# 2或5出现次数的最大值
	max_two_five_count = max(ele_count.get(2, 0), ele_count.get(5, 0))
	# 质因数去重
	primes_set = set(primes)

	# [2],[5],[2,5]
	if primes_set | two_five == two_five:
		result['类型'] = '有限小数'
		result['小数位数'] = max_two_five_count
	# 2和5都不存在
	elif not (2 in primes_set or 5 in primes_set):
		result['类型'] = '纯循环小数'
		arr, loop_part = get_loop_part(numerator, denominator, 0)
		result['循环部分'] = loop_part
	# 既存在2，5也存在其他
	else:
		result['类型'] = '混循环小数'
		result['开始循环的位数'] = max_two_five_count + 1
		arr, loop_part = get_loop_part(numerator, denominator, max_two_five_count)
		result['循环前部分'] = arr
		result['循环部分'] = loop_part
	
	return result



result = main()
print(result)

