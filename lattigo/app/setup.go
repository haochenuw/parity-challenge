package main

import (
	"app/utils"
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"os"

	"github.com/tuneinsight/lattigo/v5/core/rlwe"
	"github.com/tuneinsight/lattigo/v5/he/hefloat"
)

// go run gen_keys.go --sk sk.bin --cc cc.bin --key_public pub.bin --key_eval mult.bin --input in.bin
func main() {
	ccFile := flag.String("cc", "", "")
	skFile := flag.String("sk", "", "")
	_ = flag.String("key_public", "", "")
	evalFile := flag.String("key_eval", "", "")
	inputFile := flag.String("input", "", "")

	flag.Parse()

	paramsJSON := struct {
		LogN            int   `json:"log_n"`
		LogQ            []int `json:"log_q"`
		LogP            []int `json:"log_p"`
		LogDefaultScale int   `json:"log_default_scale"`
	}{}

	dataJSON, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("os.Open(%s): %s", "config.json", err.Error())
	}

	if err := json.Unmarshal(dataJSON, &paramsJSON); err != nil {
		log.Fatalf(err.Error())
	}

	var params hefloat.Parameters

	// 128-bit secure parameters enabling depth-7 circuits.
	// LogN:14, LogQP: 431.
	if params, err = hefloat.NewParametersFromLiteral(
		hefloat.ParametersLiteral{
			LogN:            paramsJSON.LogN,            // log2(ring degree)
			LogQ:            paramsJSON.LogQ,            // log2(primes Q) (ciphertext modulus)
			LogP:            paramsJSON.LogP,            // log2(primes P) (auxiliary modulus)
			LogDefaultScale: paramsJSON.LogDefaultScale, // log2(scale)
		}); err != nil {
		log.Fatalf(err.Error())
	}

	// Key Generator
	kgen := rlwe.NewKeyGenerator(params)

	// Secret Key
	sk := kgen.GenSecretKeyNew()

	// Encoder
	ecd := hefloat.NewEncoder(params)

	// Encryptor
	enc := rlwe.NewEncryptor(params, sk)

	rlk := kgen.GenRelinearizationKeyNew(sk)

	evk := rlwe.NewMemEvaluationKeySet(rlk)

	// Source for sampling random plaintext values (not cryptographically secure)
	/* #nosec G404 */
	r := rand.New(rand.NewSource(0))

	// Populates the vector of plaintext values
	values := make([]float64, params.MaxSlots())
	for i := range values {
		values[i] = float64(i%256) + (2*r.Float64()-1)*1e-5
	}

	pt := hefloat.NewPlaintext(params, params.MaxLevel())

	// Encodes the vector of plaintext values
	if err = ecd.Encode(values, pt); err != nil {
		log.Fatalf(err.Error())
	}

	input, err := enc.EncryptNew(pt)

	if err != nil {
		log.Fatalf(err.Error())
	}

	if err := utils.Serialize(params, *ccFile); err != nil {
		log.Fatalf(err.Error())
	}

	if err := utils.Serialize(sk, *skFile); err != nil {
		log.Fatalf(err.Error())
	}

	if err := utils.Serialize(evk, *evalFile); err != nil {
		log.Fatalf(err.Error())
	}

	if err := utils.Serialize(input, *inputFile); err != nil {
		log.Fatalf(err.Error())
	}
}
