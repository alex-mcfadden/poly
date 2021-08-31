package synthesis

import (
	"fmt"
	"strings"
	"sync"
	"testing"

	"github.com/TimothyStiles/poly/transform"
	"github.com/TimothyStiles/poly/transform/codon"
)

/******************************************************************************

Synthesis Fixer tests begin here

******************************************************************************/

var dataDir string = "../data/"

func ExampleFixCds() {
	bla := "ATGAGTATTCAACATTTCCGTGTCGCCCTTATTCCCTTTTTTGCGGCATACGGAAATGTTGAATACTCATTTTGCCTTCCTGTTTTTGCTCACCCAGAAACGCTGGTGAAAGTAAAAGATGCTGAAGATCAGTTGGGTGCACGAGTGGGTTACATCGAACTGGATCTCAACAGCGGTAAGATCCTTGAGAGTTTTCGCCCCGAAGAACGTTTTCCAATGATGAGCACTTTTAAAGTTCTGCTATGTGGCGCGGTATTATCCCGTATTGACGCCGGGCAAGAGCAACTCGGTCGCCGCATACACTATTCTCAGAATGACTTGGTTGAGTACTCACCAGTCACAGAAAAGCATCTTACGGATGGCATGACAGTAAGAGAATTATGCAGTGCTGCCATAACCATGAGTGATAACACTGCGGCCAACTTACTTCTGACAACGATCGGAGGACCGAAGGAGCTAACCGCTTTTTTGCACAACATGGGGGATCATGTAACTCGCCTTGATCGTTGGGAACCGGAGCTGAATGAAGCCATACCAAACGACGAGCGTGACACCACGATGCCTGTAGCAATGGCAACAACGTTGCGCAAACTATTAACTGGCGAACTACTTACTCTAGCTTCCCGGCAACAATTAATAGACTGGATGGAGGCGGATAAAGTTGCAGGACCACTTCTGCGCTCGGCCCTTCCGGCTGGCTGGTTTATTGCTGATAAATCTGGAGCCGGTGAGCGTGGGTCTCGCGGTATCATTGCAGCACTGGGGCCAGATGGTAAGCCCTCCCGTATCGTAGTTATCTACACGACGGGGAGTCAGGCAACTATGGATGAACGAAATAGACAGATCGCTGAGATAGGTGCCTCACTGATTAAGCATTGGTAA"

	codonTable := codon.ReadCodonJSON(dataDir + "pichiaTable.json")

	fixedSeq, changes, _ := FixCds(":memory:", bla, codonTable, []func(string, chan DnaSuggestion, *sync.WaitGroup){RemoveRepeat(20), RemoveSequence([]string{"GAAGAC", "GGTCTC", "GCGATG", "CGTCTC", "GCTCTTC", "CACCTGC"}, "TypeIIS restriction enzyme site")}, 10)
	fmt.Printf("Changed position %d from %s to %s for reason: %s. Complete sequence: %s", changes[1].Position, changes[1].From, changes[1].To, changes[1].Reason, fixedSeq)

	// Output: Changed position 245 from GGG to GGA for reason: TypeIIS restriction enzyme site. Complete sequence: ATGAGTATTCAACATTTCCGTGTCGCCCTTATTCCCTTTTTTGCGGCATATGGAAATGTTGAATACTCATTTTGCCTTCCTGTTTTTGCTCACCCAGAAACGCTGGTGAAAGTAAAAGATGCTGAAGATCAGTTGGGTGCACGAGTGGGTTACATCGAACTGGATCTCAACAGCGGTAAGATCCTTGAGAGTTTTCGCCCCGAAGAACGTTTTCCAATGATGAGCACTTTTAAAGTTCTGCTATGTGGCGCGGTATTATCCCGTATTGACGCCGGGCAAGAGCAACTCGGTCGCCGCATACACTATTCTCAGAATGACTTGGTTGAGTACTCACCAGTCACAGAAAAGCATCTTACGGATGGCATGACAGTAAGAGAATTATGCAGTGCTGCCATAACCATGAGTGATAACACTGCGGCCAACTTACTTCTGACAACGATCGGAGGACCGAAGGAGCTAACCGCTTTTTTGCACAACATGGGGGATCATGTAACTCGCCTTGATCGTTGGGAACCGGAGCTGAATGAAGCCATACCAAACGACGAGCGTGACACCACGATGCCTGTAGCAATGGCAACAACGTTGCGCAAACTATTAACTGGCGAACTACTTACTCTAGCTTCCCGGCAACAATTAATAGACTGGATGGAGGCGGATAAAGTTGCAGGACCACTTCTGCGCTCGGCCCTTCCGGCTGGCTGGTTTATTGCTGATAAATCTGGAGCCGGTGAGCGTGGATCTCGCGGTATCATTGCAGCACTGGGGCCAGATGGTAAGCCCTCCCGTATCGTAGTTATCTACACGACGGGGAGTCAGGCAACTATGGATGAACGAAATAGACAGATCGCTGAGATAGGTGCCTCACTGATTAAGCATTGGTAA
}

func ExampleFixCdsSimple() {
	bla := "ATGAAAAAAAAAAGTATTCAACATTTCCGTGTCGCCCTTATTCCCTTTTTTGCGGCATTTTGCCTTCCTGTTTTTGCTCACCCAGAAACGCTGGTGAAAGTAAAAGATGCTGAAGATCAGTTGGGTGCACGAGTGGGTTACATCGAACTGGATCTCAACAGCGGTAAGATCCTTGAGAGTTTTCGCCCCGAAGAACGTTTTCCAATGATGAGCACTTTTAAAGTTCTGCTATGTGGCGCGGTATTATCCCGTATTGACGCCGGGCAAGAGCAACTCGGTCGCCGCATACACTATTCTCAGAATGACTTGGTTGAGTACTCACCAGTCACAGAAAAGCATCTTACGGATGGCATGACAGTAAGAGAATTATGCAGTGCTGCCATAACCATGAGTGATAACACTGCGGCCAACTTACTTCTGACAACGATCGGAGGACCGAAGGAGCTAACCGCTTTTTTGCACAACATGGGGGATCATGTAACTCGCCTTGATCGTTGGGAACCGGAGCTGAATGAAGCCATACCAAACGACGAGCGTGACACCACGATGCCTGTAGCAATGGCAACAACGTTGCGCAAACTATTAACTGGCGAACTACTTACTCTAGCTTCCCGGCAACAATTAATAGACTGGATGGAGGCGGATAAAGTTGCAGGACCACTTCTGCGCTCGGCCCTTCCGGCTGGCTGGTTTATTGCTGATAAATCTGGAGCCGGTGAGCGTGGGTCTCGCGGTATCATTGCAGCACTGGGGCCAGATGGTAAGCCCTCCCGTATCGTAGTTATCTACACGACGGGGAGTCAGGCAACTATGGATGAACGAAATAGACAGATCGCTGAGATAGGTGCCTCACTGATTAAGCATTGGTAA"

	codonTable := codon.ReadCodonJSON(dataDir + "pichiaTable.json")

	fixedSeq, changes, _ := FixCdsSimple(bla, codonTable, []string{"GGTCTC"})
	fmt.Printf("Changed position %d from %s to %s for reason: %s. Complete sequence: %s", changes[0].Position, changes[0].From, changes[0].To, changes[0].Reason, fixedSeq)

	// Output: Changed position 1 from AAA to AAG for reason: Homopolymers. Complete sequence: ATGAAGAAAAAAAGTATTCAACATTTCCGTGTCGCCCTTATTCCCTTTTTTGCGGCATTTTGCCTTCCTGTTTTTGCTCACCCAGAAACGCTGGTGAAAGTAAAAGATGCTGAAGATCAGTTGGGTGCACGAGTGGGTTACATCGAACTGGATCTCAACAGCGGTAAGATCCTTGAGAGTTTTCGCCCCGAAGAACGTTTTCCAATGATGAGCACTTTTAAAGTTCTGCTATGTGGCGCGGTATTATCCCGTATTGACGCCGGGCAAGAGCAACTCGGTCGCCGCATACACTATTCTCAGAATGACTTGGTTGAGTACTCACCAGTCACAGAAAAGCATCTTACGGATGGCATGACAGTAAGAGAATTATGCAGTGCTGCCATAACCATGAGTGATAACACTGCGGCCAACTTACTTCTGACAACGATCGGAGGACCGAAGGAGCTAACCGCTTTTTTGCACAACATGGGGGATCATGTAACTCGCCTTGATCGTTGGGAACCGGAGCTGAATGAAGCCATACCAAACGACGAGCGTGACACCACGATGCCTGTAGCAATGGCAACAACGTTGCGCAAACTATTAACTGGCGAACTACTTACTCTAGCTTCCCGGCAACAATTAATAGACTGGATGGAGGCGGATAAAGTTGCAGGACCACTTCTGCGCTCGGCCCTTCCGGCTGGCTGGTTTATTGCTGATAAATCTGGAGCCGGTGAGCGTGGATCTCGCGGTATCATTGCAGCACTGGGGCCAGATGGTAAGCCCTCCCGTATCGTAGTTATCTACACGACGGGGAGTCAGGCAACTATGGATGAACGAAATAGACAGATCGCTGAGATAGGTGCCTCACTGATTAAGCATTGGTAA
}

func TestFixCdsWithAlteredCodonTable(t *testing.T) {
	bla := "ATGAAAAAAAAAAGTATTCAACATTTCCGTGTCGCCCTTATTCCCTTTTTTGCGGCATTTTGCCTTCCTGTTTTTGCTCACCCAGAAACGCTGGTGAAAGTAAAAGATGCTGAAGATCAGTTGGGTGCACGAGTGGGTTACATCGAACTGGATCTCAACAGCGGTAAGATCCTTGAGAGTTTTCGCCCCGAAGAACGTTTTCCAATGATGAGCACTTTTAAAGTTCTGCTATGTGGCGCGGTATTATCCCGTATTGACGCCGGGCAAGAGCAACTCGGTCGCCGCATACACTATTCTCAGAATGACTTGGTTGAGTACTCACCAGTCACAGAAAAGCATCTTACGGATGGCATGACAGTAAGAGAATTATGCAGTGCTGCCATAACCATGAGTGATAACACTGCGGCCAACTTACTTCTGACAACGATCGGAGGACCGAAGGAGCTAACCGCTTTTTTGCACAACATGGGGGATCATGTAACTCGCCTTGATCGTTGGGAACCGGAGCTGAATGAAGCCATACCAAACGACGAGCGTGACACCACGATGCCTGTAGCAATGGCAACAACGTTGCGCAAACTATTAACTGGCGAACTACTTACTCTAGCTTCCCGGCAACAATTAATAGACTGGATGGAGGCGGATAAAGTTGCAGGACCACTTCTGCGCTCGGCCCTTCCGGCTGGCTGGTTTATTGCTGATAAATCTGGAGCCGGTGAGCGTGGGTCTCGCGGTATCATTGCAGCACTGGGGCCAGATGGTAAGCCCTCCCGTATCGTAGTTATCTACACGACGGGGAGTCAGGCAACTATGGATGAACGAAATAGACAGATCGCTGAGATAGGTGCCTCACTGATTAAGCATTGGTAA"

	codonTable := codon.ReadCodonJSON(dataDir + "alteredPichiaTable.json")

	var functions []func(string, chan DnaSuggestion, *sync.WaitGroup)
	functions = append(functions, RemoveSequence([]string{"CGTGT"}, "Should change to CGA with the Altered Picha Table, because I choose this to be highest"))

	fixedSeq, changes, _ := FixCds(":memory:", bla, codonTable, functions, 10)
	textChange := fmt.Sprintf("Changed position %d from %s to %s for reason: %s. Complete sequence: %s", changes[0].Position, changes[0].From, changes[0].To, changes[0].Reason, fixedSeq)
	shouldChangeTo := "Changed position 9 from CGT to CGA for reason: Should change to CGA with the Altered Picha Table, because I choose this to be highest. Complete sequence: ATGAAAAAAAAAAGTATTCAACATTTCCGAGTCGCCCTTATTCCCTTTTTTGCGGCATTTTGCCTTCCTGTTTTTGCTCACCCAGAAACGCTGGTGAAAGTAAAAGATGCTGAAGATCAGTTGGGTGCACGAGTGGGTTACATCGAACTGGATCTCAACAGCGGTAAGATCCTTGAGAGTTTTCGCCCCGAAGAACGTTTTCCAATGATGAGCACTTTTAAAGTTCTGCTATGTGGCGCGGTATTATCCCGTATTGACGCCGGGCAAGAGCAACTCGGTCGCCGCATACACTATTCTCAGAATGACTTGGTTGAGTACTCACCAGTCACAGAAAAGCATCTTACGGATGGCATGACAGTAAGAGAATTATGCAGTGCTGCCATAACCATGAGTGATAACACTGCGGCCAACTTACTTCTGACAACGATCGGAGGACCGAAGGAGCTAACCGCTTTTTTGCACAACATGGGGGATCATGTAACTCGCCTTGATCGTTGGGAACCGGAGCTGAATGAAGCCATACCAAACGACGAGCGTGACACCACGATGCCTGTAGCAATGGCAACAACGTTGCGCAAACTATTAACTGGCGAACTACTTACTCTAGCTTCCCGGCAACAATTAATAGACTGGATGGAGGCGGATAAAGTTGCAGGACCACTTCTGCGCTCGGCCCTTCCGGCTGGCTGGTTTATTGCTGATAAATCTGGAGCCGGTGAGCGTGGGTCTCGCGGTATCATTGCAGCACTGGGGCCAGATGGTAAGCCCTCCCGTATCGTAGTTATCTATACGACGGGGAGTCAGGCAACTATGGATGAACGAAATAGACAGATCGCTGAGATAGGTGCCTCACTGATTAAGCATTGGTAA"
	if textChange != shouldChangeTo {
		t.Errorf("%s\nshould be\n%s", textChange, shouldChangeTo)
	}
}

func BenchmarkFixCds(b *testing.B) {
	phusion := "MGHHHHHHHHHHSSGILDVDYITEEGKPVIRLFKKENGKFKIEHDRTFRPYIYALLRDDSKIEEVKKITGERHGKIVRIVDVEKVEKKFLGKPITVWKLYLEHPQDVPTIREKVREHPAVVDIFEYDIPFAKRYLIDKGLIPMEGEEELKILAFDIETLYHEGEEFGKGPIIMISYADENEAKVITWKNIDLPYVEVVSSEREMIKRFLRIIREKDPDIIVTYNGDSFDFPYLAKRAEKLGIKLTIGRDGSEPKMQRIGDMTAVEVKGRIHFDLYHVITRTINLPTYTLEAVYEAIFGKPKEKVYADEIAKAWESGENLERVAKYSMEDAKATYELGKEFLPMEIQLSRLVGQPLWDVSRSSTGNLVEWFLLRKAYERNEVAPNKPSEEEYQRRLRESYTGGFVKEPEKGLWENIVYLDFRALYPSIIITHNVSPDTLNLEGCKNYDIAPQVGHKFCKDIPGFIPSLLGHLLEERQKIKTKMKETQDPIEKILLDYRQKAIKLLANSFYGYYGYAKARWYCKECAESVTAWGRKYIELVWKELEEKFGFKVLYIDTDGLYATIPGGESEEIKKKALEFVKYINSKLPGLLELEYEGFYKRGFFVTKKRYAVIDEEGKVITRGLEIVRRDWSEIAKETQARVLETILKHGDVEEAVRIVKEVIQKLANYEIPPEKLAIYEQITRPLHEYKAIGPHVAVAKKLAAKGVKIKPGMVIGYIVLRGDGPISNRAILAEEYDPKKHKYDAEYYIENQVLPAVLRILEGFGYRKEDLRYQKTRQVGLTSWLNIKKSGTGGGGATVKFKYKGEEKEVDISKIKKVWRVGKMISFTYDEGGGKTGRGAVSEKDAPKELLQMLEKQKK*"
	codonTable := codon.ReadCodonJSON(dataDir + "pichiaTable.json")
	var functions []func(string, chan DnaSuggestion, *sync.WaitGroup)
	functions = append(functions, RemoveSequence([]string{"GAAGAC", "GGTCTC", "GCGATG", "CGTCTC", "GCTCTTC", "CACCTGC"}, "TypeIIS restriction enzyme site."))
	for i := 0; i < b.N; i++ {
		seq, _ := codon.Optimize(phusion, codonTable)
		optimizedSeq, changes, err := FixCds(":memory:", seq, codonTable, functions, 10)
		if err != nil {
			b.Errorf("Failed to fix phusion with error: %s", err)
		}
		for _, cutSite := range []string{"GAAGAC", "GGTCTC", "GCGATG", "CGTCTC", "GCTCTTC", "CACCTGC"} {
			if strings.Contains(optimizedSeq, cutSite) {
				fmt.Println(changes)
				b.Errorf("phusion" + " contains " + cutSite)
			}
			if strings.Contains(transform.ReverseComplement(optimizedSeq), cutSite) {
				fmt.Println(changes)
				b.Errorf("phusion" + " reverse complement contains " + cutSite)
			}
		}
	}
}

func TestReversion(t *testing.T) {
	// Previously, there was an error where BsmBI could get in a loop
	// It would first change CGA -> AGA, then get stuck changing AGA -> AGA
	codonTable := codon.ReadCodonJSON(dataDir + "pichiaTable.json")
	seq := "GGACGAGACGGC"
	var functions []func(string, chan DnaSuggestion, *sync.WaitGroup)
	functions = append(functions, RemoveSequence([]string{"GGTCTC", "CGTCTC"}, "TypeIIS restriction enzyme site."))
	_, _, err := FixCds(":memory:", seq, codonTable, functions, 10)
	if err != nil {
		t.Errorf("Failed with error: %s", err)
	}
}

func TestFixCds(t *testing.T) {
	codonTable := codon.ReadCodonJSON(dataDir + "pichiaTable.json")
	phusion := "MGHHHHHHHHHHSSGILDVDYITEEGKPVIRLFKKENGKFKIEHDRTFRPYIYALLRDDSKIEEVKKITGERHGKIVRIVDVEKVEKKFLGKPITVWKLYLEHPQDVPTIREKVREHPAVVDIFEYDIPFAKRYLIDKGLIPMEGEEELKILAFDIETLYHEGEEFGKGPIIMISYADENEAKVITWKNIDLPYVEVVSSEREMIKRFLRIIREKDPDIIVTYNGDSFDFPYLAKRAEKLGIKLTIGRDGSEPKMQRIGDMTAVEVKGRIHFDLYHVITRTINLPTYTLEAVYEAIFGKPKEKVYADEIAKAWESGENLERVAKYSMEDAKATYELGKEFLPMEIQLSRLVGQPLWDVSRSSTGNLVEWFLLRKAYERNEVAPNKPSEEEYQRRLRESYTGGFVKEPEKGLWENIVYLDFRALYPSIIITHNVSPDTLNLEGCKNYDIAPQVGHKFCKDIPGFIPSLLGHLLEERQKIKTKMKETQDPIEKILLDYRQKAIKLLANSFYGYYGYAKARWYCKECAESVTAWGRKYIELVWKELEEKFGFKVLYIDTDGLYATIPGGESEEIKKKALEFVKYINSKLPGLLELEYEGFYKRGFFVTKKRYAVIDEEGKVITRGLEIVRRDWSEIAKETQARVLETILKHGDVEEAVRIVKEVIQKLANYEIPPEKLAIYEQITRPLHEYKAIGPHVAVAKKLAAKGVKIKPGMVIGYIVLRGDGPISNRAILAEEYDPKKHKYDAEYYIENQVLPAVLRILEGFGYRKEDLRYQKTRQVGLTSWLNIKKSGTGGGGATVKFKYKGEEKEVDISKIKKVWRVGKMISFTYDEGGGKTGRGAVSEKDAPKELLQMLEKQKK*"
	var functions []func(string, chan DnaSuggestion, *sync.WaitGroup)
	functions = append(functions, RemoveSequence([]string{"GAAGAC", "GGTCTC", "GCGATG", "CGTCTC", "GCTCTTC", "CACCTGC"}, "TypeIIS restriction enzyme site."))
	seq, _ := codon.Optimize(phusion, codonTable)
	optimizedSeq, _, err := FixCds(":memory:", seq, codonTable, functions, 10)
	if err != nil {
		t.Errorf("Failed with error: %s", err)
	}

	for _, cutSite := range []string{"GAAGAC", "GGTCTC", "GCGATG", "CGTCTC", "GCTCTTC", "CACCTGC"} {
		if strings.Contains(optimizedSeq, cutSite) {
			t.Errorf("phusion" + " contains " + cutSite)
		}
		if strings.Contains(transform.ReverseComplement(optimizedSeq), cutSite) {
			t.Errorf("phusion" + " reverse complement contains " + cutSite)
		}
	}

	// Does this flip back in the history?
	fixedSeq, _, _ := FixCdsSimple("ATGTATTGA", codonTable, []string{"TAT"})
	if fixedSeq != "ATGTACTGA" {
		t.Errorf("Failed to fix ATGTATTGA -> ATGTACTGA")
	}

	// Repeat checking
	blaWithRepeat := "ATGAGTATTCAACATTTCCGTGTCGCCCTTATTCCCTTTTTTGCGGCATTTTGCCTTCCTGTTTTTGCTCACCCAGAAACGCTGGTGAAAGTAAAAGATGCTGAAGATCAGTTGGGTGCACGAGTGGGTTACATCGAACTGGATCTCAACAGCGGTAAGATCCTTGAGAGTTTTCGCCCCGAAGAACGTTTTCCAATGATGAGCACTTTTAAAGTTCTGCTATGTGGCGCGGTATTATCCCGTATTGACGCCGGGCAAGAGCAACTCGGTCGCCGCATACACTATTCTCAGAATGACTTGGTTGAGTACTCACCAGTCACAGAAAAGCATCTTACGGATGGCATGACAGTAAGAGAATTATGCAGTGCTGCCATAACCATGAGTGATAACACTGCGGCCAACTTACTTCTGACAACGATCGGAGGACCGAAGGAGCTAACCGCTTTTTTGCACAACATGGGGGATCATGTAACTCGCCTTGATCGTTGGGAACCGGAGCTGAATGAAGCCATACCAAACGACGAGCGTGACACCACGATGCCTGTAGCAATGGCAACAACGTTGCGCAAACTATTAACTGGCGAACTACTTACTCTAGCTTCCCGGCAACAATTAATAGACTGGATGGAGGCGGATAAAGTTGCAGGACCACTTCTGCGCTCGGCCCTTCCGGCTGGCTGGTTTATTGCTGATAAATCTGGAGCCGGTGAGCGTGGGTCTCGCGGTATCATTGCAGCACTGGGGCCAGATGGTAAGCCCTCCCGTATCGTAGTTATCTACACGACGGGGAGTCAGGCAACTATGGATGAACGAAATAGACAGATCGCTGAGATAGGTGCCTCACTGATTAAGCATTGGGGTGCCTCACTGATTAAGCATTGGTAA"
	functions = append(functions, RemoveRepeat(20))
	blaWithoutRepeat, _, err := FixCds(":memory:", blaWithRepeat, codonTable, functions, 10)
	if err != nil {
		t.Errorf("Failed to remove repeat with error: %s", err)
	}
	targetBlaWithoutRepeat := "ATGAGTATTCAACATTTCCGTGTCGCCCTTATTCCCTTTTTTGCGGCATTTTGCCTTCCTGTTTTTGCTCACCCAGAAACGCTGGTGAAAGTAAAAGATGCTGAAGATCAGTTGGGTGCACGAGTGGGTTACATCGAACTGGATCTCAACAGCGGTAAGATCCTTGAGAGTTTTCGCCCCGAAGAACGTTTTCCAATGATGAGCACTTTTAAAGTTCTGCTATGTGGCGCGGTATTATCCCGTATTGACGCCGGGCAAGAGCAACTCGGTCGCCGCATACACTATTCTCAGAATGACTTGGTTGAGTACTCACCAGTCACAGAAAAGCATCTTACGGATGGCATGACAGTAAGAGAATTATGCAGTGCTGCCATAACCATGAGTGATAACACTGCGGCCAACTTACTTCTGACAACGATCGGAGGACCGAAGGAGCTAACCGCTTTTTTGCACAACATGGGGGATCATGTAACTCGCCTTGATCGTTGGGAACCGGAGCTGAATGAAGCCATACCAAACGACGAGCGTGACACCACGATGCCTGTAGCAATGGCAACAACGTTGCGCAAACTATTAACTGGCGAACTACTTACTCTAGCTTCCCGGCAACAATTAATAGACTGGATGGAGGCGGATAAAGTTGCAGGACCACTTCTGCGCTCGGCCCTTCCGGCTGGCTGGTTTATTGCTGATAAATCTGGAGCCGGTGAGCGTGGATCTCGCGGTATCATTGCAGCACTGGGGCCAGATGGTAAGCCCTCCCGTATCGTAGTTATCTACACGACGGGGAGTCAGGCAACTATGGATGAACGAAATAGACAGATCGCTGAGATAGGTGCCTCACTGATTAAGCATTGGGGAGCTTCACTGATTAAGCATTGGTAA"

	if blaWithoutRepeat != targetBlaWithoutRepeat {
		t.Errorf("Expected blaWithoutRepeat sequence %s, got: %s", targetBlaWithoutRepeat, blaWithoutRepeat)
	}

	// Test low and high GC content
	var gcFunctions []func(string, chan DnaSuggestion, *sync.WaitGroup)
	gcFunctions = append(gcFunctions, GcContentFixer(0.90, 0.10))
	fixedSeq, _, err = FixCds(":memory:", "GGGCCC", codonTable, gcFunctions, 10)
	if fixedSeq != "GGACCC" {
		fmt.Println(err)
		t.Errorf("Failed to fix GGGCCC -> GGACC. Got %s", fixedSeq)
	}
	fixedSeq, _, _ = FixCds(":memory:", "AAATTT", codonTable, gcFunctions, 10)
	if fixedSeq != "AAGTTT" {
		fmt.Println(err)
		t.Errorf("Failed to fix AAATTT -> AAGTTT. Got %s", fixedSeq)
	}
}

func TestFixCdsBadInput(t *testing.T) {
	// This block tests a sequence that is not divisible by 3
	codonTable := codon.ReadCodonJSON(dataDir + "pichiaTable.json")
	var functions []func(string, chan DnaSuggestion, *sync.WaitGroup)
	functions = append(functions, RemoveSequence([]string{"GAAGAC", "GGTCTC", "GCGATG", "CGTCTC", "GCTCTTC", "CACCTGC"}, "TypeIIS restriction enzyme site"))
	_, _, err := FixCds(":memory:", "AT", codonTable, functions, 10)
	if err == nil {
		t.Errorf("FixCds should fail with sequence input that is not divisible by 3")
	}

	// This block tests a sequence that has a bad GC bias
	badGcBiasFunc := func(sequence string, c chan DnaSuggestion, wg *sync.WaitGroup) {
		c <- DnaSuggestion{0, 1, "XY", 1, "this should fail"}
		wg.Done()
	}
	_, _, err = FixCds(":memory:", "ATG", codonTable, []func(string, chan DnaSuggestion, *sync.WaitGroup){badGcBiasFunc}, 10)
	if err == nil {
		t.Errorf("XY should fail as a valid GC bias")
	}

	// This block tests something with no solution space
	_, _, err = FixCds(":memory:", "GGG", codonTable, []func(string, chan DnaSuggestion, *sync.WaitGroup){GcContentFixer(0.10, 0.05)}, 10)
	if err == nil {
		t.Errorf("There should be no solution to GGG -> less than .10 gc content.")
	}

	// This block tests that any given suggestion will suggest within the confines of the sequence.
	outOfRangePosition := func(sequence string, c chan DnaSuggestion, wg *sync.WaitGroup) {
		c <- DnaSuggestion{10000000000, 1000000000, "GC", 1, "this should fail"}
		wg.Done()
	}
	_, _, err = FixCds(":memory:", "ATG", codonTable, []func(string, chan DnaSuggestion, *sync.WaitGroup){outOfRangePosition}, 10)
	if err == nil {
		t.Errorf("outOfRangePosition should fail because it is out of range of the start and end positions of the sequence.")
	}
}
