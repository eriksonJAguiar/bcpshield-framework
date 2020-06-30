
cpu = read.csv('/Users/erjulioaguiar/Documents/privacy-blockchain/results/table-results/round4-write/cpu_hprovider.csv', sep = ";")

plot(cpu$cpu, cpu$Time, type="l")

