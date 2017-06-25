import os

from threading import Thread
from multiprocessing import Process
from time import strftime, localtime, sleep
from random import choice

# 吞吐量tps(s) = 用户数 / 平均响应时间

# 是否测试
isTest = False
# 使用线程 or 进程模拟
isThread = True
# 起始人数
init_person = 10
# 每阶段增加人数
add_person = 5
# 每阶段持续时间(s)
duration = 5
# 最终人数
final_person = 20
# 添加用来请求的url信息 ('method', 'url', {'k':v})
URL = [
    # ('POST', 'http://192.168.200.65:5050/sfwa_server/admin/index.do', {}),
    ('GET', 'https://www.baidu.com/s?ie=utf8&oe=utf8&wd=python', {}),
    ('GET', 'https://www.baidu.com/s?ie=utf8&oe=utf8&wd=java', {}),
]
# http请求头信息
HEADERS = {"user-agent": "Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko"}

# 报告文件为html
html_prefix = '<title>压力测试结果报告</title><meta chararset="utf-8" /><style>*{margin:0;padding:0;font-family:consolas}tr:nth-of-type(odd){background:#e8edff}tr:nth-of-type(even){background:white}tr:nth-of-type(odd):hover{background:white}tr:nth-of-type(even):hover{background:#e8edff}table{width:100%;text-align:center;color:#669}tr:first-of-type,tr:last-of-type{font-size:18px;font-weight:900}tr:hover{cursor:pointer}td{padding:8px}</style><table><tr><td>任务</td><td>请求总数</td><td>失败总数</td><td>成功率</td><td>平均耗时(s)</td></tr>'
html_template = '<tr><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>'

# 获得当前的时间表示
now_time_str = lambda: strftime('%Y-%m-%d~%H-%M-%S', localtime())
# 输出日志信息
log = lambda *args: print(now_time_str(), *args)

# [(任务名称, 请求次数, 失败次数, 成功率, 平均耗时, 吞吐量)] 每个任务完成时, 将结果写入这个集合中, 如果是测试请求的话, 此集合长度为1
Run_Results = [None for _ in range(1 if isTest else final_person)]

try:
    #尝试引入 requests 如果没有的话, 就使用pip安装
    import requests as req
except:
    os.system('pip3 install requests')
    import requests as req

class Url():
    def __init__(self, url):
        self.isGET = True if url[0] == 'GET' else False
        self.url = url[1]
        self.data = url[2]

    def http(self, sess):
        if self.isGET:
            return sess.get(self.url, params=self.data, headers=HEADERS, timeout=5)
        else:
            return sess.post(self.url, data=self.data, headers=HEADERS, timeout=5)


# 整理一下URL
URL = [Url(u) for u in URL if u[0] in ('GET', 'POST')]


# 判断所有任务是否都已完成, 并生成结果报告, 或者退出
def generateOrExit():
    # 如果 Run_Results 没有被填满时, 则返回, 否则生成结果报告
    if len([r for r in Run_Results if r]) != len(Run_Results):
        return
    # 计算所有任务的总的数据, 非测试的情况下
    if not isTest:
        allData = ['总结', 0, 0, 0, 0]
        for r in Run_Results:
            allData[1] += r[1]  # 请求次数 相加
            allData[2] += r[2]  # 失败次数 相加
            allData[4] += r[4]  # 平均耗时 相加
        allData[3] = 1 - (allData[2] / allData[1])  # 计算出总的成功率
        allData[4] = allData[4] / final_person  # 总的平均耗时/任务数
        # 把计算的总数据放在最后
        Run_Results.append(tuple(allData))
    filename = 'report%s.html' % (now_time_str(),)
    f = open(filename, 'w')
    f.write(html_prefix)
    for r in Run_Results: f.write(html_template % r)
    f.close()
    os.system(filename)


# 结束任务的标志
isRun = True


def run(num):
    # 每一个任务中只存在一个 Http Session
    sess = req.session()
    tag = '第%s个任务' % (num,)
    print(tag, '开启')
    # 每一个任务的总耗时
    count_time = 0.0
    # 每一个任务的请求总次数
    count_frequency = 0
    # 失败总数
    count_failure = 0
    while isRun:
        try:
            # 随机从URL当中, 拿到一个值去请求
            re = choice(URL).http(sess)
            log(tag, re)
            # 只统计http请求的时间
            count_time += re.elapsed.microseconds / 10e4
            # 响应码超过400 则算失败
            if re.status_code >= 400: count_failure += 1
        # 请求失败的话直接结束任务
        except Exception as ex:
            log(tag, ex)
            break
        # http请求次数+1
        count_frequency += 1

    # 任务结束后, 汇报结果
    Run_Results[num] = (
        tag, count_frequency, count_failure, 1 - (count_failure / count_frequency), count_time / count_frequency)
    # 判断所有任务是否都已完成, 并生成结果报告, 或者退出
    generateOrExit()


def main():
    # 线程任务 or 进程任务
    Task = Thread if isThread else Process
    # 初始化 final_person 个任务, 作为任务池
    tasks = [Task(target=run, args=(i,)) for i in range(final_person)]

    # 初始化运行 init_person 个任务
    [tasks[i].start() for i in range(init_person)]

    # 初始化计数器, 用来判断任务池当中运行了多少个任务
    counter = init_person

    # 还需要 frequency 个阶段
    frequency = int((final_person - init_person) / add_person)
    for _ in range(frequency):
        # 每过 duration 秒, 就运行 add_person 个任务
        sleep(duration)
        [tasks[j].start() for j in range(counter, counter + add_person) if j < len(tasks)]
        counter += add_person

    # 最后一个阶段 一共运行 final_person 个任务 duration 秒
    sleep(duration)


def test():
    # isThread 对测试方法也有效
    (Thread if isThread else Process)(target=run, args=(0,)).start()
    # 等待 duration 秒
    sleep(duration)


# main
if __name__ == "__main__":
    (test if isTest else main)()
    isRun = False
