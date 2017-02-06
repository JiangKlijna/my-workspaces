
# 条形图
# 链接：https://zhuanlan.zhihu.com/p/25087813

import matplotlib.pyplot as plt
import numpy as np
y = [20,10,30,25,15]
index = np.arange(5)
plt.bar(left=index, height=y, color='green', width=0.5)
plt.show()
