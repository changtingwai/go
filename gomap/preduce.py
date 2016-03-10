#!/usr/bin/env python
import sys
from operator import itemgetter
#map key,list
word2count = {}
for line in sys.stdin:
    res = []
    line = line.strip()
    precmd,count_all,count_exi,Time_all = line.split("\t")
    # precmd,num = cmd.split("_")
    try:
        count_all = int(count_all)
        count_exi = int(count_exi)
        Time_all = int(Time_all)

        #get list
        res = word2count.get(precmd,res)
        if not res:
            res.append(count_all)
            res.append(count_exi)
            res.append(Time_all)
            word2count[precmd] = res
        else:
            count_all += res[0]
            count_exi += res[1]
            Time_all += res[2]
            res = []
            res.append(count_all)
            res.append(count_exi)
            res.append(Time_all)
            word2count[precmd] = res
    except ValueError:
        pass
for precmd in word2count:
    res = word2count[precmd]
    print '%s\t%s\t%s\t%s' %(precmd,res[0],res[1],res[2])
