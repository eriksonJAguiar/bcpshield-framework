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

def read_latency_time(file_out, port):
    start_t = time.time()
    lt_by_org = measure_latency(host="35.211.104.239", port=port, runs=1000, timeout=5)
    #lt_peer1_org1 = measure_latency(host="35.211.104.239", port=port, runs=1000, timeout=5)
    end_t = (time.time()) - start_t
    run_times = []
    table = pd.DataFrame()

    for t in range(1000):
        run_times.append(start_t*5)
    
    table['latancy'] = lt_by_org

    table.to_csv(file_out, sep=",")
    
def read_dicom_tps_time(file_in,file_out, ports):
    dicoms = pd.read_csv(file_in,sep=";")
    dicoms_json = dicoms.to_json(orient='split')
    start_run = time.time()
    table_tps_time = pd.DataFrame()
    for tr in range(1,30):
        transaction_size = tr*len(dicoms_json)
        for dcm in dicoms_json:
            start_latency = time.time()
            resp = requests.get("35.211.104.239:3000",params=dcm)
            end_latency = time.time() - start_latency
        end_t = time.time() - start_run
        tps = end_t/transaction_size
        table_tps_time = table_tps_time.append({"Time":  end_t, "TPS": tps}, ignore_index=True)
    
    table.to_csv("%s"%(file_out), sep=';', header=True, index=False)


if __name__ == "__main__":
    
    ports = [17056, 17051]

    
    #pid = int(sys.argv[1])

    print('Get pids')
    pids = get_pids(ports)
    
    measure_hw(ports)

