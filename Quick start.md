## Build from source

###  Install Dependencies

- **Go 1.17 or 1.18**

  Installation instructions can be found here: https://golang.org/doc/install.
  Ensure Go was installed properly and is a supported version:

  ```sh
  $ go version
  $ go env GOROOT GOPATH
  ```
  
- **Git**

  Installation instructions can be found at https://git-scm.com or  https://gitforwindows.org.
  ```shell
  $ git version
  ```

###  Install dcrnd

- **Windows Example**

  ```PowerShell
  PS> git clone https://github.com/Decred-Next/dcrnd.git $env:USERPROFILE\src\dcrnd
  PS> cd $env:USERPROFILE\src\dcrnd
  PS> go install . .\cmd\...
  ```

  Run the `dcrnd` executable now installed in `"$(go env GOPATH)\bin"`.

- **Unix Example**

  ```sh
  $ git clone https://github.com/Decred-Next/dcrnd.git $HOME/src/dcrnd
  $ (cd $HOME/src/dcrnd && go install)
  ```

  Run the `dcrnd` executable now installed in `$GOPATH/bin`.

- **macOs Example**

  ```shell
  $ git clone https://github.com/Decred-Next/dcrnd.git $HOME/src/dcrnd
  $ (cd $HOME/src/dcrnd && go install)
  ```

  Run the `dcrnd` executable now installed in `$GOPATH/bin`.

###  Install dcrctl

- **Windows Example**

  ```PowerShell
  PS> git clone https://github.com/Decred-Next/dcrnctl.git $env:USERPROFILE\src\dcrnctl
  PS> cd $env:USERPROFILE\src\dcrnd
  PS> go install . .\cmd\...
  ```

  Run the `dcrctl` executable now installed in `"$(go env GOPATH)\bin"`.

- **Unix Example**

  ```sh
  $ git clone https://github.com/Decred-Next/dcrnctl.git $HOME/src/dcrnctl
  $ (cd $HOME/src/dcrnctl && go install)
  ```

  Run the `dcrctl` executable now installed in `$GOPATH/bin`.

- **macOs Example**

  ```sh
  $ git clone https://github.com/Decred-Next/dcrnctl.git $HOME/src/dcrnctl
  $ (cd $HOME/src/dcrnctl && go install)
  ```

  Run the `dcrctl` executable now installed in `$GOPATH/bin`.

###  Install dcrwallet

- **Windows Example**

  ```PowerShell
  PS> git clone https://github.com/Decred-Next/dcrnwallet.git $env:USERPROFILE\src\dcrnwallet
  PS> cd $env:USERPROFILE\src\dcrnwallet
  PS> go install . .\cmd\...
  ```

  Run the `dcrwallet` executable now installed in `"$(go env GOPATH)\bin"`.

- **Unix Example**

  ```sh
  $ git clone https://github.com/Decred-Next/dcrnwallet.git $HOME/src/dcrnwallet
  $ (cd $HOME/src/dcrnwallet && go install)
  ```

  Run the `dcrwallet` executable now installed in `$GOPATH/bin`.

- **macOs Example**

  ```sh
  $ git clone https://github.com/Decred-Next/dcrnwallet.git $HOME/src/dcrnwallet
  $ (cd $HOME/src/dcrnwallet && go install)
  ```

  Run the `dcrwallet` executable now installed in `$GOPATH/bin`.

## Configuration

- **Configuration File Locations**

  These configuration files are located within the application home directory of the application. The location of these default home directories for Windows, macOS, and Linux are listed below:

  | OS      | dcrnd, dcrnwallet, dcrnctl App Directories |
  | ------- | ------------------------------------------ |
  | Windows | `%LOCALAPPDATA%\Dcrd\`                     |
  |         | `%LOCALAPPDATA%\Dcrwallet\`                |
  |         | `%LOCALAPPDATA%\Dcrctl\`                   |
  | macOS   | `~/Library/Application Support/Dcrd/`      |
  |         | `~/Library/Application Support/Dcrwallet/` |
  |         | `~/Library/Application Support/Dcrctl/`    |
  | Linux   | `~/.dcrd/`                                 |
  |         | `~/.dcrwallet/`                            |
  |         | `~/.dcrctl/`                               |

  Each of these folders is allowed its own `.conf` file, named after the individual application (`e.g. dcrd uses dcrd.conf`). Please also note that the `Dcrd` and `Dcrwallet` home directories are automatically created when each application is first launched.

- **Minimum Configuration**

  At the very minimum, for `dcrd`, `dcrwallet`, and `dcrctl` to be able to communicate with each other, they need to be launched with the same `rpcuser`/`rpcpass` combination.Please follow these steps:

  - If the OS-specific home directories listed in Configuration File Locations do not exist, please create them for `dcrd`, `dcrwallet`, and `dcrctl`.

  - Choose an arbitrary username and password, these will only be used for each application to communicate via remote procedure call. The easiest configuration is to set them all equal.

  - Using your favorite text editor, create a new text file and add the following lines:

     ```ini
     [Application Options]
     rpcuser=<chosen-username>
     rpcpass=<chosen-password>
     ```
     Save it as `dcrd.conf` in `dcrd`‘s home directory.

  - Create another new text file and add the following lines:

     ```ini
     [Application Options]
     username=<chosen-username>
     password=<chosen-password>
     ```
     Save it as `dcrwallet.conf` in `dcrwallet`‘s home directory.

  - Create a third text file and add the following lines:

     ```ini
     [Application Options]
     rpcuser=<chosen-username>
     rpcpass=<chosen-password>
     ```
     Save it as `dcrctl.conf` in `dcrctl`‘s home directory.

## dcrnd Setup Guide

- **Start dcrnd**

  With the correctly set configuration files, open another shell window in your Decred-Next directory . Type the command for Windows, macOS, and Linux are listed below:

  ```shell
  Windows: dcrnd.exe
  macOS: ./dcrnd
  Linux: ./dcrnd
  ```

- **Wait for dcrnd to Sync to the Decred-Next Blockchain**

  When `dcrnd` launches successfully, you should see your shell window begin to fill up with messages as the daemon connects to the network and starts processing blocks. Wait until it is completed - the entire blockchain is being downloaded into the `dcrd` data directory.

  You will see a line at the start like this:

  ```shell
  22:58:04 2016-02-09 [INF] BMGR: Syncing to block height 617 from peer 104.236.167.133:9108
  ```

  Then, as it continues to download blocks, you will see lines like this:

  ```shell
  22:58:16 2016-02-09 [INF] BMGR: Processed 321 blocks in the last 10.03s (544 transactions, height 322, 2016-02-09 09:50:34 +1000 EST)
  ```

  The blockchain will be fully synced once the most recently processed block is the current block height. 

  Note that this connection will be used in the future. You must leave this `dcrd` instance running in order to use `dcrwallet`.

## dcrnwallet Setup Guide

- **Manual Wallet Creation Command**

  If you do not already have a `wallet.db` file stored in `dcrwallet`‘s home directory, you must run the `dcrwallet --create` command. Steps for this can be found below.

  - Open a new shell window (Bash/Command Prompt/etc,..).
  - Navigate to the directory of the `dcrwallet` executable.
  - Enter the command `./dcrwallet --create` .

- **Wallet Creation Walkthrough**

  During this process, you’ll set a private passphrase, optionally set a public passphrase, and record your seed. To accomplish this, follow the steps below:

  - Set Passphrases for Your Wallet
     If the `dcrwallet --create` command successfully executed, you should be greeted by the following text:
     
     ```
     Enter the private passphrase for your new wallet:
     ```
     This first passphrase, the private passphrase, is what you will use to unlock your wallet when creating transactions or voting with Proof-of-Stake. Please use a unique and strong password. This password also protects the private keys within your wallet file, securing it from theft.
     After you’ve verified your private passphrase, you should see the following prompt:
     ```
     Do you want to add an additional layer of encryption for public data? (n/no/y/yes) [no]:
     ```
     The public passphrase is optional. It is used to encrypt all of the public data (transactions and addresses) within your wallet file so if it is stolen, an adversary can’t link you to your transactions.
     
  - Record Your Seed
     Before creating a new seed for your wallet, please review the Critical Information section above.
     After you’ve set your private passphrase and optional public passphrase, you’ll see the following prompt:
     
     ```
     Do you have an existing wallet seed you want to use? (n/no/y/yes) [no]:
     ```
     This guide assumes you do not have an existing seed, so continue by hitting `Enter` which will answer the prompt with the default `[no]`. NOTE: If you wish to restore your wallet by using your seed, you would simply enter `[yes]` here and follow the instructions on screen.
     After answering `[no]`, your seed phrase (wallet generation seed) and its hex will be displayed in the window. Please read through the IMPORTANT section displayed immediately after the hex.
     It cannot be stressed enough how important it is to save your seed phrase in a secure location, so if you haven’t committed this to memory, please review the Critical Information section above.
     Once you have written down the seed phrase and hex, type `OK` and press `Enter`. 
     After pressing `Enter`, you should see the following message:
     ```
     Creating the wallet...
     The wallet has been created successfully.
     ```
     The wallet will then be created. This might take a few minutes if you have a slow computer.

- **Launching dcrwallet**

  In order to launch `dcrwallet`, you first must have created your wallet and connected `dcrnd` to the Decred-Next network.

  With the correctly set configuration files, open another shell window in your Decred-Next directory (or use the last window if you have just created your wallet). Type the following command (review this guide’s Prerequisites to determine the right command for your OS/Shell application):

  ```shell
  ./dcrwallet
  ```
  
  Your `dcrwallet` will now connect to the network via `dcrd`. It will begin to scan the network for your active addresses which can take a few minutes on slow computers. Eventually it will start showing lines like:

  ```shell
  [INF] WLLT: Connecting block 0000000000002004ea8fa74af334cb291a22832642b5be603995683534bbb97b, height 9990
  ```
  
  This means your wallet is successfully connected to the network through your daemon.

## dcrctl Basics 

### overview

  `dcrctl` is the client that controls `dcrd` and `dcrwallet` via remote procedure call. You can use `dcrctl` for many things such as checking your balance, buying tickets, creating transactions, and viewing network information.

  `dcrctl` is not a daemon - it does not run permanently in the background - it calls the requested RPC method, prints the response, and then terminates immediately.

### Usage

- **Prerequisites:**

  1. Setup dcrnd and have it running in the background.

  2. Setup dcrwallet and have it running in the background.

- **To call `dcrd` RPC methods:**

  ```shell
  dcrctl <options> <rpc method> <rpc method args>
  ```

- **To call `dcrwallet` RPC methods:**

  ```shell
  dcrctl <options> --wallet <rpc method> <rpc method args>
  ```

- **To list available options:**

  ```shell
  ./dcrctl --help
  ```

- **To list available RPC methods:**

  ```shell
  ./dcrctl -l
  ```

------
### Example

- **Checking Your Balance**

  To send the `getbalance` command to `dcrwallet` using `dcrctl`, enter the following in your shell:

  ```
  dcrctl --wallet getbalance
  ```

  This will return all of the balances for all of your accounts.

------

- **Getting a New Receiving Address**

  To send the `getnewaddress` command to `dcrwallet` using `dcrctl`, enter the following in your shell:

  ```
  dcrctl --wallet getnewaddress
  ```

  This will generate and return a new payment address. To minimize address reuse, use this command to get a new address for each transaction you wish to receive.

------

- **Sending DCR**

  To send DCR to an address, issue the `sendtoaddress` command to `dcrwallet` using `dcrctl`. Enter the following in your shell, filling in the receiving address and amount to send:

  ```
  dcrctl --wallet sendtoaddress <address> <amount>
  ```

  If successful, `dcrctl` will return a transaction hash that can be used to watch the transaction on the official Decred-Next Block Explorer.

## Using the Block Explorer


### Overview

  All blocks and transactions on the Decred-Next blockchain are visible through the use of the block explorer, [dcrndata](https://github.com/Decred-Next/dcrndata).

  Public instances of dcrndata are available for the following networks:

- [mainnet](https://data.dcrn.xyz/)
- [testnet](https://testdata.dcrn.xyz/)

  Below is a quick review of some of the information on them.
  
  | Option         | Explanation                                               |
  | :------------- | :-------------------------------------------------------- |
  | `Height`       | The block number.                                         |
  | `Age`          | How long ago the block was added to the blockchain.       |
  | `Transactions` | The number of transactions included in the block.         |
  | `Votes`        | The number of proof-of-stake votes included in the block. |
  | `Fresh Stake`  | The number of new tickets purchased in this block.        |
  | `Size`         | The size (in bytes) of the block.                         |
  
  Under `Latest Transactions`, you can see the transaction ID (txid) and the value (in DCR) transmitted across the network.

------

### Blocks

  Blocks can be found by searching for their block height number, clicking on a `Height` value from the home page, or from their `BlockHash` value. Older blocks will have lower block numbers. The top half of a block overview shows relevant information about this specific block. This information includes: the block height, the block hash, and several key network parameters, described below:

| Option                   | Explanation                                                  |
| :----------------------- | :----------------------------------------------------------- |
| `Number of Transactions` | The number of standard transactions (DCR sent from one user to another). |
| `Height`                 | The height of the blockchain in which this block resides.    |
| `Block Reward`           | The amount of new DCR minted in this block.                  |
| `Timestamp`              | The time this block was created by a miner and was included in the blockchain. |
| `Merkle Root`            | A hash value of all the transaction hashes included in this block. |
| `Stake Root`             | A hash value of all the stake related transaction hashes in this block. This includes ticket purchases, votes, and ticket revocations. |
| `VoteBits`               | (1) Block was approved by proof-of-stake voters. (2) Block was vetoed by proof-of-stake voters and all non-stake transactions in the block were invalidated, along with the newly generated block reward for the proof-of-work miner and the Decred-Next Treasury. |
| `Final State`            | The final state of the pseudo random number generator used for ticket selection. |
| `Voters`                 | The number of successful proof-of-stake votes cast in this block. The maximum value is 5. |
| `Fresh Stake`            | The number of stake ticket purchases confirmed in this block. |
| `Revocations`            | The number of tickets that failed to vote and were revoked.  |
| `PoolSize`               | The total number of active proof-of-stake tickets.           |
| `Difficulty`             | The proof-of-work network difficulty.                        |
| `SBits`                  | The price of one proof-of-stake ticket.                      |
| `Bits`                   | A compact version of the network difficulty at the time the block was mined. |
| `Size`                   | The size of the block (in bytes).                            |
| `Version`                | The version of the block.                                    |
| `Nonce`                  | The value used by a miner to find the correct solution for this block. |

### Transactions

  This section lists all the transactions that were mined into this block. Transactions are chosen from the network mempool in order of highest fee first. All transactions in the block overview follow this order: Standard transactions (peer-to-peer transfer), proof-of-stake votes, proof-of-stake ticket purchases. The following sections will review each type of transaction.

------

#### Standard transactions

  Here’s the information included in standard Decred-Next transactions.

  | Option              | Explanation                                             |
  | :------------------ | :------------------------------------------------------ |
  | `Size`              | The size of the transaction in bytes.                   |
  | `Fee rate`          | The rate of fees collected by the network (per kB).     |
  | `Received Time`     | The time the network confirmed the transaction.         |
  | `Mined Time`        | The time a miner included the transaction in a block.   |
  | `Included in Block` | The block number that the transaction became a part of. |

  Note `Received Time`, `Mined Time`, and `Included in Block` will not have a value until a miner validates the transaction and includes it in a Decred-Next block. After being confirmed in a block, the transaction is considered complete.

------

#### Ticket purchases

  For a ticket purchase (stake submission) there are a few differences from a standard transaction shown.

  Note the difference under details: The word `Ticket` appears above the sender’s wallet address on the left, and the words `Subsidy Commitment` appear on the right. This particular user purchased a stake ticket for 8.75411638 DCR and received change in the amount of 7.15994209 DCR. The address listed on the left under `Ticket` is the address that contains the funds used to purchase this ticket. The first output on the right is the address that retains voting rights for this specific ticket. The second output, `Subsidy Commitment`, is the address where the reward will go. This is not yet shown by the block explorer at this time. The third and final output is the address where change for this transaction will be sent.

------

#### Proof-of-stake votes

  Note the identifying terms in the details section: `Vote`, `Stake Base`, `Block Commitment`, and `Vote Bits`:

  These keywords indicate that this transaction is a vote that was cast from a proof-of-stake ticket holder. In this particular example, the user had previously purchased a ticket for 8.99472311 DCR and was sent 10.82959184 DCR after the vote was cast in this transaction.
