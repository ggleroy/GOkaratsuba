# GOkaratsuba  
Karatsuba algorithm using Golang!  
A **Foundations of Algorithm Design and Analysis** assignment.  

## About  
The **Karatsuba algorithm** is a fast multiplication method using **divide-and-conquer**, reducing time complexity from **O(n²)** to **O(n^1.585)**. It splits numbers into halves, computes three intermediate products, and efficiently combines them, minimizing multiplications. This makes it ideal for **large integer arithmetic, cryptography, and computer algebra systems**.  

## Structure  
This repository consists of a single **Go** file: `main.go`.  

## How to Run  
Ensure **Go** is installed on your machine, then execute:  

```sh
go run main.go
```

## Cyclomatic Complexity  
Using the formula **𝑀 = 𝐸 − 𝑁 + 2𝑃**, where:  
- **𝐸** (edges) = 13  
- **𝑁** (nodes) = 11  
- **𝑃** (connected components) = 1  

The **cyclomatic complexity** of the Karatsuba algorithm is **4**.  

## Asymptotic Complexity  

- **Best case**: **O(1)**, when one of the numbers has only one digit.  
- **Average/Worst case**: **O(n^log(3)) ≈ O(n^1.585)**, where **n** is the number of digits in the input numbers.  
- **Space complexity**: **O(n^log(3))**  

