def printSet(set):
	for i in range(len(set)):
		print(type(set),"[",i,"] = ",set[i])


def Bubble(list):
	print("\nBubbleSort:\n")
	if type(list)!=type([]):
		return
	
	for i in range(len(list)):
		for j in range(1,len(list)-i):
			if list[j]<list[j-1]:
				temp = list[j]
				list[j] = list[j-1]
				list[j-1] = temp
	
	printSet(list)

def Select(list):
	print("\nSelectSort:\n")
	if type(list)!=type([]):
		return
	
	for i in range(1,len(list)):
		min = list[i-1];
		tag = i-1;
		for j in range(i,len(list)):
			if min>list[j]:
				min = list[j]
				tag = j
		temp = list[i-1]
		list[i-1] = list[tag]
		list[tag] = temp
	printSet(list)
	
def Insert(list):
	print("\nInsertSort:\n")
	if type(list)!=type([]):
		return
	
	for i in range(1,len(list)):
		now = list[i]
		j = i
		while j>0 and now<list[j-1]:
			list[j] = list[j-1]
			j = j-1
		list[j] = now
	printSet(list)
	
Bubble([5,3,1,9,2,6])
Select([5,3,1,9,2,6])
Insert([5,3,1,9,2,6])


