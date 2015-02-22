package ermahgerd

import "testing"

func TestGert(t *testing.T) {
	sentences := []string{"Oh my god! Sandwiches!", "Goosebumps! My favorite books!", "I liked your status on Facebook"}
	expected := []string{"ER MAH GERD! SERNDWERCHERS!", "GERSBERMS! MAH FRAVRIT BERKS!", "I LERKERD YER STERTERS ERN FERCERBERK"}

	for index, s := range sentences {
		response := Gert(s)

		if response != expected[index] {
			t.Error("Expected", expected[index], ",instead got", response)
		}
	}
}
