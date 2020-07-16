import numpy as np
import pandas as pd
import uuid
import json
import requests

from datetime import datetime
from pymongo import MongoClient

if __name__ == "__main__":
    dicoms = pd.read_csv("./patients_dicom_new.csv", sep=";")
    dicoms = dicoms.dropna(subset=['dicomID'])
    dicoms['machineModel'] = list(map(lambda x: "AXAX1E20", range(len(dicoms['patientID']))))
    dicoms['patientAge'] = dicoms['patientAge'].replace(np.nan, 0, regex=True)
    dicoms['patientAge'] = list(map(lambda x: int(x), dicoms['patientAge']))
    dicoms['patientHeigth'] = dicoms['patientHeigth'].replace(np.nan, 0.0, regex=True)
    dicoms['patientWeigth'] = dicoms['patientWeigth'].replace(np.nan, 0.0, regex=True)
    dicoms['patientInsuranceplan'] = list(map(lambda x: str(x), dicoms['patientInsuranceplan']))
    dicoms['patientTelephone'] = list(map(lambda x: str(x), dicoms['patientTelephone']))
    dicoms['patientID'] = list(map(lambda x: str(x), dicoms['patientID']))
    dicoms['patientRace'] = dicoms['patientRace'].replace(np.nan, "Black", regex=True)
    dicoms['patientRace'] = dicoms['patientRace'].replace("Not Reported", "Black", regex=True)
    dicoms['patientRace'] = list(map(lambda x: str(x), dicoms['patientRace']))
    dicoms['patientGender'] = dicoms['patientGender'].replace(np.nan, "Unknown", regex=True)
    dicoms['timestamp'] = list(map(lambda x: datetime.now(),range(len(dicoms['patientID']))))
    dicoms = dicoms.replace(np.nan, "Unknown", regex=True)
    values_json = dicoms.reset_index(drop=True).to_dict(orient='records')

    mongo_client = MongoClient('mongodb://localhost:27020')

    db = mongo_client.get_database('private')
    col = db.create_collection('privData')

    for data in values_json:
        col.insert_one(data)