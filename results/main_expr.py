from peer import Peer
from experiment import Experiment

class Server():

    def init(self):
        peer0_hprovider = Peer('127.0.0.1',7051,"hprovider")
        peer1_hprovider = Peer('127.0.0.1',8051,"hprovider")

        peer0_research = Peer('127.0.0.1',9051,"research")
        peer1_research = Peer('127.0.0.1',10051,"research")

        peer0_patient = Peer('127.0.0.1',12051,"patient")
        peer1_patient = Peer('127.0.0.1',13051,"patient")

        expr = Experiment()

        expr.add_peer(peer0_hprovider)
        expr.add_peer(peer1_hprovider)
        expr.add_peer(peer0_research)
        expr.add_peer(peer1_research)
        expr.add_peer(peer0_patient)
        expr.add_peer(peer1_patient)

        expr.run_hw_experiments()


class Client():

    def init(self, transaction_number=50):
        peer0_hprovider = Peer('35.211.244.95',7051,"hprovider")
        peer1_hprovider = Peer('35.211.244.95',8051,"hprovider")

        peer0_research = Peer('35.211.244.95',9051,"research")
        peer1_research = Peer('35.211.244.95',10051,"research")

        peer0_patient = Peer('35.211.244.95',12051,"patient")
        peer1_patient = Peer('35.211.244.95',13051,"patient")

        expr = Experiment()

        expr.add_peer(peer0_hprovider)
        expr.add_peer(peer1_hprovider)
        expr.add_peer(peer0_research)
        expr.add_peer(peer1_research)
        expr.add_peer(peer0_patient)
        expr.add_peer(peer1_patient)

        print("Start running ...")
        for i in range(1):
            print(f"Running {i+1} round")
            expr.measure_transations(transaction_number)
        
        print("Finish running")