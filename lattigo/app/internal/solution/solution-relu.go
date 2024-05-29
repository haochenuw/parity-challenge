package solution

import (
	"fmt"
	"math/big"

	"github.com/tuneinsight/lattigo/v5/core/rlwe"
	"github.com/tuneinsight/lattigo/v5/he/hefloat"
	"github.com/tuneinsight/lattigo/v5/utils/bignum"
)


func SolveTestcase(params *hefloat.Parameters, evk *rlwe.MemEvaluationKeySet, in *rlwe.Ciphertext, index int) (out *rlwe.Ciphertext, err error) {
	// Put your solution here


	// relu := func(x float64) (y float64) {
	// 	return math.Max(0,x)
	// }

	// K := 1.0
	// Chebyhsev approximation of the sigmoid in the domain [-K, K] of degree 63.
	// poly := hefloat.NewPolynomial(GetChebyshevPoly(K, 7, relu))

	// poly := GetMinimaxPoly(K, 7, relu)

	// poly := hefloat.NewPolynomial
	// prec := uint(100)

	// coeffs_float64 := []float64{0.01568820242985998, 0.5072211371781018, 2.101243237989789, -0.46788565285049355, -6.5862632081337775, 3.6456240150729107, 7.587101728020599, -5.818417658286128}
	// coeffs_float64 := []float64{0.0229645346261931, 0.4999999999999936, 1.4340123453814493, 1.0966062321883419e-10, -2.0891847536610593, -3.4105118408214585e-10, 1.1551724027032066, 2.3139697111946557e-10}
	// coeffs_float64 := []float64{0, 1} // f(x) = x
	// coeffs := make([]*big.Float, len(coeffs_float64))

	// Apply the function to each element
	// for i, v := range coeffs_float64 {
	// 	coeffs[i] = bignum.NewFloat(v, prec)
	// }
	// poly := bignum.NewPolynomial(bignum.Chebyshev, coeffs, nil)
	slots := uint64(params.MaxSlots())
	values := make([]float64, slots)

	for i := uint64(0); i < slots; i++ {
		// if i % 2 == 0 {
		// 	values[i] = 0
		// } else {
		// 	values[i] = 1
		// }
		values[i] = 0
	}
	// values[0] = 0
	// values[1] = 0
	values[index] = 1

	encoder := hefloat.NewEncoder(*params)

	plaintext := hefloat.NewPlaintext(*params, params.MaxLevel())

	encoder.Encode(values, plaintext)

	eval := hefloat.NewEvaluator(*params, evk)

	// Instantiates the polynomial evaluator
	// polyEval := hefloat.NewPolynomialEvaluator(*params, eval)
	eval.Mul(in, plaintext, in)

	eval.Rescale(in, in)

	// var ct *rlwe.Ciphertext
	// if ct, err = polyEval.Evaluate(in, poly, params.DefaultScale()); err != nil {
	// 	panic(err)
	// }
	// ciphertext, err := evaluateHelper(params, mod1Parameters, eval, in)
	// fmt.Printf("Solution time: %s \n", time.Since(before))

	return in, err
}

// func evaluateHelper(params *hefloat.Parameters, mod1Parameters hefloat.Mod1Parameters, eval *hefloat.Evaluator, in *rlwe.Ciphertext) (out *rlwe.Ciphertext, err error) {
// 	fmt.Println("Level of input at beginning", in.Level())

// 	// Scale the message to Delta = Q/MessageRatio
// 	// scale := rlwe.NewScale(math.Exp2(math.Round(math.Log2(float64(params.Q()[0]) / mod1Parameters.MessageRatio()))))
// 	// scale = scale.Div(in.Scale)
// 	// eval.ScaleUp(in, rlwe.NewScale(math.Round(scale.Float64())), in)

// 	// // fmt.Println( "Scale after step 1", in.Scale.Value.Text('f', -1))

// 	// // // Scale the message up to Sine/MessageRatio
// 	// scale = mod1Parameters.ScalingFactor().Div(in.Scale)
// 	// scale = scale.Div(rlwe.NewScale(mod1Parameters.MessageRatio()))

// 	// eval.ScaleUp(in, rlwe.NewScale(math.Round(scale.Float64())), in)

// 	// Shifting by -128, to make symmetric range.
// 	// eval.Sub(in, 128, in)
// 	fmt.Printf("Old scale \n", in.Scale.Value.Text('f', -1))

// 	// Normalization
// 	// eval.Mul(in, 1/(4*float64(mod1Parameters.K())*mod1Parameters.QDiff()), in)
// 	// eval.Rescale(in, in)
// 	fmt.Println("Level of input before mod1 evaluation", in.Level())

// 	// ciphertext := in

// 	// beforeMod1 := time.Now()

// 	K := 1.0

// 	mod1Poly := bignum.ChebyshevApproximation(cos2pi, bignum.Interval{
// 		Nodes: 60,
// 		A:     *new(big.Float).SetPrec(cosine.EncodingPrecision).SetFloat64(-K),
// 		B:     *new(big.Float).SetPrec(cosine.EncodingPrecision).SetFloat64(K),
// 	})

// 	polyEval := hefloat.NewPolynomialEvaluator(*params, eval)

// 	scale, offset := mod1Poly.ChangeOfBasis()
// 	fmt.Printf("scale = %10.10f, offset = %10.10f \n", scale, offset)
// 	// eval.Mul(in, scale, in)
// 	// eval.Rescale(in, in)
// 	fmt.Println("New scale =", in.Scale.Value.Text('f', -1))
// 	// fmt.Println("Target scale =", params.DefaultScale().Value.Text('f', -1))

// 	ciphertext, err := polyEval.Evaluate(in, mod1Poly, rlwe.NewScale(params.DefaultScale()))

// 	// ciphertext, err := hefloat.NewMod1Evaluator(eval, polyEval, mod1Parameters).EvaluateNew(in)
// 	if err != nil {
// 		fmt.Println("error during mod 1 evaluation")
// 		return nil, err
// 	}
// 	// fmt.Println("Level of Ct after mod1 evaluation", ciphertext.Level())
// 	// fmt.Printf("Mod1 time: %s \n", time.Since(beforeMod1))

// 	// before := time.Now()
// 	// eval.Mul(ciphertext, float64(6.28318), ciphertext)
// 	// // eval.Rescale(ciphertext, ciphertext)
// 	// fmt.Printf("One plain mult time: %s \n", time.Since(before))

// 	// then cipher = cipher^2
// 	// beforeCipherMult := time.Now()
// 	// eval.MulRelin(ciphertext, ciphertext, ciphertext)
// 	// eval.Rescale(ciphertext, ciphertext)
// 	// fmt.Println("Final Level of Ct ", ciphertext.Level())

// 	return ciphertext, err
// }

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

// GetChebyshevPoly returns the Chebyshev polynomial approximation of f the
// in the interval [-K, K] for the given degree.
func GetChebyshevPoly(K float64, degree int, f64 func(x float64) (y float64)) bignum.Polynomial {

	FBig := func(x *big.Float) (y *big.Float) {
		xF64, _ := x.Float64()
		return new(big.Float).SetPrec(x.Prec()).SetFloat64(f64(xF64))
	}

	var prec uint = 128

	interval := bignum.Interval{
		A:     *bignum.NewFloat(-K, prec),
		B:     *bignum.NewFloat(K, prec),
		Nodes: degree,
	}

	// Returns the polynomial.
	return bignum.ChebyshevApproximation(FBig, interval)
}

// GetMinimaxPoly returns the minimax polynomial approximation of f the
// in the interval [-K, K] for the given degree.
func GetMinimaxPoly(K float64, degree int, f64 func(x float64) (y float64)) bignum.Polynomial {

	FBig := func(x *big.Float) (y *big.Float) {
		xF64, _ := x.Float64()
		return new(big.Float).SetPrec(x.Prec()).SetFloat64(f64(xF64))
	}

	// Bit-precision of the arbitrary precision arithmetic used by the minimax solver
	var prec uint = 160

	// Minimax (Remez) approximation of sigmoid
	r := bignum.NewRemez(bignum.RemezParameters{
		// Function to Approximate
		Function: FBig,

		// Polynomial basis of the approximation
		Basis: bignum.Chebyshev,

		// Approximation in [A, B] of degree Nodes.
		Intervals: []bignum.Interval{
			{
				A:     *bignum.NewFloat(-K, prec),
				B:     *bignum.NewFloat(K, prec),
				Nodes: degree,
			},
		},

		// Bit-precision of the solver
		Prec: prec,

		// Scan step for root finding
		// ScanStep: bignum.NewFloat(1/16.0, prec),
		ScanStep: bignum.NewFloat(1/128.0, prec),
		// Optimizes the scan-step for root finding
		OptimalScanStep: true,
	})

	// Max 10 iters, and normalized min/max error of 1e-15
	fmt.Printf("Minimax Approximation of Degree %d\n", degree)
	// r.Approximate(10, 1e-15)
	r.Approximate(20, 1e-30)
	fmt.Println()

	// Shoes the coeffs with 50 decimals of precision
	fmt.Printf("Minimax Chebyshev Coefficients [%f, %f]\n", -K, K)
	r.ShowCoeffs(16)
	fmt.Println()

	// Shows the min and max error with 50 decimals of precision
	fmt.Println("Minimax Error")
	r.ShowError(16)
	fmt.Println()

	// Returns the polynomial.
	return bignum.NewPolynomial(bignum.Chebyshev, r.Coeffs, [2]float64{-K, K})
}
