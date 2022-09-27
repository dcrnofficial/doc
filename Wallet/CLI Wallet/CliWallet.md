# CLI Wallet

*Last updated for CLI Wallet release v1.5.2*


## Installation
dcrninstall is the recommended method to install the Decred-next command line interface tools dcrnd, dcrnwallet, and dcrnctl.

dcrninstall is an automatic installer and upgrader for the Decred-next software. The newest release can be found here: https://github.com/dcrnofficial/binary-release/releases. Binaries are provided for Windows, macOS and Linux.


## Operating System Differences
<span id="osd"></span>
The purpose of this page is to explain the key differences for running the cross-platform command-line applications on Windows, Linux, and macOS.

### Launch Commands
The first major difference in the command line applications (dcrnd,dcrnwallet,dcrnctl) is how you launch them from the command line. This is not as much of an operating system as it is a difference in shells. Windows comes with Command Prompt and PowerShell installed. macOS uses Bash within the Terminal application, and many Linux distributions use Bash as well. Below are examples of the basic run commands for Bash and Command Prompt.

**Command Prompt:** `dcrnd.exe`, `dcrnwallet.exe`, `dcrnctl.exe`
**Bash:** `./dcrnd`, `./dcrnwallet`, `./dcrnctl`

**Some of our guides might be OS-agnostic with the launch commands. If a guide says to run dcrnctl --wallet getbalance, it's referring to using dcrnctl.exe --wallet getbalance for Command Prompt and ./dcrnctl --wallet getbalance for Bash.**

### Application Directory Locations

The other way the command line clients differ is the location of each application directory (blocks, wallets, log files, configuration files are all stored within the data directory). Below is a table of the default application directories for each application.

<span id="adp"></span>
|  OS  | dcrnd |dcrnwallet| dcrnctl |
|:----:|:-----:|:--------:|:-------:|
|Linux|~/.dcrnd/|~/.dcrnwallet/|~/.dcrnctl/|
|Windows|%LOCALAPPDATA%\Dcrnd\ |%LOCALAPPDATA%\Dcrnwallet\ |%LOCALAPPDATA%\Dcrnctl\ |
|MacOS|~/Library/Application Support/Dcrnd/|~/Library/Application Support/Dcrnwallet/|~/Library/Application Support/Dcrnctl/|



## Startup Basics

This guide applies to command-line application users. Decrediton users can safely ignore the use of config files - Decrediton handles basic configuration automatically. It is also worth noting that some of our guides show configuration file settings and other guides show startup command flags.


### Configuration File Locations
All of the Decred software, when started, reads from a configuration file to determine which settings it should enable/disable/set during that initial load. All of the command line startup flags (e.g. dcrnwallet --testnet) can be replaced by settings within the appropriate configuration file (e.g. dcrnwallet --testnet could be replaced by testnet=1 in dcrnwallet.conf).

Each of these folders is allowed its own .conf file, named after the individual application (e.g. dcrnd uses dcrnd.conf). Please also note that the Dcrnd and Dcrnwallet home directories are automatically created when each application is first launched. You will have to manually create a Dcrnctl home directory to utilize a config file.

### Startup Command Flags
A majority of the settings you are able to set via the configuration file can also be passed to the application as parameters during launch. For example, the following OS-specific commands would open dcrnd for Testnet use, an alternative to using testnet=1 in your config file:
>Windows: dcrnd.exe --testnet
>macOS: ./dcrnd --testnet
>Linux: ./dcrnd --testnet

The above example would first look to the dcrnd configuration file for settings and then look to the executable command to enable the testnet setting.



## dcrnd Setup Guide
This guide is intended to help you setup the dcrnd application.

**Prerequisites:**
1. You have installed dcrnd. 
2. Review how the launch commands for the Command Prompt (Windows) and Bash (macOS/Linux) shells differ [here](#osd).

Dcrnd is the node daemon for Decred. A daemon is a program that works in the background that you do not interface with directly. dcrnd maintains the entire past transactional ledger (or blockchain) of Decred and allows relaying of transactions to other Decred nodes across the world. You can think of it as your own personal Decred blockchain server. The blockchain is saved in the data folder within dcrnd's home directory.

### Connect to the Decred Network

The first time launching dcrnd, it will connect to the Decred network and begin downloading the blockchain. To allow dcrnwallet and dcrnctl to communicate with dcrnd, the configuration files must have rpcuser and rpcpass settings enabled.

1. Configure RPC Username and Password
2. Start dcrnd
    Type the following command
    `dcrnd`

3. Wait for dcrnd to Sync to the Decred Blockchain
   
When dcrnd launches successfully, you should see your shell window begin to fill up with messages as the daemon connects to the network and starts processing blocks. Wait until it is completed - the entire blockchain is being downloaded into the dcrnd data directory.

You will see a line at the start like this:
>2022-05-19 10:18:53.537 [INF] BMGR: Syncing to block height 793 from peer 54.193.161.248:9108

Then, as it continues to download blocks, you will see lines like this:
>2022-05-19 10:18:58.050 [INF] BMGR: Processed 47 blocks in the last 10.07s (50 transactions, 723 tickets, 0 votes, 0 revocations, height 47, 2022-05-18 22:18:58 +0800 CST)

The blockchain will be fully synced once the most recently processed block is the current block height. You can tell by either comparing the date and time in the log message or by comparing the height of the last block processed against the last block height on the official block explorer.

Note that this connection will be used in the future. You must leave this dcrnd instance running in order to use dcrnwallet.



## dcrnwallet Setup Guide
This guide is intended to help you setup the dcrnwallet application.

**Prerequisites:**
1. You have installed dcrnwallet.
2. Setup dcrndn and have it running in the background.

Dcrnwallet is the daemon that handles Decred wallet functionality for a single user. It manages all of your accounts, addresses, and transactions; tracks balances across addresses; and allows stakeholders to participate in Proof-of-Stake voting.

In order to run dcrnwallet, a wallet.db must exist within dcrnwallet's home directory. In order for that file to exist, you must create a new wallet. dcrinstall automatically starts the creation process. If you delete your wallet.db or used another installation process, you'll have to run the manual wallet creation command.

### Manual Wallet Creation Command
<span id="mwcc"></span>
If you do not already have a wallet.db file stored in dcrnwallet's home directory, you must run the dcrnwallet --create command. Steps for this can be found below.
1. Open a new shell window (Bash/Command Prompt/etc,..).
2. Navigate to the directory of the dcrnwallet executable.
3. Enter the command `dcrnwallet --create`


### Wallet Creation Walkthrough
During this process, you'll set a private passphrase, optionally set a public passphrase, and record your seed. To accomplish this, follow the steps below:
1. **Set Passphrases for Your Wallet**
    If the `dcrnwallet --create` command successfully executed, you should be greeted by the following text:
    >Enter the private passphrase for your new wallet:

    This first passphrase, the private passphrase, is what you will use to unlock your wallet when creating transactions or voting with Proof-of-Stake. Please use a unique and strong password. This password also protects the private keys within your wallet file, securing it from theft.

    After you've verified your private passphrase, you should see the following prompt:
    >Do you want to add an additional layer of encryption for public data? (n/no/y/yes) [no]:

    The public passphrase is optional. It is used to encrypt all of the public data (transactions and addresses) within your wallet file so if it is stolen, an adversary can't link you to your transactions.

2. **Record Your Seed**
    After you've set your private passphrase and optional public passphrase, you'll see the following prompt:
    >Do you have an existing wallet seed you want to use? (n/no/y/yes) [no]:
    
    This guide assumes you do not have an existing seed, so continue by hitting Enter which will answer the prompt with the default [no]. 
    **NOTE: If you wish to restore your wallet by using your seed, you would simply enter [yes] here and follow the instructions on screen.**

    After answering [no], your seed phrase (wallet generation seed) and its hex will be displayed in the window. Please read through the IMPORTANT section displayed immediately after the hex.

    >:zap:*During the creation process for your wallet, you will be given a sequence of 33 words known as a seed phrase. This seed phrase is essentially the private key for your wallet. You will be able to use this seed phrase to restore your private keys, transaction history, and balances using any Decred wallet on any computer.
    This ultimately means that anyone who knows your seed can use it to restore your private keys, transaction history, and balances to a Decred wallet on their computer without your knowledge. For this reason, it is of utmost importance to keep your seed phrase safe. Treat this seed the same way you would treat a physical key to a safe. If you lose your seed phrase, you permanently lose access to your wallet and all funds within it. It cannot be recovered by anyone, including the Decred developers. It is recommended you write it down on paper and store that somewhere secure. If you decide to keep it on your computer, it would be best to keep it in an encrypted document (do not forget the password) in case the file or your computer is stolen.
    Every seed phrase is also associated with a 64 character seed hex. The seed hex functions the same way as the seed phrase, so it is also important to keep your seed hex secure.
    :zap:DO NOT, UNDER ANY CIRCUMSTANCES, GIVE YOUR SEED OR THE ASSOCIATED HEX KEY TO ANYONE! NOT EVEN DECRED DEVELOPERS!*
    
    Once you have written down the seed phrase and hex, type OK and press Enter. 
    **NOTE: If you did not write the phrase down before continuing, you should start this process over after deleting your wallet file.**

    After pressing Enter, you should see the following message:
    >The wallet has been created successfully.


### Launching dcrnwallet
In order to launch dcrnwallet, you first must have created your wallet and connected dcrnd to the Decred network.

1. Configure RPC Username and Password
2. Start dcrnwallet
    Type the following command
    `dcrnwallet -u <rpcuser> -P <rpcpass> --pass <yourprivatepassphrase>`

    Your dcrnwallet will now connect to the network via dcrnd. It will begin to scan the network for your active addresses which can take a few minutes on slow computers. Eventually it will start showing lines like:
    >2022-05-19 10:36:54.885 [INF] SYNC: Blockchain sync completed, wallet ready for general usage.

    This means your wallet is successfully connected to the network through your daemon.



## dcrnctl Basics
**Prerequisites:**
1. You have installed the command line tools.
2. Setup dcrnd and have it running in the background.
3. Setup dcrnwallet and have it running in the background.

Dcrnctl is the client that controls dcrnd and dcrnwallet via remote procedure call (RPC). You can use dcrnctl for many things such as checking your balance, buying tickets, creating transactions, and viewing network information.

Dcrnctl is not a daemon - it does not run permanently in the background - it calls the requested RPC method, prints the response, and then terminates immediately.

### Usage
<span id="ctlUsage"></span>
* To call dcrnd RPC methods:
    `dcrnctl <options> <rpc method> <rpc method args>`

* To call dcrnwallet RPC methods:
    `dcrnctl <options> --wallet <rpc method> <rpc method args>`

* To list available options:
    `dcrnctl --help`

* To list available RPC methods:
    `dcrnctl -l`


### Unlocking Your Wallet
Some functionality of dcrnwallet requires the wallet to be unlocked.
The command to unlock your wallet follows:
`promptsecret | dcrnctl --wallet walletpassphrase - 0`

Here, we are specifying for dcrnctl to send the command to dcrnwallet by using the `--wallet` flag. The actual command is `walletpassphrase` which accepts two parameters, your private passphrase and a time limit. Specifying `0` for a time limit unlocks dcrnwallet without a time limit. Note, however, that this only unlocks the wallet for the current session. If you close the window the wallet is running in, or stop/restart the dcrnwallet, you will need to unlock it again the next time you start it.

### Checking Your Balance
To send the getbalance command to dcrnwallet using dcrnctl, enter the following in your shell:
`dcrnctl --wallet getbalance`

This will return all of the balances for all of your accounts.

### Getting a New Receiving Address
To send the getnewaddress command to dcrnwallet using dcrnctl, enter the following in your shell:
`dcrnctl --wallet getnewaddress`

This will generate and return a new payment address. To minimize address reuse, use this command to get a new address for each transaction you wish to receive.

### Sending DCRN
To send DCRN to an address, issue the sendtoaddress command to dcrnwallet using dcrnctl. Enter the following in your shell, filling in the receiving address and amount to send:
`dcrnctl --wallet sendtoaddress <address> <amount>`

If successful, dcrnctl will return a transaction hash that can be used to watch the transaction on the official Decrned Block Explorer.

### Check Network Stats
There are many different commands to check different network stats. Here we will cover sending getinfo to dcrnd and getstakeinfo to dcrnwallet.

To get information about your local dcrnd node, issue the getinfo command to dcrnd using dcrnctl. Enter the following in your shell:
`dcrnctl getinfo`

To get staking information about the Proof-of-Stake network, issue the getstakeinfo command to dcrnwallet using dcrnctl. Enter the following in your shell:
`dcrnctl --wallet getstakeinfo`

### Additional Commands
More commands can also be found on the dcrnctl RPC Commands page.



## dcrnd CLI Arguments

*Last updated for CLI release v1.5.2*

Dcrnd daemon should work with default configuration for most users, however there is a wide variety of command line arguments to change the way it behave if required. For example, the following command can be used to change the log directory dcrnd will write to.
`dcrnd --logdir=/my/custom/log/directory`

`dcrnd` support a `--help` argument which will print details of all the arguments it support and then exit immediately.

#### full list of dcrnd CLI arguments
<span id="dcrndCli"></span>
<details>
<summary>Click to expand</summary>

|  Argument  | Description |
|:-----------|:------------|
|  /A, /appdata:               |Path to application home directory (default:~/.dcrnd)|
|  /V, /version                |Display version information and exit|
|  /C, /configfile:            |Path to configuration file (default:~/.dcrnd/dcrnd.conf)|
|  /b, /datadir:               |Directory to store data (default:~/.dcrnd/data)|
|      /logdir:                |Directory to log output (default:~/.dcrnd/logs)|
|      /nofilelogging          |Disable file logging.|
|  /a, /addpeer:               |Add a peer to connect with at startup|
|      /connect:               |Connect only to the specified peers at startup|
|      /nolisten               |Disable listening for incoming connections -- NOTE: Listening is automatically disabled if the --connect or --proxy options are used without also specifying listen interfaces via --listen|
|      /listen:                |Add an interface/port to listen for connections (default all interfaces port: 9108, testnet:19108)|
|      /maxsameip:             |Max number of connections with the same IP -- 0 to disable (default: 5)|
|      /maxpeers:              |Max number of inbound and outbound peers(default: 125)|
|      /nobanning              |Disable banning of misbehaving peers|
|      /banduration:           |How long to ban misbehaving peers.  Valid time units are {s, m, h}.  Minimum 1 second (default:24h0m0s)|
|      /banthreshold:          |Maximum allowed ban score before disconnecting and banning misbehaving peers. (default: 100)|
|      /whitelist:             |Add an IP network or IP that will not be banned.(eg. 192.168.1.0/24 or ::1)|
|  /u, /rpcuser:               |Username for RPC connections|
|  /P, /rpcpass:               |Password for RPC connections|
|      /rpclimituser:          |Username for limited RPC connections|
|      /rpclimitpass:          |Password for limited RPC connections|
|      /rpclisten:             |Add an interface/port to listen for RPC connections (default port: 9109, testnet: 19109)|
|      /rpccert:               |File containing the certificate file (default: ~/.dcrnd/rpc.cert)|
|      /rpckey:                |File containing the certificate key (default: ~/.dcrnd/rpc.key)|
|      /rpcmaxclients:         |Max number of RPC clients for standard   connections (default: 10)|
|      /rpcmaxwebsockets:      |Max number of RPC websocket connections (default: 25)|
|      /rpcmaxconcurrentreqs:  |Max number of concurrent RPC requests that may be   processed concurrently (default: 20)|
|      /norpc                  |Disable built-in RPC server -- NOTE: The RPC  rver is disabled by default if no rpcuser/rpcpass or rpclimituser/rpclimitpass is specified|
|      /notls                  |Disable TLS for the RPC server -- NOTE: This is only allowed if the RPC server is bound to localhost|
|      /nodnsseed              |Disable DNS seeding for peers|
|      /externalip:            |Add an ip to the list of local addresses we claim to listen on to peers|
|      /proxy:                 |Connect via SOCKS5 proxy (eg. 127.0.0.1:9050)|
|      /proxyuser:             |Username for proxy server|
|      /proxypass:             |Password for proxy server|
|      /onion:                 |Connect to tor hidden services via SOCKS5 proxy (eg. 127.0.0.1:9050)|
|      /onionuser:             |Username for onion proxy server|
|      /onionpass:             |Password for onion proxy server|
|      /noonion                |Disable connecting to tor hidden services|
|      /nodiscoverip           |Disable automatic network address discovery|
|      /torisolation           |Enable Tor stream isolation by randomizing user credentials for each connection.|
|      /testnet                |Use the test network|
|      /simnet                 |Use the simulation test network|
|      /regnet                 |Use the regression test network|
|      /nocheckpoints          |Disable built-in checkpoints.  Don't do this unless you know what you're doing.|
|      /dbtype:                |Database backend to use for the Block Chain(default: ffldb)|
|      /profile:               |Enable HTTP profiling on given [addr:]port --NOTE port must be between 1024 and 65536|
|      /cpuprofile:            |Write CPU profile to the specified file|
|      /memprofile:            |Write mem profile to the specified file|
|      /dumpblockchain:        |Write blockchain as a flat file of blocks for use with addblock, to the specified filename|
|      /miningtimeoffset:      |Offset the mining timestamp of a block by this many seconds (positive values are in the past)|
|  /d, /debuglevel:            |Logging level for all subsystems {trace, debug,info, warn, error, critical} -- You may also specify <subsystem>=<level>,<subsystem2>=<level>,... to set the log level for individual subsystems --Use show to list available subsystems (default:info)|
|      /upnp                   |Use UPnP to map our listening port outside of NAT|
|      /minrelaytxfee:         |The minimum transaction fee in DCR/kB to be considered a non-zero fee. (default: 0.0001)|
|      /limitfreerelay:        |Limit relay of transactions with no transaction fee to the given amount in thousands of bytes per minute (default: 15)|
|      /norelaypriority        |Do not require free or low-fee transactions to have high priority for relaying|
|      /maxorphantx:           |Max number of orphan transactions to keep in memory (default: 1000)|
|      /generate               |Generate (mine) coins using the CPU|
|      /miningaddr:            |Add the specified payment address to the list of addresses to use for generated blocks -- At least one address is required if the generate option is set|
|      /blockminsize:          |Minimum block size in bytes to be used when creating a block|
|      /blockmaxsize:          |Maximum block size in bytes to be used when creating a block (default: 375000)|
|      /blockprioritysize:     |Size in bytes for high-priority/low-fee transactions when creating a block (default:20000)|
|      /sigcachemaxsize:       |The maximum number of entries in the signature verification cache (default: 100000)|
|      /nonaggressive          |Disable mining off of the parent block of the blockchain if there aren't enough voters|
|      /nominingstatesync      |Disable synchronizing the mining state with other nodes|
|      /allowoldvotes          |Enable the addition of very old votes to the mempool|
|      /blocksonly             |Do not accept transactions from remote peers.|
|      /acceptnonstd           |Accept and relay non-standard transactions to the network regardless of the default settings for the active network.|
|      /rejectnonstd           |Reject non-standard transactions regardless of the default settings for the active network.|
|      /txindex                |Maintain a full hash-based transaction index which makes all transactions available via the getrawtransaction RPC|
|      /droptxindex            |Deletes the hash-based transaction index from the database on start up and then exits.|
|      /addrindex              |Maintain a full address-based transaction index which makes the searchrawtransactions RPC available|
|      /dropaddrindex          |Deletes the address-based transaction index from the database on start up and then exits.|
|      /noexistsaddrindex      |Disable the exists address index, which tracks whether or not an address has even been used.|
|      /dropexistsaddrindex    |Deletes the exists address index from the database on start up and then exits.|
|      /nocfilters             |Disable compact filtering (CF) support|
|      /dropcfindex            |Deletes the index used for compact filtering (CF) support from the database on start up and then exits.|
|      /piperx:                |File descriptor of read end pipe to enable parent -> child process communication|
|      /pipetx:                |File descriptor of write end pipe to enable parent <- child process communication|
|      /lifetimeevents         |Send lifetime notifications over the TX pipe|
|      /altdnsnames:           |Specify additional dns names to use when generating the rpc server certificate [%Dcrnd_ALT_DNSNAMES%]|
</details>



## dcrnwallet CLI Arguments

*Last updated for CLI release v1.5.2*

Dcrnwallet daemon should work with default configuration for most users, however there is a wide variety of command line arguments to change the way it behave if required. For example, the following command can be used to change the log directory dcrnwallet will write to.
`dcrnwallet --logdir=/my/custom/log/directory`

`dcrnwallet` support a `--help` argument which will print details of all the arguments it support and then exit immediately.

#### full list of dcrnd CLI arguments
<details>
<summary>Click to expand</summary>

|  Argument  | Description |
|:-----------|:------------|
|  /C, /configfile:                             |Path to configuration file (default: ~/.dcrnwallet/dcrnwallet.conf)|
|  /V, /version                                 |Display version information and exit|
|      /create                                  |Create new wallet|
|      /createtemp                              |Create simulation wallet in nonstandard --appdata; private passphrase is 'password'|
|      /createwatchingonly                      |Create watching wallet from account extended pubkey|
|  /A, /appdata:                                |Application data directory for wallet config, databases and logs (default: ~/.dcrnwallet)|
|      /testnet                                 |Use the test network|
|      /simnet                                  |Use the simulation test network|
|      /noinitialload                           |Defer wallet creation/opening on startup and enable loading wallets over RPC|
|  /d, /debuglevel:                             |Logging level {trace, debug, info, warn, error, critical} (default: info)|
|      /logdir:                                 |Directory to log output. (default: ~/.dcrnwallet/logs)|
|      /profile:                                |Enable HTTP profiling this interface/port|
|      /memprofile:                             |Write mem profile to the specified file|
|      /walletpass:                             |Public wallet password; required when created with one|
|      /promptpass                              |Prompt for private passphase from terminal and unlock without timeout|
|      /pass:                                   |Unlock with private passphrase|
|      /promptpublicpass                        |Prompt for public passphras from terminal|
|      /enableticketbuyer                       |Enable the automatic ticket buyer|
|      /enablevoting                            |Automatically create votes and revocations|
|      /purchaseaccount:                        |Account to autobuy tickets from (default: default)|
|      /pooladdress:                            |VSP fee address|
|      /poolfees:                               |VSP fee percentage (1.00 equals 1.00% fee)|
|      /gaplimit:                               |Allowed unused address gap between used addresses of accounts (default: 20)|
|      /stakepoolcoldextkey:                    |xpub:maxindex for fee addresses (VSP-only option)|
|      /allowhighfees                           |Do not perform high fee checks|
|      /txfee:                                  |Transaction fee per kilobyte (default: 0.0001 DCR)|
|      /accountgaplimit:                        |Allowed gap of unused accounts (default: 10)|
|      /disablecointypeupgrades                 |Never upgrade from legacy to SLIP0044 coin type keys|
|  /c, /rpcconnect:                             |Network address of dcrnd RPC server|
|      /cafile:                                 |dcrnd RPC Certificate Authority|
|      /noclienttls                             |Disable TLS for dcrnd RPC; only allowed when connecting to localhost|
|      /dcrndusername:                           |dcrnd RPC username; overrides --username|
|      /dcrndpassword:                           |dcrnd RPC password; overrides --password|
|      /proxy:                                  |Establish network connections and DNS lookups through a SOCKS5 proxy (e.g. 127.0.0.1:9050)|
|      /proxyuser:                              |Proxy server username|
|      /proxypass:                              |Proxy server password|
|      /circuitlimit:                           |Set maximum number of open Tor circuits; used only when --torisolation is enabled (default: 32)|
|      /torisolation                            |Enable Tor stream isolation by randomizing user credentials for each connection|
|      /nodcrndproxy                             |Never use configured proxy to dial dcrnd websocket connectons|
|      /spv                                     |Sync using simplified payment verification|
|      /spvconnect:                             |SPV sync only with specified peers; disables DNS seeding|
|      /rpccert:                                |RPC server TLS certificate (default: ~/.dcrnwallet/rpc.cert)|
|      /rpckey:                                 |RPC server TLS key (default: ~/.dcrnwallet/rpc.key)|
|      /tlscurve:                               |Curve to use when generating TLS keypairs (default: Ed25519)|
|      /onetimetlskey                           |Generate self-signed TLS keypairs each startup; only write certificate file|
|      /noservertls                             |Disable TLS for the RPC servers; only allowed when binding to localhost|
|      /grpclisten:                             |Listen for gRPC connections on this interface|
|      /rpclisten:                              |Listen for JSON-RPC connections on this interface|
|      /nogrpc                                  |Disable gRPC server|
|      /nolegacyrpc                             |Disable JSON-RPC server|
|      /rpcmaxclients:                          |Max JSON-RPC HTTP POST clients (default: 10)|
|      /rpcmaxwebsockets:                       |Max JSON-RPC websocket clients (default: 25)|
|  /u, /username:                               |JSON-RPC username and default dcrnd RPC username|
|  /P, /password:                               |JSON-RPC password and default dcrnd RPC password|
|      /pipetx:                                 |File descriptor or handle of write end pipe to enable child -> parent process communication /piperx: File descriptor or handle of read end pipe to enable parent -> child process communication|
|      /rpclistenerevents                       |Notify JSON-RPC and gRPC listener addresses over the TX pipe|
|      /csppserver:                             |Network address of CoinShuffle++ server|
|      /csppserver.ca:                          |CoinShuffle++ Certificate Authority|
|      /mixedaccount:                           |Account/branch used to derive CoinShuffle++ mixed outputs and voting rewards|
|      /ticketsplitaccount:                     |Account to derive fresh addresses from for mixed ticket splits; uses mixedaccount if unset|
|      /changeaccount:                          |Account used to derive unmixed CoinJoin outputs in CoinShuffle++ protocol|
|      /mixchange                               |Use CoinShuffle++ to mix change account outputs into mix account|
|      /reuseaddresses                          |DEPRECATED -- Reuse addresses for ticket purchase to cut down on address overuse|
|      /disallowfree                            |DEPRECATED -- Force transactions to always include a fee Ticket Buyer Options:|
|      /ticketbuyer.balancetomaintainabsolute:  |Amount of funds to keep in wallet when purchasing tickets (default: 0 DCR)|
|      /ticketbuyer.votingaddress:              |Purchase tickets with voting rights assigned to this address|
|      /ticketbuyer.limit:                      |Buy no more than specified number of tickets per block (0 disables limit)|
|      /ticketbuyer.votingaccount:              |Account used to derive addresses specifying voting rights|
</details>



## dcrnctl RPC Commands
<span id="ctlCommand"></span>
The dcrnd and dcrnwallet daemons have APIs that can be used to access lower-level functionality not available in their respective Command-line Interfaces (CLIs). These APIs are called using Remote Procedure Calls (RPCs). RPCs also allow for integrations with clients written in any language that supports RPCs.

`dcrnctl` support a `--l` argument which will print details of all the RPC commands it support and then exit immediately.

#### dcrnd RPC Commands
<details>
<summary>Click to expand</summary>

|  RPC Method  |    Params   |
|:-------------|:------------|
|addnode |"addr" "add remove onetry"|
|createrawssrtx |[{"amount":n.nnn,"txid":"value","vout":n,"tree":n}] (fee)|
|createrawsstx |[{"txid":"value","vout":n,"tree":n,"amt":n},...] {"address":amount} [{"addr":"value","commitamt":n,"changeaddr":"value","changeamt":n},...]|
|createrawtransaction |[{"amount":n.nnn,"txid":"value","vout":n,"tree":n},...] {"address":amount,...} (locktime expiry)|
|debuglevel |"levelspec"|
|decoderawtransaction |"hextx"|
|decodescript |"hexscript" |(version)|
|estimatefee |numblocks|
|estimatesmartfee |confirmations |(mode="conservative")|
|estimatestakediff |(tickets)|
|existsaddress |"address"|
|existsaddresses |["address",...]|
|existsexpiredtickets |["txhash",...]|
|existsliveticket |"txhash"|
|existslivetickets |["txhash",...]|
|existsmempooltxs |["txhash",...]|
|existsmissedtickets |["txhash",...]|
|generate |numblocks|
|getaddednodeinfo |dns ("node")|
|getbestblock ||
|getbestblockhash ||
|getblock |"hash" (verbose=true verbosetx=false)|
|getblockchaininfo ||
|getblockcount ||
|getblockhash |index|
|getblockheader |"hash" (verbose=true)|
|getblocksubsidy |height voters|
|getcfilterv2 |"blockhash"|
|getchaintips ||
|getcoinsupply ||
|getconnectioncount ||
|getcurrentnet ||
|getdifficulty ||
|getgenerate ||
|gethashespersec ||
|getheaders |["blocklocator",...] "hashstop"|
|getinfo ||
|getmempoolinfo ||
|getmininginfo ||
|getnettotals ||
|getnetworkhashps |(blocks=120 height=-1)|
|getnetworkinfo ||
|getpeerinfo ||
|getrawmempool |(verbose=false "txtype")|
|getrawtransaction |"txid" (verbose=0)|
|getstakedifficulty ||
|getstakeversioninfo |(count)|
|getstakeversions |"hash" count|
|getticketpoolvalue ||
|gettreasurybalance |("hash" verbose=false)|
|gettreasuryspendvotes |("block" ["tspend",...])|
|gettxout |"txid" vout tree (includemempool=true)|
|gettxoutsetinfo ||
|getvoteinfo |version|
|getwork |("data")|
|help |("command")|
|invalidateblock |"blockhash"|
|livetickets ||
|missedtickets ||
|node |"connect|remove|disconnect" "target" ("perm|temp")|
|ping ||
|reconsiderblock |"blockhash"|
|regentemplate ||
|searchrawtransactions |"address" (verbose=1 skip=0 count=100 vinextra=0 reverse=false ["filteraddr",...])|
|sendrawtransaction |"hextx" (allowhighfees=false)|
|setgenerate |generate (genproclimit=-1)|
|stop ||
|submitblock |"hexblock" ({"workid":"value"})|
|ticketfeeinfo |(blocks windows)|
|ticketsforaddress |"address"|
|ticketvwap |(start end)|
|txfeeinfo |(blocks rangestart rangeend)|
|validateaddress |"address"|
|verifychain |(checklevel=3 checkdepth=288)|
|verifymessage |"address" "signature" "message"|
|version ||
</details>


#### dcrnwallet RPC Commands
<details>
<summary>Click to expand</summary>

|  RPC Method  |    Params   |
|:-------------|:------------|
|abandontransaction |"hash"|
|accountaddressindex |"account" branch|
|accountsyncaddressindex |"account" branch index|
|accountunlocked |"account"|
|addmultisigaddress |nrequired ["key",...] ("account")|
|addtransaction |"blockhash" "transaction"|
|auditreuse |(since)|
|consolidate |inputs ("account" "address")|
|createmultisig |nrequired ["key",...]|
|createnewaccount |"account"|
|createrawtransaction |[{"amount":n.nnn,"txid":"value","vout":n,"tree":n},...] {"address":amount,...} (locktime expiry)|
|createsignature |"address" inputindex hashtype "previouspkscript" "serializedtransaction"|
|createvotingaccount |"name" "pubkey" (childindex=0)|
|disapprovepercent ||
|discoverusage |("startblock" discoveraccounts gaplimit)|
|dumpprivkey |"address"|
|fundrawtransaction |"hexstring" "fundaccount" ({"changeaddress":changeaddress,"feerate":feerate,"conftarget":conftarget})|
|generatevote |"blockhash" height "tickethash" votebits |"votebitsext"|
|getaccount |"address"|
|getaccountaddress |"account"|
|getaddressesbyaccount |"account"|
|getbalance |("account" minconf=1)|
|getbestblock ||
|getbestblockhash ||
|getblock |"hash" (verbose=true verbosetx=false)|
|getblockcount ||
|getblockhash |index|
|getblockheader |"hash" (verbose=true)|
|getcfilterv2 |"blockhash"|
|getcoinjoinsbyacct ||
|getcurrentnet ||
|getinfo ||
|getmasterpubkey |("account")|
|getmultisigoutinfo |"hash" index|
|getnewaddress |("account" "gappolicy")|
|getpeerinfo ||
|getrawchangeaddress |("account")|
|getreceivedbyaccount |"account" (minconf=1)|
|getreceivedbyaddress |"address" (minconf=1)|
|getstakeinfo ||
|gettickets |includeimmature|
|gettransaction |"txid" (includewatchonly=false)|
|gettxout |"txid" vout tree (includemempool=true)|
|getunconfirmedbalance |("account")|
|getvotechoices |("tickethash")|
|getwalletfee ||
|help |("command")|
|importcfiltersv2 |startheight ["filter",...]|
|importprivkey |"privkey" ("label" rescan=true scanfrom)|
|importscript |"hex" (rescan=true scanfrom)|
|importxpub |"name" "xpub"|
|listaccounts |(minconf=1)|
|listaddresstransactions |["address",...] ("account")|
|listalltransactions |("account")|
|listlockunspent |("account")|
|listreceivedbyaccount |(minconf=1 includeempty=false includewatchonly=false)|
|listreceivedbyaddress |(minconf=1 includeempty=false includewatchonly=false)|
|listsinceblock |("blockhash" targetconfirmations=1 includewatchonly=false)|
|listtransactions |("account" count=10 from=0 includewatchonly=false)|
|listunspent |(minconf=1 maxconf=9999999 ["address",...] "account")|
|lockaccount |"account"|
|lockunspent unlock |[{"amount":n.nnn,"txid":"value","vout":n,"tree":n},...]|
|mixaccount ||
|mixoutput |"outpoint"|
|processunmanagedticket |("tickethash")|
|purchaseticket |"fromaccount" spendlimit (minconf=1 "ticketaddress" numtickets=1 "pooladdress" poolfees expiry "comment" dontsigntx)|
|redeemmultisigout |"hash" index tree ("address")|
|redeemmultisigouts |"fromscraddress" ("toaddress" number)|
|renameaccount |"oldaccount" "newaccount"|
|rescanwallet |(beginheight=0)|
|revoketickets ||
|sendfrom |"fromaccount" "toaddress" amount (minconf=1 "comment" "commentto")|
|sendfromtreasury |"key" amounts|
|sendmany |"fromaccount" {"address":amount,...} (minconf=1 "comment")|
|sendrawtransaction |"hextx" (allowhighfees=false)|
|sendtoaddress |"address" amount ("comment" "commentto")|
|sendtomultisig |"fromaccount" amount ["pubkey",...] (nrequired=1 minconf=1 "comment")|
|sendtotreasury |amount|
|setaccountpassphrase |"account" "passphrase"|
|setdisapprovepercent percent|
|settreasurypolicy |"key" "policy" ("ticket")|
|settspendpolicy |"hash" "policy" ("ticket")|
|settxfee |amount|
|setvotechoice |"agendaid" "choiceid" ("tickethash")|
|signmessage |"address" "message"|
|signrawtransaction |"rawtx" ([{"txid":"value","vout":n,"tree":n,"scriptpubkey":"value","redeemscript":"value"},...] ["privkey",...] flags="ALL")|
|signrawtransactions |["rawtx",...] (send=true)|
|stakepooluserinfo |"user"|
|sweepaccount |"sourceaccount" "destinationaddress" (requiredconfirmations feeperkb)|
|syncstatus ||
|ticketinfo |(startheight=0)|
|ticketsforaddress |"address"|
|treasurypolicy |("key" "ticket")|
|tspendpolicy |("hash" "ticket")|
|unlockaccount |"account" "passphrase"|
|validateaddress |"address"|
|validatepredcp0005cf ||
|verifymessage |"address" "signature" "message"|
|version ||
|walletinfo ||
|walletislocked ||
|walletlock ||
|walletpassphrase |"passphrase" timeout|
|walletpassphrasechange |"oldpassphrase" "newpassphrase"|
|walletpubpassphrasechange |"oldpassphrase" "newpassphrase"|
</details>
