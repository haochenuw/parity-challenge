package solution

import (
	"fmt"
	"math"
	"time"

	"github.com/tuneinsight/lattigo/v5/core/rlwe"
	"github.com/tuneinsight/lattigo/v5/he/hefloat"
)

func SolveTestcase(params *hefloat.Parameters, evk *rlwe.MemEvaluationKeySet, in *rlwe.Ciphertext) (out *rlwe.Ciphertext, err error) {
	// Put your solution here
	evm := hefloat.Mod1ParametersLiteral{
		LevelStart:      10,
		Mod1Type:        hefloat.CosContinuous,
		LogMessageRatio: 0,
		K:               32,
		Mod1Degree:      24,
		Mod1InvDegree:   0,
		DoubleAngle:     3,
		LogScale:        60,
	}

	mod1Parameters, err := hefloat.NewMod1ParametersFromLiteral(*params, evm)

	// todo: initialize eval = hefloat.Evaluator
	eval := hefloat.NewEvaluator(*params, evk)

	fmt.Println("Scale after step 0", in.Scale.Value.Text('f', -1))

	before := time.Now()
	ciphertext, err := evaluateHelper(params, mod1Parameters, eval, in)
	fmt.Printf("Solution time: %s \n", time.Since(before))

	return ciphertext, err
}

func evaluateHelper(params *hefloat.Parameters, mod1Parameters hefloat.Mod1Parameters, eval *hefloat.Evaluator, in *rlwe.Ciphertext) (out *rlwe.Ciphertext, err error) {
	fmt.Println("Level of input at beginning", in.Level())

	// Scale the message to Delta = Q/MessageRatio
	scale := rlwe.NewScale(math.Exp2(math.Round(math.Log2(float64(params.Q()[0]) / mod1Parameters.MessageRatio()))))
	scale = scale.Div(in.Scale)
	eval.ScaleUp(in, rlwe.NewScale(math.Round(scale.Float64())), in)

	// fmt.Println("Scale after step 1", in.Scale.Value.Text('f', -1))

	// // Scale the message up to Sine/MessageRatio
	scale = mod1Parameters.ScalingFactor().Div(in.Scale)
	scale = scale.Div(rlwe.NewScale(mod1Parameters.MessageRatio()))

	eval.ScaleUp(in, rlwe.NewScale(math.Round(scale.Float64())), in)

	// Shifting by -128, to make symmetric range.
	eval.Sub(in, 128, in)

	// Normalization
	eval.Mul(in, 1/(4*float64(mod1Parameters.K())*mod1Parameters.QDiff()), in)
	eval.Rescale(in, in)
	fmt.Println("Level of input before mod1 evaluation", in.Level())

	beforeMod1 := time.Now()

	ciphertext, err := hefloat.NewMod1Evaluator(eval, hefloat.NewPolynomialEvaluator(*params, eval), mod1Parameters).EvaluateNew(in)
	if err != nil {
		fmt.Println("error during mod 1 evaluation")
		return nil, err
	}
	fmt.Println("Level of Ct after mod1 evaluation", ciphertext.Level())
	fmt.Printf("Mod1 time: %s \n", time.Since(beforeMod1))

	before := time.Now()
	eval.Mul(ciphertext, float64(6.28318530717958), ciphertext)
	// eval.Rescale(ciphertext, ciphertext)
	fmt.Printf("One plain mult time: %s \n", time.Since(before))

	// then cipher = cipher^2
	beforeCipherMult := time.Now()
	eval.MulRelin(ciphertext, ciphertext, ciphertext)
	eval.Rescale(ciphertext, ciphertext)
	fmt.Println("Final Level of Ct ", ciphertext.Level())

	return ciphertext, err
}

// // The evaluatenew function
// func EvaluateNew(eval hefloat.Mod1Evaluator, ct *rlwe.Ciphertext) (*rlwe.Ciphertext, error) {

// 	var err error

// 	evm := eval.Mod1Parameters

// 	if ct.Level() < evm.LevelStart() {
// 		return nil, fmt.Errorf("cannot Evaluate: ct.Level() < Mod1Parameters.LevelStart")
// 	}

// 	if ct.Level() > evm.LevelStart() {
// 		eval.DropLevel(ct, ct.Level()-evm.LevelStart())
// 	}

// 	// Stores default scales
// 	prevScaleCt := ct.Scale

// 	// Normalize the modular reduction to mod by 1 (division by Q)
// 	ct.Scale = evm.ScalingFactor()

// 	// Compute the scales that the ciphertext should have before the double angle
// 	// formula such that after it it has the scale it had before the polynomial
// 	// evaluation

// 	Qi := eval.GetParameters().Q()

// 	targetScale := ct.Scale
// 	for i := 0; i < evm.doubleAngle; i++ {
// 		targetScale = targetScale.Mul(rlwe.NewScale(Qi[evm.levelStart-evm.mod1Poly.Depth()-evm.doubleAngle+i+1]))
// 		targetScale.Value.Sqrt(&targetScale.Value)
// 	}

// 	// Division by 1/2^r and change of variable for the Chebyshev evaluation
// 	if evm.Mod1Type == CosDiscrete || evm.Mod1Type == CosContinuous {
// 		offset := new(big.Float).Sub(&evm.mod1Poly.B, &evm.mod1Poly.A)
// 		offset.Mul(offset, new(big.Float).SetFloat64(evm.scFac))
// 		offset.Quo(new(big.Float).SetFloat64(-0.5), offset)

// 		if err = eval.Add(ct, offset, ct); err != nil {
// 			return nil, fmt.Errorf("cannot Evaluate: %w", err)
// 		}
// 	}

// 	// Chebyshev evaluation
// 	if ct, err = eval.PolynomialEvaluator.Evaluate(ct, evm.mod1Poly, rlwe.NewScale(targetScale)); err != nil {
// 		return nil, fmt.Errorf("cannot Evaluate: %w", err)
// 	}

// 	// Double angle
// 	sqrt2pi := evm.sqrt2Pi
// 	for i := 0; i < evm.doubleAngle; i++ {
// 		sqrt2pi *= sqrt2pi

// 		if err = eval.MulRelin(ct, ct, ct); err != nil {
// 			return nil, fmt.Errorf("cannot Evaluate: %w", err)
// 		}

// 		if err = eval.Add(ct, ct, ct); err != nil {
// 			return nil, fmt.Errorf("cannot Evaluate: %w", err)
// 		}

// 		if err = eval.Add(ct, -sqrt2pi, ct); err != nil {
// 			return nil, fmt.Errorf("cannot Evaluate: %w", err)
// 		}

// 		if err = eval.Rescale(ct, ct); err != nil {
// 			return nil, fmt.Errorf("cannot Evaluate: %w", err)
// 		}
// 	}

// 	// ArcSine
// 	if evm.mod1InvPoly != nil {
// 		if ct, err = eval.PolynomialEvaluator.Evaluate(ct, *evm.mod1InvPoly, ct.Scale); err != nil {
// 			return nil, fmt.Errorf("cannot Evaluate: %w", err)
// 		}
// 	}

// 	// Multiplies back by q
// 	ct.Scale = prevScaleCt
// 	return ct, nil
// }
