package beauty

//Beautify : If the terminal size is enough, print the label PWD-SAFETY.
func Beautify() {
	zeroLine := " ========================================================\n"
	firstLine := " |  	              _            __      _            |\n"
	secondLine := " |   _ ____      ____| |___  __ _ / _| ___| |_ _   _    |\n"
	thirdLine := " |  | '_ \\ \\ /\\ / / _` / __|/ _` | |_ / _ \\ __| | | |   |\n"
	fourthLine := " |  | |_) \\ V  V / (_| \\__ \\ (_| |  _|  __/ |_| |_| |   |\n"
	fifthLine := " |  | .__/ \\_/\\_/ \\__,_|___/\\__,_|_|  \\___|\\__|\\__, |   |\n"
	sixthLine := " |  |_|                                        |___/    |\n"
	seventhLine := " |                                                      |\n"
	eigthLine := " | https://github.com/edoardottt/pwdsafety              |\n"
	ninethLine := " | edoardottt, https://www.edoardoottavianelli.it       |\n"
	tenthLine := " | v0.2                                                 |\n"
	eleventhLine := " ========================================================\n"
	beauty := zeroLine + firstLine + secondLine + thirdLine + fourthLine + fifthLine + sixthLine + seventhLine + eigthLine + ninethLine + tenthLine + eleventhLine
	firstAdvice := " -> Use a password manager\n"
	secondAdvice := " -> Don't use the same password for different services\n"
	thirdAdvice := " -> Enable 2FA when possible\n"
	println(beauty + "\n" + firstAdvice + secondAdvice + thirdAdvice)
}
