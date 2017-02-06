
# 直方图
# 链接：https://zhuanlan.zhihu.com/p/25087813

import matplotlib.pyplot as plt
import numpy as np
mu = 100
sigma = 20
x = mu + sigma * np.random.randn(20000)  # 样本数量
plt.hist(x,bins=100,color='green',normed=True)   # bins显示有几个直方,normed是否对数据进行标准化
plt.show()
