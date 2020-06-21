class Peer(object):

    def __init__(self, ip:str, port:int, org:str) -> None:
        self.__ip:str = ip
        self.__port:int = port
        self.__org:str = org
    
    @property
    def ip(self) -> str:
        return self.__ip
    
    @property
    def port(self) -> int:
        return self.__port
    
    @property
    def org(self) -> str:
        return self.__org
