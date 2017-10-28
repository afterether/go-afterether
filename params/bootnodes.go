// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	"enode://18414b19b5ec88f889760baa2d993f7ac14560de7e5ae9744427784417c04170af70974b2e5adb215b3136b6df2a20658d5c669930b3707b5602df4a5122795e@45.33.110.185:30300",
	"enode://37e322121d658fc6a3127413dadc4d356f2ef266b69cd5aa5e95326e76d58180080c54723638ae1444739e9f4c2906687afc110b336e4b0671b35f6167e59e83@45.79.98.29:30300",
	"enode://f1be54f3137fcc888b7ddd184fbd7be33eb6d41d73c7e223eb66dd2d8189b9182354f16c150fc92ace10f87567509e5f2b58ccc2a1b6b6491f605f374b2a35f3@45.79.90.221:30300",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
}

// RinkebyV5Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network for the experimental RLPx v5 topic-discovery network.
var RinkebyV5Bootnodes = []string{
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
}
