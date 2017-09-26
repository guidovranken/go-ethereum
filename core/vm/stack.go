package vm

import fuzz_helper "github.com/guidovranken/go-coverage-instrumentation/helper"

import (
	"fmt"
	"math/big"
)

// stack is an object for basic stack operations. Items popped to the stack are
// expected to be changed and modified. stack does not take care of adding newly
// initialised objects.
type Stack struct {
	data []*big.Int
}

func newstack() *Stack {
	fuzz_helper.AddCoverage(22588)
	return &Stack{data: make([]*big.Int, 0, 1024)}
}

func (st *Stack) Data() []*big.Int {
	fuzz_helper.AddCoverage(44810)
	return st.data
}

func (st *Stack) push(d *big.Int) {
	fuzz_helper.AddCoverage(5262)

	st.data = append(st.data, d)
}
func (st *Stack) pushN(ds ...*big.Int) {
	fuzz_helper.AddCoverage(17878)
	st.data = append(st.data, ds...)
}

func (st *Stack) pop() (ret *big.Int) {
	fuzz_helper.AddCoverage(45021)
	ret = st.data[len(st.data)-1]
	st.data = st.data[:len(st.data)-1]
	return
}

func (st *Stack) len() int {
	fuzz_helper.AddCoverage(39040)
	return len(st.data)
}

func (st *Stack) swap(n int) {
	fuzz_helper.AddCoverage(2095)
	st.data[st.len()-n], st.data[st.len()-1] = st.data[st.len()-1], st.data[st.len()-n]
}

func (st *Stack) dup(pool *intPool, n int) {
	fuzz_helper.AddCoverage(21668)
	st.push(pool.get().Set(st.data[st.len()-n]))
}

func (st *Stack) peek() *big.Int {
	fuzz_helper.AddCoverage(45213)
	return st.data[st.len()-1]
}

// Back returns the n'th item in stack
func (st *Stack) Back(n int) *big.Int {
	fuzz_helper.AddCoverage(16619)
	return st.data[st.len()-n-1]
}

func (st *Stack) require(n int) error {
	fuzz_helper.AddCoverage(12692)
	if st.len() < n {
		fuzz_helper.AddCoverage(6577)
		return fmt.Errorf("stack underflow (%d <=> %d)", len(st.data), n)
	} else {
		fuzz_helper.AddCoverage(17393)
	}
	fuzz_helper.AddCoverage(42483)
	return nil
}

func (st *Stack) Print() {
	fuzz_helper.AddCoverage(64174)
	fmt.Println("### stack ###")
	if len(st.data) > 0 {
		fuzz_helper.AddCoverage(35657)
		for i, val := range st.data {
			fuzz_helper.AddCoverage(30358)
			fmt.Printf("%-3d  %v\n", i, val)
		}
	} else {
		fuzz_helper.AddCoverage(23294)
		fmt.Println("-- empty --")
	}
	fuzz_helper.AddCoverage(38740)
	fmt.Println("#############")
}

var _ = fuzz_helper.AddCoverage
