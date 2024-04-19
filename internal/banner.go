/*
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU General Public License for more details.
 *
 *    You should have received a copy of the GNU General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    Author:
 *      Edoardo Ottavianelli <edoardott@gmail.com>
 */

package beauty

const (
	Version = "v0.4.0"
)

// Beautify : If the terminal size is enough, print the label PWD-SAFETY.
func Beautify() {
	firstLine := "                    _            __      _            \n"
	secondLine := "   _ ____      ____| |___  __ _ / _| ___| |_ _   _    \n"
	thirdLine := "  | '_ \\ \\ /\\ / / _` / __|/ _` | |_ / _ \\ __| | | |   \n"
	fourthLine := "  | |_) \\ V  V / (_| \\__ \\ (_| |  _|  __/ |_| |_| |   \n"
	fifthLine := "  | .__/ \\_/\\_/ \\__,_|___/\\__,_|_|  \\___|\\__|\\__, |   \n"
	sixthLine := "  |_|                                        |___/   " + Version + "\n"
	seventhLine := " https://github.com/edoardottt/pwdsafety              \n"
	eigthLine := " edoardottt, https://www.edoardoottavianelli.it"
	beauty := firstLine + secondLine + thirdLine + fourthLine + fifthLine + sixthLine +
		seventhLine + eigthLine
	firstAdvice := " > Use a password manager\n"
	secondAdvice := " > Don't use the same password for different services\n"
	thirdAdvice := " > Enable 2FA wherever possible\n"
	println(beauty + firstAdvice + secondAdvice + thirdAdvice)
}
