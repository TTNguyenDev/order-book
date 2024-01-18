startMars:
	ignite chain serve -c mars.yml
startVenus:
	ignite chain serve -c venus.yml
relayerCfg:
	ignite relayer configure -a \
		--source-rpc "http://0.0.0.0:26657" \
		--source-faucet "http://0.0.0.0:4500" \
		--source-port "dex" \
		--source-version "dex-1" \
		--source-gasprice "0.0000025stake" \
		--source-prefix "cosmos" \
		--source-gaslimit 300000 \
		--target-rpc "http://0.0.0.0:26659" \
		--target-faucet "http://0.0.0.0:4501" \
		--target-port "dex" \
		--target-version "dex-1" \
		--target-gasprice "0.0000025stake" \
		--target-prefix "cosmos" \
		--target-gaslimit 300000
relayerConnect:
	ignite relayer connect
