from peer import Peer
from tcp_latency import measure_latency
from ipfs_dicom import IpfsDicom
from gevent import monkey
monkey.patch_all(thread=False,select=False)


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
import uuid
import queue
import statistics
import random
import grequests

# Sharing values
transactions_number = 1
flag = 1
ls_trans_amount = list()
ls_trans_time = list()
data_id = list()
general_time = None
throughput_general = list()
size_send = 0


class Experiment(object):

    def __init__(self):
        self.__peers: list(Peer) = list()
        self.__thread_request: list(threading.Thread) = list()
        self.__global_time: float = 0.0
        self.__writer_file_lock = threading.Lock()
        #os.system("mkdir ~/.ipfs-temp")

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
        os.system(
            "sudo lsof -n -i :%d  | awk '/LISTEN/{print $2}' >> pid_%d.txt" % (port, port))

        pid: int

        with open("pid_%d.txt" % (port), "r") as f:
            pid = int(f.readline())

        os.system("rm pid_%d.txt" % (port))

        self.__writer_file_lock.release()

        return pid

    def __measure_memory_per_time(self, peer: Peer) -> None:
        """function to mensure a memory usage per time

        Args:
            peer (Peer): peer to access ports
        """
        try:
            pid: int = self.get_pid(peer.port)
            process: psutil.Process = psutil.Process(pid)
            time.sleep(2)
            init_t = time.time()
            memory_percent = round(process.memory_percent(), 4)
            time_exec = round(time.time() - init_t, 4)
        except:
            pass

        with self.__writer_file_lock:
            with open("cpu_%s.csv" % (peer.org), mode="a+") as csv_file:
                f = csv.writer(csv_file, delimiter=";")
                f.writerow([time_exec, memory_percent])

    def __measure_cpu_per_time(self, peer: Peer) -> None:
        """function to mensure a cpu usage per time

        Args:
            peer (Peer): peer to access ports
        """
        try:
            pid: int = self.get_pid(peer.port)
            process: psutil.Process = psutil.Process(pid)
            init_t = time.time()
            cpu_percent = round(process.cpu_percent(interval=5), 4)
            time_exec = round((time.time() - init_t)/5, 4)
        except:
            pass

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
            try:
                for peer in self.__peers:
                    aux_thread_cpu = threading.Thread(
                        target=self.__measure_cpu_per_time, args=(peer,))
                    aux_thread_memory = threading.Thread(
                        target=self.__measure_memory_per_time, args=(peer,))
                    aux_thread_cpu.start()
                    aux_thread_memory.start()
                    sleep(1)
                    aux_thread_cpu.join()
                    aux_thread_memory.join()
            except:
                aux_thread_cpu.join()
                aux_thread_memory.join()

            time.sleep(2)

    def __measure_latency_per_tps(self, peer: Peer) -> None:
        """Function to measure latency

        Args:
            peer (Peer): Peer to measure latency
        """
        init_time: float = time.time()
        org_latency = measure_latency(
            host=peer.ip, port=peer.port, runs=1000, timeout=5)
        end_time: float = time.time()

        run_times: list(float) = list()

        for i in range(1, 1000):
            run_times.append((end_time-init_time)/(5*i))

        # while self.__writer_file_lock.locked():
        #    continue

        # self.__writer_file_lock.acquire()

        with self.__writer_file_lock:
            file_exists = os.path.exists("latency_%s.csv" % (peer.org))
            with open("latency_%s.csv" % (peer.org), mode="a+") as file_csv:
                f = csv.DictWriter(file_csv, delimiter=";",
                                fieldnames=['Time', 'Latency'])

                if not file_exists:
                    f.writeheader()

                for (t, lt) in zip(run_times, org_latency):
                    f.writerow({'Time': t, 'Latency': lt})

        # self.__writer_file_lock.release()

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

    def measure_transations(self, limit_transactions, requestType, method):
        
        global general_time
        global ls_trans_time
        global throughput_general
        global data_id
        global transactions_number

        general_time = time.time()    
        aux_tran_time = 1
        tps_list = list()
        lat_avg = list()
        thoughput_avg = list()

        try:
            while transactions_number <= limit_transactions:

                method.send_request()
                print(f"Transactions sent { transactions_number }")

                tps = transactions_number/(time.time() - general_time)
                tps_list.append(tps)
                aux_tran_time += 1
                lat_avg.append(statistics.mean(ls_trans_time))
                thoughput_avg.append(statistics.mean(throughput_general))
               
                time.sleep(0.3)
        except:
            print("Error")
            pass
                
        #Grava transactions per time for write
        fl_true = os.path.exists("transactions_per_time_%s.csv"%(requestType))
        with open("transactions_per_time_%s.csv"%(requestType), mode="a+") as file_csv:
            f = csv.DictWriter(file_csv, delimiter=";", fieldnames=['Time', 'TrAmount'])

            if not fl_true:
                f.writeheader()
                                          
            for (t,am) in zip(ls_trans_time, ls_trans_amount):
                f.writerow({'Time': t, 'TrAmount': am})
        
        #Grava Latency per tps
        fl_true = os.path.exists("latency_avg_per_tps_%s.csv"%(requestType))
        with open("latency_avg_per_tps_%s.csv"%(requestType), mode="a+") as file_csv:
            f = csv.DictWriter(file_csv, delimiter=";", fieldnames=['Latency', 'TPS'])

            if not fl_true:
                f.writeheader()
                                          
            for (lt, tps) in zip(lat_avg,  tps_list):
                f.writerow({'Latency': lt, 'TPS':  tps})
        
        #Grava Throughput per tps
        fl_true = os.path.exists("throughput_avg_per_tps_%s.csv"%(requestType))
        with open("throughput_avg_per_tps_%s.csv"%(requestType), mode="a+") as file_csv:
            f = csv.DictWriter(file_csv, delimiter=";", fieldnames=['Throughput', 'TPS'])

            if not fl_true:
                f.writeheader()
                                          
            for (th, tps) in zip(lat_avg,  tps_list):
                f.writerow({'Throughput': th, 'TPS':  tps})
        
        #Grava Ids
        if requestType ==  "write" or requestType ==  "write-priv":
            with open("dicom_ids_%s.txt"%(requestType), mode="a+") as file_txt:
                for d in data_id:
                    file_txt.write(d)
                    file_txt.write("\n")

        
        print("Finished!")
            
class RequestGetAsset(object):
    """Class to Strategy pattern for describes the method send_request GET elements

    """
    def __init__(self):
        self.__global_time = None
        self.__writer_file_lock = threading.Lock()
    
    def send_request(self, file_in: str, ip_api: str, port_api: str) -> None:       
        """Send requets to server to aim measure metrics and evaluate the system

        Args:
            file_in (str): input file json to represents 
            ip_api (str): API request ip 
            port_api (str) API request port
        """
        dicoms: pd.DataFrame = pd.read_csv(file_in, sep=";")
        dicoms = dicoms.dropna(subset=['dicomID'])
        dicoms_dict: list(dict) = dicoms.to_dict(orient='records')

        # self.__global_time = time.time()
        table_tps_time: pd.DataFrame = pd.DataFrame()

        print("Iniciando GET Asset ...")
        # time.sleep(10)
        for tr in range(1, 4):
            # dicomId = list(map(lambda d: d+str(tr),dicoms_dict['dicomID']))
            transaction_size: int = tr*len(dicoms_dict)
            for dcm in dicoms_dict:
                try:
                    print(dcm)
                    init_t = time.time()
                    dcm_id = str(dcm['dicomID'])+str(tr)
                    req_json = {
                        'dicomId': dcm_id,
                        'user': "erikson"
                    }
                    print(req_json)
                    url = "http://%s:%s/api/getAsset"%(ip_api, port_api)
                    payload = json.dumps(req_json)
                    headers = { 'Content-Type': "application/json" }
                    resp: requests.Response = requests.request("GET", url, data=payload, headers=headers)
                    print(resp)
                    
                    end_t: float = round(time.time() - init_t, 4)
                    tps: float = round(end_t/transaction_size, 4)
                    
                    with self.__writer_file_lock:
                        file_exists = os.path.exists("tps_per_time_read.csv")
                        with open("tps_per_time_read.csv", mode="a+") as file_csv:
                            f = csv.DictWriter(file_csv, delimiter=";",
                                            fieldnames=['Time', 'TPS'])

                            if not file_exists:
                                f.writeheader()
                           
                            f.writerow({'Time': end_t, 'TPS': tps})
                        
                        res_json_file = "./dicoms_result_anonimity.json"
                        if not os.path.exists(res_json_file):
                            with open(res_json_file, "w") as f:
                                json.dump(f, [], indent=4)
                        
                        with open(req_json_file) as f:
                            resp_aux = json.load(f)
                        
                        resp_aux.append(resp.json())

                        with open(res_json_file, "w") as f:
                            json.dump(f, resp_aux, indent=4)
                    
                    time.sleep(3)
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
        self.__writer_file_lock = threading.Lock()


    def send_request(self, file_in: str, ip_api: str, port_api: str) -> None:
        """Send requets to server to aim measure metrics and evaluate the system

        Args:
            file_in (str): input file json to represents 
            ip_api (str): API request ip 
            port_api (str) API request port
        """
        dicoms: pd.DataFrame = pd.read_csv(file_in, sep=";")
        dicoms = dicoms.dropna(subset=['dicomID'])
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
        
        print("Send files ...")

        for tr in range(1, 6):
            dicoms['dicomID'] = list(map(lambda d: d+str(tr), dicoms['dicomID'].values))
            dicoms_dict: list(dict) = dicoms.to_dict(orient='records')
            
            transaction_size: int = tr*len(dicoms_dict)
            
            for dcm in dicoms_dict:
                try:
                    init_t = time.time()
                    dcm['dicomID'] = dcm['dicomID']+str(tr)
                    url = "http://%s:%s/api/addAsset" % (ip_api, port_api)
                    payload = json.dumps(dcm)
                    headers = {'Content-Type': "application/json"}
                    resp: requests.Response = requests.request(
                        "POST", url, data=payload, headers=headers)
                    print(resp.json())
                    
                    end_t: float = round(time.time() - init_t, 4)
                    tps: float = round(end_t/transaction_size, 4)
                    
                    with self.__writer_file_lock:
                        file_exists = os.path.exists("tps_per_time.csv")
                        with open("tps_per_time.csv", mode="a+") as file_csv:
                            f = csv.DictWriter(file_csv, delimiter=";",
                                            fieldnames=['Time', 'TPS'])

                            if not file_exists:
                                f.writeheader()
                            
                           
                            f.writerow({'Time': end_t, 'TPS': tps})

                    time.sleep(2)
                except:
                    pass

        # Send signal to finsh request and experiments
        client_request = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        client_request.connect((ip_api, 4000))
        client_request.sendall(b'True')
        client_request.close()

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
        dicoms = dicoms.dropna(subset=['dicomID'])
        dicoms_dict: str = dicoms.to_dict(orient='records')

        self.__global_time = time.time()
        table_tps_time: pd.DataFrame = pd.DataFrame()
        anonymized_files = list()

        print("Iniciando envio ...")
        # time.sleep(10)
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
                    
                    # ? Add metrics tps
                    end_t: float = round(time.time() - self.__global_time, 4)
                    tps: float = round(end_t/transaction_size, 4)
                    table_tps_time = table_tps_time.append(
                        {"Time":  end_t, "TPS": tps}, ignore_index=True)
                    
                    # ? Save files recovered
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

class MensurePostSimple(object):
    """Class to Strategy pattern for describes the method send_request POST elements

    """
    def __init__(self):
        pass

    def send_request(self) -> None:
        """Send write requets to server to aim measure metrics and evaluate the system
        """
        
        global flag
        global transactions_number
        global ls_trans_amount
        global ls_trans_time
        global data_id
        global throughput_general
        global size_send
        
        data_values = list()
        data_values_json = list()


        for i in range(20*flag):
            data = {
                "patientHeigth": 1.75,
                "patientID": "11110",
                "patientRace": "White",
                "patientGender": "male",
                "patientWeigth": 70.5,
                "patientFirstname": "Erikson",
                "patientTelephone": "(43) 0000-0000",
                "machineModel": "AMX",
                "patientOrganization": "USP",
                "dicomID": str(uuid.uuid4()).join(flag),
                "patientAge": 23,
                "patientAddress": "ASasasasasas",
                "patientInsuranceplan": "IIIIIAAAA",
                "user": "erikson",
                "patientLastname": "Aguiar"
            }
            data_values_json.append(data)
            aux_d = json.dumps(data)
            data_values.append(aux_d)
            transactions_number += 1
            
        flag += 1
        
        start_send = time.time()
        url = "http://%s:%d/api/addAsset" % ("35.211.244.95",3000)
        headers = {'Content-Type': "application/json"}
        rs = (grequests.post(url, headers=headers,data=d) for d in data_values)
        grequests.map(rs)
        end_send_time = time.time() - start_send
        size_send += len(json.dumps(data_values_json).encode('utf8'))
        

        ls_trans_amount.append(transactions_number)
        ls_trans_time.append(end_send_time)
        throughput_general.append(size_send)
        data_id += list(map(lambda d: d['dicomID'], data_values_json))

class MensureGetSimple(object):
    """Class to Strategy pattern for describes the method send_request GET elements

    """
    def __init__(self):
        pass

    def send_request(self) -> None:
        """Send read requets to server to aim measure metrics and evaluate the system
        """
        
        global flag
        global transactions_number
        global ls_trans_amount
        global ls_trans_time
        global throughput_general
        global size_send
        
        data_values = list()
        data_values_json = list()

        dicomIds = list()

        with open("./table-results/write-results/dicom_ids_write.txt", "r") as f:
           for line in f:
               l = line.rstrip('\n')
               dicomIds.append(l)


        for i in range(50*flag):
            random.seed(time.time()*i)
            index = random.randint(0,(len(dicomIds)-1))
            data = {
                "user": "erikson",
                "dicomId": dicomIds[index]
            }
            data_values_json.append(data)
            aux_d = json.dumps(data)
            data_values.append(aux_d)
            transactions_number += 1
            
        flag += 1
        
        start_send = time.time()
        url = "http://%s:%d/api/getAsset" % ("35.211.244.95",3000)
        headers = {'Content-Type': "application/json"}
        rs = (grequests.get(url, headers=headers,data=d) for d in data_values)
        grequests.map(rs)
        end_send_time = time.time() - start_send
        size_send += len(json.dumps(data_values_json).encode('utf8'))
        

        ls_trans_amount.append(transactions_number)
        ls_trans_time.append(end_send_time)
        throughput_general.append(size_send)
        
class MensurePostPriv(object):
    """Class to Strategy pattern for describes the method send_request POST elements 
        using Private Method for HFL

    """
    def __init__(self):
        pass

    def send_request(self) -> None:
        """Send write Priv requets to server to aim measure metrics and evaluate the system
        """

        global flag
        global transactions_number
        global ls_trans_amount
        global ls_trans_time
        global data_id
        global throughput_general
        global size_send
        
        data_values = list()
        data_values_json = list()


        for i in range(50*flag):
            data = {
                "patientHeigth": 1.75,
                "patientID": "11110",
                "patientRace": "White",
                "patientGender": "male",
                "patientWeigth": 70.5,
                "patientFirstname": "Erikson",
                "patientTelephone": "(43) 0000-0000",
                "machineModel": "AMX",
                "patientOrganization": "USP",
                "dicomID": str(uuid.uuid4()),
                "patientAge": 23,
                "patientAddress": "ASasasasasas",
                "patientInsuranceplan": "IIIIIAAAA",
                "user": "erikson",
                "patientLastname": "Aguiar"
            }
            data_values_json.append(data)
            aux_d = json.dumps(data)
            data_values.append(aux_d)
            transactions_number += 1
            
        flag += 1
        
        start_send = time.time()
        url = "http://%s:%d/api/addAssetPriv" % ("35.211.244.95",3000)
        headers = {'Content-Type': "application/json"}
        rs = (grequests.post(url, headers=headers,data=d) for d in data_values)
        grequests.map(rs)
        end_send_time = time.time() - start_send
        size_send += len(json.dumps(data_values_json).encode('utf8'))
        

        ls_trans_amount.append(transactions_number)
        ls_trans_time.append(end_send_time)
        throughput_general.append(size_send)
        data_id += list(map(lambda d: d['dicomID'], data_values_json))

class MensureGetPriv(object):
    """Class to Strategy pattern for describes the method send_request GET elements
        using priv HLF

    """
    def __init__(self):
        pass

    def send_request(self) -> None:
        """Send read requets to server to aim measure metrics and evaluate the system
        """
        
        global flag
        global transactions_number
        global ls_trans_amount
        global ls_trans_time
        global throughput_general
        global size_send
        
        data_values = list()
        data_values_json = list()

        dicomIds = list()

        with open("./table-results/write-results/dicom_ids_write-priv.txt", "r") as f:
           for line in f:
               l = line.rstrip('\n')
               dicomIds.append(l)


        for i in range(50*flag):
            random.seed(time.time()*i)
            index = random.randint(0,(len(dicomIds)-1))
            data = {
                "user": "erikson",
                "dicomId": dicomIds[index]
            }
            data_values_json.append(data)
            aux_d = json.dumps(data)
            data_values.append(aux_d)
            transactions_number += 1
            
        flag += 1
        
        start_send = time.time()
        url = "http://%s:%d/api/getAssetPriv" % ("35.211.244.95",3000)
        headers = {'Content-Type': "application/json"}
        rs = (grequests.get(url, headers=headers,data=d) for d in data_values)
        grequests.map(rs)
        end_send_time = time.time() - start_send
        size_send += len(json.dumps(data_values_json).encode('utf8'))
        

        ls_trans_amount.append(transactions_number)
        ls_trans_time.append(end_send_time)
        throughput_general.append(size_send)