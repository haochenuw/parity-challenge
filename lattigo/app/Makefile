# Flags default values
# sk = artifacts/sk.bin
cc = artefacts/context
key_public = artefacts/public_key
key_eval = artefacts/relinearization_key
input = artefacts/ciphertext_0_0
output = artefacts/out.bin

my_cc = temps/context
my_sk = temps/sk.bin
my_key_public = temps/pk.bin
my_key_eval = temps/evk.bin
my_input = temps/in.bin
my_output = temps/out.bin

test-all:
	go run setup.go --sk=$(my_sk) --cc=$(my_cc) --key_public=$(my_key_public) --key_eval=$(my_key_eval) --input=$(my_input)
	go run main.go --cc=$(my_cc) --key_public=$(my_key_public) --key_eval=$(my_key_eval) --input=$(my_input) --output=$(my_output)
	go run verify.go --sk=$(my_sk) --cc=$(my_cc) --input=$(my_input) --output=$(my_output)
	go run clean.go
	go clean

main:
	go run main.go --cc=$(cc) --key_public=$(key_public) --key_eval=$(key_eval) --input=$(input) --output=$(output)

setup:
	go run setup.go --sk=$(sk) --cc=$(cc) --key_public=$(key_public) --key_eval=$(key_eval) --input=$(input)

solution:
	go run main.go --cc=$(cc) --key_public=$(key_public) --key_eval=$(key_eval) --input=$(input) --output=$(output)
	go run verify.go --sk=$(sk) --cc=$(cc) --input=$(input) --output=$(output)

clean:
	go run clean.go
	go clean
