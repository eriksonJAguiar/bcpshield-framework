from peer import Peer
from tcp_latency import measure_latency
from ipfs_dicom import IpfsDicom

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
import hashlib



class Experiment(object):

    def __init__(self):
        self.__peers: list(Peer) = list()
        self.__thread_request: list(threading.Thread) = list()
        self.__global_time: float = 0.0
        self.__writer_file_lock = threading.Lock()
        os.system("mkdir ~/.ipfs-temp")

    def add_peer(self, peer: Peer) -> None:
        """Function to add peer on experiments

        Args:
            peer (Peer): peer to add
        """
        self.__peers.append(peer)

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
        org_latency = measure_latency(host=peer.ip, port=peer.port, runs=1000, timeout=5)
        end_time: float = time.time()

        run_times: list(float) = list()

        for i in range(1, 1000):
            run_times.append((end_time-self.__global_time)/(5*i))

        while self.__writer_file_lock.locked():
            continue
        
        self.__writer_file_lock.acquire()

        with open("latency_%s.csv" % (peer.org), mode="a+") as file_csv:
            f = csv.DictWriter(file_csv, delimiter=";",
                               fieldnames=['Time', 'Latency'])

            f.writeheader()
            for (t, lt) in zip(run_times, org_latency):
                f.writerow({'Time': t, 'Latency': lt})
        
        self.__writer_file_lock.release()

    def __measure_throughput_per_tps(self, peer: Peer) -> float:
        pass
       
    def run_network_experiments(self, file_test_json: str, ip_api: str, port_api: str, method) -> None:
        """Send requets to server to aim measure network metrics and evaluate the system
            We using the design pattern strategy

        Args:
            file_in (str): input file json to represents 
            ip_api (str): API request ip 
            port_api (str): API request port
            method (class): Represent the class will be used to acess send_request
        """
        req_thr: threading.Thread = threading.Thread(
            target=method.send_request, args=(file_test_json, ip_api, port_api))

        req_thr.start()

        while req_thr.is_alive():
            for peer in self.__peers:
                thr: threading.Thread = threading.Thread(
                    target=self.__measure_latency_per_tps, args=(peer,))

                thr.start()
            
            time.sleep(2)
        

class RequestGetAsset(object):
    """Class to Strategy pattern for describes the method send_request GET elements

    """
    def __init__(self):
        self.__global_time = None
    
    def send_request(self, file_in: str, ip_api: str, port_api: str) -> None:       
        """Send requets to server to aim measure metrics and evaluate the system

        Args:
            file_in (str): input file json to represents 
            ip_api (str): API request ip 
            port_api (str) API request port
        """
        dicoms: pd.DataFrame = pd.read_csv(file_in, sep=";")
        dicoms_dict: str = dicoms.to_dict(orient='records')

        self.__global_time = time.time()
        table_tps_time: pd.DataFrame = pd.DataFrame()
        print("Iniciando GET Asset ...")
        #time.sleep(10)
        for tr in range(1, 6):
            dicomId = list(map(lambda d: d+str(tr),dicoms_dict['dicomID']))
            transaction_size: int = tr*len(dicoms_dict)
            for dcm_id in dicomId:
                try:
                    req_json = {
                        'dicomId': dcm_id,
                        'user': "erikson"
                    }
                    url = "http://%s:%s/api/getAsset"%(ip_api, port_api)
                    payload = json.dumps(req_json)
                    headers = { 'Content-Type': "application/json" }
                    resp: requests.Response = requests.request("GET", url, data=payload, headers=headers)
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

        table_tps_time.to_csv("tps_per_time_%d_GET.csv"%(port_api),
                              sep=';', header=True, index=False)

class RequestPost(object):
    """Class to Strategy pattern for describes the method send_request POST elements 
        on blockchain

    """

    def __init__(self):
        self.__global_time = None

    def send_request(self, file_in: str, ip_api: str, port_api: str) -> None:
        """Send requets to server to aim measure metrics and evaluate the system

        Args:
            file_in (str): input file json to represents 
            ip_api (str): API request ip 
            port_api (str) API request port
        """
        dicoms: pd.DataFrame = pd.read_csv(file_in, sep=";")
        dicoms['user'] = list(
            map(lambda x: "erikson", range(len(dicoms['patientID']))))
        dicoms['machineModel'] = list(
            map(lambda x: "AXAX1E20", range(len(dicoms['patientID']))))
        dicoms['patientAge'] = dicoms['patientAge'].replace(
            np.nan, "0", regex=True)
        dicoms['patientAge'] = list(
            map(lambda x: int(x), dicoms['patientAge']))
        dicoms['patientHeigth'] = dicoms['patientHeigth'].replace(
            np.nan, "0.0", regex=True)
        dicoms['patientWeigth'] = dicoms['patientWeigth'].replace(
            np.nan, "0.0", regex=True)
        dicoms['patientInsuranceplan'] = list(
            map(lambda x: str(x), dicoms['patientInsuranceplan']))
        dicoms['patientID'] = list(map(lambda x: str(x), dicoms['patientID']))
        dicoms['patientTelephone'] = list(
            map(lambda x: str(x), dicoms['patientTelephone']))
        dicoms['patientAge'] = list(
            map(lambda x: str(x), dicoms['patientAge']))
        dicoms['patientHeigth'] = list(
            map(lambda x: str(x), dicoms['patientHeigth']))
        dicoms['patientWeigth'] = list(
            map(lambda x: str(x), dicoms['patientWeigth']))

        dicoms = dicoms.replace(np.nan, " ", regex=True)
        dicoms_dict: list(dict) = dicoms.to_dict(orient='records')
        self.__global_time = time.time()
        table_tps_time: pd.DataFrame = pd.DataFrame()
        print("Send files ...")
        #time.sleep(10)
        for tr in range(1, 6):
            transaction_size: int = tr*len(dicoms_dict)
            for dcm in dicoms_dict:
                try:
                    dcm['dicomID'] = dcm['dicomID']+str(tr)
                    url = "http://%s:%s/api/addAsset" % (ip_api, port_api)
                    payload = json.dumps(dcm)
                    headers = {'Content-Type': "application/json"}
                    resp: requests.Response = requests.request(
                        "POST", url, data=payload, headers=headers)
                    print(resp.json())
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

class RequestGetKAnonymity(object):
    """Class to Strategy pattern for describes the method send_request GET elements applying 
       K-anonymity privacy

    """
    def __init__(self):
        self.__global_time = None
    
    def send_request(self, file_in: str, ip_api: str, port_api: str) -> None:       
        """Send requets to server to aim measure metrics and evaluate the system

        Args:
            file_in (str): input file json to represents 
            ip_api (str): API request ip 
            port_api (str) API request port
        """
        dicoms: pd.DataFrame = pd.read_csv(file_in, sep=";")
        dicoms_dict: str = dicoms.to_dict(orient='records')

        self.__global_time = time.time()
        table_tps_time: pd.DataFrame = pd.DataFrame()
        anonymized_files = list()

        print("Iniciando envio ...")
        #time.sleep(10)
        for tr in range(1, 6):
            #!Convert values 
            dicomsId: list(str) = list(map(lambda d: d+str(tr), dicoms['dicomID'].tolist()))
            patientId: list(str) = dicoms['patientID'].tolist()
            transaction_size: int = tr*len(dicoms_dict)
            for dcm_id, pat_id in zip(dicomsId, patientId):
                try:
                    #! Generate token
                    hl: hashlib = hashlib.sha256()
                    value: str = dcm_id+pat_id+str(time.time())
                    hl.update(value.encode())
                    token: str = hl.hexdigest()
                    
                    #! manage IPFS network
                    ipfs: IpfsDicom = IpfsDicom('../../../Downloads/dataset-resultados/dicom-dataset/CPTAC-LSCC/', "35.233.252.12")
                    ipfs_resp: str = ipfs.send_dicom(dcm_id, token)
                                   
                    #!Send asset to a doctor
                    req_share = {
                        "dicomID": dcm_id,
                        "patientID": pat_id,
                        "doctorID": "1100",
                        "user": "erikson",
                        "hashIPFS": ipfs_resp
                    }
                    payload_share: str = json.dumps(req_share)
                    url_post: str = "http://%s:%s/api/shareAssetWithDoctor"%(ip_api, port_api)
                    headers: dict = { 'Content-Type': "application/json" }
                    share_resp: requests.Response = requests.request("POST", url_post, data=payload_share, headers=headers)
                    share_resp_json = share_resp.json()

                    #!Get imaging using K-anonymity
                    req_json = {
	                    "user": "erikson",
	                    "hashIPFS": share_resp_json['id']
                    }
                    url_get = "http://%s:%s/api/getSharedAssetWithDoctor"%(ip_api, port_api)
                    payload_get = json.dumps(req_json)
                    resp_get: requests.Response = requests.request("GET", url_get, data=payload_get, headers=headers)
                    
                    #? Add metrics tps
                    end_t: float = round(time.time() - self.__global_time, 4)
                    tps: float = round(end_t/transaction_size, 4)
                    table_tps_time = table_tps_time.append(
                        {"Time":  end_t, "TPS": tps}, ignore_index=True)
                    
                    #? Save files recovered
                    anonymized_files.append(resp_get.json())                    
                    
                    time.sleep(5)
                except:
                    pass

        # !Send signal to finsh request and experiments
        client_request = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        client_request.connect((ip_api, 4000))
        client_request.sendall(b'True')
        client_request.close()



        with open("./dicoms_result_anonimity.json", "w") as file_out:
                json.dump(anonymized_files, file_out)

        table_tps_time.to_csv("tps_per_time_%d_GET.csv"%(port_api),
                              sep=';', header=True, index=False)