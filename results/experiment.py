from peer import Peer
from tcp_latency import measure_latency

import pandas as pd
import numpy as np

import threading
import psutil
import time
import requests
import socket
import queue
import csv
import json
import os


class Experiment(object):

    def __init__(self) -> None:
        self.__peers: list(Peer) = list()
        self.__thread_request: list(threading.Thread) = list()
        self.__global_time: float = 0.0
        self.__writer_file_lock = threading.Lock()

    def add_peer(self, peer: Peer) -> None:
        """Function to add peer on experiments

        Args:
            peer (Peer): peer to add
        """
        self.__peers.append(peer)

    # def __get_pid(self, port: int) -> int:
    #     """function to get pid of the process

    #     Args:
    #         port (int): port to indate the process

    #     Returns:
    #         int: pid of the process
    #     """
    #     connections: list = psutil.net_connections()
    #     pid: int = None
    #     pid = list(filter(lambda c: c.pid if c.laddr[1] == port else None, connections))

    #     return pid[0].pid

    def get_pid(self, port: int) -> int:
        """function to get pid of the process

        Args:
            port (int): port to indate the process

        Returns:
            int: pid of the process
        """
        while self.__writer_file_lock.locked():
            continue

        self.__writer_file_lock.acquire()
        os.system("sudo lsof -n -i :%d  | awk '/LISTEN/{print $2}' >> pid_%d.txt"%(port,port))
        
        pid: int

        with open("pid_%d.txt"%(port), "r") as f:
            pid = int(f.readline())
  
        os.system("rm pid_%d.txt"%(port))

        self.__writer_file_lock.release()

        return pid

    def __measure_memory_per_time(self, peer: Peer) -> None:
        """function to mensure a memory usage per time

        Args:
            peer (Peer): peer to access ports
        """
        pid: int = self.get_pid(peer.port)
        process: psutil.Process = psutil.Process(pid)
        time.sleep(2)
        memory_percent = round(process.memory_percent(), 4)
        time_exec = round(time.time() - self.__global_time, 4)

        with self.__writer_file_lock:
            with open("cpu_%s.csv" % (peer.org), mode="a+") as csv_file:
                f = csv.writer(csv_file, delimiter=";")
                f.writerow([time_exec, memory_percent])

    def __measure_cpu_per_time(self, peer: Peer) -> None:
        """function to mensure a cpu usage per time

        Args:
            peer (Peer): peer to access ports
        """

        pid: int = self.get_pid(peer.port)
        process: psutil.Process = psutil.Process(pid)
        cpu_percent = round(process.cpu_percent(interval=5), 4)
        time_exec = round(time.time() - self.__global_time, 4)

        with self.__writer_file_lock:
            with open("memory_%s.csv" % (peer.org), mode="a+") as csv_file:
                f = csv.writer(csv_file, delimiter=";")
                f.writerow([time_exec, cpu_percent])

    def __socket_stop_experiments(self) -> None:
        """Socket to listen when experiments send a message end
        """
        server_experiments = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        server_experiments.bind((socket.gethostname(), 4000))
        server_experiments.listen(2)

        conn, addr = server_experiments.accept()

        end_experiments: bool = False
        while end_experiments is False:
            data = conn.recv(1024)
            self.__end_experiments = bool(data)

        server_experiments.close()

    def run_hw_experiments(self) -> None:
        """function to mesure hardware performance on experiments
        """
        thr_end_experiments = threading.Thread(
            target=self.__socket_stop_experiments, args=())
        thr_end_experiments.start()

        while thr_end_experiments.is_alive():
            for peer in self.__peers:
                aux_thread_cpu = threading.Thread(
                    target=self.__measure_cpu_per_time, args=(peer,))
                aux_thread_memory = threading.Thread(
                    target=self.__measure_memory_per_time, args=(peer,))
                aux_thread_cpu.start()
                aux_thread_memory.start()

            time.sleep(2)

    def __measure_latency_per_tps(self, peer: Peer) -> None:
        """Function to measure latency

        Args:
            peer (Peer): Peer to measure latency
        """
        org_latency = measure_latency(
            host=peer.ip, port=peer.port, runs=1000, timeout=5)
        end_time: float = time.time()

        run_times: list(float) = list()

        for i in range(1, 1000):
            run_times.append((end_time-self.__global_time)*5)

        with open("latency_%s.csv" % (peer.org), mode="a+") as file_csv:
            f = csv.DictWriter(file_csv, delimiter=";",
                               fieldnames=['Time', 'Latency'])

            f.writeheader()
            for (t, lt) in zip(run_times, org_latency):
                f.writerow({'Time': t, 'Latency': lt})

    def __measure_throughput_per_tps(self, peer: Peer) -> float:
        pass

    def send_request(self, file_in: str, ip_api: str, port_api: str) -> None:       
        """Send requets to server to aim measure metrics and evaluate the system

        Args:
            file_in (str): input file json to represents 
            ip_api (str): API request ip 
            port_api (str) API request port
        """
        dicoms: pd.DataFrame = pd.read_csv(file_in, sep=";")
        dicoms['user'] = list(map(lambda x: "erikson", range(len(dicoms['patientID']))))
        dicoms['machineModel'] = list(map(lambda x: "AXAX1E20", range(len(dicoms['patientID']))))
        dicoms['patientAge'] =  dicoms['patientAge'].replace(np.nan, "0", regex=True)
        dicoms['patientHeigth'] =  dicoms['patientHeigth'].replace(np.nan, "0", regex=True)
        dicoms['patientWeigth'] =  dicoms['patientWeigth'].replace(np.nan, "0", regex=True)
        dicoms['patientInsuranceplan'] = list(map(lambda x: str(x), dicoms['patientInsuranceplan']))
        dicoms['patientID'] = list(map(lambda x: str(x), dicoms['patientID']))
        dicoms['patientTelephone'] = list(map(lambda x: str(x), dicoms['patientTelephone']))
        dicoms['patientAge'] = list(map(lambda x: str(x), dicoms['patientAge']))
        dicoms['patientHeigth'] = list(map(lambda x: str(x), dicoms['patientHeigth']))
        dicoms['patientWeigth'] = list(map(lambda x: str(x), dicoms['patientWeigth']))

        dicoms =  dicoms.replace(np.nan, " ", regex=True)
        dicoms_dict: str = dicoms.to_dict(orient='records')
        self.__global_time = time.time()
        table_tps_time: pd.DataFrame = pd.DataFrame()
        print("Iniciando envio ...")
        #time.sleep(10)
        for tr in range(1, 30):
            transaction_size: int = tr*len(dicoms_dict)
            for dcm in dicoms_dict:
                try:
                    url = "http://%s:%s/api/addAsset"%(ip_api, port_api)
                    payload = json.dumps(dcm)
                    headers = { 'Content-Type': "application/json" }
                    resp: requests.Response = requests.request("POST", url, data=payload, headers=headers)
                    end_t: float = round(time.time() - self.__global_time, 4)
                    tps: float = round(end_t/transaction_size, 4)
                    table_tps_time = table_tps_time.append(
                        {"Time":  end_t, "TPS": tps}, ignore_index=True)
                    time.sleep(5)
                except:
                    pass

        # Send signal to finsh request and experiments
        client_request = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        client_request.connect((ip_api, 4000))
        client_request.sendall(b'True')
        client_request.close()

        table_tps_time.to_csv("tps_per_time_api.csv",
                              sep=';', header=True, index=False)

    def run_network_experiments(self, file_test_json: str, ip_api: str, port_api: str) -> None:

        req_thr: threading.Thread = threading.Thread(
            target=self.send_request, args=(file_test_json, ip_api, port_api))
        req_thr.start()

        while req_thr.is_alive():
            for peer in self.__peers:
                thr: threading.Thread = threading.Thread(
                    target=self.__measure_latency_per_tps, args=(peer,))

                thr.start()

            time.sleep(2)
