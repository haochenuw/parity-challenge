package main

import (
	"app/utils"
	"flag"
	"fmt"
	"log"
	"math"

	"github.com/tuneinsight/lattigo/v5/core/rlwe"
	"github.com/tuneinsight/lattigo/v5/he/hefloat"
)

func main() {
	ccFile := flag.String("cc", "", "")
	skFile := flag.String("sk", "", "")
	inputFile := flag.String("input", "", "")
	outputFile := flag.String("output", "", "")

	flag.Parse()

	params := new(hefloat.Parameters)
	if err := utils.Deserialize(params, *ccFile); err != nil {
		log.Fatalf(err.Error())
	}

	sk := new(rlwe.SecretKey)
	if err := utils.Deserialize(sk, *skFile); err != nil {
		log.Fatalf(err.Error())
	}

	in := new(rlwe.Ciphertext)
	if err := utils.Deserialize(in, *inputFile); err != nil {
		log.Fatalf(err.Error())
	}

	out := new(rlwe.Ciphertext)
	if err := utils.Deserialize(out, *outputFile); err != nil {
		log.Fatalf(err.Error())
	}

	dec := rlwe.NewDecryptor(*params, sk)
	ecd := hefloat.NewEncoder(*params)

	have := make([]float64, out.Slots())
	if err := ecd.Decode(dec.DecryptNew(out), have); err != nil {
		log.Fatalf("%T.Decode: %s", ecd, err.Error())
	}

	want := make([]float64, in.Slots())
	if err := ecd.Decode(dec.DecryptNew(in), want); err != nil {
		log.Fatalf("%T.Decode: %s", ecd, err.Error())
	}

	for i := range have {
		// want[i] = float64(int(math.Round(want[i])) & 1)
		// want[i] = math.Cos(2*math.Pi*want[i])
		want[i] = math.Max(0, want[i])
	}

	fmt.Println("Have: -- ", have[:10])
	fmt.Println("Want: -- ", want[:10])

	maxDelta := 0.0
	avgDelta := 0.0
	for i, _ := range have {
		delta := math.Abs(have[i] - want[i])
		maxDelta = math.Max(delta, maxDelta)
		avgDelta += delta
	}
	fmt.Println("maxDelta: -- ", maxDelta)
	fmt.Println("avgDelta: -- ", avgDelta / float64(len(have)))



	fmt.Println(hefloat.GetPrecisionStats(*params, ecd, nil, have, want, 0, false).String())
}
