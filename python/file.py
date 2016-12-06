import os
import re

class Vi:
	def __init__(self):
		self.size = 0
		self.show_re = re.compile('^(s|show)\\d*$')
		self.del_re = re.compile('^(d|del)\\d*$')
		self.mod_re = re.compile('^(m|mod)\\d*$')
		self.back_re = re.compile('^(b|back)\\d*$')

	def inputFilename(self):
		while True:
			self.fname = input()
			if os.path.exists(self.fname):
				print("ERROR: '%s' already exists" %(self.fname))
			else:
				print("\n")
				break

	def inputLines(self):
		self.lines = []
		while True:
			try:
				line = input(str(self.size) + '> ')
			except EOFError:
				self.inputCommand()
				break
			if line == '.':
			    break
			else:
				self.lines.append(line)
				self.size += 1

	def inputCommand(self):
		try:
			cmd = input('\n-> ')
			if self.show_re.match(cmd):
				self.show()
			elif self.del_re.match(cmd):
				self.delete()
			elif self.mod_re.match(cmd):
				self.modify()
			elif self.back_re.match(cmd):
				self.back()
			else:
				self.exit()
		except EOFError:
			os._exit()

	def show(self, size=-1):
		pass

	def delete(self, size=-1):
		pass

	def modify(self, size=-1):
		pass

	def back(self, size=-1):
		pass

	def exit(self):
		self.save()
		os._exit()

	def save(self):
		file = open(self.fname, 'w')
		for line in self.lines:
			file.writelines(line + "\n")
		file.close()

	def main(self):
		self.inputFilename()
		self.inputLines()

if __name__ == '__main__':
	Vi().main()
