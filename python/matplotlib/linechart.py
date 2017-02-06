
# 折线图
# 链接：https://zhuanlan.zhihu.com/p/25087813

import matplotlib.pyplot as plt
import numpy as np
x = np.linspace(-10,10,100)
y = x**3
plt.plot(x,y,linestyle='--',color='green',marker='<')
plt.show()
