# Voting Service Provider (VSP)

**Non-custodial services that can be authorized to vote on behalf of a ticket, usually providing a number of geographically distributed servers to reduce the chance of missed tickets.**

**Using a VSP does not give the VSP access to your funds. All you are doing is granting voting rights to the VSP.**

In order to support network decentralization, it is recommended that you join a smaller VSP with fewer live tickets. As VSPs control tickets delegated to them, they could in theory vote those tickets in a way which contradicts the expressed wishes of the ticket owners. This could be easily detected, and any VSP which attempted it would likely be abandoned by the stakeholder community. However, it is good practice to limit the power of individual VSPs, to limit the potential for damage from this kind of attack.

Unlike Proof-of-Work (PoW) mining pools, VSPs do not pool work or rewards. The number of tickets a VSP has does not affect how regularly one’s tickets will be called and rewards received.

All VSPs run the same basic code, but they may differ in the amount of redundancy available. More redundancy equals less chance of missed votes, although all VSPs will have some missed votes as many missed votes are caused by PoW miners. This is because sometimes miners will find a solution to the next block so quickly that votes haven’t had time to propagate around the network.

PoS voting using a VSP can be done using Decredniton or dcrnwallet.

* Decredniton - GUI wallet for Windows/macOS/Linux. The Tickets menu, Purchase tab is used to buy tickets.
* dcrnwallet - CLI wallet for Windows/macOS/Linux. The Buying Tickets with dcrnwallet guide explains how to purchase tickets via command line.


## Frequently asked questions
---
1. **What is a Voting Service Provider?**
A Voting Service Provider (VSP) is similar in some ways to a PoW mining pool, but for the PoS system. Through the options in your Decred-Next wallet, you can share your voting rights with a VSP. If your ticket is selected to vote, the VSP will cast the vote for you and you are rewarded with the PoS reward. Unlike mining pools, the PoS reward is not split amongst the users of the VSP. The full reward goes to the owner of the specific ticket that voted.
A VSP allows you to buy tickets and vote without the requirement of maintaining an online and unlocked wallet. It is important to note that your funds never leave your wallet. You are not sending anything to the VSP, just giving it authority to vote on your behalf. A VSP cannot access your funds.
VSPs will usually implement multi-wallet redundancy by having many wallets physically distributed around the globe. This means there’s less chance of a vote being missed because one wallet is down. It also reduces latency between the wallet and network which can reduce the chance of a vote being missed.
---
2. **Are there any other benefits to a Voting Service Provider (VSP) other than not needing to run a full node and keeping your wallet unlocked? For example, will it have a better chance of voting, or vote quicker?**
A VSP cannot increase a tickets chance of voting, or cause it to vote faster. Ticket selection is a random process which cannot be influenced.
VSPs will usually implement multi-wallet redundancy by having many wallets physically distributed around the globe. This means there’s less chance of a vote being missed because one wallet is down. It also reduces latency between the wallet and network which can reduce the chance of a vote being missed.
---
3. **Does a Voting Service Provider (VSP) split the reward between all participants (% based on the amount of tickets you submitted to the VSP)?**
While it is technically possible to create a VSP that supports this type of proportional reward splitting, the current VSP reference implementation doesn’t enable this.
---
4. **I have to run my wallet to buy tickets, but will they vote properly without me if I shut it down and the Voting Service Provider (VSP) votes for me instead?**
Yes, that is correct. You only need to run your wallet in order to spend your coins to purchase the ticket which delegates your voting rights to the VSP which will then vote on your behalf. The reward address is a consensus-enforced commitment in the ticket purchase for one of your own addresses for which only you have the private key.
---
5. **Are there any issues that could arise from a Voting Service Provider (VSP) having too many people. For example, force voting a block in or out?**
It is certainly possible, but one of the things that all VSPs should support is allowing each user to select their individual voting preferences. That way, whenever their ticket comes up and the VSP votes on their behalf, it will vote according to their preferences.
---
6. **What safeguards are in place to stop Voting Service Provider (VSP) owners disappearing with the funds in the tickets they vote for?**
The current VSP design is such that the VSP can NOT steal the funds. You are only sharing the ability to vote on tickets, and not the ability to spend any of your funds. The ticket purchase contains a consensus-enforced commitment for the final subsidy address, so there is simply no way for the VSP to steal the funds.
The worst that would happen if a VSP owner disappears is the votes will be missed which results in the ticket being revoked, which in turn causes the original coins to go back to the original coin owner (minus the initial transaction fee). However, it is possible for the ticket owner to run their own wallet in order to vote should the VSP owner disappear.
---
7. **What happens if my Voting Service Provider (VSP) goes down?**
In the unlikely scenario your VSP goes down permanently, you can still vote your own tickets, as well as revoke any missed or expired tickets.
---

## Recommended VSP list

* https://dcrnvsp.com/
* https://stakedcrn.net/
  
---

## Installation

dcrstakepool is a web application which coordinates generating 1-of-2 multisig
addresses on a pool of [dcrnwallet](https://github.com/Decred-Next/dcrnwallet) servers
so users can purchase proof-of-stake tickets
on the Decred-Next network and have the pool of wallet servers
vote on their behalf when the ticket is selected.

**NB:** a proposal
was approved by stakeholders to rename "Stakepool" to "Voting Service Provider", a.k.a. "VSP".
These names are used interchangably in this repository.


### Requirements

- [Go](https://golang.org) 1.12 or newer (1.13 is recommended).
- MySQL
- Nginx or other web server to proxy to dcrstakepool

### Build from source

Building or updating from source requires the following build dependencies:

- **Go 1.12 or 1.13**

Building or updating from source requires only an installation of Go
([instructions](https://golang.org/doc/install)). It is recommended to add
`$GOPATH/bin` to your `PATH` at this point.

To build and install from a checked-out repo, run `go install . ./backend/stakepoold`
in the repo's root directory.

* Set the `GO111MODULE=on` environment variable if building from within
  `GOPATH`.

### Pre-requisites

These instructions assume you are familiar with dcrnd/dcrwallet.

- Create basic dcrnd/dcrwallet/dcrctl config files with usernames, passwords,
  rpclisten, and network set appropriately within them or run example commands
  with additional flags as necessary.

- Build/install dcrnd and dcrwallet from latest master.

- Run dcrnd instances and let them fully sync.


### Voting service fees/cold wallet

- Setup a new wallet for receiving payment for voting service fees.  **This should
  be completely separate from the voting service infrastructure.**
- From your local machine...

```bash
$ dcrwallet --create
$ dcrwallet
```

- Get the master pubkey for the account you wish to use. This will be needed to
  configure dcrwallet and dcrstakepool.

```bash
$ dcrctl --wallet createnewaccount stakepoolfees
$ dcrctl --wallet getmasterpubkey stakepoolfees
```

- Mark 10000 addresses in use for the account so the wallet will recognize
  transactions to those addresses. Fees from UserId 1 will go to address 1,
  UserId 2 to address 2, and so on.

```bash
$ dcrctl --wallet accountsyncaddressindex teststakepoolfees 0 10000
```

### Voting service voting wallets

- Create the wallets.  All wallets should have the same seed.  **Backup the seed
  for disaster recovery!**
- Log into wallet servers separately and create wallets one at a time using the
  same seed.

```bash
$ ssh walletserver1
$ dcrwallet --create
```

- Start a properly configured dcrwallet and unlock it. See
  sample-dcrwallet.conf.
- From your local machine...

```bash
$ cp sample-dcrwallet.conf dcrwallet.conf
$ vim dcrwallet.conf
$ scp dcrwallet.conf walletserver1:~/.dcrwallet/
$ ssh walletserver1
$ dcrwallet
```

- Get the master pubkey from the default account.  This will be used for
  votingwalletextpub in dcrstakepool.conf.

```bash
$ ssh walletserver1
$ dcrctl --wallet getmasterpubkey default
```

### MySQL

- Log into your frontend
- Install, configure, and start MySQL
- Add stakepool user and create the stakepool database

```bash
$ ssh frontendserver
$ mysql -uroot -p

MySQL> CREATE USER 'stakepool'@'localhost' IDENTIFIED BY 'password';
MySQL> GRANT ALL PRIVILEGES ON *.* TO 'stakepool'@'localhost' WITH GRANT OPTION;
MySQL> FLUSH PRIVILEGES;
MySQL> CREATE DATABASE stakepool;
```

### stakepoold setup

- Copy sample config and edit appropriately.
- From your local machine...

```bash
$ mkdir .stakepoold
$ cp sample-stakepoold.conf .stakepoold/stakepoold.conf
$ vim .stakepoold/stakepoold.conf
$ scp -r .stakepoold walletserver1:~/
$ scp -r .stakepoold walletserver2:~/
```

- Build and copy the stakepoold executable to each wallet server.
- From your local machine...

```bash
$ cd backend/stakepoold/
$ go build
$ scp stakepoold walletserver1:~/
$ scp stakepoold walletserver2:~/
```

### dcrnstakepool setup

- Create the .dcrstakepool directory and copy dcrwallet certs to it:

```bash
$ ssh frontendserver
$ mkdir ~/.dcrstakepool
$ cd ~/.dcrstakepool
$ scp walletserver1:~/.dcrwallet/rpc.cert wallet1.cert
$ scp walletserver2:~/.dcrwallet/rpc.cert wallet2.cert
$ scp walletserver1:~/.stakepoold/rpc.cert stakepoold1.cert
$ scp walletserver2:~/.stakepoold/rpc.cert stakepoold2.cert
```

- Copy sample config and edit appropriately.
- From your local machine...

```bash
$ cp sample-dcrstakepool.conf dcrstakepool.conf
$ vim dcrstakepool.conf
$ scp dcrstakepool.conf frontendserver:~/.dcrstakepool/
```
- Build and copy the entire dcrnstakepool folder to your frontend.
- From your local machine...

```bash
$ go build
$ scp -r ../dcrnstakepool frontendserver:~/
```

## Running

### stakepoold

Log into all servers and run stakepoold one at a time.

```bash
$ ssh walletserver1
$ ./stakepoold
```

### dcrnstakepool

Log into your frontend and run dcrstakepool

```bash
$ ssh frontendserver
$ cd dcrnstakepool
$ ./dcrstakepool
```
To run `dcrstakepool` from another folder, such as `/opt/dcrstakepool`, it is
necessary to copy (1) the `dcrstakepool` executable generated by `go build`, (2)
the `public` folder, and (3) the `views` folder into the other folder.

By default, `dcrnstakepool` looks for the `public` and `views` folders in the
same parent directory as the `dcrstakepool` executable. If you wish to run
dcrstakepool from a different directory you will need to change **publicpath**
and **templatepath** from their relative paths to an absolute path.