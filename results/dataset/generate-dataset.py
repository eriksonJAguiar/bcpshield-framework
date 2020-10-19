import pandas as pd
import json

def read_tabula_data(n=30000):
    dt_attibutes = pd.read_csv('./DatabsetExperimentosNYSDOH/PNDS_Individual_Q416_v2.csv', sep=",", nrows=n)
    new_database = pd.DataFrame()
    new_database['firstname'] = dt_attibutes['fname']
    new_database['lastname'] = dt_attibutes['lname']
    new_database['planid'] = dt_attibutes['medsid']
    new_database['patientID'] = dt_attibutes['npi']
    new_database['phone'] = dt_attibutes['phone']
    new_database['gender'] = list(map(lambda x: 'male' if x == 1 else 'female' , dt_attibutes['gender']))
    new_database['organization'] = dt_attibutes['Organization']
    new_database['address'] = dt_attibutes['staddres'] 
    new_database['city'] = dt_attibutes['city']
    new_database['state'] =  dt_attibutes['state']

    return new_database

def read_dicom_json():
    with open('imagens-data/LSCC.json', 'r') as json_file:
        json_data = json_file.read()
    
    data = json.loads(json_data)
    new_database = pd.DataFrame()

    dicomid = [] 
    age = [] 
    race = [] 
    height = [] 
    weight = []

    for json_row in data:
        
        if 'case_id' in json_row.keys():
             dicomid.append(json_row['case_id'])
        else:
            dicomid.append('NaN')
        
        if 'age' in json_row.keys():
            age.append(json_row['age'])
        else:
            age.append('NaN')

        if 'race' in json_row.keys():
            race.append(json_row['race'])
        else:
            race.append('NaN')
        
        if 'height_in_cm' in json_row.keys():
            height.append(json_row['height_in_cm'])
        else:
            height.append('NaN')
        
        if 'weight_in_kg' in json_row.keys():
            weight.append(json_row['weight_in_kg'])
        else:
            weight.append('NaN')

    new_database['dicomID'] = dicomid
    new_database['age'] = age
    new_database['race'] =  race
    new_database['height'] = height
    new_database['weight'] = weight


    return new_database


if __name__ == "__main__":

    database_generated = read_tabula_data(300)
    aux_json_data = read_dicom_json()


    database_generated['dicomID'] = aux_json_data['dicomID']
    database_generated['age'] = aux_json_data['age']
    database_generated['race'] = aux_json_data['race']
    database_generated['height'] = aux_json_data['height']
    database_generated['weight'] = aux_json_data['weight']

    database_generated.to_csv('patients_dicom.csv',sep=";")






    
