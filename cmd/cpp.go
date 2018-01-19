// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"os"

	"text/template"
	"time"
	"math/rand"
)

type program struct {
	Input string
	Bytes []byte
}

const tplCpp = `
// encrypted strProtect [C/C++]
// value = "{{.Input}}"
unsigned char value[{{len .Bytes}}] = { {{ range $key, $value := .Bytes }} {{ $value }}, {{ end }} };

int ii = 0;
for (unsigned int i = 0, v = 0; i < {{len .Bytes}}; i += 2, ii++) {
	v = value[i] ^ value[i + 1];
	v = v / 2;
	value[ii] = v;
}
value[ii] = '\0';
`
var input string
var randomByte byte

func init() {
	rootCmd.AddCommand(cppCmd)
	cppCmd.Flags().StringVarP(&input, "i", "l", "", "Input that needs to be protected")
	rand.Seed(time.Now().Unix())
	randomByte = uint8(rand.Int())
}

// cppCmd represents the cpp command
var cppCmd = &cobra.Command{
	Use:   "cpp",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Create a new template and parse the letter into it.
		t := template.Must(template.New("C++").Parse(tplCpp))
		t.Execute(os.Stdout, newProgram(input))
	},
}

func newProgram(input string) *program {
	return &program{
		Input: input,
		Bytes: protect(input),
	}
}


func protect(input string) []byte {
	result := make([]byte, len(input) * 2)
	now := time.Now().Unix()
	ii := 0
	for i := 0; i < len(result); i += 2 {
		rand.Seed(now + int64(i))
		random := uint8(rand.Int())
		value := input[ii]
		//Add more entropy
		value = value * 2

		protectedValue := value ^ random
		result[i+1] = random
		result[i] = protectedValue
		ii++
	}
	return result
}

func unProtect(input []byte) string {
	ii := 0
	for i := 0; i < len(input); i += 2 {
		random := input[i + 1]
		protectedValue := input[i]
		value := protectedValue ^ random
		//Add more entropy
		value = value / 2

		input[ii] = byte(value)
		ii++
	}
	//input[ii] = '\x00'
	return string(input[:ii])
}
