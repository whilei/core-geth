// Copyright 2019 The go-ethereum Authors
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

package openrpc

// This file contains a string constant containing the JSON schema data for OpenRPC.

// OpenRPCSchema defines the default full suite of possibly available go-ethereum RPC
// methods.
const OpenRPCSchema = `
{
    "openrpc": "1.0.0",
    "info": {
      "version": "1.0.10",
      "title": "Ethereum JSON-RPC",
      "description": "This API lets you interact with an EVM-based client via JSON-RPC",
      "license": {
        "name": "Apache 2.0",
        "url": "https://www.apache.org/licenses/LICENSE-2.0.html"
      }
    },
    "methods": [
      {
        "name": "web3_clientVersion",
        "description": "Returns the version of the current client",
        "summary": "current client version",
        "params": [],
        "result": {
          "name": "clientVersion",
          "description": "client version",
          "schema": {
            "title": "clientVersion",
            "type": "string"
          }
        }
      },
      {
        "name": "web3_sha3",
        "summary": "Hashes data",
        "description": "Hashes data using the Keccak-256 algorithm",
        "params": [
          {
            "name": "data",
            "description": "data to hash using the Keccak-256 algorithm",
            "summary": "data to hash",
            "schema": {
              "title": "data",
              "type": "string",
              "pattern": "^0x[a-fA-F\\d]+$"
            }
          }
        ],
        "result": {
          "name": "hashedData",
          "description": "Keccak-256 hash of the given data",
          "schema": {
            "$ref": "#/components/schemas/Keccak"
          }
        },
        "examples": [
          {
            "name": "sha3Example",
            "params": [
              {
                "name": "sha3ParamExample",
                "value": "0x68656c6c6f20776f726c64"
              }
            ],
            "result": {
              "name": "sha3ResultExample",
              "value": "0x47173285a8d7341e5e972fc677286384f802f8ef42a5ec5f03bbfa254cb01fad"
            }
          }
        ]
      },
      {
        "name": "net_listening",
        "summary": "returns listening status",
        "description": "Determines if this client is listening for new network connections.",
        "params": [],
        "result": {
          "name": "netListeningResult",
          "description": "` + "`" + `true` + "`" + ` if listening is active or ` + "`" + `false` + "`" + ` if listening is not active",
          "schema": {
            "title": "isNetListening",
            "type": "boolean"
          }
        },
        "examples": [
          {
            "name": "netListeningTrueExample",
            "description": "example of true result for net_listening",
            "params": [],
            "result": {
              "name": "netListeningExampleFalseResult",
              "value": true
            }
          }
        ]
      },
      {
        "name": "net_peerCount",
        "summary": "number of peers",
        "description": "Returns the number of peers currently connected to this client.",
        "params": [],
        "result": {
          "name": "quantity",
          "description": "number of connected peers.",
          "schema": {
            "title": "numConnectedPeers",
            "description": "Hex representation of number of connected peers",
            "type": "string"
          }
        }
      },
      {
        "name": "net_version",
        "summary": "Network identifier associated with network",
        "description": "Returns the network ID associated with the current network.",
        "params": [],
        "result": {
          "name": "networkId",
          "description": "Network ID associated with the current network",
          "schema": {
            "title": "networkId",
            "type": "string",
            "pattern": "^[\\d]+$"
          }
        }
      },
      {
        "name": "eth_blockNumber",
        "summary": "Returns the number of most recent block.",
        "params": [],
        "result": {
          "$ref": "#/components/contentDescriptors/BlockNumber"
        }
      },
      {
        "name": "eth_call",
        "summary": "Executes a new message call (locally) immediately without creating a transaction on the block chain.",
        "params": [
          {
            "$ref": "#/components/contentDescriptors/Transaction"
          },
          {
            "$ref": "#/components/contentDescriptors/BlockNumber"
          }
        ],
        "result": {
          "name": "returnValue",
          "description": "The return value of the executed contract",
          "schema": {
            "$ref": "#/components/schemas/Bytes"
          }
        }
      },
      {
        "name": "eth_chainId",
        "summary": "Returns the currently configured chain id",
        "description": "Returns the currently configured chain id, a value used in replay-protected transaction signing as introduced by [EIP-155](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-155.md).",
        "params": [],
        "result": {
          "name": "chainId",
          "description": "hex format integer of the current chain id. Defaults are ETC=61, ETH=1, Morden=62.",
          "schema": {
            "title": "chainId",
            "type": "string",
            "pattern": "^0x[a-fA-F\\d]+$"
          }
        }
      },
      {
        "name": "eth_coinbase",
        "summary": "Returns the client coinbase address.",
        "params": [],
        "result": {
          "name": "address",
          "description": "The address owned by the client that is used as default for things like the mining reward",
          "schema": {
            "$ref": "#/components/schemas/Address"
          }
        }
      },
      {
        "name": "eth_estimateGas",
        "summary": "Generates and returns an estimate of how much gas is necessary to allow the transaction to complete. The transaction will not be added to the blockchain. Note that the estimate may be significantly more than the amount of gas actually used by the transaction, for a variety of reasons including EVM mechanics and node performance.",
        "params": [
          {
            "$ref": "#/components/contentDescriptors/Transaction"
          }
        ],
        "result": {
          "name": "gasUsed",
          "description": "The amount of gas used",
          "schema": {
            "$ref": "#/components/schemas/Integer"
          }
        }
      },
      {
        "name": "eth_gasPrice",
        "summary": "Returns the current price per gas in wei",
        "params": [],
        "result": {
          "$ref": "#/components/contentDescriptors/GasPrice"
        }
      },
      {
        "name": "eth_getBalance",
        "summary": "Returns Ether balance of a given or account or contract",
        "params": [
          {
            "name": "address",
            "required": true,
            "description": "The address of the account or contract",
            "schema": {
              "$ref": "#/components/schemas/Address"
            }
          },
          {
            "name": "blockNumber",
            "description": "A BlockNumber at which to request the balance",
            "schema": {
              "$ref": "#/components/schemas/BlockNumber"
            }
          }
        ],
        "result": {
          "name": "getBalanceResult",
          "schema": {
            "$ref": "#/components/schemas/IntegerOrNull"
          }
        }
      },
      {
        "name": "eth_getBlockByHash",
        "summary": "Gets a block for a given hash",
        "params": [
          {
            "name": "blockHash",
            "required": true,
            "schema": {
              "$ref": "#/components/schemas/BlockHash"
            }
          },
          {
            "name": "includeTransactions",
            "description": "If ` + "`" + `true` + "`" + ` it returns the full transaction objects, if ` + "`" + `false` + "`" + ` only the hashes of the transactions.",
            "required": true,
            "schema": {
              "title": "isTransactionsIncluded",
              "type": "boolean"
            }
          }
        ],
        "result": {
          "name": "getBlockByHashResult",
          "schema": {
            "$ref": "#/components/schemas/BlockOrNull"
          }
        }
      },
      {
        "name": "eth_getBlockByNumber",
        "summary": "Gets a block for a given number",
        "params": [
          {
            "$ref": "#/components/contentDescriptors/BlockNumber"
          },
          {
            "name": "includeTransactions",
            "description": "If ` + "`" + `true` + "`" + ` it returns the full transaction objects, if ` + "`" + `false` + "`" + ` only the hashes of the transactions.",
            "required": true,
            "schema": {
              "title": "isTransactionsIncluded",
              "type": "boolean"
            }
          }
        ],
        "result": {
          "name": "getBlockByNumberResult",
          "schema": {
            "$ref": "#/components/schemas/BlockOrNull"
          }
        }
      },
      {
        "name": "eth_getBlockTransactionCountByHash",
        "summary": "Returns the number of transactions in a block from a block matching the given block hash.",
        "params": [
          {
            "$ref": "#/components/contentDescriptors/BlockHash"
          }
        ],
        "result": {
          "name": "blockTransactionCountByHash",
          "description": "The Number of total transactions in the given block",
          "schema": {
            "$ref": "#/components/schemas/IntegerOrNull"
          }
        }
      },
      {
        "name": "eth_getBlockTransactionCountByNumber",
        "summary": "Returns the number of transactions in a block from a block matching the given block number.",
        "params": [
          {
            "$ref": "#/components/contentDescriptors/BlockNumber"
          }
        ],
        "result": {
          "name": "blockTransactionCountByHash",
          "description": "The Number of total transactions in the given block",
          "schema": {
            "$ref": "#/components/schemas/IntegerOrNull"
          }
        }
      },
      {
        "name": "eth_getCode",
        "summary": "Returns code at a given contract address",
        "params": [
          {
            "name": "address",
            "required": true,
            "description": "The address of the contract",
            "schema": {
              "$ref": "#/components/schemas/Address"
            }
          },
          {
            "name": "blockNumber",
            "description": "A BlockNumber of which the code existed",
            "schema": {
              "$ref": "#/components/schemas/BlockNumber"
            }
          }
        ],
        "result": {
          "name": "bytes",
          "schema": {
            "$ref": "#/components/schemas/Bytes"
          }
        }
      },
      {
        "name": "eth_getFilterChanges",
        "summary": "Polling method for a filter, which returns an array of logs which occurred since last poll.",
        "params": [
          {
            "name": "filterId",
            "required": true,
            "schema": {
              "$ref": "#/components/schemas/FilterId"
            }
          }
        ],
        "result": {
          "name": "logResult",
          "schema": {
            "title": "logResult",
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/Log"
            }
          }
        }
      },
      {
        "name": "eth_getFilterLogs",
        "summary": "Returns an array of all logs matching filter with given id.",
        "params": [
          {
            "name": "filterId",
            "required": true,
            "schema": {
              "$ref": "#/components/schemas/FilterId"
            }
          }
        ],
        "result": {
          "$ref": "#/components/contentDescriptors/Logs"
        }
      },
      {
        "name": "eth_getRawTransactionByHash",
        "summary": "Returns raw transaction data of a transaction with the given hash.",
        "params": [
          {
            "$ref": "#/components/contentDescriptors/TransactionHash"
          }
        ],
        "result": {
          "name": "rawTransactionByHash",
          "description": "The raw transaction data",
          "schema": {
            "$ref": "#/components/schemas/Bytes"
          }
        }
      },
      {
        "name": "eth_getRawTransactionByBlockHashAndIndex",
        "summary": "Returns raw transaction data of a transaction with the given hash.",
        "params": [
          {
            "$ref": "#/components/contentDescriptors/BlockHash"
          },
          {
            "name": "index",
            "description": "The ordering in which a transaction is mined within its block.",
            "required": true,
            "schema": {
              "$ref": "#/components/schemas/Integer"
            }
          }
        ],
        "result": {
          "name": "rawTransaction",
          "description": "The raw transaction data",
          "schema": {
            "$ref": "#/components/schemas/Bytes"
          }
        }
      },
      {
        "name": "eth_getRawTransactionByBlockNumberAndIndex",
        "summary": "Returns raw transaction data of a transaction with the given hash.",
        "params": [
          {
            "$ref": "#/components/contentDescriptors/BlockNumber"
          },
          {
            "name": "index",
            "description": "The ordering in which a transaction is mined within its block.",
            "required": true,
            "schema": {
              "$ref": "#/components/schemas/Integer"
            }
          }
        ],
        "result": {
          "name": "rawTransaction",
          "description": "The raw transaction data",
          "schema": {
            "$ref": "#/components/schemas/Bytes"
          }
        }
      },
      {
        "name": "eth_getLogs",
        "summary": "Returns an array of all logs matching a given filter object.",
        "params": [
          {
            "$ref": "#/components/contentDescriptors/Filter"
          }
        ],
        "result": {
          "$ref": "#/components/contentDescriptors/Logs"
        }
      },
      {
        "name": "eth_getStorageAt",
        "summary": "Gets a storage value from a contract address, a position, and an optional blockNumber",
        "params": [
          {
            "$ref": "#/components/contentDescriptors/Address"
          },
          {
            "$ref": "#/components/contentDescriptors/Position"
          },
          {
            "$ref": "#/components/contentDescriptors/BlockNumber"
          }
        ],
        "result": {
          "name": "dataWord",
          "schema": {
            "$ref": "#/components/schemas/DataWord"
          }
        }
      },
      {
        "name": "eth_getTransactionByBlockHashAndIndex",
        "summary": "Returns the information about a transaction requested by the block hash and index of which it was mined.",
        "params": [
          {
            "$ref": "#/components/contentDescriptors/BlockHash"
          },
          {
            "name": "index",
            "description": "The ordering in which a transaction is mined within its block.",
            "required": true,
            "schema": {
              "$ref": "#/components/schemas/Integer"
            }
          }
        ],
        "result": {
          "$ref": "#/components/contentDescriptors/TransactionResult"
        },
        "examples": [
          {
            "name": "nullExample",
            "params": [
              {
                "name": "blockHashExample",
                "value": "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"
              },
              {
                "name": "indexExample",
                "value": "0x0"
              }
            ],
            "result": {
              "name": "nullResultExample",
              "value": null
            }
          }
        ]
      },
      {
        "name": "eth_getTransactionByBlockNumberAndIndex",
        "summary": "Returns the information about a transaction requested by the block hash and index of which it was mined.",
        "params": [
          {
            "$ref": "#/components/contentDescriptors/BlockNumber"
          },
          {
            "name": "index",
            "description": "The ordering in which a transaction is mined within its block.",
            "required": true,
            "schema": {
              "$ref": "#/components/schemas/Integer"
            }
          }
        ],
        "result": {
          "$ref": "#/components/contentDescriptors/TransactionResult"
        }
      },
      {
        "name": "eth_getTransactionByHash",
        "summary": "Returns the information about a transaction requested by transaction hash.",
        "params": [
          {
            "$ref": "#/components/contentDescriptors/TransactionHash"
          }
        ],
        "result": {
          "$ref": "#/components/contentDescriptors/TransactionResult"
        }
      },
      {
        "name": "eth_getTransactionCount",
        "summary": "Returns the number of transactions sent from an address",
        "params": [
          {
            "$ref": "#/components/contentDescriptors/Address"
          },
          {
            "$ref": "#/components/contentDescriptors/BlockNumber"
          }
        ],
        "result": {
          "name": "transactionCount",
          "schema": {
            "title": "nonceOrNull",
            "oneOf": [
              {
                "$ref": "#/components/schemas/Nonce"
              },
              {
                "$ref": "#/components/schemas/Null"
              }
            ]
          }
        }
      },
      {
        "name": "eth_getTransactionReceipt",
        "summary": "Returns the receipt information of a transaction by its hash.",
        "params": [
          {
            "$ref": "#/components/contentDescriptors/TransactionHash"
          }
        ],
        "result": {
          "name": "transactionReceiptResult",
          "description": "returns either a receipt or null",
          "schema": {
            "title": "transactionReceiptOrNull",
            "oneOf": [
              {
                "$ref": "#/components/schemas/Receipt"
              },
              {
                "$ref": "#/components/schemas/Null"
              }
            ]
          }
        }
      },
      {
        "name": "eth_getUncleByBlockHashAndIndex",
        "summary": "Returns information about a uncle of a block by hash and uncle index position.",
        "params": [
          {
            "$ref": "#/components/contentDescriptors/BlockHash"
          },
          {
            "name": "index",
            "description": "The ordering in which a uncle is included within its block.",
            "required": true,
            "schema": {
              "$ref": "#/components/schemas/Integer"
            }
          }
        ],
        "result": {
          "name": "uncle",
          "schema": {
            "$ref": "#/components/schemas/UncleOrNull"
          }
        }
      },
      {
        "name": "eth_getUncleByBlockNumberAndIndex",
        "summary": "Returns information about a uncle of a block by hash and uncle index position.",
        "params": [
          {
            "name": "uncleBlockNumber",
            "description": "The block in which the uncle was included",
            "required": true,
            "schema": {
              "$ref": "#/components/schemas/BlockNumber"
            }
          },
          {
            "name": "index",
            "description": "The ordering in which a uncle is included within its block.",
            "required": true,
            "schema": {
              "$ref": "#/components/schemas/Integer"
            }
          }
        ],
        "result": {
          "name": "uncleResult",
          "description": "returns an uncle or null",
          "schema": {
            "$ref": "#/components/schemas/UncleOrNull"
          }
        },
        "examples": [
          {
            "name": "nullResultExample",
            "params": [
              {
                "name": "uncleBlockNumberExample",
                "value": "0x0"
              },
              {
                "name": "uncleBlockNumberIndexExample",
                "value": "0x0"
              }
            ],
            "result": {
              "name": "nullResultExample",
              "value": null
            }
          }
        ]
      },
      {
        "name": "eth_getUncleCountByBlockHash",
        "summary": "Returns the number of uncles in a block from a block matching the given block hash.",
        "params": [
          {
            "$ref": "#/components/contentDescriptors/BlockHash"
          }
        ],
        "result": {
          "name": "uncleCountResult",
          "description": "The Number of total uncles in the given block",
          "schema": {
            "$ref": "#/components/schemas/IntegerOrNull"
          }
        }
      },
      {
        "name": "eth_getUncleCountByBlockNumber",
        "summary": "Returns the number of uncles in a block from a block matching the given block number.",
        "params": [
          {
            "$ref": "#/components/contentDescriptors/BlockNumber"
          }
        ],
        "result": {
          "$ref": "#/components/contentDescriptors/UncleCountResult"
        }
      },
      {
        "name": "eth_getProof",
        "summary": "Returns the account- and storage-values of the specified account including the Merkle-proof.",
        "params": [
          {
            "name": "address",
            "description": "The address of the account or contract",
            "required": true,
            "schema": {
              "$ref": "#/components/schemas/Address"
            }
          },
          {
            "name": "storageKeys",
            "required": true,
            "schema": {
              "title": "storageKeys",
              "description": "A storage key is indexed from the solidity compiler by the order it is declared. For mappings it uses the keccak of the mapping key with its position (and recursively for X-dimensional mappings)",
              "items": {
                "$ref": "#/components/schemas/StorageProofKey"
              }
            }
          },
          {
            "$ref": "#/components/contentDescriptors/BlockNumber"
          }
        ],
        "result": {
          "name": "account",
          "schema": {
            "title": "proofAccountOrNull",
            "oneOf": [
              {
                "title": "proofAccount",
                "type": "object",
                "description": "The merkle proofs of the specified account connecting them to the blockhash of the block specified",
                "properties": {
                  "address": {
                    "title": "proofAccountAddress",
                    "description": "The address of the account or contract of the request",
                    "$ref": "#/components/schemas/Address"
                  },
                  "accountProof": {
                    "$ref": "#/components/schemas/AccountProof"
                  },
                  "balance": {
                    "title": "proofAccountBalance",
                    "description": "The Ether balance of the account or contract of the request",
                    "$ref": "#/components/schemas/Integer"
                  },
                  "codeHash": {
                    "title": "proofAccountCodeHash",
                    "description": "The code hash of the contract of the request (keccak(NULL) if external account)",
                    "$ref": "#/components/schemas/Keccak"
                  },
                  "nonce": {
                    "title": "proofAccountNonce",
                    "description": "The transaction count of the account or contract of the request",
                    "$ref": "#/components/schemas/Nonce"
                  },
                  "storageHash": {
                    "title": "proofAccountStorageHash",
                    "description": "The storage hash of the contract of the request (keccak(rlp(NULL)) if external account)",
                    "$ref": "#/components/schemas/Keccak"
                  },
                  "storageProof": {
                    "$ref": "#/components/schemas/StorageProof"
                  }
                }
              },
              {
                "$ref": "#/components/schemas/Null"
              }
            ]
          }
        }
      },
      {
        "name": "eth_getWork",
        "summary": "Returns the hash of the current block, the seedHash, and the boundary condition to be met ('target').",
        "params": [],
        "result": {
          "name": "work",
          "schema": {
            "title": "getWorkResults",
            "type": "array",
            "items": [
              {
                "$ref": "#/components/schemas/PowHash"
              },
              {
                "$ref": "#/components/schemas/SeedHash"
              },
              {
                "$ref": "#/components/schemas/Difficulty"
              }
            ]
          }
        }
      },
      {
        "name": "eth_hashrate",
        "summary": "Returns the number of hashes per second that the node is mining with.",
        "params": [],
        "result": {
          "name": "hashesPerSecond",
          "description": "Integer of the number of hashes per second",
          "schema": {
            "$ref": "#/components/schemas/Integer"
          }
        }
      },
      {
        "name": "eth_mining",
        "summary": "Returns true if client is actively mining new blocks.",
        "params": [],
        "result": {
          "name": "mining",
          "description": "Whether or not the client is mining",
          "schema": {
            "type": "boolean"
          }
        }
      },
      {
        "name": "eth_newBlockFilter",
        "summary": "Creates a filter in the node, to notify when a new block arrives. To check if the state has changed, call eth_getFilterChanges.",
        "params": [],
        "result": {
          "$ref": "#/components/contentDescriptors/FilterId"
        }
      },
      {
        "name": "eth_newFilter",
        "summary": "Creates a filter object, based on filter options, to notify when the state changes (logs). To check if the state has changed, call eth_getFilterChanges.",
        "params": [
          {
            "$ref": "#/components/contentDescriptors/Filter"
          }
        ],
        "result": {
          "name": "filterId",
          "description": "The filter ID for use in ` + "`" + `eth_getFilterChanges` + "`" + `",
          "schema": {
            "$ref": "#/components/schemas/Integer"
          }
        }
      },
      {
        "name": "eth_newPendingTransactionFilter",
        "summary": "Creates a filter in the node, to notify when new pending transactions arrive. To check if the state has changed, call eth_getFilterChanges.",
        "params": [],
        "result": {
          "$ref": "#/components/contentDescriptors/FilterId"
        }
      },
      {
        "name": "eth_pendingTransactions",
        "summary": "Returns the pending transactions list",
        "params": [],
        "result": {
          "name": "pendingTransactions",
          "schema": {
            "$ref": "#/components/schemas/Transactions"
          }
        }
      },
      {
        "name": "eth_protocolVersion",
        "summary": "Returns the current ethereum protocol version.",
        "params": [],
        "result": {
          "name": "protocolVersion",
          "description": "The current ethereum protocol version",
          "schema": {
            "$ref": "#/components/schemas/Integer"
          }
        }
      },
      {
        "name": "eth_sendRawTransaction",
        "summary": "Creates new message call transaction or a contract creation for signed transactions.",
        "params": [
          {
            "name": "signedTransactionData",
            "required": true,
            "description": "The signed transaction data",
            "schema": {
              "$ref": "#/components/schemas/Bytes"
            }
          }
        ],
        "result": {
          "name": "transactionHash",
          "description": "The transaction hash, or the zero hash if the transaction is not yet available.",
          "schema": {
            "$ref": "#/components/schemas/Keccak"
          }
        }
      },
      {
        "name": "eth_submitHashrate",
        "deprecated": true,
        "summary": "Used for submitting mining hashrate.",
        "params": [
          {
            "name": "hashRate",
            "required": true,
            "schema": {
              "$ref": "#/components/schemas/DataWord"
            }
          },
          {
            "name": "id",
            "required": true,
            "description": "String identifying the client",
            "schema": {
              "$ref": "#/components/schemas/DataWord"
            }
          }
        ],
        "result": {
          "name": "submitHashRateSuccess",
          "description": "whether of not submitting went through successfully",
          "schema": {
            "type": "boolean"
          }
        }
      },
      {
        "name": "eth_submitWork",
        "summary": "Used for submitting a proof-of-work solution.",
        "params": [
          {
            "$ref": "#/components/contentDescriptors/Nonce"
          },
          {
            "name": "powHash",
            "required": true,
            "schema": {
              "$ref": "#/components/schemas/PowHash"
            }
          },
          {
            "name": "mixHash",
            "required": true,
            "schema": {
              "$ref": "#/components/schemas/MixHash"
            }
          }
        ],
        "result": {
          "name": "solutionValid",
          "description": "returns true if the provided solution is valid, otherwise false.",
          "schema": {
            "type": "boolean"
          }
        },
        "examples": [
          {
            "name": "submitWorkExample",
            "params": [
              {
                "name": "nonceExample",
                "description": "example of a number only used once",
                "value": "0x0000000000000001"
              },
              {
                "name": "powHashExample",
                "description": "proof of work to submit",
                "value": "0x6bf2cAE0dE3ec3ecA5E194a6C6e02cf42aADfe1C2c4Fff12E5D36C3Cf7297F22"
              },
              {
                "name": "mixHashExample",
                "description": "the mix digest example",
                "value": "0xD1FE5700000000000000000000000000D1FE5700000000000000000000000000"
              }
            ],
            "result": {
              "name": "solutionInvalidExample",
              "description": "this example should return ` + "`" + `false` + "`" + ` as it is not a valid pow to submit",
              "value": false
            }
          }
        ]
      },
      {
        "name": "eth_syncing",
        "summary": "Returns an object with data about the sync status or false.",
        "params": [],
        "result": {
          "name": "syncing",
          "schema": {
            "title": "isSyncingResult",
            "oneOf": [
              {
                "title": "syncingData",
                "description": "An object with sync status data",
                "type": "object",
                "properties": {
                  "startingBlock": {
                    "title": "syncingDataStartingBlock",
                    "description": "Block at which the import started (will only be reset, after the sync reached his head)",
                    "$ref": "#/components/schemas/Integer"
                  },
                  "currentBlock": {
                    "title": "syncingDataCurrentBlock",
                    "description": "The current block, same as eth_blockNumber",
                    "$ref": "#/components/schemas/Integer"
                  },
                  "highestBlock": {
                    "title": "syncingDataHighestBlock",
                    "description": "The estimated highest block",
                    "$ref": "#/components/schemas/Integer"
                  },
                  "knownStates": {
                    "title": "syncingDataKnownStates",
                    "description": "The known states",
                    "$ref": "#/components/schemas/Integer"
                  },
                  "pulledStates": {
                    "title": "syncingDataPulledStates",
                    "description": "The pulled states",
                    "$ref": "#/components/schemas/Integer"
                  }
                }
              },
              {
                "type": "boolean"
              }
            ]
          }
        }
      },
      {
        "name": "eth_uninstallFilter",
        "summary": "Uninstalls a filter with given id. Should always be called when watch is no longer needed. Additionally Filters timeout when they aren't requested with eth_getFilterChanges for a period of time.",
        "params": [
          {
            "name": "filterId",
            "required": true,
            "schema": {
              "$ref": "#/components/schemas/FilterId"
            }
          }
        ],
        "result": {
          "name": "filterUninstalledSuccess",
          "description": "returns true if the filter was successfully uninstalled, false otherwise.",
          "schema": {
            "type": "boolean"
          }
        }
      }
    ],
    "components": {
      "schemas": {
        "ProofNode": {
          "title": "proofNode",
          "type": "string",
          "description": "An individual node used to prove a path down a merkle-patricia-tree",
          "$ref": "#/components/schemas/Bytes"
        },
        "AccountProof": {
          "$ref": "#/components/schemas/ProofNodes"
        },
        "StorageProofKey": {
          "title": "storageProofKey",
          "description": "The key used to get the storage slot in its account tree.",
          "$ref": "#/components/schemas/Integer"
        },
        "StorageProof": {
          "title": "storageProofSet",
          "type": "array",
          "description": "Current block header PoW hash.",
          "items": {
            "title": "storageProof",
            "type": "object",
            "description": "Object proving a relationship of a storage value to an account's storageHash.",
            "properties": {
              "key": {
                "$ref": "#/components/schemas/StorageProofKey"
              },
              "value": {
                "title": "storageProofValue",
                "description": "The value of the storage slot in its account tree",
                "$ref": "#/components/schemas/Integer"
              },
              "proof": {
                "$ref": "#/components/schemas/ProofNodes"
              }
            }
          }
        },
        "ProofNodes": {
          "title": "proofNodes",
          "type": "array",
          "description": "The set of node values needed to traverse a patricia merkle tree (from root to leaf) to retrieve a value",
          "items": {
            "$ref": "#/components/schemas/ProofNode"
          }
        },
        "PowHash": {
          "title": "powHash",
          "description": "Current block header PoW hash.",
          "$ref": "#/components/schemas/DataWord"
        },
        "SeedHash": {
          "title": "seedHash",
          "description": "The seed hash used for the DAG.",
          "$ref": "#/components/schemas/DataWord"
        },
        "MixHash": {
          "title": "mixHash",
          "description": "The mix digest.",
          "$ref": "#/components/schemas/DataWord"
        },
        "Difficulty": {
          "title": "difficulty",
          "description": "The boundary condition ('target'), 2^256 / difficulty.",
          "$ref": "#/components/schemas/DataWord"
        },
        "FilterId": {
          "title": "filterId",
          "type": "string",
          "description": "An identifier used to reference the filter."
        },
        "BlockHash": {
          "title": "blockHash",
          "type": "string",
          "pattern": "^0x[a-fA-F\\d]{64}$",
          "description": "The hex representation of the Keccak 256 of the RLP encoded block"
        },
        "BlockNumber": {
          "title": "blockNumber",
          "type": "string",
          "description": "The hex representation of the block's height",
          "$ref": "#/components/schemas/Integer"
        },
        "BlockNumberTag": {
          "title": "blockNumberTag",
          "type": "string",
          "description": "The optional block height description",
          "enum": [
            "earliest",
            "latest",
            "pending"
          ]
        },
        "BlockOrNull": {
          "title": "blockOrNull",
          "oneOf": [
            {
              "$ref": "#/components/schemas/Block"
            },
            {
              "$ref": "#/components/schemas/Null"
            }
          ]
        },
        "IntegerOrNull": {
          "title": "integerOrNull",
          "oneOf": [
            {
              "$ref": "#/components/schemas/Integer"
            },
            {
              "$ref": "#/components/schemas/Null"
            }
          ]
        },
        "UncleOrNull": {
          "title": "uncleOrNull",
          "oneOf": [
            {
              "$ref": "#/components/schemas/Uncle"
            },
            {
              "$ref": "#/components/schemas/Null"
            }
          ]
        },
        "AddressOrNull": {
          "title": "addressOrNull",
          "oneOf": [
            {
              "$ref": "#/components/schemas/Address"
            },
            {
              "$ref": "#/components/schemas/Null"
            }
          ]
        },
        "Receipt": {
          "title": "receipt",
          "type": "object",
          "description": "The receipt of a transaction",
          "required": [
            "blockHash",
            "blockNumber",
            "contractAddress",
            "cumulativeGasUsed",
            "from",
            "gasUsed",
            "logs",
            "logsBloom",
            "to",
            "transactionHash",
            "transactionIndex"
          ],
          "properties": {
            "blockHash": {
              "$ref": "#/components/schemas/BlockHash"
            },
            "blockNumber": {
              "$ref": "#/components/schemas/BlockNumber"
            },
            "contractAddress": {
              "title": "ReceiptContractAddress",
              "description": "The contract address created, if the transaction was a contract creation, otherwise null",
              "$ref": "#/components/schemas/AddressOrNull"
            },
            "cumulativeGasUsed": {
              "title": "ReceiptCumulativeGasUsed",
              "description": "The gas units used by the transaction",
              "$ref": "#/components/schemas/Integer"
            },
            "from": {
              "$ref": "#/components/schemas/From"
            },
            "gasUsed": {
              "title": "ReceiptGasUsed",
              "description": "The total gas used by the transaction",
              "$ref": "#/components/schemas/Integer"
            },
            "logs": {
              "title": "logs",
              "type": "array",
              "description": "An array of all the logs triggered during the transaction",
              "items": {
                "$ref": "#/components/schemas/Log"
              }
            },
            "logsBloom": {
              "$ref": "#/components/schemas/BloomFilter"
            },
            "to": {
              "$ref": "#/components/schemas/To"
            },
            "transactionHash": {
              "$ref": "#/components/schemas/TransactionHash"
            },
            "transactionIndex": {
              "$ref": "#/components/schemas/TransactionIndex"
            },
            "postTransactionState": {
              "title": "ReceiptPostTransactionState",
              "description": "The intermediate stateRoot directly after transaction execution.",
              "$ref": "#/components/schemas/Keccak"
            },
            "status": {
              "title": "ReceiptStatus",
              "description": "Whether or not the transaction threw an error.",
              "type": "boolean"
            }
          }
        },
        "BloomFilter": {
          "title": "bloomFilter",
          "type": "string",
          "description": "A 2048 bit bloom filter from the logs of the transaction. Each log sets 3 bits though taking the low-order 11 bits of each of the first three pairs of bytes in a Keccak 256 hash of the log's byte series"
        },
        "Log": {
          "title": "log",
          "type": "object",
          "description": "An indexed event generated during a transaction",
          "properties": {
            "address": {
              "title": "LogAddress",
              "description": "Sender of the transaction",
              "$ref": "#/components/schemas/Address"
            },
            "blockHash": {
              "$ref": "#/components/schemas/BlockHash"
            },
            "blockNumber": {
              "$ref": "#/components/schemas/BlockNumber"
            },
            "data": {
              "title": "LogData",
              "description": "The data/input string sent along with the transaction",
              "$ref": "#/components/schemas/Bytes"
            },
            "logIndex": {
              "title": "LogIndex",
              "description": "The index of the event within its transaction, null when its pending",
              "$ref": "#/components/schemas/Integer"
            },
            "removed": {
              "title": "logIsRemoved",
              "description": "Whether or not the log was orphaned off the main chain",
              "type": "boolean"
            },
            "topics": {
              "$ref": "#/components/schemas/Topics"
            },
            "transactionHash": {
              "$ref": "#/components/schemas/TransactionHash"
            },
            "transactionIndex": {
              "$ref": "#/components/schemas/TransactionIndex"
            }
          }
        },
        "Topics": {
          "title": "LogTopics",
          "description": "Topics are order-dependent. Each topic can also be an array of DATA with 'or' options.",
          "type": "array",
          "items": {
            "$ref": "#/components/schemas/Topic"
          }
        },
        "Topic": {
          "title": "topic",
          "description": "32 Bytes DATA of indexed log arguments. (In solidity: The first topic is the hash of the signature of the event (e.g. Deposit(address,bytes32,uint256))",
          "$ref": "#/components/schemas/DataWord"
        },
        "TransactionIndex": {
          "title": "transactionIndex",
          "description": "The index of the transaction. null when its pending",
          "$ref": "#/components/schemas/IntegerOrNull"
        },
        "BlockNumberOrNull": {
          "title": "blockNumberOrNull",
          "description": "The block number or null when its the pending block",
          "oneOf": [
            { "$ref": "#/components/schemas/BlockNumber" },
            { "$ref": "#/components/schemas/Null" }
          ]
        },
        "BlockHashOrNull": {
          "title": "blockNumberOrNull",
          "description": "The block hash or null when its the pending block",
          "$ref": "#/components/schemas/KeccakOrPending"
        },
        "NonceOrNull": {
          "title": "nonceOrNull",
          "description": "Randomly selected number to satisfy the proof-of-work or null when its the pending block",
          "oneOf": [
            { "$ref": "#/components/schemas/Nonce" },
            { "$ref": "#/components/schemas/Null" }
          ]
        },
        "From": {
          "title": "From",
          "description": "The sender of the transaction",
          "$ref": "#/components/schemas/Address"
        },
        "To": {
          "title": "To",
          "description": "Destination address of the transaction. Null if it was a contract create.",
          "oneOf": [
            { "$ref": "#/components/schemas/Address" },
            { "$ref": "#/components/schemas/Null" }
          ]
        },
        "LightBlock": {
          "title": "lightBlock",
          "type": "object",
          "properties": {
            "number": {
              "$ref": "#/components/schemas/BlockNumberOrNull"
            },
            "hash": {
              "$ref": "#/components/schemas/BlockHashOrNull"
            },
            "parentHash": {
              "$ref": "#/components/schemas/BlockHash"
            },
            "nonce": {
              "$ref": "#/components/schemas/NonceOrNull"
            },
            "sha3Uncles": {
              "title": "blockShaUncles",
              "description": "Keccak hash of the uncles data in the block",
              "$ref": "#/components/schemas/Keccak"
            },
            "logsBloom": {
              "title": "blockLogsBloom",
              "type": "string",
              "description": "The bloom filter for the logs of the block or null when its the pending block",
              "pattern": "^0x[a-fA-F\\d]+$"
            },
            "transactionsRoot": {
              "title": "blockTransactionsRoot",
              "description": "The root of the transactions trie of the block.",
              "$ref": "#/components/schemas/Keccak"
            },
            "stateRoot": {
              "title": "blockStateRoot",
              "description": "The root of the final state trie of the block",
              "$ref": "#/components/schemas/Keccak"
            },
            "receiptsRoot": {
              "title": "blockReceiptsRoot",
              "description": "The root of the receipts trie of the block",
              "$ref": "#/components/schemas/Keccak"
            },
            "miner": {
              "$ref": "#/components/schemas/AddressOrNull"
            },
            "difficulty": {
              "title": "blockDifficulty",
              "type": "string",
              "description": "Integer of the difficulty for this block"
            },
            "totalDifficulty": {
              "title": "blockTotalDifficulty",
              "description": "Integer of the total difficulty of the chain until this block",
              "$ref": "#/components/schemas/IntegerOrNull"
            },
            "extraData": {
              "title": "blockExtraData",
              "type": "string",
              "description": "The 'extra data' field of this block"
            },
            "size": {
              "title": "blockSize",
              "type": "string",
              "description": "Integer the size of this block in bytes"
            },
            "gasLimit": {
              "title": "blockGasLimit",
              "type": "string",
              "description": "The maximum gas allowed in this block"
            },
            "gasUsed": {
              "title": "blockGasUsed",
              "type": "string",
              "description": "The total used gas by all transactions in this block"
            },
            "timestamp": {
              "title": "blockTimeStamp",
              "type": "string",
              "description": "The unix timestamp for when the block was collated"
            }
          }
        },
        "Uncle": {
          "title": "uncle",
          "type": "object",
          "description": "Orphaned blocks that can be included in the chain but at a lower block reward. NOTE: An uncle doesn’t contain individual transactions.",
          "allOf": [
            {
              "$ref": "#/components/schemas/LightBlock"
            },
            {
              "title": "uncleData",
              "type": "object",
              "properties": {
                "uncles": {
                  "title": "uncleHashes",
                  "description": "Array of uncle hashes",
                  "type": "array",
                  "items": {
                    "title": "uncleHash",
                    "description": "Block hash of the RLP encoding of an uncle block",
                    "$ref": "#/components/schemas/Keccak"
                  }
                }
              }
            }
          ]
        },
        "Block": {
          "title": "block",
          "type": "object",
          "allOf": [
            {
              "$ref": "#/components/schemas/LightBlock"
            },
            {
              "$ref": "#/components/schemas/Uncle"
            },
            {
              "title": "blockData",
              "type": "object",
              "properties": {
                "transactions": {
                  "title": "transactionsOrHashes",
                  "description": "Array of transaction objects, or 32 Bytes transaction hashes depending on the last given parameter",
                  "type": "array",
                  "items": {
                    "oneOf": [
                      {
                        "$ref": "#/components/schemas/Transaction"
                      },
                      {
                        "$ref": "#/components/schemas/TransactionHash"
                      }
                    ]
                  }
                }
              }
            }
          ]
        },
        "Transaction": {
          "title": "transaction",
          "type": "object",
          "required": [
            "gas",
            "gasPrice",
            "nonce"
          ],
          "properties": {
            "blockHash": {
              "$ref": "#/components/schemas/BlockHashOrNull"
            },
            "blockNumber": {
              "$ref": "#/components/schemas/BlockNumberOrNull"
            },
            "from": {
              "$ref": "#/components/schemas/From"
            },
            "gas": {
              "title": "transactionGas",
              "type": "string",
              "description": "The gas limit provided by the sender in Wei"
            },
            "gasPrice": {
              "title": "transactionGasPrice",
              "type": "string",
              "description": "The gas price willing to be paid by the sender in Wei"
            },
            "hash": {
              "$ref": "#/components/schemas/TransactionHash"
            },
            "input": {
              "title": "transactionInput",
              "type": "string",
              "description": "The data field sent with the transaction"
            },
            "nonce": {
              "title": "transactionNonce",
              "description": "The total number of prior transactions made by the sender",
              "$ref": "#/components/schemas/Nonce"
            },
            "to": {
              "$ref": "#/components/schemas/To"
            },
            "transactionIndex": {
              "$ref": "#/components/schemas/TransactionIndex"
            },
            "value": {
              "title": "transactionValue",
              "description": "Value of Ether being transferred in Wei",
              "$ref": "#/components/schemas/Keccak"
            },
            "v": {
              "title": "transactionSigV",
              "type": "string",
              "description": "ECDSA recovery id"
            },
            "r": {
              "title": "transactionSigR",
              "type": "string",
              "description": "ECDSA signature r"
            },
            "s": {
              "title": "transactionSigS",
              "type": "string",
              "description": "ECDSA signature s"
            }
          }
        },
        "Transactions": {
          "title": "transactions",
          "description": "An array of transactions",
          "type": "array",
          "items": {
            "$ref": "#/components/schemas/Transaction"
          }
        },
        "TransactionHash": {
          "title": "transactionHash",
          "type": "string",
          "description": "Keccak 256 Hash of the RLP encoding of a transaction",
          "$ref": "#/components/schemas/Keccak"
        },
        "KeccakOrPending": {
          "title": "keccakOrPending",
          "oneOf": [
            {
              "$ref": "#/components/schemas/Keccak"
            },
            {
              "$ref": "#/components/schemas/Null"
            }
          ]
        },
        "Keccak": {
          "title": "keccak",
          "type": "string",
          "description": "Hex representation of a Keccak 256 hash",
          "pattern": "^0x[a-fA-F\\d]{64}$"
        },
        "Nonce": {
          "title": "nonce",
          "description": "A number only to be used once",
          "$ref": "#/components/schemas/Integer"
        },
        "Null": {
          "title": "null",
          "type": "null",
          "description": "Null"
        },
        "Integer": {
          "title": "integer",
          "type": "string",
          "pattern": "^0x[a-fA-F0-9]+$",
          "description": "Hex representation of the integer"
        },
        "Address": {
          "title": "address",
          "type": "string",
          "pattern": "^0x[a-fA-F\\d]{40}$"
        },
        "Addresses": {
          "title": "addresses",
          "type": "array",
          "description": "List of contract addresses from which to monitor events",
          "items": {
            "$ref": "#/components/schemas/Address"
          }
        },
        "Position": {
          "title": "position",
          "type": "string",
          "description": "Hex representation of the storage slot where the variable exists",
          "pattern": "^0x([a-fA-F0-9]?)+$"
        },
        "DataWord": {
          "title": "dataWord",
          "type": "string",
          "description": "Hex representation of a 256 bit unit of data",
          "pattern": "^0x([a-fA-F\\d]{64})?$"
        },
        "Bytes": {
          "title": "bytes",
          "type": "string",
          "description": "Hex representation of a variable length byte array",
          "pattern": "^0x([a-fA-F0-9]?)+$"
        }
      },
      "contentDescriptors": {
        "Block": {
          "name": "block",
          "summary": "A block",
          "description": "A block object",
          "schema": {
            "$ref": "#/components/schemas/Block"
          }
        },
        "Null": {
          "name": "Null",
          "description": "JSON Null value",
          "summary": "Null value",
          "schema": {
            "$ref": "#/components/schemas/Null"
          }
        },
        "Signature": {
          "name": "signature",
          "summary": "The signature.",
          "required": true,
          "schema": {
            "title": "signatureBytes",
            "type": "string",
            "description": "Hex representation of byte array between 2 and 65 chars long",
            "pattern": "0x^([A-Fa-f0-9]{2}){65}$"
          }
        },
        "GasPrice": {
          "name": "gasPrice",
          "required": true,
          "schema": {
            "description": "Integer of the current gas price",
            "$ref": "#/components/schemas/Integer"
          }
        },
        "Transaction": {
          "required": true,
          "name": "transaction",
          "schema": {
            "$ref": "#/components/schemas/Transaction"
          }
        },
        "TransactionResult": {
          "name": "transactionResult",
          "description": "Returns a transaction or null",
          "schema": {
            "oneOf": [
              {
                "$ref": "#/components/schemas/Transaction"
              },
              {
                "$ref": "#/components/schemas/Null"
              }
            ]
          }
        },
        "UncleCountResult": {
          "name": "uncleCountResult",
          "description": "The Number of total uncles in the given block",
          "schema": {
            "$ref": "#/components/schemas/IntegerOrNull"
          }
        },
        "Message": {
          "name": "message",
          "required": true,
          "schema": {
            "$ref": "#/components/schemas/Bytes"
          }
        },
        "Filter": {
          "name": "filter",
          "required": true,
          "schema": {
            "title": "filter",
            "type": "object",
            "description": "A filter used to monitor the blockchain for log/events",
            "properties": {
              "fromBlock": {
                "$ref": "#/components/schemas/BlockNumber"
              },
              "toBlock": {
                "$ref": "#/components/schemas/BlockNumber"
              },
              "address": {
                "title": "oneOrArrayOfAddresses",
                "oneOf": [
                  {
                    "$ref": "#/components/schemas/Address"
                  },
                  {
                    "$ref": "#/components/schemas/Addresses"
                  }
                ]
              },
              "topics": {
                "$ref": "#/components/schemas/Topics"
              }
            }
          }
        },
        "Address": {
          "name": "address",
          "required": true,
          "schema": {
            "$ref": "#/components/schemas/Address"
          }
        },
        "BlockHash": {
          "name": "blockHash",
          "required": true,
          "schema": {
            "$ref": "#/components/schemas/BlockHash"
          }
        },
        "Nonce": {
          "name": "nonce",
          "required": true,
          "schema": {
            "$ref": "#/components/schemas/Nonce"
          }
        },
        "Position": {
          "name": "key",
          "required": true,
          "schema": {
            "$ref": "#/components/schemas/Position"
          }
        },
        "Logs": {
          "name": "logs",
          "description": "An array of all logs matching filter with given id.",
          "schema": {
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/Log"
            }
          }
        },
        "FilterId": {
          "name": "filterId",
          "schema": {
            "$ref": "#/components/schemas/FilterId"
          }
        },
        "BlockNumber": {
          "name": "blockNumber",
          "required": true,
          "schema": {
            "oneOf": [
              {
                "$ref": "#/components/schemas/BlockNumber"
              },
              {
                "$ref": "#/components/schemas/BlockNumberTag"
              }
            ]
          }
        },
        "TransactionHash": {
          "name": "transactionHash",
          "required": true,
          "schema": {
            "$ref": "#/components/schemas/TransactionHash"
          }
        }
      }
    }
  }

`
