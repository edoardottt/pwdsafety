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

package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"io"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

//GetMD4Hash : Return MD4 hash of input
func GetMD4Hash(text string) string {
	h := md4.New()
	io.WriteString(h, text)
	return hex.EncodeToString(h.Sum(nil))
}

//GetMD5Hash : Return MD5 hash of input
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

//GetSHA1Hash : Return SHA1 hash of input
func GetSHA1Hash(text string) string {
	hash := sha1.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

//GetRipemd160Hash : Return Ripemd160 hash of input
func GetRipemd160Hash(text string) string {
	hasher := ripemd160.New()
	hasher.Write([]byte(text))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash[:])
}

//GetSHA224Hash : Return SHA224 hash of input
func GetSHA224Hash(text string) string {
	hash := sha256.Sum224([]byte(text))
	return hex.EncodeToString(hash[:])
}

//GetSHA256Hash : Return SHA256 hash of input
func GetSHA256Hash(text string) string {
	hash := sha256.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}

//GetSHA384Hash : Return SHA384 hash of input
func GetSHA384Hash(text string) string {
	hash := sha512.Sum384([]byte(text))
	return hex.EncodeToString(hash[:])
}

//GetSHA512Hash : Return SHA512 hash of input
func GetSHA512Hash(text string) string {
	hash := sha512.Sum512([]byte(text))
	return hex.EncodeToString(hash[:])
}

//GetSHA3224Hash : Return SHA3 224 hash of input
func GetSHA3224Hash(text string) string {
	hash := sha3.Sum224([]byte(text))
	return hex.EncodeToString(hash[:])
}

//GetSHA3256Hash : Return SHA3 256 hash of input
func GetSHA3256Hash(text string) string {
	hash := sha3.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}

//GetSHA3384Hash : Return SHA3 384 hash of input
func GetSHA3384Hash(text string) string {
	hash := sha3.Sum384([]byte(text))
	return hex.EncodeToString(hash[:])
}

//GetSHA3512Hash : Return SHA3 512 hash of input
func GetSHA3512Hash(text string) string {
	hash := sha3.Sum512([]byte(text))
	return hex.EncodeToString(hash[:])
}

//GetBlake2b256Hash : Return Blake2b256 hash of input
func GetBlake2b256Hash(text string) string {
	hash := blake2b.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}

//GetBlake2b384Hash : Return Blake2b384 hash of input
func GetBlake2b384Hash(text string) string {
	hash := blake2b.Sum384([]byte(text))
	return hex.EncodeToString(hash[:])
}

//GetBlake2b512Hash : Return Blake2b512 hash of input
func GetBlake2b512Hash(text string) string {
	hash := blake2b.Sum512([]byte(text))
	return hex.EncodeToString(hash[:])
}

//GetBlake2s256Hash : Return Blake2s256 hash of input
func GetBlake2s256Hash(text string) string {
	hash := blake2s.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}
