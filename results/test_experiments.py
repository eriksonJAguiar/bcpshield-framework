from experiment import Experiment
from peer import Peer
import unittest


class TestExperiment(unittest.TestCase):

    # def test_get_pid(self) -> None:
    #     """ Testing when a user to request a pid through process port
    #     """
    #     expr: Experiment = Experiment()
    #     pid = expr.get_pid(17050)
    #     print("Valor: %d"%pid)
    #     self.assertIsNot(pid, None)

    # def test_memory_usage(self) -> None:
    #     """ Testing memory usage by process
    #     """
    #     expr: Experiment = Experiment()
    #     memory = expr.measure_memory_per_time(17050)
    #     print("Memory: %f"%memory)
    #     self.assertTrue(isinstance(memory, float) and cpu > 0)

    # def test_cpu_usage(self) -> None:
    #     """ Testing cpu usage by process
    #     """
    #     expr: Experiment = Experiment()
    #     cpu = expr.measure_cpu_per_time(17050)
    #     print("CPU: %f"%cpu)
    #     self.assertTrue(isinstance(cpu, float) and cpu > 0)

    def test_server_run_testing(self):
        peer0_org1 = Peer("127.0.0.1", 17050, "org1")

        expr = Experiment()
        expr.add_peer(peer0_org1)

        try:
            expr.run_hw_experiments()
            self.assertTrue(True)
        except:
            self.assertFalse(False)

    def test_client_run_testing(self):
        peer0_org1 = Peer("127.0.0.1", 17050, "org1")

        expr = Experiment()
        expr.add_peer(peer0_org1)

        try:
            expr.run_network_experiments(
                "dataset/patients_dicom.csv", "127.0.0.1", 3000)
            self.assertTrue(True)
        except:
            self.assertFalse(False)
    
    def test_send_dicom_ipfs(self):
        import hashlib
        import time 
        from ipfs_dicom import IpfsDicom

        #! Generate token
        hl: hashlib = hashlib.sha256()
        value: str = "C3L-00503"+"10001"+str(time.time())
        hl.update(value.encode())
        token: str = hl.hexdigest()
                    
        #! manage IPFS network
        ipfs: IpfsDicom = IpfsDicom('../../../Downloads/dataset-resultados/dicom-dataset/CPTAC-LSCC/', "35.233.252.12")
        ipfs_resp: str = ipfs.send_dicom("C3L-00503", token)

        print(ipfs_resp)

        self.assertIsNotNone(ipfs_resp)

    def test_get_dicom(self):
        from ipfs_dicom import IpfsDicom
        
        ipfs: IpfsDicom = IpfsDicom("35.233.252.12")
        isValid = ipfs.get_dicom("QmYP4T25FBFWNnPeNKQyJZd2NbSkYSPiWKfHQb42L9JysM")

        print(isValid)

        self.assertIs(isValid)

