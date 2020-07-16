from bs4 import BeautifulSoup
import pandas as pd
import numpy as np
import os


with open("../caliper-bcshield/caliper-benchmarks/reports/Simple/report-5.html") as fp:
    soup = BeautifulSoup(fp)

# table_summary = soup.find('div', attrs={'id':'benchmarksummary'})

# #table_summary = soup.find_all('table', limit=1)

# data = []

# rows = table_summary.find_all('tr')
# for row in rows:
#     cols = row.find_all('td')
#     cols = [ele.text.strip() for ele in cols]
#     data.append([ele for ele in cols if ele])


# th_value = table_summary.find_all('th')
# labels = [ele.text.strip() for ele in th_value]
    

# data = data [1:]
# d = np.array(data)
# dt = pd.DataFrame(d, columns=labels)


# if not os.path.exists("./reports-network-simple.csv"):
#     dt.to_csv("./reports-network-simple.csv", mode="a", index=False, sep=";")
# else:
#     dt.to_csv("./reports-network-simple.csv", mode="a", index=False, header=False, sep=";")


def get_tables(id_tb):
    div = soup.find('div', attrs={'id':id_tb})

    data = []

    rows = div.find_all('tr')
    for row in rows:
        cols = row.find_all('td')
        cols = [ele.text.strip() for ele in cols]
        vals = [ele for ele in cols if ele]
        data.append(vals)


    #Remove empty values
    data = [d for d in data if len(d) > 0]
    network = data[0]
    resources = data[1:]

    return network, resources


def get_label(id_tb):
    
    div = soup.find('div', attrs={'id':id_tb})
    
    th_value = div.find_all('th')
    label = [ele.text.strip() for ele in th_value]

    return label

def save_values(id_table):
    net, perf = get_tables(id_table)
    label = get_label(id_table)
    label_net = label[:len(net)]
    label_perf = label[len(net):]

    perf = perf[1:]
    #dt_net = np.array(net)
    df_net = pd.DataFrame([net], columns=label_net)
    
    #dt_perf = np.array(perf)
    df_perf = pd.DataFrame(perf, columns=label_perf)
    tps_rate = int(id_table.split("-")[1])
    df_perf['TPS_Rate'] = [tps_rate for i in range(len(df_perf['Name']))]

    if not os.path.exists("./reports-network-simple.csv"):
        df_net.to_csv("./reports-network-simple.csv", mode="a", index=False, sep=";")
    else:
        df_net.to_csv("./reports-network-simple.csv", mode="a", index=False, header=False, sep=";")
    

    if not os.path.exists("./reports-perform-simple.csv"):
        df_perf.to_csv("./reports-perform-simple.csv", mode="a", index=False, sep=";")
    else:
        df_perf.to_csv("./reports-perform-simple.csv", mode="a", index=False, header=False, sep=";")



if __name__ == "__main__":
    values = ['50', '100', '150', '200', '250']

    for v in values:
        save_values('create-%s'%(v))
        save_values('query-%s'%(v))