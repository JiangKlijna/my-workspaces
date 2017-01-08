
# author jiangKlijna
import os, threading
from tkinter import *
from tkinter import ttk
import tkinter.messagebox as messagebox
from tkinter.filedialog import askdirectory

class Download(threading.Thread):
	def __init__(self, dir_, url, lb):
		super(Download, self).__init__()
		self.lb = lb
		self.dir = dir_
		self.url = url

	def run(self):
		shell = 'you-get -o %s %s' % (self.dir, self.url)
		print(shell)
		if os.system(shell) == 0:
			self.lb.insert(0, '下载完成...' + self.url[:28])
		else:
			self.lb.insert(0, '下载失败...' + self.url[:28])

class App(Frame):
	def __init__(self, master=None):
		Frame.__init__(self, master)
		self.init_setting()
		self.init_stringvar()
		self.init_widgets()

	def init_setting(self):
		self.pack()
		self.master.title('you-get download')
		self.master.minsize(600, 300)
		self.master.maxsize(600, 300)
		self.tasklist = []

	def init_stringvar(self):
		self.dir_sv = StringVar()
		self.dir_sv.set('F:/')
		self.url_sv = StringVar()

	def init_widgets(self):
		table = ttk.Frame(self, padding="3 3 12 12")
		table.grid(column=0, row=0, sticky=(N, W, E, S))
		table.columnconfigure(0, weight=1)
		table.rowconfigure(0, weight=1)

		ttk.Button(table, textvariable=self.dir_sv, command=self.onSelectDir).grid(column=1, row=1, sticky=W)
		ttk.Label(table, text="请在下方(ctrl+v)输入视频的链接地址").grid(column=1, row=2, sticky=W)
		ttk.Button(table, text='点我下载', command=self.onDownload, width=80).grid(column=1, row=4, sticky=W)
		ttk.Entry(table, textvariable=self.url_sv, width=80).grid(column=1, row=3, sticky=W)
		self.lb = Listbox(table, width=80)
		self.lb.grid(column=1, row=5, sticky=W)

	def updateListBox(self):
		self.delete()

	def onDownload(self):
		url = self.url_sv.get().strip()
		if url == '':
			messagebox.showinfo('Error', 'url can not be space')
		else :
			self.url_sv.set('')
			self.lb.insert(0, "正在下载..." + url[:28])
			Download(self.dir_sv.get(), url, self.lb).start()

	def onSelectDir(self):
		path = askdirectory()
		self.dir_sv.set(path)


def main():
	app = App()
	app.mainloop()

if __name__ == '__main__':
	main()
