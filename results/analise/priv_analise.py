import json
import base64
import numpy as np
import distance
import struct
import statistics
from pyitlib import discrete_random_variable as drv
from scipy.spatial import distance as dist
from sklearn.metrics import mutual_info_score
from collections import Counter

def read_json(path):
    data = None
    with open(path) as json_file:
        data = json.load(json_file)

    return data


def str_to_binary(data):
    bin_str = ''.join(format(ord(x), 'b') for x in data)

    bin_array = [int(b) for b in bin_str]

    return bin_array

def float_to_bin(num):
    bin_str = format(struct.unpack('!I', struct.pack('!f', num))[0], '032b')

    bin_array = [int(b) for b in bin_str]

    return bin_array

def dict_to_list(data):
    dictlist = list()
    for _, value in data.items():
        dictlist.append(value)
    
    return dictlist

#def diff_to_list(data):


def entropy_shannon(np_data):
    px = np_data/np_data.sum()
    entr = -np.nansum(px*np.log2(px))
    
    return entr

# def mutual_information(np_data1, np_data2):
#     px = 

def bow_convert(dataset, value):
    dt_general = list()
    for d in dataset:
        attrs = [at for at in dict_to_list(d) if not at == ' ' and not at == '']
        dt_general += attrs
    
    counted = dict(Counter(dt_general))
    
    bow_array = list()
    for v in value:
        bow_array.append(counted[v])
    

    return np.array(bow_array)
        



def entropy_calculate(data, file_name):
    entr_final = list()

    for i in range(len(data)):
        attrs = [at for at in dict_to_list(data[i]) if not at == ' ' and not at == '']
        bow_data = bow_convert(data, attrs)
        entr = entropy_shannon(bow_data)

        entr_final.append(entr)

    with open("%s.txt"%(file_name), mode="a") as f:
        for e in entr_final:
            f.write("%f\n"%(e))

def mutual_info_diff(original, private, file_name):
    original_attrs = [at for at in dict_to_list(original) if not at == ' ' and not at == '']
    bow_orignal = bow_convert([original], original)
    for p in private:
        attrs = [at for at in dict_to_list(p) if not at == ' ' and not at == '']
        bow_data = bow_convert(private, attrs)


def entropy_calculate_diff(data, file_name):
    entr_final = list()

    for i in range(len(data)):
        attrs = [at for at in dict_to_list(data[i]) if not at == ' ' and not at == '']
        np_data = np.array(attrs)
        entr = entropy_shannon(np_data)

        entr_final.append(entr)
    
    with open("%s.txt"%(file_name), mode="a") as f:
            for e in entr_final:
                f.write("%f\n"%(e))

def similarity_calculate(original, private, file_name):
    kano_sim = list()
    for i in range(len(private)):
        orig = dict_to_list(original[i])
        k_an = [og for og in dict_to_list(
            private[i]) if not og == ' ' and not og == '']
        aux_dist = list()

        print(orig)
        print(k_an)

        for j in range(len(k_an)):
            bow_o = bow_convert(original, orig[j])
            bow_k = bow_convert(private, k_an[j])
            d = dist.cosine(bow_o, bow_k)
            aux_dist.append(d)

        kano_sim += aux_dist

    # with open('%s.txt'%(file_name), mode="a") as f:
    #     for s in kano_sim:
    #         f.write("%f\n" % s)

def similarity_calculate_diff(original, private, file_name):
    kano_sim = list()
    for i in range(len(private)):
        orig = dict_to_list(original)
        k_an = [og for og in dict_to_list(
            private[i]) if not og == ' ' and not og == '']
        
        aux_dist = list()
        for j in range(len(k_an)):
            bin_o = float_to_bin(orig[j])
            bin_k = float_to_bin(k_an[j])
            d = dist.cosine(bin_o, bin_k)
            aux_dist.append(d)
        
        kano_sim += aux_dist

    with open('%s.txt'%(file_name), mode="a") as f:
        for s in kano_sim:
            f.write("%f\n" % s)


def diff():
    #[0.0001, 0.001, 0.01, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0]
    dt = read_json("../priv-diff-values/py-diff-priv-final.json")
    entropy_calculate_diff(dt, "../entropy-values/py-diff-entropia-final")

def diff_dist():
    diff_original = read_json("../py-diff-orginal.json")
    dt = read_json("../priv-diff-values/py-diff-priv-final.json")
    similarity_calculate_diff(diff_original, dt, "../entropy-values/py-diff-dist-final")


if __name__ == "__main__":
    diff_dist()
    #original_kano = read_json(
    #  "../privacy-experiments/resultados/data-original-kanon.json")
    #kano = read_json("../privacy-experiments/resultados/data-priv-kanon.json")
    #bin_k_original = str_to_binary(json.dumps(original_kano[0]))
    #bin_kano = str_to_binary(json.dumps(kano[0]))
    
    #diff_original = read_json("../py-diff-orginal.json")
    #diff = read_json("../py-diff-priv-range.json")

    #similarity_calculate_diff(diff_original, diff, 'py-priv-diff')
    #entropy_calculate_diff(diff,"py-diff-entropia-range")
    #entropy_calculate(kano, 'k-anony-entr')
    #entropy_calculate(diff, "py-diff-entropia")
    #entropy_calculate_diff(diff, "py-diff-entr")
    #entropy_calculate(original_kano, 'k-anony-original-entr')
    #entropy_calculate(diff, 'diff-entr')
    #entropy_calculate(diff_original, 'diff-original-entr')
    #similarity_calculate(original_kano, kano, "../entropy-values/kanonimity-similarity")