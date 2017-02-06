
# 饼状图
# 链接：https://zhuanlan.zhihu.com/p/25087813

import matplotlib.pyplot as plt
import numpy as np
labels = 'A','B','C','D'
fracs = [15,30,45,10]
plt.axes(aspect=1)  #使x y轴比例相同
explode = [0,0.05,0,0]  # 突出某一部分区域
plt.pie(x=fracs, labels=labels, autopct='%.0f%%', explode=explode)  #autopct显示百分比
plt.show()
