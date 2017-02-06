
# 网格图
# 链接：https://zhuanlan.zhihu.com/p/25087813

import matplotlib.pyplot as plt
import numpy as np
y=np.arange(1,5)
plt.plot(y,y*2)
plt.grid(True,color='g',linestyle='--',linewidth='1')
plt.show()
