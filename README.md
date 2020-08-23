[![Actions Status](https://github.com/adolsalamanca/vending-machine/workflows/Go/badge.svg)](https://github.com/adolsalamanca/vending-machine/actions)

## Vending Machine Coding Kata

The goal of this program is to model a vending machine and the state it must maintain during its operation.

The machine works like all vending machines: it takes money then gives you items. The vending machine accepts money in the form of 0.05, 0.10, 0.25 and 1

It should have at least have 3 primary items that cost 0.65, 1.00, and 1.50. Also user may hit the button “return coin” to get back the money they’ve entered so far, If you put more money in than the item price, you get the item and change back.

## Specification

### Valid set of actions on the vending machine are:

* Insert Money  -  0.05, 0.10, 0.25, 1
* Return Coins -  returns all inserted money
* GET Item     -  GET Water, GET Juice, GET Soda - order item (Water = 0.65, Juice = 1.00, Soda = 1.50)
* SERVICE      -  a service person opens the machine and set the available change and how many items we have.

### Valid set of responses on the vending machine are:

* Return Coin - 0.05, 0.10, 0.25
* Sold Item   - Water,  Juice, Soda - vend item

### Vending machine must track the following state:

* Available items         - each item has a count, a price and selector
* Available change        - amount os each coins available
* Balance                 - currently inserted money
 
## Examples 
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




