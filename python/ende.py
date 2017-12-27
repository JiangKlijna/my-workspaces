from time import strftime, localtime, mktime
from datetime import datetime

format = '%S%M%H%d%m%Y'

#data type is int
#return type is str
en = lambda data: strftime(format, localtime(data))

#data type is str
#return type is int
de = lambda data: int(mktime(datetime.strptime(data, format).timetuple()))

if __name__ == '__main__':
	print(en(331024))
	print(de(en(331024)))
