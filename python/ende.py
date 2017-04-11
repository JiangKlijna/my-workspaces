import time, datetime

format = '%S%M%H%d%m%Y'

def en(data):
	#data type is int
	#return type is str
	return time.strftime(format, time.localtime(data))

def de(data):
	#data type is str
	#return type is int
	dt = datetime.datetime.strptime(data, format)
	return int(time.mktime(dt.timetuple()))

print(en(331024))
print(de(en(331024)))
