

hardware <- read.csv("/Users/erjulioaguiar/Documents/privacy-blockchain/results/reports-perform-simple.csv", sep=";")


cpu_peer0_hprovider <- subset(hardware , (Name == 'peer0.hprovider.healthcare.com'), select=c(CPU..avg.,TPS_Rate))
cpu_peer1_hprovider <- subset(hardware , (Name == 'peer1.hprovider.healthcare.com'), select=c(CPU..avg.,TPS_Rate))
cpu_peer0_research <- subset(hardware , (Name == 'peer0.research.healthcare.com'), select=c(CPU..avg.,TPS_Rate))
cpu_peer1_research <- subset(hardware , (Name == 'peer1.research.healthcare.com'), select=c(CPU..avg.,TPS_Rate))
cpu_peer0_patient <- subset(hardware , (Name == 'peer0.patient.healthcare.com'), select=c(CPU..avg.,TPS_Rate))
cpu_peer1_patient <- subset(hardware , (Name == 'peer1.patient.healthcare.com'), select=c(CPU..avg.,TPS_Rate))

rates <- c(50, 100,150,200)
tps_rate <- rep(c(50, 100,150,200), each=3)
tps_rate_50 <- rep(c(50), each=20)
tps_rate_50 <- rep(c(50, 100,150,200), each=6)
tps_rate_50 <- rep(c(50, 100,150,200), each=6)

cpu_hprovider <- c(cpu_peer0_hprovider$CPU..avg.,cpu_peer1_hprovider$CPU..avg.)
cpu_research <- c(cpu_peer0_research$CPU..avg.,cpu_peer1_research$CPU..avg.)
cpu_patient <- c(cpu_peer0_patient$CPU..avg.,cpu_peer1_patient$CPU..avg.)
tps_values <-c(cpu_peer0_hprovider$TPS_Rate, cpu_peer1_hprovider$TPS_Rate)

avg_hprovider <- mean(cpu_hprovider)
avg_research <- mean(cpu_research)
avg_patient <- mean(cpu_patient)

cpu <- c(mean(cpu_peer0_hprovider$CPU..avg.), mean(cpu_peer1_hprovider$CPU..avg.), mean(cpu_peer0_research$CPU..avg.),
         mean(cpu_peer1_research$CPU..avg.),mean(cpu_peer0_patient$CPU..avg.),mean(cpu_peer1_patient$CPU..avg.))

barplot(cpu,space = 0.5, main = "Uso de CPU Blockchain", ylab = "CPU(%) Média", 
        xlab = "Organizações utilizadas", col = c("#0362fc", "#8db4fc"), 
        names.arg = c("SProvedor", "SProvedor", "Pesquisador", "Pesquisador", "Paciente", "Paciente"), 
        beside = TRUE, ylim = range(pretty(c(0, cpu+2))))


legend("topright",
       c("Par0","Par1"),
       fill = c("#0362fc", "#8db4fc")
)

cpu_tps_50 <- subset(cpu_peer0_hprovider, (TPS_Rate == 50), select = c(CPU..avg.,TPS_Rate))
#cpu_tps_50 <- mean(cpu_tps_50$CPU..avg.)

cpu_tps_100 <- subset(cpu_peer0_hprovider, (TPS_Rate == 100), select = c(CPU..avg.,TPS_Rate))
#cpu_tps_100 <- mean(cpu_tps_100$CPU..avg.)

cpu_tps_150 <- subset(cpu_peer0_hprovider, (TPS_Rate == 150), select = c(CPU..avg.,TPS_Rate))
#cpu_tps_150 <- mean(cpu_tps_150$CPU..avg.)

cpu_tps_200 <- subset(cpu_peer0_hprovider, (TPS_Rate == 200), select = c(CPU..avg.,TPS_Rate))
#cpu_tps_200 <- mean(cpu_tps_200$CPU..avg.)

cpu_tps_250 <- subset(cpu_peer0_hprovider, (TPS_Rate == 250), select = c(CPU..avg.,TPS_Rate))
#cpu_tps_250 <- mean(cpu_tps_250$CPU..avg.)
