# Using the DCRN Explorer

## Overview
All blocks and transactions on the Decred-Next blockchain are visible through the use of the DCRN explorer, dcrndata.

Public instances of dcrndata are available for the following networks:
* mainnet
* testnet

| Option |  Explanation  |
|:-------|:--------------|
|Height|The block number.|
|Age|How long ago the block was added to the blockchain.|
|Transactions|The number of transactions included in the block.|
|Votes|The number of proof-of-stake votes included in the block.|
|Tickets|The number of new tickets purchased in this block.|
|Size|The size (in bytes) of the block.|



## Blocks
Blocks can be found by searching for their block height number, clicking on a Height value from the home page, or from their BlockHash value. Older blocks will have lower block numbers. The top half of a block overview shows relevant information about this specific block. This information includes: the block height, the block hash, and several key network parameters, described below:

| Option |  Explanation  |
|:-------|:--------------|
|Height|The height of the blockchain in which this block resides.|
|Block Reward|The amount of new DCRN minted in this block.|
|Voters|The number of successful proof-of-stake votes cast in this block. The maximum value is 5.|
|Tickets|The number of tickets mined this block.|
|Transactions|The number of standard transactions (DCRN sent from one user to another).|
|Total sent|A sum of DCRN of all transactions in this block.|
|Size|The size of the block (in bytes).|
|Block Time|The time this block was created by a miner and was included in the blockchain.|
|Revocations|The number of tickets that failed to vote and were revoked.|
|Ticket Price|The price of one proof-of-stake ticket.|
|PoolSize|The total number of active proof-of-stake tickets.|
|PoW Difficulty|The proof-of-work network difficulty.|
|Nonce|The value used by a miner to find the correct solution for this block.|
|Final State|The final state of the pseudo random number generator used for ticket selection.|
|Vote Bits|(1) Block was approved by proof-of-stake voters. (2) Block was vetoed by proof-of-stake voters and all non-stake transactions in the block were invalidated, along with the newly generated block reward for the proof-of-work miner and the Decred Treasury.|
|Merkle Root|A hash value of all the transaction hashes included in this block.|
|Stake Root|A hash value of all the stake related transaction hashes in this block. This includes ticket purchases, votes, and ticket revocations.|



## Voting
Here’s the information included in Voting.

|   Option   |    Explanation    |
|:-----------|:------------------|
|Current Ticket Price|The amount of DCRN one must time-lock in order to buy a ticket now. |
|Next Ticket Price|The ticket price is algorithmically adjusted with the goal of keeping the ticket pool at an optimal size of 40,960 tickets. The  amount of DCRN one must time-lock if you buy a tickey in next cycle.|
|Ticket Pool Size|The pool of live tickets that are available to be called to vote. The target size for the ticket pool is 40,960.|
|Next Ticket Price Change|The number of blocks until the next price adjustment.|
|Vote Reward|The amount of voting reward(PoS reward) and voting cycle.|
|Total Staked DCRN|The sum of DCRN be locked throught buy tickets.|



## Mining
Here’s the information included in Mining.
|   Option   |    Explanation    |
|:-----------|:------------------|
|Difficulty|Difficulty is a measure of how difficult it is to mine a new block |
|Hashrate|The number of hashes per second computed by miners on the network.|
|Pow Reward|Miners receive block rewards after mined a new block.|



## Mempool
Mempool is a pool of transactions waiting to be mined.

###  Transactions
This section lists all the transactions that were mined into this block. Transactions are chosen from the network mempool in order of highest fee first. All transactions in the block overview follow this order: Standard transactions (peer-to-peer transfer), proof-of-stake ticket purchases, proof-of-stake votes. The following sections will review each type of transaction.

Here’s the information included in standard Decred transactions.

| Option |  Explanation  |
|:-------|:--------------|
|TransactionID|Unique identifier for this transaction.|
|Total DCRN|The total number of DCRN in this transcation.|
|Fee|The total fee of this transaction (Fee rate*Size).|
|Fee rate|The rate of fees collected by the network (per kB).|
|Size|The size of the transaction in bytes.|


## APIs
The dcrndata block explorer is exposed by two APIs: a Decred-Next implementation of the [Insight API](./InsightAPI#InsightAPI), and its own JSON HTTP API. The Insight API uses the path prefix `/insight/api`. The [Dcrndata API](./DcrndDataAPI#DcrndDataAPI) uses the path prefix `/api`. 