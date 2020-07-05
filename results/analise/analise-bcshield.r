
tr_tps = read.csv('/Users/erjulioaguiar/Documents/privacy-blockchain/results/table-results/write-results/round-others/transactions_per_time_write2.csv', sep = ";")
lat_tps = read.csv('/Users/erjulioaguiar/Documents/privacy-blockchain/results/latency_avg_per_tps_read.csv', sep = ";")

plot(cpu_time$X0.0006, cpu_time$X0.5861, type="l")

barplot(lat_tps$TPS, lat_tps$Latency)



hist(tr_tps$TrAmount)