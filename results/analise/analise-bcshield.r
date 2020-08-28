
####-- Hardware experimentos --####
hardware <- read.csv("/Users/erjulioaguiar/Documents/privacy-blockchain/results/table-results/reports-perform-bc.csv", sep=";")


#####-------- Calcular CPU por organização - Create #################
cpu_peer0_hprovider <- subset(hardware , (Name == 'peer0.hprovider.healthcare.com' & Query == 'create'), select=c(CPU..avg.,TPS_Rate))
cpu_peer1_hprovider <- subset(hardware , (Name == 'peer1.hprovider.healthcare.com' & Query == 'create'), select=c(CPU..avg.,TPS_Rate))
cpu_peer0_research <- subset(hardware , (Name == 'peer0.research.healthcare.com' & Query == 'create'), select=c(CPU..avg.,TPS_Rate))
cpu_peer1_research <- subset(hardware , (Name == 'peer1.research.healthcare.com' & Query == 'create'), select=c(CPU..avg.,TPS_Rate))
cpu_peer0_patient <- subset(hardware , (Name == 'peer0.patient.healthcare.com' & Query == 'create'), select=c(CPU..avg.,TPS_Rate))
cpu_peer1_patient <- subset(hardware , (Name == 'peer1.patient.healthcare.com' & Query == 'create'), select=c(CPU..avg.,TPS_Rate))


cpu_hprovider <- c(cpu_peer0_hprovider$CPU..avg.,cpu_peer1_hprovider$CPU..avg.)
cpu_research <- c(cpu_peer0_research$CPU..avg.,cpu_peer1_research$CPU..avg.)
cpu_patient <- c(cpu_peer0_patient$CPU..avg.,cpu_peer1_patient$CPU..avg.)
cpu_tps_hprovider <-c(cpu_peer0_hprovider$TPS_Rate, cpu_peer1_hprovider$TPS_Rate)
cpu_tps_researcher <-c(cpu_peer0_research$TPS_Rate, cpu_peer1_research$TPS_Rate)
cpu_tps_patient <-c(cpu_peer0_patient$TPS_Rate, cpu_peer1_patient$TPS_Rate)

cpu <- c(cpu_hprovider, cpu_research, cpu_patient)
tps_cpu <- c(cpu_tps_hprovider,cpu_tps_researcher,cpu_tps_patient)
label_cpu <-c(rep("Provedor-saude", length(cpu_hprovider)), 
              rep("Pesquisador",length(cpu_research)), rep("Paciente", length(cpu_patient)))

data_cpu <- data.frame(label_cpu, cpu, tps_cpu)

graph_cpu <- ggplot(data_cpu, aes(fill=label_cpu, y=cpu, x=tps_cpu)) + 
        geom_bar(position="dodge", stat="identity") +
        guides(fill=guide_legend(title="Organização")) +
        xlab("Transações/s (TPS)") + ylab("Uso de CPU (%)") +
        scale_fill_manual(values = c("#0C1D40", "#5E88BF", "#F2AE2E"))

png("/Users/erjulioaguiar/Documents/privacy-blockchain/results/analise/cpu_uso_criar_pt.png")
print(graph_cpu)
dev.off()

label_cpu <-c(rep("H-Provider", length(cpu_hprovider)), 
              rep("Researcher",length(cpu_research)), rep("Patient", length(cpu_patient)))

data_cpu <- data.frame(label_cpu, cpu, tps_cpu)

graph_cpu <- ggplot(data_cpu, aes(fill=label_cpu, y=cpu, x=tps_cpu)) + 
        geom_bar(position="dodge", stat="identity") +
        guides(fill=guide_legend(title="Orgs")) +
        xlab("Transaction per sec (TPS)") + ylab("CPU usage (%)") +
        scale_fill_manual(values = c("#0C1D40", "#5E88BF", "#F2AE2E"))

png("/Users/erjulioaguiar/Documents/privacy-blockchain/results/analise/cpu_uso_criar_en.png")
print(graph_cpu)
dev.off()


#####-------- Calcular CPU por organização - Get #################
cpu_peer0_hprovider <- subset(hardware , (Name == 'peer0.hprovider.healthcare.com' & Query == 'query'), select=c(CPU..avg.,TPS_Rate))
cpu_peer1_hprovider <- subset(hardware , (Name == 'peer1.hprovider.healthcare.com' & Query == 'query'), select=c(CPU..avg.,TPS_Rate))
cpu_peer0_research <- subset(hardware , (Name == 'peer0.research.healthcare.com' & Query == 'query'), select=c(CPU..avg.,TPS_Rate))
cpu_peer1_research <- subset(hardware , (Name == 'peer1.research.healthcare.com' & Query == 'query'), select=c(CPU..avg.,TPS_Rate))
cpu_peer0_patient <- subset(hardware , (Name == 'peer0.patient.healthcare.com' & Query == 'query'), select=c(CPU..avg.,TPS_Rate))
cpu_peer1_patient <- subset(hardware , (Name == 'peer1.patient.healthcare.com' & Query == 'query'), select=c(CPU..avg.,TPS_Rate))


cpu_hprovider <- c(cpu_peer0_hprovider$CPU..avg.,cpu_peer1_hprovider$CPU..avg.)
cpu_research <- c(cpu_peer0_research$CPU..avg.,cpu_peer1_research$CPU..avg.)
cpu_patient <- c(cpu_peer0_patient$CPU..avg.,cpu_peer1_patient$CPU..avg.)
cpu_tps_hprovider <-c(cpu_peer0_hprovider$TPS_Rate, cpu_peer1_hprovider$TPS_Rate)
cpu_tps_researcher <-c(cpu_peer0_research$TPS_Rate, cpu_peer1_research$TPS_Rate)
cpu_tps_patient <-c(cpu_peer0_patient$TPS_Rate, cpu_peer1_patient$TPS_Rate)

cpu <- c(cpu_hprovider, cpu_research, cpu_patient)
tps_cpu <- c(cpu_tps_hprovider,cpu_tps_researcher,cpu_tps_patient)
label_cpu <-c(rep("Provedor-saude", length(cpu_hprovider)), 
              rep("Pesquisador",length(cpu_research)), rep("Paciente", length(cpu_patient)))

data_cpu <- data.frame(label_cpu, cpu, tps_cpu)

graph_cpu <- ggplot(data_cpu, aes(fill=label_cpu, y=cpu, x=tps_cpu)) + 
        geom_bar(position="dodge", stat="identity") +
        guides(fill=guide_legend(title="Organização")) +
        xlab("Transações/s (TPS)") + ylab("Uso de CPU (%)") +
        scale_fill_manual(values = c("#0C1D40", "#5E88BF", "#F2AE2E"))

png("/Users/erjulioaguiar/Documents/privacy-blockchain/results/analise/cpu_uso_get_pt.png")
print(graph_cpu)
dev.off()

data_cpu <- data.frame(label_cpu, cpu, tps_cpu)
label_cpu <-c(rep("H-Provider", length(cpu_hprovider)), 
              rep("Researcher",length(cpu_research)), rep("Patient", length(cpu_patient)))
graph_cpu <- ggplot(data_cpu, aes(fill=label_cpu, y=cpu, x=tps_cpu)) + 
        geom_bar(position="dodge", stat="identity") +
        guides(fill=guide_legend(title="Orgs")) +
        labs(title = "CPU usage on blockchain - Get") +
        xlab("Transaction per sec (TPS)") + ylab("CPU usage (%)") +
        scale_fill_manual(values = c("#0C1D40", "#5E88BF", "#F2AE2E"))

png("/Users/erjulioaguiar/Documents/privacy-blockchain/results/analise/cpu_uso_get_en.png")
print(graph_cpu)
dev.off()

#### Memoria #######

#####-------- Calcular Memoria por organização - Create ################
mem_peer0_hprovider <- subset(hardware , (Name == 'peer0.hprovider.healthcare.com' & Query == 'create'), select=c(Memory.avg...MB.,TPS_Rate))
mem_peer1_hprovider <- subset(hardware , (Name == 'peer1.hprovider.healthcare.com' & Query == 'create'), select=c(Memory.avg...MB.,TPS_Rate))
mem_peer0_research <- subset(hardware , (Name == 'peer0.research.healthcare.com' & Query == 'create'), select=c(Memory.avg...MB.,TPS_Rate))
mem_peer1_research <- subset(hardware , (Name == 'peer1.research.healthcare.com' & Query == 'create'), select=c(Memory.avg...MB.,TPS_Rate))
mem_peer0_patient <- subset(hardware , (Name == 'peer0.patient.healthcare.com' & Query == 'create'), select=c(Memory.avg...MB.,TPS_Rate))
mem_peer1_patient <- subset(hardware , (Name == 'peer1.patient.healthcare.com' & Query == 'create'), select=c(Memory.avg...MB.,TPS_Rate))


mem_hprovider <- c(mem_peer0_hprovider$Memory.avg...MB.,mem_peer1_hprovider$Memory.avg...MB.)
mem_research <- c(mem_peer0_research$Memory.avg...MB.,mem_peer1_research$Memory.avg...MB.)
mem_patient <- c(mem_peer0_patient$Memory.avg...MB.,mem_peer1_patient$Memory.avg...MB.)
mem_tps_hprovider <-c(mem_peer0_hprovider$TPS_Rate, mem_peer1_hprovider$TPS_Rate)
mem_tps_researcher <-c(mem_peer0_research$TPS_Rate, mem_peer1_research$TPS_Rate)
mem_tps_patient <-c(mem_peer0_patient$TPS_Rate, mem_peer1_patient$TPS_Rate)

mem <- c(mem_hprovider, mem_research, mem_patient)
tps_mem <- c(mem_tps_hprovider,mem_tps_researcher,mem_tps_patient)
label_mem <-c(rep("Provedor-saude", length(mem_hprovider)), 
              rep("Pesquisador",length(mem_research)), rep("Paciente", length(mem_patient)))

data_mem <- data.frame(label_mem, mem, tps_mem)

graph_mem <- ggplot(data_mem, aes(fill=label_mem, y=mem, x=tps_mem)) + 
        geom_bar(position="dodge", stat="identity") +
        guides(fill=guide_legend(title="Organização")) +
        xlab("Transações/s (TPS)") + ylab("Uso de Memória (MB)") +
        scale_fill_manual(values = c("#0C1D40", "#5E88BF", "#F2AE2E"))

png("/Users/erjulioaguiar/Documents/privacy-blockchain/results/analise/mem_uso_criar_pt.png")
print(graph_mem)
dev.off()

label_mem <-c(rep("H-Provider", length(mem_hprovider)), 
              rep("Researcher",length(mem_research)), rep("Patient", length(mem_patient)))

data_mem <- data.frame(label_mem, mem, tps_mem)

graph_mem <- ggplot(data_mem, aes(fill=label_mem, y=mem, x=tps_mem)) + 
        geom_bar(position="dodge", stat="identity") +
        guides(fill=guide_legend(title="Orgs")) +
        xlab("Transaction per sec (TPS)") + ylab("Memory usage (MB)") +
        scale_fill_manual(values = c("#0C1D40", "#5E88BF", "#F2AE2E"))

png("/Users/erjulioaguiar/Documents/privacy-blockchain/results/analise/mem_uso_criar_en.png")
print(graph_mem)
dev.off()


#####-------- Calcular Memoria por organização - Get #################
mem_peer0_hprovider <- subset(hardware , (Name == 'peer0.hprovider.healthcare.com' & Query == 'query'), select=c(Memory.avg...MB.,TPS_Rate))
mem_peer1_hprovider <- subset(hardware , (Name == 'peer1.hprovider.healthcare.com' & Query == 'query'), select=c(Memory.avg...MB.,TPS_Rate))
mem_peer0_research <- subset(hardware , (Name == 'peer0.research.healthcare.com' & Query == 'query'), select=c(Memory.avg...MB.,TPS_Rate))
mem_peer1_research <- subset(hardware , (Name == 'peer1.research.healthcare.com' & Query == 'query'), select=c(Memory.avg...MB.,TPS_Rate))
mem_peer0_patient <- subset(hardware , (Name == 'peer0.patient.healthcare.com' & Query == 'query'), select=c(Memory.avg...MB.,TPS_Rate))
mem_peer1_patient <- subset(hardware , (Name == 'peer1.patient.healthcare.com' & Query == 'query'), select=c(Memory.avg...MB.,TPS_Rate))


mem_hprovider <- c(mem_peer0_hprovider$Memory.avg...MB.,mem_peer1_hprovider$Memory.avg...MB.)
mem_research <- c(mem_peer0_research$Memory.avg...MB.,mem_peer1_research$Memory.avg...MB.)
mem_patient <- c(mem_peer0_patient$Memory.avg...MB.,mem_peer1_patient$Memory.avg...MB.)
mem_tps_hprovider <-c(mem_peer0_hprovider$TPS_Rate, mem_peer1_hprovider$TPS_Rate)
mem_tps_researcher <-c(mem_peer0_research$TPS_Rate, mem_peer1_research$TPS_Rate)
mem_tps_patient <-c(mem_peer0_patient$TPS_Rate, mem_peer1_patient$TPS_Rate)

mem <- c(mem_hprovider, mem_research, mem_patient)
tps_mem <- c(mem_tps_hprovider,mem_tps_researcher,mem_tps_patient)
label_mem <-c(rep("Provedor-saude", length(mem_hprovider)), 
              rep("Pesquisador",length(mem_research)), rep("Paciente", length(mem_patient)))

data_mem <- data.frame(label_mem, mem, tps_mem)

graph_mem <- ggplot(data_mem, aes(fill=label_mem, y=mem, x=tps_mem)) + 
        geom_bar(position="dodge", stat="identity") +
        guides(fill=guide_legend(title="Organização")) +
        xlab("Transações/s (TPS)") + ylab("Uso de memória (MB)") +
        scale_fill_manual(values = c("#0C1D40", "#5E88BF", "#F2AE2E"))

png("/Users/erjulioaguiar/Documents/privacy-blockchain/results/analise/mem_uso_get_pt.png")
print(graph_mem)
dev.off()

label_mem <-c(rep("H-Provider", length(mem_hprovider)), 
              rep("Researcher",length(mem_research)), rep("Patient", length(mem_patient)))

data_mem <- data.frame(label_mem, mem, tps_mem)

graph_mem <- ggplot(data_mem, aes(fill=label_mem, y=mem, x=tps_mem)) + 
        geom_bar(position="dodge", stat="identity") +
        guides(fill=guide_legend(title="Orgs")) +
        xlab("Transaction per sec (TPS)") + ylab("Memory usage (MB)") +
        scale_fill_manual(values = c("#0C1D40", "#5E88BF", "#F2AE2E"))

png("/Users/erjulioaguiar/Documents/privacy-blockchain/results/analise/mem_uso_get_en.png")
print(graph_mem)
dev.off()


### Rede #####

network <- read.csv("/Users/erjulioaguiar/Documents/privacy-blockchain/results/table-results/reports-network-bc.csv", sep=";")

#####-------- Latencia blockchain #################
lat_add <- subset(network , (Query == 'create'), select=c(Avg.Latency..s.,TPS_Rate))
lat_consulta <- subset(network , (Query == 'query'), select=c(Avg.Latency..s.,TPS_Rate))

lat_avg <- c(lat_add$Avg.Latency..s., lat_consulta$Avg.Latency..s.)
lat_tps <- c(lat_add$TPS_Rate, lat_consulta$TPS_Rate)

lat_label <- c(rep("Adição", length(lat_add$Avg.Latency..s)), 
               rep("Consulta", length(lat_consulta$Avg.Latency..s.)))

data_lat <- data.frame(lat_label, lat_avg, lat_tps)

graph_lat <- ggplot(data_lat, aes(fill=lat_label, y=lat_avg, x=lat_tps)) + 
        geom_bar(position="dodge", stat="identity") +
        guides(fill=guide_legend(title="Operação")) +
        xlab("Transações/s (TPS)") + ylab("Latência (s)") +
        scale_fill_manual(values = c("#0099c6", "#f4b401"))

png("/Users/erjulioaguiar/Documents/privacy-blockchain/results/analise/latencia_pt.png")
print(graph_lat)
dev.off()

lat_label <- c(rep("Add", length(lat_add$Avg.Latency..s)), 
               rep("Get", length(lat_consulta$Avg.Latency..s.)))

data_lat <- data.frame(lat_label, lat_avg, lat_tps)

graph_lat <- ggplot(data_lat, aes(fill=lat_label, y=lat_avg, x=lat_tps)) + 
        geom_bar(position="dodge", stat="identity") +
        guides(fill=guide_legend(title="Operação")) +
        xlab("Transactions per sec (TPS)") + ylab("Latency (s)") +
        scale_fill_manual(values = c("#0099c6", "#f4b401"))

png("/Users/erjulioaguiar/Documents/privacy-blockchain/results/analise/latencia_en.png")
print(graph_lat)
dev.off()

#####-------- throughput blockchain #################
vazao_add <- subset(network , (Query == 'create'), select=c(Throughput..TPS.,TPS_Rate))
vazao_consulta <- subset(network , (Query == 'query'), select=c(Throughput..TPS.,TPS_Rate))

vazao_avg <- c(vazao_add$Throughput..TPS., vazao_consulta$Throughput..TPS.)
vazao_tps <- c(vazao_add$TPS_Rate, vazao_consulta$TPS_Rate)

vazao_label <- c(rep("Adição", length(vazao_add$Throughput..TPS.)), 
               rep("Consulta", length(vazao_consulta$Throughput..TPS.)))

data_vazao <- data.frame(vazao_label, vazao_avg, vazao_tps)

graph_vazao <- ggplot(data_vazao, aes(fill=vazao_label, y=vazao_avg, x=vazao_tps)) + 
        geom_bar(position="dodge", stat="identity") +
        guides(fill=guide_legend(title="Operação")) +
        xlab("Transações/s (TPS)") + ylab("Vazão (KB/s)") +
        scale_fill_manual(values = c("#0099c6", "#f4b401"))

png("/Users/erjulioaguiar/Documents/privacy-blockchain/results/analise/throughput_pt.png")
print(graph_vazao)
dev.off()


vazao_label <- c(rep("Add", length(vazao_add$Throughput..TPS.)), 
                 rep("Get", length(vazao_consulta$Throughput..TPS.)))

data_vazao <- data.frame(vazao_label, vazao_avg, vazao_tps)

graph_vazao <- ggplot(data_vazao, aes(fill=vazao_label, y=vazao_avg, x=vazao_tps)) + 
        geom_bar(position="dodge", stat="identity") +
        guides(fill=guide_legend(title="Operation")) +
        xlab("Transactions per sec (TPS)") + ylab("Throughput (KB/s)") +
        scale_fill_manual(values = c("#0099c6", "#f4b401"))

png("/Users/erjulioaguiar/Documents/privacy-blockchain/results/analise/throughput_en.png")
print(graph_vazao)
dev.off()

### Metricas Jmeter ######

get_simple <- read.csv("/Users/erjulioaguiar/Documents/privacy-blockchain/results/Jmeter-results/new/tabela-resultados-simple.csv", sep=",")
get_kano <- read.csv("/Users/erjulioaguiar/Documents/privacy-blockchain/results/Jmeter-results/new/tabela-resultados-kano.csv", sep=",")
get_diff <- read.csv("/Users/erjulioaguiar/Documents/privacy-blockchain/results/Jmeter-results/new/tabela-resultados-diffpriv.csv", sep=",")

simple_aggr <- read.csv("/Users/erjulioaguiar/Documents/privacy-blockchain/results/Jmeter-results/new/gafico-aggr-kano.csv", sep=";")

get_simple <- subset(get_simple, (responseCode == 200))

time_simple <- as.POSIXct(get_simple$timeStamp, origin="1970-01-01")
time_simple <- format.Date(time_simple, "%H:%M:%S")
numeric_time <- (as.numeric(as.POSIXct(paste("2014-01-01", time_simple))) - 
              as.numeric(as.POSIXct("2014-01-01 0:0:0")))/60
lat_simple <- get_simple$Latency/60


lat_med_simple <- mean(get_simple$Latency)
tr_med_simple <- simple_aggr$Vazão[0:1]
label_graph <- c("Latência", "Throughput")

dt <- data.frame(lat_med_simple, tr_med_simple, label_graph)

ggplot(data = dt, aes(x = label_graph, y = lat_med_simple, group = tr_med_simple)) +
        geom_line(aes(color = Party, alpha = 1), size = 2) +
        geom_point(aes(color = Party, alpha = 1), size = 4) +
        # move the x axis labels up top
        scale_x_discrete(position = "top") +
        theme_bw() +
        # Format tweaks
        # Remove the legend
        theme(legend.position = "none") +
        # Remove the panel border
        theme(panel.border     = element_blank()) +
        #  Labelling as desired
        labs(
                title = "Voter's stated preferences for June 7 elections in Ontario",
                subtitle = "(Mainstreet Research)"
        )

#### Tempo e anonimização ####

tempo_diff <- read.delim2("/Users/erjulioaguiar/Documents/privacy-blockchain/results/privacy-experiments/resultados/time-priv-diff.txt", sep = "\n")
t_diff<- gsub('[a-zA-Z]', '', tempo_diff$tempo)
t_diff <- as.double(t_diff)

tempo_kano <- read.delim2("/Users/erjulioaguiar/Documents/privacy-blockchain/results/privacy-experiments/resultados/time-priv-kanon.txt", sep = "\n")
t_kano<- gsub('[a-zA-Z]', '', tempo_kano$tempo)
t_kano <- as.double(t_kano)

#### Entropia #####

entropia_diff <- read.delim2("/Users/erjulioaguiar/Documents/privacy-blockchain/results/entropy-values/py-diff-entropia.txt", sep="\n") 
entr_diff = read.delim2("/Users/erjulioaguiar/Documents/privacy-blockchain/results/analise/py-diff-entropia-range.txt", sep="\n")
entr_kano = read.delim2("/Users/erjulioaguiar/Documents/privacy-blockchain/results/analise/entropy-kano.txt", sep="\n")
epsilon <- c(0.0001, 0.001, 0.01, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0)
dist_diff <- read.delim2("/Users/erjulioaguiar/Documents/privacy-blockchain/results/entropy-values/py-diff-distancia.txt", sep="\n")


## Line graph for Entropy for entropy and distance 

priv_entropy_norm <- normalize(as.double(entropia_diff$entropy))
priv_dist_norm <- normalize(as.double(dist_diff$dist))

data_line <- c(priv_entropy_norm, priv_dist_norm)

label <- c(rep("Entropia", (length(data_line)/2)), 
           rep("Utilidade", length(data_line)/2))

ep <- c(epsilon,epsilon)

df_priv <- data.frame("quant"=data_line, "epsilon"= ep, "Métricas"=label)

png("/Users/erjulioaguiar/Documents/privacy-blockchain/results/analise/privacidade/epsilon_vs_entropy.png")
ggplot(data=df_priv, aes(x=epsilon, y=quant, group=label, colour=Métricas)) +
        geom_line(size=1.2) +
        geom_point() +
        xlab("Epsilon") + ylab("Quantidade") +
        #scale_colour_manual(name = "Métricas", values = c("#2CCBBE", "#606F9E")) +
        guides(group=guide_legend(title="Operation")) +
        theme_classic()
dev.off()

# Boxplot da entropia entre os métodos
entr_diff_num <- as.numeric(entropia_diff$entropy)
entr_diff_norm <- normalize(entr_diff_num[entr_diff_num >= 0])

entr_kano <- as.numeric(entr_kano$entropy)
entr_kano_norm <- normalize(entr_kano)

entropy <- c(entr_kano_norm, entr_diff_norm)
entr <- normalize(entropy)

label_entropy <- c(rep("K-Anomimato", length(entr_kano_norm)), 
               rep("Privacidade Diferencial", length(entr_diff_norm)))

data_entr <- data.frame("entropy"=entropy, "label"=label_entropy)

png("/Users/erjulioaguiar/Documents/privacy-blockchain/results/analise/privacidade/boxplot_entropia_final.png")
ggplot(data_entr, aes(x=label, y=entropy)) + 
        geom_boxplot() +
        theme_classic() +
        xlab("Modelo de privacidade") + ylab("Entropia - H(X)")
dev.off()



data <- c(entr_diff_conv$entropy,entr_kano_norm)
label <- c(rep("Privacidade Diferencial", length(entr_diff_conv$entropy)), 
           rep("K-Anonimato", length(entr_kano_norm)))

df_entr <- data_frame(data, label)

png("/Users/erjulioaguiar/Documents/privacy-blockchain/results/analise/privacidade/boxplot_entropia.png")
ggplot(df_entr, es(x=label, y=data)) + 
        geom_boxplot() +
        theme_classic() +
        xlab("Modelo de privacidade") + ylab("Entropia - H(X)")
dev.off()

d <- data.frame(episilon$epsilon, entr_diff_num)

gpathEntropyDiff <- ggplot(d,aes(y=entr_diff_norm, x=episilon.epsilon)) + 
        geom_bar(position="dodge", stat="identity") +
        guides(fill=guide_legend(title="Operation")) +
        labs(title = "Throughput on blockchain") +
        xlab("Transactions per sec (TPS)") + ylab("Throughput (KB/s)") +
        scale_fill_manual(values = c("#0099c6", "#f4b401"))
print(gpathEntropyDiff)


### Similaridade ###

episilon <- read.delim2("/Users/erjulioaguiar/Documents/privacy-blockchain/results/py-episilon.txt",sep = "\n")
dist_diff_str <- read.delim2("/Users/erjulioaguiar/Documents/privacy-blockchain/results/entropy-values/py-diff-dist-final.txt", sep="\n")
dist_kano_str <- read.delim2("/Users/erjulioaguiar/Documents/privacy-blockchain/results/entropy-values/kanonimity-similarity.txt", sep="\n")


dist_diff <- as.double(dist_diff_str$dist)
dist_kano <- as.double(dist_kano_str$dist)
dist_diff_norm <- normalize(dist_diff)
dist_kano_norm <- normalize(dist_kano)

data_dist <- c(dist_kano ,dist_diff)
data_dist <- normalize(data_dist)

label_dist <- c(rep("K-Anonimato", length(dist_kano)),
                rep("Privacidade Diferencial", length(dist_diff)))

dt <- data.frame("distancia"=data_dist, "label"=label_dist)

png("/Users/erjulioaguiar/Documents/privacy-blockchain/results/analise/privacidade/boxplot_distancia.png")
ggplot(dt, aes(x=label, y=distancia)) + 
        geom_boxplot() +
        theme_classic() +
        xlab("Modelo de privacidade") + ylab("Distância") +
        scale_fill_manual(values = c("#0099c6", "#f4b401"))
dev.off()


#### Funcoes de densidade para similaridade ####

data_curves <- c(dist_kano, dist_diff)
label_curves <- c(rep("K-Anonimato", length(dist_kano)), 
                  rep("Privacidade diferencial", length(dist_diff)))

d <- data.frame("sim"=data_curves, "label"=label_curves)

ci_kano <- ci(dist_kano, ci = 0.95)
ci_diff <- ci(dist_diff, ci = 0.95)

# PDF
png("/Users/erjulioaguiar/Documents/privacy-blockchain/results/analise/privacidade/pdf_similaridade.png")
ggplot(data=d, aes(x=sim, group=label, fill=label)) +
        geom_density(alpha=0.4, size = 1) + 
        xlab("Distância") + ylab("Densidade") +
        guides(fill=guide_legend(title="Modelos")) +
        geom_vline(xintercept=ci_kano$CI_low, color="red", size=0.8, linetype = "dashed") +
        geom_vline(xintercept=ci_kano$CI_high, color="red", size=0.8, linetype = "dashed") +
        geom_vline(xintercept=ci_diff$CI_low, color="blue", size=0.8, linetype = "dashed") +
        geom_vline(xintercept=ci_diff$CI_high, color="blue", size=0.8, linetype = "dashed") +
        theme_classic() +
        theme(legend.position = c(0.75, 0.8))
dev.off()

# -------------------------- CDF Similaridade K-Anon -----------------------------------

normal_dist <- rnorm(dist_kano)
chi_square <- rchisq(dist_kano, df=2)
gamm_dist <- rgamma(dist_kano, shape = 2)
exp_dist <- rexp(dist_kano)
tstu_dist <- rt(dist_kano, df=2)
lognor_dsit <- rlnorm(dist_kano)


#Comparaçao entre as curvas - teste KS
ks.test(dist_kano, gamm_dist)
ks.test(dist_kano, normal_dist)
ks.test(dist_kano, chi_square)
ks.test(dist_kano, exp_dist)
ks.test(dist_kano, tstu_dist)
ks.test(dist_kano, lognor_dsit)

data_array_kano <- c(dist_kano, normal_dist, exp_dist, tstu_dist)

label_edf <- c(rep("K-Anonimato", length(dist_kano)),
               rep("Normal", length(normal_dist)), 
               rep("Exponencial", length(lognor_dsit)),
               rep("T-Student", length(tstu_dist))) 


data_edf_kano <- data.frame("similaridade"=data_array_kano,"Funções"=label_edf)

png("/Users/erjulioaguiar/Documents/privacy-blockchain/results/analise/privacidade/cdf_similaridade_kanon.png")
ggplot(data=data_edf_kano, aes(x=similaridade, group=Funções, fill=Funções)) +
        stat_ecdf(aes(color = Funções, linetype = Funções), geom = "step", size = 1) + 
        xlab("Distância") + ylab("Probabilidade acumulada") +
        theme_classic() +
        theme(legend.position = c(0.8, 0.3))
dev.off()


# -------------------------- CDF Similaridade Diff Priv -----------------------------------

normal_dist <- rnorm(dist_diff)
chi_square <- rchisq(dist_diff, df=2)
gamm_dist <- rgamma(dist_diff, shape = 2)
exp_dist <- rexp(dist_diff)
tstu_dist <- rt(dist_diff, df=2)
lognor_dsit <- rlnorm(dist_diff)

#Comparaçao entre as curvas - teste KS
ks.test(dist_diff, gamm_dist)
ks.test(dist_diff, normal_dist)
ks.test(dist_diff, chi_square)
ks.test(dist_diff, exp_dist)
ks.test(dist_diff, tstu_dist)
ks.test(dist_diff, lognor_dsit)

data_array_diff <- c(dist_diff, exp_dist, chi_square, normal_dist)

label_edf <- c(rep("Priv diferencial", length(dist_diff)),
               rep("Exponencial", length(exp_dist)),
               rep("Chi-Quadrada", length(chi_square)),
               rep("Normal", length(normal_dist))) 


data_edf_kano <- data.frame("similaridade"=data_array_diff,"Funções"=label_edf)

png("/Users/erjulioaguiar/Documents/privacy-blockchain/results/analise/privacidade/cdf_similaridade_diffpriv.png")
ggplot(data=data_edf_kano, aes(x=similaridade, group=Funções, fill=Funções)) +
        stat_ecdf(aes(color = Funções, linetype = Funções), geom = "step", size = 1) + 
        xlab("Distância") + ylab("Probabilidade acumulada") +
        theme_classic() +
        theme(legend.position = c(0.8, 0.3))
dev.off()


#### Funcoes de densidade para Entropia ####

data_curves <- c(entr_kano_norm, entr_diff_norm)
label_curves <- c(rep("K-Anonimato", length(entr_kano_norm)),
                  rep("Privacidade Diferencial", length(entr_diff_norm)))

d <- data.frame("sim"=data_curves, "label"=label_curves)

ci_kano <- ci(entr_kano_norm, ci = 0.95)
ci_diff <- ci(entr_diff_norm, ci = 0.95)

# PDF
png("/Users/erjulioaguiar/Documents/privacy-blockchain/results/analise/privacidade/pdf_entropia.png")
ggplot(data=d, aes(x=sim, group=label, fill=label)) +
        geom_density(alpha=0.4, size = 1) + 
        xlab("Entropia - H(X)") + ylab("Densidade") +
        guides(fill=guide_legend(title="Modelos")) +
        geom_vline(xintercept=ci_kano$CI_low, color="red", size=0.8, linetype = "dashed") +
        geom_vline(xintercept=ci_kano$CI_high, color="red", size=0.8, linetype = "dashed") +
        geom_vline(xintercept=ci_diff$CI_low, color="blue", size=0.8, linetype = "dashed") +
        geom_vline(xintercept=ci_diff$CI_high, color="blue", size=0.8, linetype = "dashed") +
        theme_classic() +
        theme(legend.position = c(0.2, 0.8))
dev.off()


# -------------------------- CDF Entropia K-Anon -----------------------------------

normal_entr <- rnorm(entr_kano_norm)
chi_square <- rchisq(entr_kano_norm, df=2)
gamm_entr <- rgamma(entr_kano_norm, shape = 2)
exp_entr <- rexp(entr_kano_norm)
tstu_entr <- rt(entr_kano_norm, df=2)
lognor_entr <- rlnorm(entr_kano_norm)

# Teste K-S
ks.test(entr_kano_norm, gamm_entr)
ks.test(entr_kano_norm, normal_entr)
ks.test(entr_kano_norm, chi_square)
ks.test(entr_kano_norm, exp_entr)
ks.test(entr_kano_norm, lognor_entr)
ks.test(entr_kano_norm, tstu_entr)



data_array_kano <- c(entr_kano_norm, normal_entr, exp_entr, tstu_entr)

label_edf <- c(rep("K-Anonimato", length(entr_kano_norm)),
               rep("Normal", length(normal_entr)), 
               rep("Exponencial", length(exp_entr)),
               rep("T-Student", length(tstu_entr))) 


data_edf_kano <- data.frame("entropia"=data_array_kano,"Funções"=label_edf)

png("/Users/erjulioaguiar/Documents/privacy-blockchain/results/analise/privacidade/cdf_entropia_kanon.png")
ggplot(data=data_edf_kano, aes(x=entropia, group=Funções, fill=Funções)) +
        stat_ecdf(aes(color = Funções, linetype = Funções), geom = "step", size = 1) + 
        xlab("Entropia - H(X)") + ylab("Probabilidade acumulada") +
        theme_classic() +
        theme(legend.position = c(0.8, 0.3))
dev.off()


# -------------------------- CDF Entropia Diff Priv -----------------------------------

normal_entr <- rnorm(entr_diff_norm)
chi_square <- rchisq(entr_diff_norm, df=2)
gamm_entr <- rgamma(entr_diff_norm, shape = 2)
exp_entr <- rexp(entr_diff_norm)
tstu_entr <- rt(entr_diff_norm, df=2)
lognor_entr <- rlnorm(entr_diff_norm)


#Comparaçao entre as curvas Test-KS
ks.test(entr_diff_norm, gamm_entr)
ks.test(entr_diff_norm, normal_entr)
ks.test(entr_diff_norm, chi_square)
ks.test(entr_diff_norm, exp_entr)
ks.test(entr_diff_norm, lognor_entr)
ks.test(entr_diff_norm, tstu_entr)


data_array_diff <- c(entr_diff_norm, gamm_entr, exp_entr, lognor_entr)

label_edf <- c(rep("Priv Diferencial", length(entr_diff_norm)),
               rep("Gamma", length(gamm_entr)), 
               rep("Exponencial", length(exp_entr)),
               rep("Log Normal", length(lognor_entr))) 


data_edf_kano <- data.frame("entropia"=data_array_diff,"Funções"=label_edf)

png("/Users/erjulioaguiar/Documents/privacy-blockchain/results/analise/privacidade/cdf_entropia_diff.png")
ggplot(data=data_edf_kano, aes(x=entropia, group=Funções, fill=Funções)) +
        stat_ecdf(aes(color = Funções, linetype = Funções), geom = "step", size = 1) + 
        xlab("Entropia - H(X)") + ylab("Probabilidade acumulada") +
        theme_classic() +
        theme(legend.position = c(0.8, 0.3))
dev.off()

