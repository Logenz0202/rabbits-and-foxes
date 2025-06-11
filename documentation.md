# Documentation

## 1. Introduction

Brief description, project structure, and system requirements are presented [here](README.md).

## 2. Key Algorithms and Mechanisms

### 2.1. Energy System
- Each animal has an energy level that decreases over time
- Passive energy consumption increases with the animal's age
- Rabbits regain energy by eating grass
- Foxes regain energy by hunting rabbits
- The energy system creates natural pressure on animals to search for food

### 2.2. Reproduction Mechanism
- Animals can reproduce only when:
    - They have sufficient energy level
    - The regeneration period after previous reproduction has passed
    - They are in direct proximity to another individual of the same species
- The regeneration period prevents excessive population explosion

### 2.3. Movement System
- Animals move randomly within adjacent fields
- Implementation considers collisions - animals cannot occupy the same field
- Map boundary system ensures animals remain within the simulation area

### 2.4. Grass Growth and Spreading
- Grass grows gradually to maximum value
- Grass spreads to neighboring fields with specific probability
- The mechanism ensures regeneration of food resources for rabbits

## 3. Observations and Conclusions

### 3.1. Population Dynamics

- Cyclic fluctuations in population numbers were observed
- Typical cycle:
    1. Rabbit population growth when grass is available
    2. Fox population growth in response to high rabbit numbers
    3. Decline in rabbit population due to predation
    4. Decline in fox population due to lack of food
- Most common end cases:
    - Foxes eat all rabbits and die of starvation
    - One rabbit remains running until it dies of old age

### 3.2. Interesting Phenomena
- Formation of "hot spots" - areas of increased activity
- Periodic population extinction under extreme parameters
- System self-regulation with appropriate parameter values (haven't managed to find these parameters yet)

### 3.3. Critical Points
- Too small initial rabbit population can lead to fox extinction
- Too large fox population can lead to collapse of entire ecosystem

## 4. Evaluation and Conclusions

### 4.1. Simulation Strengths
- Realistic representation of basic ecosystem mechanisms
- Data and graph enable easy observation

### 4.2. Potential Extensions
- Implementation of more advanced animal behaviors (e.g., prey tracking by predator)
