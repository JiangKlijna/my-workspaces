
# 箱形图
# 链接：https://zhuanlan.zhihu.com/p/25087813

import matplotlib.pyplot as plt
import numpy as np
np.random.seed(100)
data = np.random.normal(size=(1000,4),loc=0,scale=1)
labels = ['A','B','C','D']
plt.boxplot(data,labels=labels)
plt.show()
