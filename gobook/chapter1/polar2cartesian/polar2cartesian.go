package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
)

type polar struct {
	radius float64
	θ      float64
}

type cartesian struct {
	x float64
	y float64
}

var prompt = "Enter a radius and an angle(in degrees), eg: 12.5 90 or %s to quit!"

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else {
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}

func createSolver(questions chan polar) chan cartesian {
	answers := make(chan cartesian)
	go func() {
		for {
			polarCoord := <-questions
			θ := polarCoord.θ * math.Pi / 180.0
			x := polarCoord.radius * math.Cos(θ)
			y := polarCoord.radius * math.Sin(θ)
			answers <- cartesian{x, y}
		}
	}()
	return answers
}

var result = "Polar radius=%.02f θ=%.02f° → Cartesian x=%.02f y=%.02f\n"

func interact(questions chan polar, answers chan cartesian) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	for {
		fmt.Println("Radius and angle: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(os.Stderr, "fail read from stdin ....")
			break
		}
		fmt.Println(line)
		var radius, θ float64
		if _, err := fmt.Sscanf(line, "%f %f", &radius, &θ); err != nil {
			fmt.Println(os.Stderr, "invalid input")
			continue
		}
		questions <- polar{radius, θ}
		coord := <-answers
		fmt.Printf(result, radius, θ, coord.x, coord.y)
	}
	fmt.Println()
}

func main() {
	questions := make(chan polar)
	defer close(questions)
	answers := createSolver(questions)
	defer close(answers)
	interact(questions, answers)
}
