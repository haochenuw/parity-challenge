# FHERMA Parity Challenge


This challenge was developed by [IBM Research](https://research.ibm.com). 

The objective of the challenge is to design an algorithm that evaluates the function `parity(x)` under CKKS.

The function `parity(x)` gets an integer and returns its least significant bit (LSB). In other words,

$$parity(x) = x \mod 2,$$

where $x \in \mathbb{Z}.$

The `parity` function is closely related to the bit extraction problem, where given an integer $x$ the goal is to find its bit representation $x =\sum 2^i b_i$ which is useful in many cases, e.g., comparisons.
An efficient implementation to `parity(x)` would lead to an efficient implementation of bit extraction.

Evaluating this function under FHE is not trivial because it is not polynomial. The approximated nature of CKKS adds an additional layer of complexity as the noise levels need to be accounted for.

## Content
* `openfhe` - template for c++ openfhe based solution
* `openfhe-python` - template fo openfhe-python based solution
* `lattigo` - template for lattigo based solution
* `testcase.json` - testcase example(could be used with docker validator)

### How to validate solution locally
Once solution is developed a participants could use a docker image to validate their solution locally.
Put your solution with the test case json file in the local directory, link directory to the docker container and specify path to the project and test case json(keep in mind that path in arguments should be a path inside a docker container)

Example: local folder with the solution is `~/user/tmp/parity/app`
```sh
$ docker run -ti -v ~/user/tmp/parity:/parity yashalabinc/fherma-validator --project-folder=/parity/app --testcase=/parity/testcase.json
```

Once validation is completed you'll see `result.json` file in the project folder. This file is exactly the same as we used to score uploaded solutions on the FHERMA platform
