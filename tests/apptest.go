package tests

import (
	"encoding/json"
	"math/big"

	"github.com/revel/revel/testing"
)

// Must Embed `testing.TestSuite`
type MyAppTest struct {
	testing.TestSuite
}

type Fruit struct {
	ID            int64   `db:"id, primarykey" json:"id"`
	Fruit_label   string  `db:"fruit_label" json:"email"`
	Fruit_name    string  `db:"fruit_name" json:"fruit_name"`
	Fruit_subtype string  `db:"fruit_subtype" json:"fruit_subtype"`
	Mass          float64 `db:"mass" json:"mass"`
	Width         float64 `db:"width" json:"width"`
	Height        float64 `db:"height" json:"height"`
	Color_score   float64 `db:"color_score" json:"color_score"`
}

// Run this before a request
func (t *MyAppTest) Before() {
	println("Set up")
}

// Check hello world
func (t *MyAppTest) TestHelloPage() {
	t.Get("/hello")
	t.AssertOk()
	t.Assert(string(t.ResponseBody) == "Hello World")
}

// Check database operation
func (t *MyAppTest) TestDatabaseOperation() {
	t.Get("/database")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")

	var fruit []Fruit
	json.Unmarshal(t.ResponseBody, &fruit)
	t.Assert(len(fruit) == 59)

}

// Check Fibonacci
func (t *MyAppTest) TestFibonacci() {
	t.Get("/fibonacci")
	t.AssertOk()

	var numbers []big.Int
	json.Unmarshal(t.ResponseBody, &numbers)

	t.Assert(len(numbers) == 5000)
}

// Run this after request
func (t *MyAppTest) After() {
	println("Tear down")
}
