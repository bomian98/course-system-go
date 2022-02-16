import pandas as pd
import numpy as np
import time
import random

import random
l = []
for s in range(1,2001): # 2000个学生
    for _ in range(10):  # 每个学生访问10次
        c = random.randint(1,30) # 随机选择课程
        l.append([str(s),str(c)])

# 判断是否存在同一个人选择同一课程两次
#for s in range(1, 2001):
#    tl = [item for item in l if item[0] == s]
#    print(tl)
#    break
#    l1 = len(tl)
#    l2 = len(set([item[1] for item in tl]))
#    if l1 != l2:
#        print(s, tl)

pd.DataFrame(l, columns=['stu','cos']).sample(frac=1).to_csv("stu_cou.csv", index =False)

df = pd.read_csv("stu_cou.csv")
df = df.iloc[:4000]
# 不重复选课
np.array(df.groupby('cos').count()['stu'])
# 去重复选课
df2 = df.drop_duplicates(['stu','cos'])
np.array(df2.groupby('cos').count()['stu'])