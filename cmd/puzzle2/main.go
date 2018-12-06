package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	compute(data)
	fmt.Printf("result = %+v\n", computePartTwo(data))
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func computePartTwo(data string) string {

	a := strings.Split(data, "\n")
	for i, curWord := range a {
		if i == len(a)-1 {
			break
		}
		for _, cmpWord := range a[i+1:] {
			diff := 0
			for k, curLetter := range curWord {
				if curLetter != rune(cmpWord[k]) {
					diff++
				}
				if diff > 1 {
					break
				}
			}
			if diff == 1 {
				return "\n" + curWord + "\n" + cmpWord
			}
		}
	}
	return "not found"
}

func compute(data string) int {
	a := strings.Split(data, "\n")
	twice := 0
	three := 0
	for _, line := range a {
		line = SortString(line)
		breakTwice := false
		breakThree := false
		for {
			if len(line) < 2 {
				break
			}
			cur := rune(line[0])
			cnt := 1
			j := 1
			for _, comp := range line[j:] {
				j++
				if comp != cur {
					break
				}
				cnt++
			}
			if cnt == 2 && !breakTwice {
				twice++
				breakTwice = true
			} else if cnt == 3 && !breakThree {
				three++
				breakThree = true
			}
			line = line[j-1:]
		}
	}
	fmt.Printf("total: %d\n", twice*three)
	return twice * three
}

var data = `xdmgyjkpruszabaqwficevtjeo
xdmgybkgwuszlbaqwfichvtneo
xdmgyjkpruszlbcwwfichvtndo
xdmgcjkprusyibaqwfichvtneo
xdmgyjktruszlbwqwficuvtneo
xdmgxjkpruszlbaqyfichvtnvo
xdmgytkpruszlbaqwficuvtnlo
xdmgydkpruszlbaqwfijhvtnjo
xfmgyjkmruszlbaqwfichvtnes
xdmgyrktruszlraqwfichvtneo
xdmgyjkihuszlbaqdfichvtneo
hdmgyjkpruszeiaqwfichvtneo
xdmzyjkpruszlbaqwgichvtnxo
xdmgyjknquszlbpqwfichvtneo
idmgyjrpruszlbtqwfichvtneo
xkmgyjkpruuzlbaqwfichvfneo
xdmgyjkpruszlfaqwficnvtner
xdmgyjkpruszlbpqwficwvteeo
xdmgyjkpwuszlbiqwfhchvtneo
xdmgyjkpruszwbaqwfichrtnbo
xdpgyjkprusblbaqwfgchvtneo
xdmryjkcruszlbaqwfichvtnee
xwmgylkpruszlbaqwfcchvtneo
xdmgyjkpruszflaqwfixhvtneo
xdmgyjkmruszloaqwfichvteeo
xvmgrjkpruszlbaqwfichvsneo
xdmvyjkprusmlbaqwfichvtnes
xdmgyjkpruszlbaqwfichkgbeo
xdmgyikpruxzlbaqwfichvtnei
xdmgyjkprugzlbaqhfichvtveo
xdmgyjkpruszlbaqjaichftneo
xdmzijkpruszlbaqwwichvtneo
xdmgyjkprsszlbaqwfihhvlneo
xdmgyjkprusqlwaqzfichvtneo
ximgyjkpruszlbawwfichvtnen
xsmgyjzpruszlbaqwfichvaneo
xdmgyjkpruszlcaoyfichvtneo
xdmgyjkprusmlbaqvnichvtneo
xdmgyjkvruszmbaqwfichvtueo
xdmgyjppuuszleaqwfichvtneo
xddgyjkprubzlbaqwfichvaneo
xdmgwjkpruszebaswfichvtneo
xdogyjkpruszlblqwfichvdneo
xdkgyjgpruszlbaqwfizhvtneo
xdvgyjkpruszlbdqwfichvtqeo
xdmgyjlpruszlbapwficgvtneo
xdmgyjkpruszlbaqofickvtngo
xdmgyjkprqszliaywfichvtneo
xdqgyjkpruszlbcqwficnvtneo
xdmgdjkpruszlbaqwxichvtseo
xdmgyjkpruczlbaqwfichdtnfo
xdmgyjkpruszluaqwficzvtnjo
xdmgyjkproszlbaqwfacevtneo
xfmgijkpruszlbrqwfichvtneo
odmgyjkpluszlbaqwfichvuneo
xdmgyjkpruszlbaqwwichukneo
xdmgdjkpruszwbaqwfichvtnet
xdmgyjkzrusvlbaqwrichvtneo
xdmgylkprutzlbaqwfichvtnbo
xdmgyjkpruszsbaqwfijtvtneo
xdmgyjkproszlbjqwfichntneo
xdmgyhkpluszlbaqwfichvtnlo
xdmgyjhprushlbaqwfichvtnzo
gdmoyjkpruszlbarwfichvtneo
cdmgyjkpruszlbaqwfcchvtned
xgmgyjkpruszlbaqwfschvtnek
xdmgyjkprusnlzamwfichvtneo
xdmgyjkprgszlbaxwfichvuneo
txmgyjksruszlbaqwfichvtneo
xdmgyjkprusbbbpqwfichvtneo
xdmoyjkpruszlbaqwfighvtxeo
xdmgyjkpruslhbaqwfichptneo
xdmgzjkpruszlbaqwffcmvtneo
xdmgyjkiruszlbaqgficuvtneo
vdbgyjkpruszlbaqwfichvtnek
xdmgyjspruszlbaqwfochvtney
xdmgyjkpruszibaqwfivhvteeo
xdmgyjkpruszfbaqwficbvtgeo
xdmgyjkprystlbaqwxichvtneo
xdmfyjkpryszlxaqwfichvtneo
xdmgyjgpruspybaqwfichvtneo
xdmgyjklruszlbjqwdichvtneo
xdmgyjkzruszltaqwfichvtnek
xdmgqjkpruszlzaqwfichvtneh
xdmgyjhnruszlbaqwficqvtneo
xdmgyjkproszlbaqweichvtnez
xdmgyjkprurzlbaawfichytneo
xdmgyfkpruszlbaqwfschutneo
xdmnyjkpruszlbaawjichvtneo
xdmgyjkpybszlbaqwfichvwneo
xdmgtjkhruszlbaqwfichatneo
xamgyjkprurzlbaqwfichvaneo
xdmgyjkpruszlbaqwgichvtnqv
ndmgyjkpruszlsaqwfuchvtneo
xdmgygkpgusrlbaqwfichvtneo
xdmgyjkpruszfbaqwfichvtnmy
xdmgyjkprupflbaqwfichvjneo
ndmgyjkpruszlbagwfichvtnxo
xdmgyjkpruszlbafwfilhvcneo
xdmgyjkpruszlbaqwfichvjsea
xebgyjkpruszlbaqafichvtneo
xdmkyjdpruszlbaqwfichvtnei
xomgyjkprufzlbaqwfochvtneo
xdmgyjkprfsllbaqwfiihvtneo
xdmyyjkpruszebaqwficmvtneo
xdmnyjkpruczlbarwfichvtneo
xdmgyjkpruszcbaqwbichvtneg
xdmgxjkpluszlbapwfichvtneo
xgrlyjkpruszlbaqwfichvtneo
xdmgyjkpruszlraqwxcchvtneo
xdmhyjupruszlbaqafichvtneo
xdmgnjkpruszlbkqwfjchvtneo
xdmgyjkpruszlwaqwfichvtndg
xdmgfjkpruvqlbaqwfichvtneo
xdmgejkptuszlbdqwfichvtneo
xlmgyjkpruszlnaqwfochvtneo
xdmgcjkpruszlbaqwfiqhvaneo
xdmgyjupruyzlbaywfichvtneo
gdmgyjkpruyzlbaqwficevtneo
xdmgyjkaruazlbapwfichvtneo
xsmiyjkpruszlbaqwfichvtveo
xdmiyjkprukzlbaqwfichvtnea
xdbgmjkxruszlbaqwfichvtneo
xdmgyjkpruskvbaqwfichdtneo
xdmgyjkprusznbaqwficshtneo
xdmgyjkprusrlbaqwfzchetneo
xdmgyrkpruszzbaqwfichvtned
xdmgyjkprusolbacwmichvtneo
xdmgypkpruszlbaqwfichvtmgo
xdmgyjkprumzlbhqwfichttneo
xdmgydkprusglbaqwfichvtnei
xdmuyjkpruszlbpqwfichvtyeo
xdmtymkprusslbaqwfichvtneo
xdmgyjjprkszlbaqwfqchvtneo
xdmgvjdpruszlbaqwfichgtneo
xdtgyjkpruwzlbaqwfjchvtneo
xdmgyjkpruszlbafseichvtneo
xdmgvjkpruszlraawfichvtneo
xdmgyukprgszlbatwfichvtneo
xhmgyjkpruszliaqwnichvtneo
xdmgyjspruszlbwqyfichvtneo
xdmgyjkjruszlqaqwfichvtnvo
xdmgyjkiruszlbnqwfichmtneo
ximgyjkpruszlbaqwfvcevtneo
xdmdyjkpruszlbaqwsithvtneo
ndmgyjkpruszlbaqwfilhatneo
xdmgyjkpruszlbaqwfinhvcnez
xdmgypkpsuszlbajwfichvtneo
xdpgmjkpluszlbaqwfichvtneo
xdmgyjnprupzlbaqwfichvtnel
xbmgyjkprmszlfaqwfichvtneo
xdmgyjkpausllbaqwfichvtseo
xdmgyjkpruszlbaqwfqchttnes
xgmgyjkpruszlbaxwfichvtneb
xdmgyjkpruszabqqwfichvineo
xdmgpjkpquszlbaqwfichvdneo
xdmgyjkeruszlbaqdficbvtneo
xdmaujkpruszlbaqwfichvteeo
xdmgyjkpruszlbaqwrirhvtnev
xdmgyjkpsugzllaqwfichvtneo
xdmgyjkpruszlbaqwfichctnlm
xdmeyjkpruszlbacwfiwhvtneo
xdmgyjkpiuhzlbaqwfijhvtneo
xdmgyjkpruszlbmqhfiohvtneo
xdegyjkpbuszlbbqwfichvtneo
xdmggxkpruszlbaqwfirhvtneo
xdmgojkpruszlbaqvfichvtteo
xdmgyjhtruszlbaqwmichvtneo
rdmgyjkpruszlbaqwfichvthek
xdlgyjqpruszlbaqwfbchvtneo
xdmgyjspriszlbavwfichvtneo
rdkgyjkpruszlbaqwfichvtnuo
tdmgyjkuruszlbaqwfichvtnev
xdmgyjkpxuszlbaqwfkchvtnso
xdegyjkpruszlbbqxfichvtneo
xdmgyjkpruszlbaqwficpvtket
xdmgyjkpruszliaqwfnchvtnec
xdmgyjkpreszlbaqwficdvtdeo
rdmgyjkpruszlbaywfychvtneo
xdmgywkpruszlbaqwficrvtaeo
xdmgyjkpruszlbanwflchvoneo
xdmgyjkpruyzlbaqufychvtneo
symgyjkpruszlbaqwfichvtqeo
xdmgyjkpruszlbaqwfichvbzqo
xzfgyjkpruszlbaqwfichvtveo
udmgyjepruszlbaqwfichbtneo
xhmgyjkpruszlbaqwfjchvtnef
xdhgyjkpruszlbaqaftchvtneo
xdmzyjkjruszlbaqwfichvtnwo
xdmgyjepruszlbaqwffchvtnef
xdmgyjkprurzlbaqwfikhvtneq
xomoyjkpruszkbaqwfichvtneo
xdmgyjkpiuszubaqwfichktneo
xdmgyjkprusdlbaqwhihhvtneo
xdmgyjkpruszlbaqwwirhvxneo
xdmgyjkpruszlbaqwficgitzeo
xdmgyjlpruszlbaqwfichpjneo
xjmgyjkpxuszlbaqwfichatneo
xdmgylkpruszlbaqwfiehvtnez
xdmgbjkpruszmbaqwfihhvtneo
xdmgyjkprubzlwaqwfichvtxeo
xdmgyjhlrustlbaqwfichvtneo
xdmmyjkpruszlbaqwfdchitneo
xdmgyjkpruszlbaqwoichhtbeo
xdzgyjkprvszlcaqwfichvtneo
ndmgyjkpruszlbaqwficavxneo
xdmgyjfpruszxbaqzfichvtneo
xdmgyjkpeuszlbaqzficdvtneo
xdmgyjkpruszlbmqffidhvtneo
xdnvyjkpruszlbafwfichvtneo
xdygyjkpruszlbljwfichvtneo
xdigyjkpruszlbeqwfuchvtneo
xdmgyjkpruszlbpzwfichvteeo
bdmgyjbpruszldaqwfichvtneo
xdmgyjkprrszlbaqmpichvtneo
idmgyjkpruszlbaqyfichvtkeo
xdmgyjkmrqsclbaqwfichvtneo
xdmgyjkpruazlbeqwfichvtxeo
ddmgyjkpruszlbsqwfichotneo
xdmgyqkpruszjbaqwfxchvtneo
xdmnyjkpruozlbaqwfichvtreo
edmgyjkpruszlbuqwwichvtneo
xdmgyjkprmshlbaqwfichctneo
xdmgyjkpruszlbaqwffghotneo
xdmcyjkprfszlbaqnfichvtneo
xdmgyjypruszhbaqwficyvtneo
xdmgyjkprzszlyaqwficmvtneo
xlmgyjkprzszlbaqwficyvtneo
xdmgyjkprutulbaqwfithvtneo
xdygyjkpruszlbpqwfichvpneo
xdmgsjkpoumzlbaqwfichvtneo
xdmgyjkpyuszlbaqdfnchvtneo
xdxgyjkpruszlbaqwfizhvtnjo
xdmgyjkpruszlbaqwfschvkndo
xdmgpjkprnszlcaqwfichvtneo
xhmgyjkpruszlbaqwficgvtnet
xdmgyjkpruswlbaqwfichvtqer
ddmgyjkprcszlbaqwfichqtneo
xdmgyjkpruhhlbaqwpichvtneo
xdmgyjkeraszlbaqwfichvtnso
nomgyjkpruszlbaqwficavxneo
xdmgyjkprdszlbaqwfobhvtneo
xdmgyjkprgszlbaqwfichvtdao
xomgyjspruswlbaqwfichvtneo
xdzgyjkpruszlbaqwficwvpneo
admgejkpruszlbaqwfimhvtneo
xdtgyjkpruszlmaqwfiqhvtneo
xdmgymkprusqlbaqwtichvtneo
xdmgyjkpluszlbaqwfidhvtnea
ztmgyjjpruszlbaqwfichvtneo`
