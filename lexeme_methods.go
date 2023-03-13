package dorklang

import (
	"strconv"
	"strings"
)

func (lexeme lexeme) String() string {
	var builder strings.Builder

	switch lexeme {
	case invalidLexeme:
		builder.WriteString("INVALID")
	case startProgramLexeme:
		builder.WriteString("START-PROGRAM")
	case endProgramLexeme:
		builder.WriteString("END-PROGRAM")
	case startAdditionSectionLexeme:
		builder.WriteString("START-ADD-SECT")
	case endAdditionSectionLexeme:
		builder.WriteString("END-ADD-SECT")
	case startSubtractionSectionLexeme:
		builder.WriteString("START-SUB-SECT")
	case endSubtractionSectionLexeme:
		builder.WriteString("END-SUB-SECT")
	case startMultiplicationSectionLexeme:
		builder.WriteString("START-MULT-SECT")
	case endMultiplicationSectionLexeme:
		builder.WriteString("END-MULT-SECT")
	case startDivisionSectionLexeme:
		builder.WriteString("START-DIV-SECT")
	case endDivisionSectionLexeme:
		builder.WriteString("END-DIV-SECT")
	case startJumpIfPositiveSectionLexeme:
		builder.WriteString("START-JMP-IF-POS-SECT")
	case endJumpIfPositiveSectionLexeme:
		builder.WriteString("END-JMP-IF-POS-SECT")
	case startJumpIfZeroSectionLexeme:
		builder.WriteString("START-JMP-IF-ZERO-SECT")
	case endJumpIfZeroSectionLexeme:
		builder.WriteString("END-JMP-IF-ZERO-SECT")
	case startCommentSectionLexeme:
		builder.WriteString("START-CMNT-SECT")
	case endCommentSectionLexeme:
		builder.WriteString("END-CMNT-SECT")
	case startReadFileSectionLexeme:
		builder.WriteString("START-READ-FILE-SECT")
	case endReadFileSectionLexeme:
		builder.WriteString("END-READ-FILE-SECT")
	case addOneLexeme:
		builder.WriteString("ADD-ONE")
	case addEightLexeme:
		builder.WriteString("ADD-EIGHT")
	case addStackPairLexeme:
		builder.WriteString("ADD-STACK-PAIR")
	case addStackWholeLexeme:
		builder.WriteString("ADD-STACK-WHOLE")
	case subtractOneLexeme:
		builder.WriteString("SUB-ONE")
	case subtractEightLexeme:
		builder.WriteString("SUB-EIGHT")
	case subtractStackPairLexeme:
		builder.WriteString("SUB-STACK-PAIR")
	case subtractStackWholeLexeme:
		builder.WriteString("SUB-STACK-WHOLE")
	case multiplyTwoLexeme:
		builder.WriteString("MULT-TWO")
	case multiplyEightLexeme:
		builder.WriteString("MULT-EIGHT")
	case multiplyStackPairLexeme:
		builder.WriteString("MULT-STACK-PAIR")
	case multiplyStackWholeLexeme:
		builder.WriteString("MULT-STACK-WHOLE")
	case divideTwoLexeme:
		builder.WriteString("DIV-TWO")
	case divideEightLexeme:
		builder.WriteString("DIV-EIGHT")
	case divideStackPairLexeme:
		builder.WriteString("DIV-STACK-PAIR")
	case divideStackWholeLexeme:
		builder.WriteString("DIV-STACK-WHOLE")
	case squareLexeme:
		builder.WriteString("SQUARE")
	case cubeLexeme:
		builder.WriteString("CUBE")
	case setZeroLexeme:
		builder.WriteString("SET-ZERO")
	case setOneByteLexeme:
		builder.WriteString("SET-ONE-BYTE")
	case setEightByteLexeme:
		builder.WriteString("SET-EIGHT-BYTE")
	case setOneKibibyteLexeme:
		builder.WriteString("SET-ONE-KIBI")
	case setEightKibibyteLexeme:
		builder.WriteString("SET-EIGHT-KIBI")
	case setOneMebibyteLexeme:
		builder.WriteString("SET-ONE-MEBI")
	case setEightMebibyteLexeme:
		builder.WriteString("SET-EIGHT-MEBI")
	case setOneGibibyteLexeme:
		builder.WriteString("SET-ONE-GIBI")
	case setEightGibibyteLexeme:
		builder.WriteString("SET-EIGHT-GIBI")
	case setRandomByteLexeme:
		builder.WriteString("SET-RAND-BYTE")
	case setRandomMaxLexeme:
		builder.WriteString("SET-RAND-MAX")
	case setSecondTimestampLexeme:
		builder.WriteString("SET-SEC-TIME")
	case setNanosecondTimestampLexeme:
		builder.WriteString("SET-NANO-TIME")
	case printCharacterLexeme:
		builder.WriteString("PRINT-CHAR")
	case printNumberLexeme:
		builder.WriteString("PRINT-NUM")
	case inputCharacterLexeme:
		builder.WriteString("INPUT-CHAR")
	case inputNumberLexeme:
		builder.WriteString("INPUT-NUM")
	case logicalAndStackPairLexeme:
		builder.WriteString("LOGIC-AND-STACK-PAIR")
	case logicalAndStackWholeLexeme:
		builder.WriteString("LOGIC-AND-STACK-WHOLE")
	case iotaFromZeroLexeme:
		builder.WriteString("IOTA-ZERO")
	case iotaFromOneLexeme:
		builder.WriteString("IOTA-ONE")
	case writeStackToFileLexeme:
		builder.WriteString("WRITE-STACK-FILE")
	case readStackFromFileLexeme:
		builder.WriteString("READ-STACK-FILE")
	case deleteFileLexeme:
		builder.WriteString("DELETE-STACK-FILE")
	case clearStackLexeme:
		builder.WriteString("CLEAR-STACK")
	case pushStackLexeme:
		builder.WriteString("PUSH-STACK")
	case popStackLastLexeme:
		builder.WriteString("POP-STACK-LAST")
	case popStackRandomLexeme:
		builder.WriteString("POP-STACK-RAND")
	case saveStackUseIndexZeroLexeme:
		builder.WriteString("USE-STACK-ZERO")
	case saveStackUseIndexOneLexeme:
		builder.WriteString("USE-STACK-ONE")
	case hashStackOneByteLexeme:
		builder.WriteString("HASH-STACK-ONE-BYTE")
	case hashStackEightByteLexeme:
		builder.WriteString("HASH-STACK-EIGHT-BYTE")
	case sortStackAscendingLexeme:
		builder.WriteString("SORT-STACK-ASC")
	case sortStackDescendingLexeme:
		builder.WriteString("SORT-STACK-DESC")
	case shuffleStackLexeme:
		builder.WriteString("SHUFFLE-STACK")
	case swapStackTopLexeme:
		builder.WriteString("SWAP-STACK-TOP")
	case reverseStackLexeme:
		builder.WriteString("REVERSE-STACK")
	case invertLexeme:
		builder.WriteString("INVERT")
	case modifierLexeme:
		builder.WriteString("MODIFIER")
	case plaintextLexeme:
		builder.WriteString("PLAINTEXT")
	case separatorLexeme:
		builder.WriteString("SEP")
	case emptyLexeme:
		builder.WriteString("EMPTY")
	default:
		builder.WriteString("UNKNOWN")
	}

	builder.WriteByte(' ')
	builder.WriteByte('[')
	builder.WriteString(strconv.FormatUint(uint64(lexeme), 10))
	builder.WriteByte(']')

	return builder.String()
}
