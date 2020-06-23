from peer import Peer
from experiment import Experiment

class Server():

    def init(self):
        peer0_hprovider = Peer('35.211.104.239',7051,"hprovider")
        peer1_hprovider = Peer('35.211.104.239',8051,"hprovider")

        peer0_research = Peer('35.211.104.239',9051,"research")
        peer1_research = Peer('35.211.104.239',10051,"research")

        peer0_patient = Peer('35.211.104.239',12051,"patient")
        peer1_patient = Peer('35.211.104.239',13051,"patient")

        expr = Experiment()

        expr.add_peer(peer0_hprovider)
        expr.add_peer(peer1_hprovider)
        expr.add_peer(peer0_research)
        expr.add_peer(peer1_research)
        expr.add_peer(peer0_patient)
        expr.add_peer(peer1_patient)

        expr.run_hw_experiments()


class Client():

    def init(self):
        peer0_hprovider = Peer('35.211.104.239',7051,"hprovider")
        peer1_hprovider = Peer('35.211.104.239',8051,"hprovider")

        peer0_research = Peer('35.211.104.239',9051,"research")
        peer1_research = Peer('35.211.104.239',10051,"research")

        peer0_patient = Peer('35.211.104.239',12051,"patient")
        peer1_patient = Peer('35.211.104.239',13051,"patient")

        expr = Experiment()

        expr.add_peer(peer0_hprovider)
        expr.add_peer(peer1_hprovider)
        expr.add_peer(peer0_research)
        expr.add_peer(peer1_research)
        expr.add_peer(peer0_patient)
        expr.add_peer(peer1_patient)

        expr.run_network_experiments('./dataset/patients_dicom.csv',"35.211.104.239",3000)