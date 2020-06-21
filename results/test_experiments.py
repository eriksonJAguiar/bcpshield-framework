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
