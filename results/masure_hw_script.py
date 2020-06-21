import pandas as pd
import os
import time
import statistics
from tcp_latency import measure_latency
import pandas as pd
import datetime
import psutil
import sys
from _thread import *
import contextlib
import requests

global time_state

@contextlib.contextmanager
def atomic_overwrite(filename):
    temp = filename + '~'
    with open(temp, "w") as f:
        yield f
    os.rename(temp, filename)

def get_cpu(pids):
    cpus = []
    for pid in pids:
        p = psutil.Process(pid)
        cpu = p.cpu_percent(interval=None)
        cpus.append(cpu)
    
    return statistics.mean(cpus)

def get_memory(pids):
    mem = []
    for pid in pids:
        p = psutil.Process(pid)
        m = p.memory_full_info().vms/(1024.0 ** 3)
        mem.append(m)

    return statistics.mean(mem)

def get_pids(ports):
    pids = []
    pcss = psutil.net_connections()
    for p in pcss:
        if p.laddr[1] in ports:
            pids.append(p.pid)

    return pids

def measure_hw(ports):
    start = time.time()
    finish = 0
    table = pd.DataFrame()
    times = 0
    print('Started collect')
    while finish <= 30:
        processTime = times
        processCpu =  psutil.cpu_percent(interval=5)  #mensure_cpu(pids)
        processMem =  psutil.virtual_memory().percent
        times += 1
        finish += int((time.time() - start)/3600)
        time.sleep(20)
        print("Mem: {1}, CPU: {0}".format(processCpu, processMem))
        table = table.append({"Time":  processTime, "UsageCPU": processCpu, "UsageMem": processMem}, ignore_index=True)
        table.to_csv('./table_hw.csv', sep=';', header=False, index=False)
 

if __name__ == "__main__":
    
    ports = [17056, 17051]

    
    #pid = int(sys.argv[1])

    print('Get pids')
    pids = get_pids(ports)
    
    measure_hw(ports)

