# Day-2-Matrix-chain-multiplication
---

Matrix multiplication is an associative operation. (AB)C is equal to A(BC) for matrices A, B, C, and it doesn’t matter which order of pair multiplication you choose.
<br />
Unfortunately, that’s not true from computational perspective. If dimensions of matrices are A=[10, 20], B=[20, 30], C=[30, 40], numbers of scalar multiplications differ significantly:
<br />
(AB)C = 10*20*30 + 10*30*40 = 18000
A(BC) = 20*30*40 + 10*20*40 = 32000
<br />
The best ordering can be found using recursive relationship. Let MCM denotes a function that returns a minimum number of scalar multiplications. Then MCM can be defined as the best split among all possible choices.