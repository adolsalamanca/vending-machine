[![Actions Status](https://github.com/adolsalamanca/vending-machine/workflows/Go/badge.svg)](https://github.com/adolsalamanca/vending-machine/actions)

# Vending Machine Coding Kata

## Description
The goal of this program is to model a vending machine and the state it must maintain during its operation.

The machine works like all vending machines: it takes money then gives you items. The vending machine accepts money in the form of 0.05, 0.10, 0.25 and 1

It should have at least have 3 primary items that cost 0.65, 1.00, and 1.50. Also user may hit the button “return coin” to get back the money they’ve entered so far, If you put more money in than the item price, you get the item and change back.

## Specification
### Valid set of actions on the vending machine are:

 | Type of action               |               Command |
 |:-----------------------------|----------------------:|
 | Insert Coins                 |  0.05, 0.10, 0.25, 1  |
 | Return Coins                 |  Return-coin          |
 | Buy Item                     |  Get-Soda             |
 | Service/Reload stock&items   |  Service items        |

## Real examples 
```
Example 1: Buy Soda with exact change
1, 0.25, 0.25, GET-SODA
-> SODA

Example 2: Start adding money, but user ask for return coin
0.10, 0.10, RETURN-COIN
-> 0.10, 0.10

Example 3: Buy Water without exact change
1, GET-WATER
-> WATER, 0.25, 0.10
```

## Implementation details
There are a few entities identified to fulfill the problem logic:
  * Coin: it represent a unit of money, it contains a category or type of coin and also a value using decimal.Decimal type not to lose precission.
  * Item: represents a type of item, it contains selector as an identifier, name, and price.[adolsalamanca](https://github.com/adolsalamanca)
  * ItemCatalogue: it is a type that holds an array of Items, and those are all the types of items that could be stored in the machine.
  * CashEngine: it contains the data to perform money management (coinsForExchange, validCoins, currentServiceCoins, currentServiceBalance)and also an important part of the 	behaviour.
  * Machine: the main entity of the program, it contains a CashEngine, ItemCatalogue, ItemStock and Logger. It is the responsible of interpret the commands that are coming from the main and execute the expected logic in the CashEngine making use of all the required entities.

I have no problem with standard go testing but as you can see, I have used Ginkgo to develop tests, it is a BDD style testing framework, I believe it is very comfortable to work with but also really powerfull.


The proposed solution uses original kata specifications in terms of products (Water,Juice and Soda) and also regarding accepted coins (0.05, 0.10, 0.25 and 1) but it is prepared to easily add new items or more coin types in initializeMachine function.
Most of the problem has been developed using TDD, making red, green, refactor cycles.


## Tests
- [x] Coin
- [x] Item
- [x] Engine
- [x] Machine



## How to run tests

Just decompress the file into the desired folder, enter the root folder of the project (where go.mod file is), and run this command in your terminal:
``
go test ./...
``

## Build & Run the simulation

Again, from the root folder of the project and run the following command in your terminal:
``
go run cmd/vending-machine/main.go
``

You will see all the reports of Drones visiting around Stations and also a Shutdown message.

## Additional info

You can follow the different phases of build&test of the application development [here](https://github.com/adolsalamanca/vending-machine/actions)

## Author

* **Adolfo Rodriguez** - *vending-machine* - [adolsalamanca](https://github.com/adolsalamanca)

 

