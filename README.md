# E-Pizza Simulator

This project simulates a real Italian Pizzeria. 
Users can: 
- Order pizzas
- Pay orders
- Track how much time is left for their order
- Remove or add ingredients from the pizzas.
- When a User Logs In can see their previous orders

## Microservices
- Chef --> it receives orders and delivers finished pizza, it is only triggered when There are Pizzas in the queue
- Client --> can place orders and pays trough server APIs
- PizzaRestaurant --> Exposes the Api to Order pizza and pay, it sends the order to the kitchen and emails the customer when the pizza is ready


## Table of Contents

- [Installation](#installation)
- [Usage](#usage)

## Installation

Explain how to install the project, including any dependencies. Provide code examples if necessary.

## Usage

Provide examples and explanations of how to use the project. Include screenshots or code snippets if applicable.