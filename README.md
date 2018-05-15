# AfterEther cryptocurrency project

AfterEther is a project to scale Ethereum by the method of blockchain clustering which consist of creating many Ethereum networks working in parallel.

http://afterether.org

## Building instructions


To build 

1. Download these 2 files:

	Ethereum geth project : https://github.com/ethereum/go-ethereum/archive/v1.7.2.tar.gz

	AfterEther patch : https://github.com/afterether/go-afterether-patch/releases/download/v1.7.2/go-afterether-v1.7.2.diff

2. Uncompress Ethereum:

	tar zxf go-afterether-v1.7.2-bin.tar.gz
	cd go-ethereum

3. Apply AfterEther patch:

	patch -p1 < go-afterether-v1.7.2.diff

4. Build 

	make all

	If you are compiling with Go v1.10 , change the line #180 in build/ci.go to this:

		`if runtime.Version() < "go1.10.0" && !strings.Contains(runtime.Version(), "devel") {`

	Otherwise you will get an error of incompatible Go version

6. Create an account

	build/bin/geth account new

7. Start `geth`
	
	build/bin/geth console

8. Verify that the genesis block is correct

	eth.getBlock(0)
	
	The hash of the block must be 0x13771466fee63d1916818f25c5e2ed5e09330b9fccc8da760c172d09412680d4

9. Mine

	\>eth.coinbase=[YOUR ACCOUNT ADDRESS]

	\>miner.start(1)




	
