# FHERMA Parity Challenge


This challenge was developed by [IBM Research](https://research.ibm.com). 

The objective of the challenge is to design an algorithm that evaluates the function `parity(x)` under CKKS.

The function `parity(x)` gets an integer and returns its least significant bit (LSB). In other words,

$$parity(x) = x \mod 2,$$

where $x \in \mathbb{Z}.$

The `parity` function is closely related to the bit extraction problem, where given an integer $x$ the goal is to find its bit representation $x =\sum 2^i b_i$ which is useful in many cases, e.g., comparisons.
An efficient implementation to `parity(x)` would lead to an efficient implementation of bit extraction.

Evaluating this function under FHE is not trivial because it is not polynomial. The approximated nature of CKKS adds an additional layer of complexity as the noise levels need to be accounted for.