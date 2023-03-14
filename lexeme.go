package dorklang

type lexeme uint64

const (
	invalidLexeme lexeme = iota
	startProgramLexeme
	endProgramLexeme
	startAdditionSectionLexeme
	endAdditionSectionLexeme
	startSubtractionSectionLexeme
	endSubtractionSectionLexeme
	startMultiplicationSectionLexeme
	endMultiplicationSectionLexeme
	startDivisionSectionLexeme
	endDivisionSectionLexeme
	startJumpIfPositiveSectionLexeme
	endJumpIfPositiveSectionLexeme
	startJumpIfZeroSectionLexeme
	endJumpIfZeroSectionLexeme
	startCommentSectionLexeme
	endCommentSectionLexeme
	startReadFileSectionLexeme
	endReadFileSectionLexeme
	addOneLexeme
	addEightLexeme
	addStackPairLexeme
	addStackWholeLexeme
	subtractOneLexeme
	subtractEightLexeme
	subtractStackPairLexeme
	subtractStackWholeLexeme
	multiplyTwoLexeme
	multiplyEightLexeme
	multiplyStackPairLexeme
	multiplyStackWholeLexeme
	divideTwoLexeme
	divideEightLexeme
	divideStackPairLexeme
	divideStackWholeLexeme
	squareLexeme
	cubeLexeme
	setZeroLexeme
	setOneByteLexeme
	setEightByteLexeme
	setOneKibibyteLexeme
	setEightKibibyteLexeme
	setOneMebibyteLexeme
	setEightMebibyteLexeme
	setOneGibibyteLexeme
	setEightGibibyteLexeme
	setRandomByteLexeme
	setRandomMaxLexeme
	setSecondTimestampLexeme
	setNanosecondTimestampLexeme
	printCharacterLexeme
	printNumberLexeme
	inputCharacterLexeme
	inputNumberLexeme
	iotaFromZeroLexeme
	iotaFromOneLexeme
	logicalAndStackPairLexeme
	logicalAndStackWholeLexeme
	writeStackToFileLexeme
	readStackFromFileLexeme
	deleteFileLexeme
	clearStackLexeme
	resetStateLexeme
	pushStackLexeme
	countStackLexeme
	popStackLastLexeme
	popStackRandomLexeme
	useStackIndexZeroLexeme
	useStackIndexOneLexeme
	useStackIndexSwappedLexeme
	hashStackOneByteLexeme
	hashStackEightByteLexeme
	sortStackAscendingLexeme
	sortStackDescendingLexeme
	shuffleStackLexeme
	swapStackTopLexeme
	reverseStackLexeme
	filePathLexeme
	invertLexeme
	modifierLexeme
	changeDirLexeme
	separatorLexeme // used for whitespace
	emptyLexeme     // used by cleanTokens to replace unnecessary tokens
	parentLexeme    // used by cleanTokens to hold a child tokenCollection
)
