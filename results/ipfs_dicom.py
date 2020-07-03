from pathlib import Path
from shutil import copy2
from pydicom import dcmread
from requests import request

import ipfshttpclient
import glob
import random
import os


class IpfsDicom(object):
    """Interation between IFPS and Dicom imaging to insert token and get files on internal database user

    """

    def __init__(self, root: str = None, ipfs_ip: str = None):
        self.__ipfs_ip = ipfs_ip
        if not root is None:
            self.__root = root
            self.__ipfs_path = os.path.join(Path.home(),".bcshield-temp")
            if not os.path.exists(self.__ipfs_path):
                os.makedirs(self.__ipfs_path)
        
    
    def __search_files(self, file_name: str) -> str:
        """Search local files on respository

        Args:
            file_id (str): Id file to share
            token (str): Token generate from request
        Returns:
            (str): new file path that will be insert token
        """

        files_path: list(str) = list(Path(self.__root).rglob("*.dcm"))
        amount_files: int = len(files_path)
        index_file: int = random.randint(0,(amount_files-1))
        file_chosen: str = files_path[index_file]
        new_file_path = os.path.join(self.__ipfs_path,file_chosen.name)
        copy2(file_chosen, new_file_path)
        
        return new_file_path

    def __insert_token(self, token: str, file_path: str) -> None:
        """insert token from bcshield on dicom imaging

        Args:
            token (str): Token generate from request
            file_name (str): name files will be send 
        """

        ds = dcmread(file_path)


        #? Add tag token
        ds.add_new([0x08,0x09], 'LO', token)

        ds.save_as(file_path)
 
    def send_dicom(self, file_name: str, token: str) -> str:
        """Send file to IPFS network and insert token on dicom files

        Args:
            file_name (str): file name that will be shared
            token (str): token will be insert on dicom

        Returns:
            str: hash to acess image on IPFS network
        """
        try:
            #! manage dicom images
            url_file = self.__search_files(file_name)
            self.__insert_token(token, url_file)

            #! Enviar para o ifps
            IPFS_API = ipfshttpclient.Client(f"/ip4/{self.__ipfs_ip}/tcp/5001/http")
            ifps_resp: dict = IPFS_API.add(url_file)
            IPFS_API.close()

            return ifps_resp['Hash']
        except:
            return None
        
    #QmS27FN1oGLqNdMZrEUaaN65vPQkDxnP1NuiBjUXwShNGL
    #QmYP4T25FBFWNnPeNKQyJZd2NbSkYSPiWKfHQb42L9JysM
    def get_dicom(self, hash_value: str) -> bool:
        try:
            url_path: str = os.path.join(Path.home(),"bcshield-recovered")
            
            if not os.path.exists(url_path):
                os.makedirs(url_path)

            IPFS_API = ipfshttpclient.Client(f"/ip4/{self.__ipfs_ip}/tcp/5001/http")
            IPFS_API.get(hash_value,url_path)
            IPFS_API.close()

            # params: dict = {
            #     "arg": hash_value,
            #     "output": url_path+"new_dicom"
            # }

            # resp_value = request("POST", url=f"http://{self.__ipfs_ip}:5001/api/v0/get", params=params)

            #print(resp_value)

            return True
        except:
            return False
            
        
        