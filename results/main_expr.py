from peer import Peer
from experiment import Experiment, MensurePostSimple, MensureGetSimple, MensurePostPriv, MensureGetPriv

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

    def init(self, method_type = "read", limit_transactions=50, round=1):

        method = None

        if method_type == "write":
            method = MensurePostSimple()
        elif method_type == "read":
            method = MensureGetSimple()
        elif method_type == "write-priv":
            method = MensurePostPriv()
        elif method_type == "read-priv":
            method = MensureGetPriv()

        expr = Experiment()
        print(f"Start running for {limit_transactions} transactions ...")
        
        for i in range(round):
            print(f"Running round {i+1} ...")
            expr.measure_transations(limit_transactions, method_type, method)
            print(f"End round {i+1} ...")

        print("Finish running")

if __name__ == "__main__":
    c = Client()
    c.init("read", 1000)