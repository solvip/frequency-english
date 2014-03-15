package main

import (
	"crypto/rand"
	"fmt"
	"github.com/solvip/frequency"
	"io/ioutil"
	"os"
)

const statePath = "state.gob"

func main() {
	a := frequency.NewAnalyzer()

	/* Try to restore previous state */
	if err := a.Restore(statePath); err != nil {
		/* None found - reread corpus */
		contents, err := ioutil.ReadFile("corpus.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not read corpus: %s\n", err.Error())

			os.Exit(1)
		}

		a.Feed(contents)
		if err := a.Save(statePath); err != nil {
			fmt.Fprintf(os.Stderr, "Could not save analyzer state: %s\n", err.Error())
		}
	}

	if readme, err := ioutil.ReadFile("README.md"); err == nil {
		fmt.Printf("Score of README.md: %f\n", a.Score(readme))
	}

	if license, err := ioutil.ReadFile("LICENSE"); err == nil {
		fmt.Printf("Score of LICENSE: %f\n", a.Score(license))
	}

	if corpus, err := ioutil.ReadFile("corpus.txt"); err == nil {
		fmt.Printf("Score of corpus.txt: %f\n", a.Score(corpus))
	}

	random := make([]byte, 500)
	rand.Read(random)
	fmt.Printf("Score of 500 random bytes: %f\n", a.Score(random))

	os.Exit(0)
}
