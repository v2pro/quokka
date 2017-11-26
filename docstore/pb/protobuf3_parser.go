// Generated from Protobuf3.g4 by ANTLR 4.7.

package pb // Protobuf3

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 59, 406,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 7, 2, 91, 10, 2, 12, 2, 14, 2, 94, 11, 2, 3, 2, 3, 2, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 5, 4, 105, 10, 4, 3, 4, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 5, 7, 125, 10, 7, 3, 7, 3, 7, 7, 7, 129, 10, 7, 12, 7, 14,
	7, 132, 11, 7, 3, 8, 3, 8, 3, 8, 5, 8, 137, 10, 8, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10,
	152, 10, 10, 12, 10, 14, 10, 155, 11, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 167, 10, 12, 12, 12, 14,
	12, 170, 11, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 5, 13, 177, 10, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 184, 10, 13, 12, 13, 14, 13,
	187, 11, 13, 3, 13, 3, 13, 5, 13, 191, 10, 13, 3, 13, 3, 13, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 205,
	10, 15, 12, 15, 14, 15, 208, 11, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16,
	3, 16, 5, 16, 216, 10, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 223,
	10, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 230, 10, 16, 12, 16,
	14, 16, 233, 11, 16, 3, 16, 3, 16, 5, 16, 237, 10, 16, 3, 17, 3, 17, 3,
	17, 5, 17, 242, 10, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 7, 18, 249,
	10, 18, 12, 18, 14, 18, 252, 11, 18, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19,
	258, 10, 19, 3, 20, 3, 20, 3, 20, 7, 20, 263, 10, 20, 12, 20, 14, 20, 266,
	11, 20, 3, 21, 3, 21, 5, 21, 270, 10, 21, 3, 22, 3, 22, 3, 23, 5, 23, 275,
	10, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23,
	285, 10, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 7, 24, 292, 10, 24, 12,
	24, 14, 24, 295, 11, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 7, 26, 306, 10, 26, 12, 26, 14, 26, 309, 11, 26, 3, 26, 3,
	26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 321,
	10, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28,
	3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 5, 28, 338, 10, 28, 3, 28, 3,
	28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 7, 30, 347, 10, 30, 12, 30, 14,
	30, 350, 11, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34,
	3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 5, 39, 369,
	10, 39, 3, 39, 3, 39, 7, 39, 373, 10, 39, 12, 39, 14, 39, 376, 11, 39,
	3, 39, 3, 39, 3, 40, 5, 40, 381, 10, 40, 3, 40, 3, 40, 7, 40, 385, 10,
	40, 12, 40, 14, 40, 388, 11, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 3,
	42, 5, 42, 396, 10, 42, 3, 42, 3, 42, 5, 42, 400, 10, 42, 3, 42, 3, 42,
	5, 42, 404, 10, 42, 3, 42, 2, 2, 43, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
	58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 2, 8, 3, 2, 18, 19,
	4, 2, 20, 20, 36, 36, 8, 2, 3, 5, 7, 9, 11, 12, 26, 29, 31, 31, 34, 35,
	8, 2, 3, 3, 7, 8, 11, 12, 26, 29, 31, 31, 34, 35, 3, 2, 54, 55, 3, 2, 40,
	41, 2, 418, 2, 84, 3, 2, 2, 2, 4, 97, 3, 2, 2, 2, 6, 102, 3, 2, 2, 2, 8,
	109, 3, 2, 2, 2, 10, 113, 3, 2, 2, 2, 12, 124, 3, 2, 2, 2, 14, 136, 3,
	2, 2, 2, 16, 138, 3, 2, 2, 2, 18, 142, 3, 2, 2, 2, 20, 158, 3, 2, 2, 2,
	22, 162, 3, 2, 2, 2, 24, 173, 3, 2, 2, 2, 26, 194, 3, 2, 2, 2, 28, 198,
	3, 2, 2, 2, 30, 211, 3, 2, 2, 2, 32, 238, 3, 2, 2, 2, 34, 245, 3, 2, 2,
	2, 36, 257, 3, 2, 2, 2, 38, 259, 3, 2, 2, 2, 40, 269, 3, 2, 2, 2, 42, 271,
	3, 2, 2, 2, 44, 274, 3, 2, 2, 2, 46, 288, 3, 2, 2, 2, 48, 296, 3, 2, 2,
	2, 50, 300, 3, 2, 2, 2, 52, 312, 3, 2, 2, 2, 54, 324, 3, 2, 2, 2, 56, 341,
	3, 2, 2, 2, 58, 343, 3, 2, 2, 2, 60, 351, 3, 2, 2, 2, 62, 353, 3, 2, 2,
	2, 64, 355, 3, 2, 2, 2, 66, 357, 3, 2, 2, 2, 68, 359, 3, 2, 2, 2, 70, 361,
	3, 2, 2, 2, 72, 363, 3, 2, 2, 2, 74, 365, 3, 2, 2, 2, 76, 368, 3, 2, 2,
	2, 78, 380, 3, 2, 2, 2, 80, 391, 3, 2, 2, 2, 82, 403, 3, 2, 2, 2, 84, 92,
	5, 4, 3, 2, 85, 91, 5, 6, 4, 2, 86, 91, 5, 8, 5, 2, 87, 91, 5, 10, 6, 2,
	88, 91, 5, 14, 8, 2, 89, 91, 5, 80, 41, 2, 90, 85, 3, 2, 2, 2, 90, 86,
	3, 2, 2, 2, 90, 87, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 90, 89, 3, 2, 2, 2,
	91, 94, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 95, 3,
	2, 2, 2, 94, 92, 3, 2, 2, 2, 95, 96, 7, 2, 2, 3, 96, 3, 3, 2, 2, 2, 97,
	98, 7, 32, 2, 2, 98, 99, 7, 56, 2, 2, 99, 100, 9, 2, 2, 2, 100, 101, 7,
	51, 2, 2, 101, 5, 3, 2, 2, 2, 102, 104, 7, 10, 2, 2, 103, 105, 9, 3, 2,
	2, 104, 103, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106,
	107, 7, 41, 2, 2, 107, 108, 7, 51, 2, 2, 108, 7, 3, 2, 2, 2, 109, 110,
	7, 17, 2, 2, 110, 111, 5, 58, 30, 2, 111, 112, 7, 51, 2, 2, 112, 9, 3,
	2, 2, 2, 113, 114, 7, 16, 2, 2, 114, 115, 5, 12, 7, 2, 115, 116, 7, 56,
	2, 2, 116, 117, 5, 82, 42, 2, 117, 118, 7, 51, 2, 2, 118, 11, 3, 2, 2,
	2, 119, 125, 7, 37, 2, 2, 120, 121, 7, 43, 2, 2, 121, 122, 5, 58, 30, 2,
	122, 123, 7, 44, 2, 2, 123, 125, 3, 2, 2, 2, 124, 119, 3, 2, 2, 2, 124,
	120, 3, 2, 2, 2, 125, 130, 3, 2, 2, 2, 126, 127, 7, 53, 2, 2, 127, 129,
	7, 37, 2, 2, 128, 126, 3, 2, 2, 2, 129, 132, 3, 2, 2, 2, 130, 128, 3, 2,
	2, 2, 130, 131, 3, 2, 2, 2, 131, 13, 3, 2, 2, 2, 132, 130, 3, 2, 2, 2,
	133, 137, 5, 16, 9, 2, 134, 137, 5, 20, 11, 2, 135, 137, 5, 28, 15, 2,
	136, 133, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 136, 135, 3, 2, 2, 2, 137,
	15, 3, 2, 2, 2, 138, 139, 7, 14, 2, 2, 139, 140, 5, 60, 31, 2, 140, 141,
	5, 18, 10, 2, 141, 17, 3, 2, 2, 2, 142, 153, 7, 45, 2, 2, 143, 152, 5,
	44, 23, 2, 144, 152, 5, 20, 11, 2, 145, 152, 5, 16, 9, 2, 146, 152, 5,
	10, 6, 2, 147, 152, 5, 50, 26, 2, 148, 152, 5, 54, 28, 2, 149, 152, 5,
	32, 17, 2, 150, 152, 5, 80, 41, 2, 151, 143, 3, 2, 2, 2, 151, 144, 3, 2,
	2, 2, 151, 145, 3, 2, 2, 2, 151, 146, 3, 2, 2, 2, 151, 147, 3, 2, 2, 2,
	151, 148, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2, 151, 150, 3, 2, 2, 2, 152,
	155, 3, 2, 2, 2, 153, 151, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 156,
	3, 2, 2, 2, 155, 153, 3, 2, 2, 2, 156, 157, 7, 46, 2, 2, 157, 19, 3, 2,
	2, 2, 158, 159, 7, 6, 2, 2, 159, 160, 5, 62, 32, 2, 160, 161, 5, 22, 12,
	2, 161, 21, 3, 2, 2, 2, 162, 168, 7, 45, 2, 2, 163, 167, 5, 10, 6, 2, 164,
	167, 5, 24, 13, 2, 165, 167, 5, 80, 41, 2, 166, 163, 3, 2, 2, 2, 166, 164,
	3, 2, 2, 2, 166, 165, 3, 2, 2, 2, 167, 170, 3, 2, 2, 2, 168, 166, 3, 2,
	2, 2, 168, 169, 3, 2, 2, 2, 169, 171, 3, 2, 2, 2, 170, 168, 3, 2, 2, 2,
	171, 172, 7, 46, 2, 2, 172, 23, 3, 2, 2, 2, 173, 174, 7, 37, 2, 2, 174,
	176, 7, 56, 2, 2, 175, 177, 7, 54, 2, 2, 176, 175, 3, 2, 2, 2, 176, 177,
	3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 190, 7, 38, 2, 2, 179, 180, 7, 47,
	2, 2, 180, 185, 5, 26, 14, 2, 181, 182, 7, 52, 2, 2, 182, 184, 5, 26, 14,
	2, 183, 181, 3, 2, 2, 2, 184, 187, 3, 2, 2, 2, 185, 183, 3, 2, 2, 2, 185,
	186, 3, 2, 2, 2, 186, 188, 3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 188, 189,
	7, 48, 2, 2, 189, 191, 3, 2, 2, 2, 190, 179, 3, 2, 2, 2, 190, 191, 3, 2,
	2, 2, 191, 192, 3, 2, 2, 2, 192, 193, 7, 51, 2, 2, 193, 25, 3, 2, 2, 2,
	194, 195, 5, 12, 7, 2, 195, 196, 7, 56, 2, 2, 196, 197, 5, 82, 42, 2, 197,
	27, 3, 2, 2, 2, 198, 199, 7, 25, 2, 2, 199, 200, 5, 72, 37, 2, 200, 206,
	7, 45, 2, 2, 201, 205, 5, 10, 6, 2, 202, 205, 5, 30, 16, 2, 203, 205, 5,
	80, 41, 2, 204, 201, 3, 2, 2, 2, 204, 202, 3, 2, 2, 2, 204, 203, 3, 2,
	2, 2, 205, 208, 3, 2, 2, 2, 206, 204, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2,
	207, 209, 3, 2, 2, 2, 208, 206, 3, 2, 2, 2, 209, 210, 7, 46, 2, 2, 210,
	29, 3, 2, 2, 2, 211, 212, 7, 24, 2, 2, 212, 213, 5, 74, 38, 2, 213, 215,
	7, 43, 2, 2, 214, 216, 7, 30, 2, 2, 215, 214, 3, 2, 2, 2, 215, 216, 3,
	2, 2, 2, 216, 217, 3, 2, 2, 2, 217, 218, 5, 76, 39, 2, 218, 219, 7, 44,
	2, 2, 219, 220, 7, 23, 2, 2, 220, 222, 7, 43, 2, 2, 221, 223, 7, 30, 2,
	2, 222, 221, 3, 2, 2, 2, 222, 223, 3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224,
	225, 5, 76, 39, 2, 225, 236, 7, 44, 2, 2, 226, 231, 7, 45, 2, 2, 227, 230,
	5, 10, 6, 2, 228, 230, 5, 80, 41, 2, 229, 227, 3, 2, 2, 2, 229, 228, 3,
	2, 2, 2, 230, 233, 3, 2, 2, 2, 231, 229, 3, 2, 2, 2, 231, 232, 3, 2, 2,
	2, 232, 234, 3, 2, 2, 2, 233, 231, 3, 2, 2, 2, 234, 237, 7, 46, 2, 2, 235,
	237, 7, 51, 2, 2, 236, 226, 3, 2, 2, 2, 236, 235, 3, 2, 2, 2, 237, 31,
	3, 2, 2, 2, 238, 241, 7, 22, 2, 2, 239, 242, 5, 34, 18, 2, 240, 242, 5,
	38, 20, 2, 241, 239, 3, 2, 2, 2, 241, 240, 3, 2, 2, 2, 242, 243, 3, 2,
	2, 2, 243, 244, 7, 51, 2, 2, 244, 33, 3, 2, 2, 2, 245, 250, 5, 36, 19,
	2, 246, 247, 7, 52, 2, 2, 247, 249, 5, 36, 19, 2, 248, 246, 3, 2, 2, 2,
	249, 252, 3, 2, 2, 2, 250, 248, 3, 2, 2, 2, 250, 251, 3, 2, 2, 2, 251,
	35, 3, 2, 2, 2, 252, 250, 3, 2, 2, 2, 253, 258, 7, 38, 2, 2, 254, 255,
	7, 38, 2, 2, 255, 256, 7, 33, 2, 2, 256, 258, 7, 38, 2, 2, 257, 253, 3,
	2, 2, 2, 257, 254, 3, 2, 2, 2, 258, 37, 3, 2, 2, 2, 259, 264, 7, 41, 2,
	2, 260, 261, 7, 52, 2, 2, 261, 263, 7, 41, 2, 2, 262, 260, 3, 2, 2, 2,
	263, 266, 3, 2, 2, 2, 264, 262, 3, 2, 2, 2, 264, 265, 3, 2, 2, 2, 265,
	39, 3, 2, 2, 2, 266, 264, 3, 2, 2, 2, 267, 270, 9, 4, 2, 2, 268, 270, 5,
	78, 40, 2, 269, 267, 3, 2, 2, 2, 269, 268, 3, 2, 2, 2, 270, 41, 3, 2, 2,
	2, 271, 272, 7, 38, 2, 2, 272, 43, 3, 2, 2, 2, 273, 275, 7, 21, 2, 2, 274,
	273, 3, 2, 2, 2, 274, 275, 3, 2, 2, 2, 275, 276, 3, 2, 2, 2, 276, 277,
	5, 40, 21, 2, 277, 278, 5, 66, 34, 2, 278, 279, 7, 56, 2, 2, 279, 284,
	5, 42, 22, 2, 280, 281, 7, 47, 2, 2, 281, 282, 5, 46, 24, 2, 282, 283,
	7, 48, 2, 2, 283, 285, 3, 2, 2, 2, 284, 280, 3, 2, 2, 2, 284, 285, 3, 2,
	2, 2, 285, 286, 3, 2, 2, 2, 286, 287, 7, 51, 2, 2, 287, 45, 3, 2, 2, 2,
	288, 293, 5, 48, 25, 2, 289, 290, 7, 52, 2, 2, 290, 292, 5, 48, 25, 2,
	291, 289, 3, 2, 2, 2, 292, 295, 3, 2, 2, 2, 293, 291, 3, 2, 2, 2, 293,
	294, 3, 2, 2, 2, 294, 47, 3, 2, 2, 2, 295, 293, 3, 2, 2, 2, 296, 297, 5,
	12, 7, 2, 297, 298, 7, 56, 2, 2, 298, 299, 5, 82, 42, 2, 299, 49, 3, 2,
	2, 2, 300, 301, 7, 15, 2, 2, 301, 302, 5, 68, 35, 2, 302, 307, 7, 45, 2,
	2, 303, 306, 5, 52, 27, 2, 304, 306, 5, 80, 41, 2, 305, 303, 3, 2, 2, 2,
	305, 304, 3, 2, 2, 2, 306, 309, 3, 2, 2, 2, 307, 305, 3, 2, 2, 2, 307,
	308, 3, 2, 2, 2, 308, 310, 3, 2, 2, 2, 309, 307, 3, 2, 2, 2, 310, 311,
	7, 46, 2, 2, 311, 51, 3, 2, 2, 2, 312, 313, 5, 40, 21, 2, 313, 314, 5,
	66, 34, 2, 314, 315, 7, 56, 2, 2, 315, 320, 5, 42, 22, 2, 316, 317, 7,
	47, 2, 2, 317, 318, 5, 46, 24, 2, 318, 319, 7, 48, 2, 2, 319, 321, 3, 2,
	2, 2, 320, 316, 3, 2, 2, 2, 320, 321, 3, 2, 2, 2, 321, 322, 3, 2, 2, 2,
	322, 323, 7, 51, 2, 2, 323, 53, 3, 2, 2, 2, 324, 325, 7, 13, 2, 2, 325,
	326, 7, 49, 2, 2, 326, 327, 5, 56, 29, 2, 327, 328, 7, 52, 2, 2, 328, 329,
	5, 40, 21, 2, 329, 330, 7, 50, 2, 2, 330, 331, 5, 70, 36, 2, 331, 332,
	7, 56, 2, 2, 332, 337, 5, 42, 22, 2, 333, 334, 7, 47, 2, 2, 334, 335, 5,
	46, 24, 2, 335, 336, 7, 48, 2, 2, 336, 338, 3, 2, 2, 2, 337, 333, 3, 2,
	2, 2, 337, 338, 3, 2, 2, 2, 338, 339, 3, 2, 2, 2, 339, 340, 7, 51, 2, 2,
	340, 55, 3, 2, 2, 2, 341, 342, 9, 5, 2, 2, 342, 57, 3, 2, 2, 2, 343, 348,
	7, 37, 2, 2, 344, 345, 7, 53, 2, 2, 345, 347, 7, 37, 2, 2, 346, 344, 3,
	2, 2, 2, 347, 350, 3, 2, 2, 2, 348, 346, 3, 2, 2, 2, 348, 349, 3, 2, 2,
	2, 349, 59, 3, 2, 2, 2, 350, 348, 3, 2, 2, 2, 351, 352, 7, 37, 2, 2, 352,
	61, 3, 2, 2, 2, 353, 354, 7, 37, 2, 2, 354, 63, 3, 2, 2, 2, 355, 356, 7,
	37, 2, 2, 356, 65, 3, 2, 2, 2, 357, 358, 7, 37, 2, 2, 358, 67, 3, 2, 2,
	2, 359, 360, 7, 37, 2, 2, 360, 69, 3, 2, 2, 2, 361, 362, 7, 37, 2, 2, 362,
	71, 3, 2, 2, 2, 363, 364, 7, 37, 2, 2, 364, 73, 3, 2, 2, 2, 365, 366, 7,
	37, 2, 2, 366, 75, 3, 2, 2, 2, 367, 369, 7, 53, 2, 2, 368, 367, 3, 2, 2,
	2, 368, 369, 3, 2, 2, 2, 369, 374, 3, 2, 2, 2, 370, 371, 7, 37, 2, 2, 371,
	373, 7, 53, 2, 2, 372, 370, 3, 2, 2, 2, 373, 376, 3, 2, 2, 2, 374, 372,
	3, 2, 2, 2, 374, 375, 3, 2, 2, 2, 375, 377, 3, 2, 2, 2, 376, 374, 3, 2,
	2, 2, 377, 378, 5, 60, 31, 2, 378, 77, 3, 2, 2, 2, 379, 381, 7, 53, 2,
	2, 380, 379, 3, 2, 2, 2, 380, 381, 3, 2, 2, 2, 381, 386, 3, 2, 2, 2, 382,
	383, 7, 37, 2, 2, 383, 385, 7, 53, 2, 2, 384, 382, 3, 2, 2, 2, 385, 388,
	3, 2, 2, 2, 386, 384, 3, 2, 2, 2, 386, 387, 3, 2, 2, 2, 387, 389, 3, 2,
	2, 2, 388, 386, 3, 2, 2, 2, 389, 390, 5, 64, 33, 2, 390, 79, 3, 2, 2, 2,
	391, 392, 7, 51, 2, 2, 392, 81, 3, 2, 2, 2, 393, 404, 5, 58, 30, 2, 394,
	396, 9, 6, 2, 2, 395, 394, 3, 2, 2, 2, 395, 396, 3, 2, 2, 2, 396, 397,
	3, 2, 2, 2, 397, 404, 7, 38, 2, 2, 398, 400, 9, 6, 2, 2, 399, 398, 3, 2,
	2, 2, 399, 400, 3, 2, 2, 2, 400, 401, 3, 2, 2, 2, 401, 404, 7, 39, 2, 2,
	402, 404, 9, 7, 2, 2, 403, 393, 3, 2, 2, 2, 403, 395, 3, 2, 2, 2, 403,
	399, 3, 2, 2, 2, 403, 402, 3, 2, 2, 2, 404, 83, 3, 2, 2, 2, 42, 90, 92,
	104, 124, 130, 136, 151, 153, 166, 168, 176, 185, 190, 204, 206, 215, 222,
	229, 231, 236, 241, 250, 257, 264, 269, 274, 284, 293, 305, 307, 320, 337,
	348, 368, 374, 380, 386, 395, 399, 403,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'bool'", "'bytes'", "'double'", "'enum'", "'fixed32'", "'fixed64'",
	"'float'", "'import'", "'int32'", "'int64'", "'map'", "'message'", "'oneof'",
	"'option'", "'package'", "'\"proto3\"'", "''proto3''", "'public'", "'repeated'",
	"'reserved'", "'returns'", "'rpc'", "'service'", "'sfixed32'", "'sfixed64'",
	"'sint32'", "'sint64'", "'stream'", "'string'", "'syntax'", "'to'", "'uint32'",
	"'uint64'", "'weak'", "", "", "", "", "", "", "'('", "')'", "'{'", "'}'",
	"'['", "']'", "'<'", "'>'", "';'", "','", "'.'", "'-'", "'+'", "'='",
}
var symbolicNames = []string{
	"", "BOOL", "BYTES", "DOUBLE", "ENUM", "FIXED32", "FIXED64", "FLOAT", "IMPORT",
	"INT32", "INT64", "MAP", "MESSAGE", "ONEOF", "OPTION", "PACKAGE", "PROTO3_DOUBLE",
	"PROTO3_SINGLE", "PUBLIC", "REPEATED", "RESERVED", "RETURNS", "RPC", "SERVICE",
	"SFIXED32", "SFIXED64", "SINT32", "SINT64", "STREAM", "STRING", "SYNTAX",
	"TO", "UINT32", "UINT64", "WEAK", "Ident", "IntLit", "FloatLit", "BoolLit",
	"StrLit", "Quote", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK",
	"LCHEVR", "RCHEVR", "SEMI", "COMMA", "DOT", "MINUS", "PLUS", "ASSIGN",
	"WS", "COMMENT", "LINE_COMMENT",
}

var ruleNames = []string{
	"proto", "syntax", "importStatement", "packageStatement", "option", "optionName",
	"topLevelDef", "message", "messageBody", "enumDefinition", "enumBody",
	"enumField", "enumValueOption", "service", "rpc", "reserved", "pbranges",
	"pbrange", "fieldNames", "pbtype", "fieldNumber", "field", "fieldOptions",
	"fieldOption", "oneof", "oneofField", "mapField", "keyType", "fullIdent",
	"messageName", "enumName", "messageOrEnumName", "fieldName", "oneofName",
	"mapName", "serviceName", "rpcName", "messageType", "messageOrEnumType",
	"emptyStatement", "constant",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type Protobuf3Parser struct {
	*antlr.BaseParser
}

func NewProtobuf3Parser(input antlr.TokenStream) *Protobuf3Parser {
	this := new(Protobuf3Parser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Protobuf3.g4"

	return this
}

// Protobuf3Parser tokens.
const (
	Protobuf3ParserEOF           = antlr.TokenEOF
	Protobuf3ParserBOOL          = 1
	Protobuf3ParserBYTES         = 2
	Protobuf3ParserDOUBLE        = 3
	Protobuf3ParserENUM          = 4
	Protobuf3ParserFIXED32       = 5
	Protobuf3ParserFIXED64       = 6
	Protobuf3ParserFLOAT         = 7
	Protobuf3ParserIMPORT        = 8
	Protobuf3ParserINT32         = 9
	Protobuf3ParserINT64         = 10
	Protobuf3ParserMAP           = 11
	Protobuf3ParserMESSAGE       = 12
	Protobuf3ParserONEOF         = 13
	Protobuf3ParserOPTION        = 14
	Protobuf3ParserPACKAGE       = 15
	Protobuf3ParserPROTO3_DOUBLE = 16
	Protobuf3ParserPROTO3_SINGLE = 17
	Protobuf3ParserPUBLIC        = 18
	Protobuf3ParserREPEATED      = 19
	Protobuf3ParserRESERVED      = 20
	Protobuf3ParserRETURNS       = 21
	Protobuf3ParserRPC           = 22
	Protobuf3ParserSERVICE       = 23
	Protobuf3ParserSFIXED32      = 24
	Protobuf3ParserSFIXED64      = 25
	Protobuf3ParserSINT32        = 26
	Protobuf3ParserSINT64        = 27
	Protobuf3ParserSTREAM        = 28
	Protobuf3ParserSTRING        = 29
	Protobuf3ParserSYNTAX        = 30
	Protobuf3ParserTO            = 31
	Protobuf3ParserUINT32        = 32
	Protobuf3ParserUINT64        = 33
	Protobuf3ParserWEAK          = 34
	Protobuf3ParserIdent         = 35
	Protobuf3ParserIntLit        = 36
	Protobuf3ParserFloatLit      = 37
	Protobuf3ParserBoolLit       = 38
	Protobuf3ParserStrLit        = 39
	Protobuf3ParserQuote         = 40
	Protobuf3ParserLPAREN        = 41
	Protobuf3ParserRPAREN        = 42
	Protobuf3ParserLBRACE        = 43
	Protobuf3ParserRBRACE        = 44
	Protobuf3ParserLBRACK        = 45
	Protobuf3ParserRBRACK        = 46
	Protobuf3ParserLCHEVR        = 47
	Protobuf3ParserRCHEVR        = 48
	Protobuf3ParserSEMI          = 49
	Protobuf3ParserCOMMA         = 50
	Protobuf3ParserDOT           = 51
	Protobuf3ParserMINUS         = 52
	Protobuf3ParserPLUS          = 53
	Protobuf3ParserASSIGN        = 54
	Protobuf3ParserWS            = 55
	Protobuf3ParserCOMMENT       = 56
	Protobuf3ParserLINE_COMMENT  = 57
)

// Protobuf3Parser rules.
const (
	Protobuf3ParserRULE_proto             = 0
	Protobuf3ParserRULE_syntax            = 1
	Protobuf3ParserRULE_importStatement   = 2
	Protobuf3ParserRULE_packageStatement  = 3
	Protobuf3ParserRULE_option            = 4
	Protobuf3ParserRULE_optionName        = 5
	Protobuf3ParserRULE_topLevelDef       = 6
	Protobuf3ParserRULE_message           = 7
	Protobuf3ParserRULE_messageBody       = 8
	Protobuf3ParserRULE_enumDefinition    = 9
	Protobuf3ParserRULE_enumBody          = 10
	Protobuf3ParserRULE_enumField         = 11
	Protobuf3ParserRULE_enumValueOption   = 12
	Protobuf3ParserRULE_service           = 13
	Protobuf3ParserRULE_rpc               = 14
	Protobuf3ParserRULE_reserved          = 15
	Protobuf3ParserRULE_pbranges          = 16
	Protobuf3ParserRULE_pbrange           = 17
	Protobuf3ParserRULE_fieldNames        = 18
	Protobuf3ParserRULE_pbtype            = 19
	Protobuf3ParserRULE_fieldNumber       = 20
	Protobuf3ParserRULE_field             = 21
	Protobuf3ParserRULE_fieldOptions      = 22
	Protobuf3ParserRULE_fieldOption       = 23
	Protobuf3ParserRULE_oneof             = 24
	Protobuf3ParserRULE_oneofField        = 25
	Protobuf3ParserRULE_mapField          = 26
	Protobuf3ParserRULE_keyType           = 27
	Protobuf3ParserRULE_fullIdent         = 28
	Protobuf3ParserRULE_messageName       = 29
	Protobuf3ParserRULE_enumName          = 30
	Protobuf3ParserRULE_messageOrEnumName = 31
	Protobuf3ParserRULE_fieldName         = 32
	Protobuf3ParserRULE_oneofName         = 33
	Protobuf3ParserRULE_mapName           = 34
	Protobuf3ParserRULE_serviceName       = 35
	Protobuf3ParserRULE_rpcName           = 36
	Protobuf3ParserRULE_messageType       = 37
	Protobuf3ParserRULE_messageOrEnumType = 38
	Protobuf3ParserRULE_emptyStatement    = 39
	Protobuf3ParserRULE_constant          = 40
)

// IProtoContext is an interface to support dynamic dispatch.
type IProtoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProtoContext differentiates from other interfaces.
	IsProtoContext()
}

type ProtoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProtoContext() *ProtoContext {
	var p = new(ProtoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_proto
	return p
}

func (*ProtoContext) IsProtoContext() {}

func NewProtoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProtoContext {
	var p = new(ProtoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_proto

	return p
}

func (s *ProtoContext) GetParser() antlr.Parser { return s.parser }

func (s *ProtoContext) Syntax() ISyntaxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISyntaxContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISyntaxContext)
}

func (s *ProtoContext) EOF() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserEOF, 0)
}

func (s *ProtoContext) AllImportStatement() []IImportStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportStatementContext)(nil)).Elem())
	var tst = make([]IImportStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportStatementContext)
		}
	}

	return tst
}

func (s *ProtoContext) ImportStatement(i int) IImportStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportStatementContext)
}

func (s *ProtoContext) AllPackageStatement() []IPackageStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPackageStatementContext)(nil)).Elem())
	var tst = make([]IPackageStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPackageStatementContext)
		}
	}

	return tst
}

func (s *ProtoContext) PackageStatement(i int) IPackageStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPackageStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPackageStatementContext)
}

func (s *ProtoContext) AllOption() []IOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptionContext)(nil)).Elem())
	var tst = make([]IOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptionContext)
		}
	}

	return tst
}

func (s *ProtoContext) Option(i int) IOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *ProtoContext) AllTopLevelDef() []ITopLevelDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITopLevelDefContext)(nil)).Elem())
	var tst = make([]ITopLevelDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITopLevelDefContext)
		}
	}

	return tst
}

func (s *ProtoContext) TopLevelDef(i int) ITopLevelDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITopLevelDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITopLevelDefContext)
}

func (s *ProtoContext) AllEmptyStatement() []IEmptyStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEmptyStatementContext)(nil)).Elem())
	var tst = make([]IEmptyStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEmptyStatementContext)
		}
	}

	return tst
}

func (s *ProtoContext) EmptyStatement(i int) IEmptyStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEmptyStatementContext)
}

func (s *ProtoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProtoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProtoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterProto(s)
	}
}

func (s *ProtoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitProto(s)
	}
}

func (p *Protobuf3Parser) Proto() (localctx IProtoContext) {
	localctx = NewProtoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Protobuf3ParserRULE_proto)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Syntax()
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Protobuf3ParserENUM)|(1<<Protobuf3ParserIMPORT)|(1<<Protobuf3ParserMESSAGE)|(1<<Protobuf3ParserOPTION)|(1<<Protobuf3ParserPACKAGE)|(1<<Protobuf3ParserSERVICE))) != 0) || _la == Protobuf3ParserSEMI {
		p.SetState(88)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case Protobuf3ParserIMPORT:
			{
				p.SetState(83)
				p.ImportStatement()
			}

		case Protobuf3ParserPACKAGE:
			{
				p.SetState(84)
				p.PackageStatement()
			}

		case Protobuf3ParserOPTION:
			{
				p.SetState(85)
				p.Option()
			}

		case Protobuf3ParserENUM, Protobuf3ParserMESSAGE, Protobuf3ParserSERVICE:
			{
				p.SetState(86)
				p.TopLevelDef()
			}

		case Protobuf3ParserSEMI:
			{
				p.SetState(87)
				p.EmptyStatement()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(93)
		p.Match(Protobuf3ParserEOF)
	}

	return localctx
}

// ISyntaxContext is an interface to support dynamic dispatch.
type ISyntaxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSyntaxContext differentiates from other interfaces.
	IsSyntaxContext()
}

type SyntaxContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySyntaxContext() *SyntaxContext {
	var p = new(SyntaxContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_syntax
	return p
}

func (*SyntaxContext) IsSyntaxContext() {}

func NewSyntaxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SyntaxContext {
	var p = new(SyntaxContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_syntax

	return p
}

func (s *SyntaxContext) GetParser() antlr.Parser { return s.parser }
func (s *SyntaxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SyntaxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SyntaxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterSyntax(s)
	}
}

func (s *SyntaxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitSyntax(s)
	}
}

func (p *Protobuf3Parser) Syntax() (localctx ISyntaxContext) {
	localctx = NewSyntaxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Protobuf3ParserRULE_syntax)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(Protobuf3ParserSYNTAX)
	}
	{
		p.SetState(96)
		p.Match(Protobuf3ParserASSIGN)
	}
	p.SetState(97)
	_la = p.GetTokenStream().LA(1)

	if !(_la == Protobuf3ParserPROTO3_DOUBLE || _la == Protobuf3ParserPROTO3_SINGLE) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(98)
		p.Match(Protobuf3ParserSEMI)
	}

	return localctx
}

// IImportStatementContext is an interface to support dynamic dispatch.
type IImportStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportStatementContext differentiates from other interfaces.
	IsImportStatementContext()
}

type ImportStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStatementContext() *ImportStatementContext {
	var p = new(ImportStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_importStatement
	return p
}

func (*ImportStatementContext) IsImportStatementContext() {}

func NewImportStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStatementContext {
	var p = new(ImportStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_importStatement

	return p
}

func (s *ImportStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStatementContext) StrLit() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserStrLit, 0)
}

func (s *ImportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterImportStatement(s)
	}
}

func (s *ImportStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitImportStatement(s)
	}
}

func (p *Protobuf3Parser) ImportStatement() (localctx IImportStatementContext) {
	localctx = NewImportStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Protobuf3ParserRULE_importStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(Protobuf3ParserIMPORT)
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserPUBLIC || _la == Protobuf3ParserWEAK {
		p.SetState(101)
		_la = p.GetTokenStream().LA(1)

		if !(_la == Protobuf3ParserPUBLIC || _la == Protobuf3ParserWEAK) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	}
	{
		p.SetState(104)
		p.Match(Protobuf3ParserStrLit)
	}
	{
		p.SetState(105)
		p.Match(Protobuf3ParserSEMI)
	}

	return localctx
}

// IPackageStatementContext is an interface to support dynamic dispatch.
type IPackageStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPackageStatementContext differentiates from other interfaces.
	IsPackageStatementContext()
}

type PackageStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageStatementContext() *PackageStatementContext {
	var p = new(PackageStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_packageStatement
	return p
}

func (*PackageStatementContext) IsPackageStatementContext() {}

func NewPackageStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageStatementContext {
	var p = new(PackageStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_packageStatement

	return p
}

func (s *PackageStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageStatementContext) FullIdent() IFullIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *PackageStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterPackageStatement(s)
	}
}

func (s *PackageStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitPackageStatement(s)
	}
}

func (p *Protobuf3Parser) PackageStatement() (localctx IPackageStatementContext) {
	localctx = NewPackageStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Protobuf3ParserRULE_packageStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.Match(Protobuf3ParserPACKAGE)
	}
	{
		p.SetState(108)
		p.FullIdent()
	}
	{
		p.SetState(109)
		p.Match(Protobuf3ParserSEMI)
	}

	return localctx
}

// IOptionContext is an interface to support dynamic dispatch.
type IOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionContext differentiates from other interfaces.
	IsOptionContext()
}

type OptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionContext() *OptionContext {
	var p = new(OptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_option
	return p
}

func (*OptionContext) IsOptionContext() {}

func NewOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionContext {
	var p = new(OptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_option

	return p
}

func (s *OptionContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionContext) OptionName() IOptionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *OptionContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *OptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterOption(s)
	}
}

func (s *OptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitOption(s)
	}
}

func (p *Protobuf3Parser) Option() (localctx IOptionContext) {
	localctx = NewOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, Protobuf3ParserRULE_option)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(Protobuf3ParserOPTION)
	}
	{
		p.SetState(112)
		p.OptionName()
	}
	{
		p.SetState(113)
		p.Match(Protobuf3ParserASSIGN)
	}
	{
		p.SetState(114)
		p.Constant()
	}
	{
		p.SetState(115)
		p.Match(Protobuf3ParserSEMI)
	}

	return localctx
}

// IOptionNameContext is an interface to support dynamic dispatch.
type IOptionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionNameContext differentiates from other interfaces.
	IsOptionNameContext()
}

type OptionNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionNameContext() *OptionNameContext {
	var p = new(OptionNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_optionName
	return p
}

func (*OptionNameContext) IsOptionNameContext() {}

func NewOptionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionNameContext {
	var p = new(OptionNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_optionName

	return p
}

func (s *OptionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionNameContext) AllIdent() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserIdent)
}

func (s *OptionNameContext) Ident(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIdent, i)
}

func (s *OptionNameContext) FullIdent() IFullIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *OptionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterOptionName(s)
	}
}

func (s *OptionNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitOptionName(s)
	}
}

func (p *Protobuf3Parser) OptionName() (localctx IOptionNameContext) {
	localctx = NewOptionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, Protobuf3ParserRULE_optionName)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(122)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserIdent:
		{
			p.SetState(117)
			p.Match(Protobuf3ParserIdent)
		}

	case Protobuf3ParserLPAREN:
		{
			p.SetState(118)
			p.Match(Protobuf3ParserLPAREN)
		}
		{
			p.SetState(119)
			p.FullIdent()
		}
		{
			p.SetState(120)
			p.Match(Protobuf3ParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserDOT {
		{
			p.SetState(124)
			p.Match(Protobuf3ParserDOT)
		}
		{
			p.SetState(125)
			p.Match(Protobuf3ParserIdent)
		}

		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITopLevelDefContext is an interface to support dynamic dispatch.
type ITopLevelDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTopLevelDefContext differentiates from other interfaces.
	IsTopLevelDefContext()
}

type TopLevelDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopLevelDefContext() *TopLevelDefContext {
	var p = new(TopLevelDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_topLevelDef
	return p
}

func (*TopLevelDefContext) IsTopLevelDefContext() {}

func NewTopLevelDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelDefContext {
	var p = new(TopLevelDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_topLevelDef

	return p
}

func (s *TopLevelDefContext) GetParser() antlr.Parser { return s.parser }

func (s *TopLevelDefContext) Message() IMessageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageContext)
}

func (s *TopLevelDefContext) EnumDefinition() IEnumDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumDefinitionContext)
}

func (s *TopLevelDefContext) Service() IServiceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServiceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IServiceContext)
}

func (s *TopLevelDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopLevelDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterTopLevelDef(s)
	}
}

func (s *TopLevelDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitTopLevelDef(s)
	}
}

func (p *Protobuf3Parser) TopLevelDef() (localctx ITopLevelDefContext) {
	localctx = NewTopLevelDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, Protobuf3ParserRULE_topLevelDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(134)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserMESSAGE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)
			p.Message()
		}

	case Protobuf3ParserENUM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(132)
			p.EnumDefinition()
		}

	case Protobuf3ParserSERVICE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(133)
			p.Service()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMessageContext is an interface to support dynamic dispatch.
type IMessageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageContext differentiates from other interfaces.
	IsMessageContext()
}

type MessageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageContext() *MessageContext {
	var p = new(MessageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_message
	return p
}

func (*MessageContext) IsMessageContext() {}

func NewMessageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageContext {
	var p = new(MessageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_message

	return p
}

func (s *MessageContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageContext) MessageName() IMessageNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageNameContext)
}

func (s *MessageContext) MessageBody() IMessageBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageBodyContext)
}

func (s *MessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMessage(s)
	}
}

func (s *MessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMessage(s)
	}
}

func (p *Protobuf3Parser) Message() (localctx IMessageContext) {
	localctx = NewMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, Protobuf3ParserRULE_message)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(Protobuf3ParserMESSAGE)
	}
	{
		p.SetState(137)
		p.MessageName()
	}
	{
		p.SetState(138)
		p.MessageBody()
	}

	return localctx
}

// IMessageBodyContext is an interface to support dynamic dispatch.
type IMessageBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageBodyContext differentiates from other interfaces.
	IsMessageBodyContext()
}

type MessageBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageBodyContext() *MessageBodyContext {
	var p = new(MessageBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageBody
	return p
}

func (*MessageBodyContext) IsMessageBodyContext() {}

func NewMessageBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageBodyContext {
	var p = new(MessageBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_messageBody

	return p
}

func (s *MessageBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageBodyContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *MessageBodyContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *MessageBodyContext) AllEnumDefinition() []IEnumDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumDefinitionContext)(nil)).Elem())
	var tst = make([]IEnumDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumDefinitionContext)
		}
	}

	return tst
}

func (s *MessageBodyContext) EnumDefinition(i int) IEnumDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumDefinitionContext)
}

func (s *MessageBodyContext) AllMessage() []IMessageContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMessageContext)(nil)).Elem())
	var tst = make([]IMessageContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMessageContext)
		}
	}

	return tst
}

func (s *MessageBodyContext) Message(i int) IMessageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMessageContext)
}

func (s *MessageBodyContext) AllOption() []IOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptionContext)(nil)).Elem())
	var tst = make([]IOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptionContext)
		}
	}

	return tst
}

func (s *MessageBodyContext) Option(i int) IOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *MessageBodyContext) AllOneof() []IOneofContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOneofContext)(nil)).Elem())
	var tst = make([]IOneofContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOneofContext)
		}
	}

	return tst
}

func (s *MessageBodyContext) Oneof(i int) IOneofContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOneofContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOneofContext)
}

func (s *MessageBodyContext) AllMapField() []IMapFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMapFieldContext)(nil)).Elem())
	var tst = make([]IMapFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMapFieldContext)
		}
	}

	return tst
}

func (s *MessageBodyContext) MapField(i int) IMapFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMapFieldContext)
}

func (s *MessageBodyContext) AllReserved() []IReservedContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IReservedContext)(nil)).Elem())
	var tst = make([]IReservedContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IReservedContext)
		}
	}

	return tst
}

func (s *MessageBodyContext) Reserved(i int) IReservedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReservedContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IReservedContext)
}

func (s *MessageBodyContext) AllEmptyStatement() []IEmptyStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEmptyStatementContext)(nil)).Elem())
	var tst = make([]IEmptyStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEmptyStatementContext)
		}
	}

	return tst
}

func (s *MessageBodyContext) EmptyStatement(i int) IEmptyStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEmptyStatementContext)
}

func (s *MessageBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMessageBody(s)
	}
}

func (s *MessageBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMessageBody(s)
	}
}

func (p *Protobuf3Parser) MessageBody() (localctx IMessageBodyContext) {
	localctx = NewMessageBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, Protobuf3ParserRULE_messageBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Match(Protobuf3ParserLBRACE)
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Protobuf3ParserBOOL)|(1<<Protobuf3ParserBYTES)|(1<<Protobuf3ParserDOUBLE)|(1<<Protobuf3ParserENUM)|(1<<Protobuf3ParserFIXED32)|(1<<Protobuf3ParserFIXED64)|(1<<Protobuf3ParserFLOAT)|(1<<Protobuf3ParserINT32)|(1<<Protobuf3ParserINT64)|(1<<Protobuf3ParserMAP)|(1<<Protobuf3ParserMESSAGE)|(1<<Protobuf3ParserONEOF)|(1<<Protobuf3ParserOPTION)|(1<<Protobuf3ParserREPEATED)|(1<<Protobuf3ParserRESERVED)|(1<<Protobuf3ParserSFIXED32)|(1<<Protobuf3ParserSFIXED64)|(1<<Protobuf3ParserSINT32)|(1<<Protobuf3ParserSINT64)|(1<<Protobuf3ParserSTRING))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(Protobuf3ParserUINT32-32))|(1<<(Protobuf3ParserUINT64-32))|(1<<(Protobuf3ParserIdent-32))|(1<<(Protobuf3ParserSEMI-32))|(1<<(Protobuf3ParserDOT-32)))) != 0) {
		p.SetState(149)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case Protobuf3ParserBOOL, Protobuf3ParserBYTES, Protobuf3ParserDOUBLE, Protobuf3ParserFIXED32, Protobuf3ParserFIXED64, Protobuf3ParserFLOAT, Protobuf3ParserINT32, Protobuf3ParserINT64, Protobuf3ParserREPEATED, Protobuf3ParserSFIXED32, Protobuf3ParserSFIXED64, Protobuf3ParserSINT32, Protobuf3ParserSINT64, Protobuf3ParserSTRING, Protobuf3ParserUINT32, Protobuf3ParserUINT64, Protobuf3ParserIdent, Protobuf3ParserDOT:
			{
				p.SetState(141)
				p.Field()
			}

		case Protobuf3ParserENUM:
			{
				p.SetState(142)
				p.EnumDefinition()
			}

		case Protobuf3ParserMESSAGE:
			{
				p.SetState(143)
				p.Message()
			}

		case Protobuf3ParserOPTION:
			{
				p.SetState(144)
				p.Option()
			}

		case Protobuf3ParserONEOF:
			{
				p.SetState(145)
				p.Oneof()
			}

		case Protobuf3ParserMAP:
			{
				p.SetState(146)
				p.MapField()
			}

		case Protobuf3ParserRESERVED:
			{
				p.SetState(147)
				p.Reserved()
			}

		case Protobuf3ParserSEMI:
			{
				p.SetState(148)
				p.EmptyStatement()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(154)
		p.Match(Protobuf3ParserRBRACE)
	}

	return localctx
}

// IEnumDefinitionContext is an interface to support dynamic dispatch.
type IEnumDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumDefinitionContext differentiates from other interfaces.
	IsEnumDefinitionContext()
}

type EnumDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumDefinitionContext() *EnumDefinitionContext {
	var p = new(EnumDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumDefinition
	return p
}

func (*EnumDefinitionContext) IsEnumDefinitionContext() {}

func NewEnumDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDefinitionContext {
	var p = new(EnumDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumDefinition

	return p
}

func (s *EnumDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDefinitionContext) EnumName() IEnumNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumNameContext)
}

func (s *EnumDefinitionContext) EnumBody() IEnumBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *EnumDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEnumDefinition(s)
	}
}

func (s *EnumDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEnumDefinition(s)
	}
}

func (p *Protobuf3Parser) EnumDefinition() (localctx IEnumDefinitionContext) {
	localctx = NewEnumDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, Protobuf3ParserRULE_enumDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Match(Protobuf3ParserENUM)
	}
	{
		p.SetState(157)
		p.EnumName()
	}
	{
		p.SetState(158)
		p.EnumBody()
	}

	return localctx
}

// IEnumBodyContext is an interface to support dynamic dispatch.
type IEnumBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumBodyContext differentiates from other interfaces.
	IsEnumBodyContext()
}

type EnumBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumBodyContext() *EnumBodyContext {
	var p = new(EnumBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumBody
	return p
}

func (*EnumBodyContext) IsEnumBodyContext() {}

func NewEnumBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumBodyContext {
	var p = new(EnumBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumBody

	return p
}

func (s *EnumBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumBodyContext) AllOption() []IOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptionContext)(nil)).Elem())
	var tst = make([]IOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptionContext)
		}
	}

	return tst
}

func (s *EnumBodyContext) Option(i int) IOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *EnumBodyContext) AllEnumField() []IEnumFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumFieldContext)(nil)).Elem())
	var tst = make([]IEnumFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumFieldContext)
		}
	}

	return tst
}

func (s *EnumBodyContext) EnumField(i int) IEnumFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumFieldContext)
}

func (s *EnumBodyContext) AllEmptyStatement() []IEmptyStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEmptyStatementContext)(nil)).Elem())
	var tst = make([]IEmptyStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEmptyStatementContext)
		}
	}

	return tst
}

func (s *EnumBodyContext) EmptyStatement(i int) IEmptyStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEmptyStatementContext)
}

func (s *EnumBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEnumBody(s)
	}
}

func (s *EnumBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEnumBody(s)
	}
}

func (p *Protobuf3Parser) EnumBody() (localctx IEnumBodyContext) {
	localctx = NewEnumBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, Protobuf3ParserRULE_enumBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(Protobuf3ParserLBRACE)
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserOPTION || _la == Protobuf3ParserIdent || _la == Protobuf3ParserSEMI {
		p.SetState(164)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case Protobuf3ParserOPTION:
			{
				p.SetState(161)
				p.Option()
			}

		case Protobuf3ParserIdent:
			{
				p.SetState(162)
				p.EnumField()
			}

		case Protobuf3ParserSEMI:
			{
				p.SetState(163)
				p.EmptyStatement()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(169)
		p.Match(Protobuf3ParserRBRACE)
	}

	return localctx
}

// IEnumFieldContext is an interface to support dynamic dispatch.
type IEnumFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumFieldContext differentiates from other interfaces.
	IsEnumFieldContext()
}

type EnumFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumFieldContext() *EnumFieldContext {
	var p = new(EnumFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumField
	return p
}

func (*EnumFieldContext) IsEnumFieldContext() {}

func NewEnumFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumFieldContext {
	var p = new(EnumFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumField

	return p
}

func (s *EnumFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumFieldContext) Ident() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIdent, 0)
}

func (s *EnumFieldContext) IntLit() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIntLit, 0)
}

func (s *EnumFieldContext) AllEnumValueOption() []IEnumValueOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumValueOptionContext)(nil)).Elem())
	var tst = make([]IEnumValueOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumValueOptionContext)
		}
	}

	return tst
}

func (s *EnumFieldContext) EnumValueOption(i int) IEnumValueOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValueOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumValueOptionContext)
}

func (s *EnumFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEnumField(s)
	}
}

func (s *EnumFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEnumField(s)
	}
}

func (p *Protobuf3Parser) EnumField() (localctx IEnumFieldContext) {
	localctx = NewEnumFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, Protobuf3ParserRULE_enumField)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.Match(Protobuf3ParserIdent)
	}
	{
		p.SetState(172)
		p.Match(Protobuf3ParserASSIGN)
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserMINUS {
		{
			p.SetState(173)
			p.Match(Protobuf3ParserMINUS)
		}

	}
	{
		p.SetState(176)
		p.Match(Protobuf3ParserIntLit)
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserLBRACK {
		{
			p.SetState(177)
			p.Match(Protobuf3ParserLBRACK)
		}
		{
			p.SetState(178)
			p.EnumValueOption()
		}
		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == Protobuf3ParserCOMMA {
			{
				p.SetState(179)
				p.Match(Protobuf3ParserCOMMA)
			}
			{
				p.SetState(180)
				p.EnumValueOption()
			}

			p.SetState(185)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(186)
			p.Match(Protobuf3ParserRBRACK)
		}

	}
	{
		p.SetState(190)
		p.Match(Protobuf3ParserSEMI)
	}

	return localctx
}

// IEnumValueOptionContext is an interface to support dynamic dispatch.
type IEnumValueOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumValueOptionContext differentiates from other interfaces.
	IsEnumValueOptionContext()
}

type EnumValueOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueOptionContext() *EnumValueOptionContext {
	var p = new(EnumValueOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumValueOption
	return p
}

func (*EnumValueOptionContext) IsEnumValueOptionContext() {}

func NewEnumValueOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueOptionContext {
	var p = new(EnumValueOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumValueOption

	return p
}

func (s *EnumValueOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueOptionContext) OptionName() IOptionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *EnumValueOptionContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *EnumValueOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEnumValueOption(s)
	}
}

func (s *EnumValueOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEnumValueOption(s)
	}
}

func (p *Protobuf3Parser) EnumValueOption() (localctx IEnumValueOptionContext) {
	localctx = NewEnumValueOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, Protobuf3ParserRULE_enumValueOption)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.OptionName()
	}
	{
		p.SetState(193)
		p.Match(Protobuf3ParserASSIGN)
	}
	{
		p.SetState(194)
		p.Constant()
	}

	return localctx
}

// IServiceContext is an interface to support dynamic dispatch.
type IServiceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServiceContext differentiates from other interfaces.
	IsServiceContext()
}

type ServiceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceContext() *ServiceContext {
	var p = new(ServiceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_service
	return p
}

func (*ServiceContext) IsServiceContext() {}

func NewServiceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceContext {
	var p = new(ServiceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_service

	return p
}

func (s *ServiceContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceContext) ServiceName() IServiceNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServiceNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IServiceNameContext)
}

func (s *ServiceContext) AllOption() []IOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptionContext)(nil)).Elem())
	var tst = make([]IOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptionContext)
		}
	}

	return tst
}

func (s *ServiceContext) Option(i int) IOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *ServiceContext) AllRpc() []IRpcContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRpcContext)(nil)).Elem())
	var tst = make([]IRpcContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRpcContext)
		}
	}

	return tst
}

func (s *ServiceContext) Rpc(i int) IRpcContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRpcContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRpcContext)
}

func (s *ServiceContext) AllEmptyStatement() []IEmptyStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEmptyStatementContext)(nil)).Elem())
	var tst = make([]IEmptyStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEmptyStatementContext)
		}
	}

	return tst
}

func (s *ServiceContext) EmptyStatement(i int) IEmptyStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEmptyStatementContext)
}

func (s *ServiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterService(s)
	}
}

func (s *ServiceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitService(s)
	}
}

func (p *Protobuf3Parser) Service() (localctx IServiceContext) {
	localctx = NewServiceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, Protobuf3ParserRULE_service)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(Protobuf3ParserSERVICE)
	}
	{
		p.SetState(197)
		p.ServiceName()
	}
	{
		p.SetState(198)
		p.Match(Protobuf3ParserLBRACE)
	}
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserOPTION || _la == Protobuf3ParserRPC || _la == Protobuf3ParserSEMI {
		p.SetState(202)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case Protobuf3ParserOPTION:
			{
				p.SetState(199)
				p.Option()
			}

		case Protobuf3ParserRPC:
			{
				p.SetState(200)
				p.Rpc()
			}

		case Protobuf3ParserSEMI:
			{
				p.SetState(201)
				p.EmptyStatement()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(207)
		p.Match(Protobuf3ParserRBRACE)
	}

	return localctx
}

// IRpcContext is an interface to support dynamic dispatch.
type IRpcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRpcContext differentiates from other interfaces.
	IsRpcContext()
}

type RpcContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcContext() *RpcContext {
	var p = new(RpcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_rpc
	return p
}

func (*RpcContext) IsRpcContext() {}

func NewRpcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcContext {
	var p = new(RpcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_rpc

	return p
}

func (s *RpcContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcContext) RpcName() IRpcNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRpcNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRpcNameContext)
}

func (s *RpcContext) AllMessageType() []IMessageTypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMessageTypeContext)(nil)).Elem())
	var tst = make([]IMessageTypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMessageTypeContext)
		}
	}

	return tst
}

func (s *RpcContext) MessageType(i int) IMessageTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageTypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMessageTypeContext)
}

func (s *RpcContext) AllOption() []IOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptionContext)(nil)).Elem())
	var tst = make([]IOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptionContext)
		}
	}

	return tst
}

func (s *RpcContext) Option(i int) IOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *RpcContext) AllEmptyStatement() []IEmptyStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEmptyStatementContext)(nil)).Elem())
	var tst = make([]IEmptyStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEmptyStatementContext)
		}
	}

	return tst
}

func (s *RpcContext) EmptyStatement(i int) IEmptyStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEmptyStatementContext)
}

func (s *RpcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterRpc(s)
	}
}

func (s *RpcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitRpc(s)
	}
}

func (p *Protobuf3Parser) Rpc() (localctx IRpcContext) {
	localctx = NewRpcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, Protobuf3ParserRULE_rpc)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		p.Match(Protobuf3ParserRPC)
	}
	{
		p.SetState(210)
		p.RpcName()
	}
	{
		p.SetState(211)
		p.Match(Protobuf3ParserLPAREN)
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserSTREAM {
		{
			p.SetState(212)
			p.Match(Protobuf3ParserSTREAM)
		}

	}
	{
		p.SetState(215)
		p.MessageType()
	}
	{
		p.SetState(216)
		p.Match(Protobuf3ParserRPAREN)
	}
	{
		p.SetState(217)
		p.Match(Protobuf3ParserRETURNS)
	}
	{
		p.SetState(218)
		p.Match(Protobuf3ParserLPAREN)
	}
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserSTREAM {
		{
			p.SetState(219)
			p.Match(Protobuf3ParserSTREAM)
		}

	}
	{
		p.SetState(222)
		p.MessageType()
	}
	{
		p.SetState(223)
		p.Match(Protobuf3ParserRPAREN)
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserLBRACE:
		{
			p.SetState(224)
			p.Match(Protobuf3ParserLBRACE)
		}
		p.SetState(229)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == Protobuf3ParserOPTION || _la == Protobuf3ParserSEMI {
			p.SetState(227)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case Protobuf3ParserOPTION:
				{
					p.SetState(225)
					p.Option()
				}

			case Protobuf3ParserSEMI:
				{
					p.SetState(226)
					p.EmptyStatement()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(231)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(232)
			p.Match(Protobuf3ParserRBRACE)
		}

	case Protobuf3ParserSEMI:
		{
			p.SetState(233)
			p.Match(Protobuf3ParserSEMI)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IReservedContext is an interface to support dynamic dispatch.
type IReservedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReservedContext differentiates from other interfaces.
	IsReservedContext()
}

type ReservedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReservedContext() *ReservedContext {
	var p = new(ReservedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_reserved
	return p
}

func (*ReservedContext) IsReservedContext() {}

func NewReservedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReservedContext {
	var p = new(ReservedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_reserved

	return p
}

func (s *ReservedContext) GetParser() antlr.Parser { return s.parser }

func (s *ReservedContext) Pbranges() IPbrangesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPbrangesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPbrangesContext)
}

func (s *ReservedContext) FieldNames() IFieldNamesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNamesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNamesContext)
}

func (s *ReservedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReservedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReservedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterReserved(s)
	}
}

func (s *ReservedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitReserved(s)
	}
}

func (p *Protobuf3Parser) Reserved() (localctx IReservedContext) {
	localctx = NewReservedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, Protobuf3ParserRULE_reserved)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.Match(Protobuf3ParserRESERVED)
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserIntLit:
		{
			p.SetState(237)
			p.Pbranges()
		}

	case Protobuf3ParserStrLit:
		{
			p.SetState(238)
			p.FieldNames()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(241)
		p.Match(Protobuf3ParserSEMI)
	}

	return localctx
}

// IPbrangesContext is an interface to support dynamic dispatch.
type IPbrangesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPbrangesContext differentiates from other interfaces.
	IsPbrangesContext()
}

type PbrangesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPbrangesContext() *PbrangesContext {
	var p = new(PbrangesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_pbranges
	return p
}

func (*PbrangesContext) IsPbrangesContext() {}

func NewPbrangesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PbrangesContext {
	var p = new(PbrangesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_pbranges

	return p
}

func (s *PbrangesContext) GetParser() antlr.Parser { return s.parser }

func (s *PbrangesContext) AllPbrange() []IPbrangeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPbrangeContext)(nil)).Elem())
	var tst = make([]IPbrangeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPbrangeContext)
		}
	}

	return tst
}

func (s *PbrangesContext) Pbrange(i int) IPbrangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPbrangeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPbrangeContext)
}

func (s *PbrangesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PbrangesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PbrangesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterPbranges(s)
	}
}

func (s *PbrangesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitPbranges(s)
	}
}

func (p *Protobuf3Parser) Pbranges() (localctx IPbrangesContext) {
	localctx = NewPbrangesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, Protobuf3ParserRULE_pbranges)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Pbrange()
	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserCOMMA {
		{
			p.SetState(244)
			p.Match(Protobuf3ParserCOMMA)
		}
		{
			p.SetState(245)
			p.Pbrange()
		}

		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPbrangeContext is an interface to support dynamic dispatch.
type IPbrangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPbrangeContext differentiates from other interfaces.
	IsPbrangeContext()
}

type PbrangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPbrangeContext() *PbrangeContext {
	var p = new(PbrangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_pbrange
	return p
}

func (*PbrangeContext) IsPbrangeContext() {}

func NewPbrangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PbrangeContext {
	var p = new(PbrangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_pbrange

	return p
}

func (s *PbrangeContext) GetParser() antlr.Parser { return s.parser }

func (s *PbrangeContext) AllIntLit() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserIntLit)
}

func (s *PbrangeContext) IntLit(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIntLit, i)
}

func (s *PbrangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PbrangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PbrangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterPbrange(s)
	}
}

func (s *PbrangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitPbrange(s)
	}
}

func (p *Protobuf3Parser) Pbrange() (localctx IPbrangeContext) {
	localctx = NewPbrangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, Protobuf3ParserRULE_pbrange)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(251)
			p.Match(Protobuf3ParserIntLit)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(252)
			p.Match(Protobuf3ParserIntLit)
		}
		{
			p.SetState(253)
			p.Match(Protobuf3ParserTO)
		}
		{
			p.SetState(254)
			p.Match(Protobuf3ParserIntLit)
		}

	}

	return localctx
}

// IFieldNamesContext is an interface to support dynamic dispatch.
type IFieldNamesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNamesContext differentiates from other interfaces.
	IsFieldNamesContext()
}

type FieldNamesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNamesContext() *FieldNamesContext {
	var p = new(FieldNamesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldNames
	return p
}

func (*FieldNamesContext) IsFieldNamesContext() {}

func NewFieldNamesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNamesContext {
	var p = new(FieldNamesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fieldNames

	return p
}

func (s *FieldNamesContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNamesContext) AllStrLit() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserStrLit)
}

func (s *FieldNamesContext) StrLit(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserStrLit, i)
}

func (s *FieldNamesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNamesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNamesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterFieldNames(s)
	}
}

func (s *FieldNamesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitFieldNames(s)
	}
}

func (p *Protobuf3Parser) FieldNames() (localctx IFieldNamesContext) {
	localctx = NewFieldNamesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, Protobuf3ParserRULE_fieldNames)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		p.Match(Protobuf3ParserStrLit)
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserCOMMA {
		{
			p.SetState(258)
			p.Match(Protobuf3ParserCOMMA)
		}
		{
			p.SetState(259)
			p.Match(Protobuf3ParserStrLit)
		}

		p.SetState(264)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPbtypeContext is an interface to support dynamic dispatch.
type IPbtypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPbtypeContext differentiates from other interfaces.
	IsPbtypeContext()
}

type PbtypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPbtypeContext() *PbtypeContext {
	var p = new(PbtypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_pbtype
	return p
}

func (*PbtypeContext) IsPbtypeContext() {}

func NewPbtypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PbtypeContext {
	var p = new(PbtypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_pbtype

	return p
}

func (s *PbtypeContext) GetParser() antlr.Parser { return s.parser }

func (s *PbtypeContext) MessageOrEnumType() IMessageOrEnumTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageOrEnumTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageOrEnumTypeContext)
}

func (s *PbtypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PbtypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PbtypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterPbtype(s)
	}
}

func (s *PbtypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitPbtype(s)
	}
}

func (p *Protobuf3Parser) Pbtype() (localctx IPbtypeContext) {
	localctx = NewPbtypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, Protobuf3ParserRULE_pbtype)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(267)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserBOOL, Protobuf3ParserBYTES, Protobuf3ParserDOUBLE, Protobuf3ParserFIXED32, Protobuf3ParserFIXED64, Protobuf3ParserFLOAT, Protobuf3ParserINT32, Protobuf3ParserINT64, Protobuf3ParserSFIXED32, Protobuf3ParserSFIXED64, Protobuf3ParserSINT32, Protobuf3ParserSINT64, Protobuf3ParserSTRING, Protobuf3ParserUINT32, Protobuf3ParserUINT64:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(265)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Protobuf3ParserBOOL)|(1<<Protobuf3ParserBYTES)|(1<<Protobuf3ParserDOUBLE)|(1<<Protobuf3ParserFIXED32)|(1<<Protobuf3ParserFIXED64)|(1<<Protobuf3ParserFLOAT)|(1<<Protobuf3ParserINT32)|(1<<Protobuf3ParserINT64)|(1<<Protobuf3ParserSFIXED32)|(1<<Protobuf3ParserSFIXED64)|(1<<Protobuf3ParserSINT32)|(1<<Protobuf3ParserSINT64)|(1<<Protobuf3ParserSTRING))) != 0) || _la == Protobuf3ParserUINT32 || _la == Protobuf3ParserUINT64) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	case Protobuf3ParserIdent, Protobuf3ParserDOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(266)
			p.MessageOrEnumType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFieldNumberContext is an interface to support dynamic dispatch.
type IFieldNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNumberContext differentiates from other interfaces.
	IsFieldNumberContext()
}

type FieldNumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNumberContext() *FieldNumberContext {
	var p = new(FieldNumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldNumber
	return p
}

func (*FieldNumberContext) IsFieldNumberContext() {}

func NewFieldNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNumberContext {
	var p = new(FieldNumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fieldNumber

	return p
}

func (s *FieldNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNumberContext) IntLit() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIntLit, 0)
}

func (s *FieldNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterFieldNumber(s)
	}
}

func (s *FieldNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitFieldNumber(s)
	}
}

func (p *Protobuf3Parser) FieldNumber() (localctx IFieldNumberContext) {
	localctx = NewFieldNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, Protobuf3ParserRULE_fieldNumber)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(269)
		p.Match(Protobuf3ParserIntLit)
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Pbtype() IPbtypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPbtypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPbtypeContext)
}

func (s *FieldContext) FieldName() IFieldNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *FieldContext) FieldNumber() IFieldNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNumberContext)
}

func (s *FieldContext) FieldOptions() IFieldOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitField(s)
	}
}

func (p *Protobuf3Parser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, Protobuf3ParserRULE_field)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserREPEATED {
		{
			p.SetState(271)
			p.Match(Protobuf3ParserREPEATED)
		}

	}
	{
		p.SetState(274)
		p.Pbtype()
	}
	{
		p.SetState(275)
		p.FieldName()
	}
	{
		p.SetState(276)
		p.Match(Protobuf3ParserASSIGN)
	}
	{
		p.SetState(277)
		p.FieldNumber()
	}
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserLBRACK {
		{
			p.SetState(278)
			p.Match(Protobuf3ParserLBRACK)
		}
		{
			p.SetState(279)
			p.FieldOptions()
		}
		{
			p.SetState(280)
			p.Match(Protobuf3ParserRBRACK)
		}

	}
	{
		p.SetState(284)
		p.Match(Protobuf3ParserSEMI)
	}

	return localctx
}

// IFieldOptionsContext is an interface to support dynamic dispatch.
type IFieldOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldOptionsContext differentiates from other interfaces.
	IsFieldOptionsContext()
}

type FieldOptionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldOptionsContext() *FieldOptionsContext {
	var p = new(FieldOptionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldOptions
	return p
}

func (*FieldOptionsContext) IsFieldOptionsContext() {}

func NewFieldOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOptionsContext {
	var p = new(FieldOptionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fieldOptions

	return p
}

func (s *FieldOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldOptionsContext) AllFieldOption() []IFieldOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldOptionContext)(nil)).Elem())
	var tst = make([]IFieldOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldOptionContext)
		}
	}

	return tst
}

func (s *FieldOptionsContext) FieldOption(i int) IFieldOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldOptionContext)
}

func (s *FieldOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterFieldOptions(s)
	}
}

func (s *FieldOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitFieldOptions(s)
	}
}

func (p *Protobuf3Parser) FieldOptions() (localctx IFieldOptionsContext) {
	localctx = NewFieldOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, Protobuf3ParserRULE_fieldOptions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		p.FieldOption()
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserCOMMA {
		{
			p.SetState(287)
			p.Match(Protobuf3ParserCOMMA)
		}
		{
			p.SetState(288)
			p.FieldOption()
		}

		p.SetState(293)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFieldOptionContext is an interface to support dynamic dispatch.
type IFieldOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldOptionContext differentiates from other interfaces.
	IsFieldOptionContext()
}

type FieldOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldOptionContext() *FieldOptionContext {
	var p = new(FieldOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldOption
	return p
}

func (*FieldOptionContext) IsFieldOptionContext() {}

func NewFieldOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOptionContext {
	var p = new(FieldOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fieldOption

	return p
}

func (s *FieldOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldOptionContext) OptionName() IOptionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *FieldOptionContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *FieldOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterFieldOption(s)
	}
}

func (s *FieldOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitFieldOption(s)
	}
}

func (p *Protobuf3Parser) FieldOption() (localctx IFieldOptionContext) {
	localctx = NewFieldOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, Protobuf3ParserRULE_fieldOption)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(294)
		p.OptionName()
	}
	{
		p.SetState(295)
		p.Match(Protobuf3ParserASSIGN)
	}
	{
		p.SetState(296)
		p.Constant()
	}

	return localctx
}

// IOneofContext is an interface to support dynamic dispatch.
type IOneofContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOneofContext differentiates from other interfaces.
	IsOneofContext()
}

type OneofContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneofContext() *OneofContext {
	var p = new(OneofContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_oneof
	return p
}

func (*OneofContext) IsOneofContext() {}

func NewOneofContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofContext {
	var p = new(OneofContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_oneof

	return p
}

func (s *OneofContext) GetParser() antlr.Parser { return s.parser }

func (s *OneofContext) OneofName() IOneofNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOneofNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOneofNameContext)
}

func (s *OneofContext) AllOneofField() []IOneofFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOneofFieldContext)(nil)).Elem())
	var tst = make([]IOneofFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOneofFieldContext)
		}
	}

	return tst
}

func (s *OneofContext) OneofField(i int) IOneofFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOneofFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOneofFieldContext)
}

func (s *OneofContext) AllEmptyStatement() []IEmptyStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEmptyStatementContext)(nil)).Elem())
	var tst = make([]IEmptyStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEmptyStatementContext)
		}
	}

	return tst
}

func (s *OneofContext) EmptyStatement(i int) IEmptyStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEmptyStatementContext)
}

func (s *OneofContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneofContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterOneof(s)
	}
}

func (s *OneofContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitOneof(s)
	}
}

func (p *Protobuf3Parser) Oneof() (localctx IOneofContext) {
	localctx = NewOneofContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, Protobuf3ParserRULE_oneof)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		p.Match(Protobuf3ParserONEOF)
	}
	{
		p.SetState(299)
		p.OneofName()
	}
	{
		p.SetState(300)
		p.Match(Protobuf3ParserLBRACE)
	}
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Protobuf3ParserBOOL)|(1<<Protobuf3ParserBYTES)|(1<<Protobuf3ParserDOUBLE)|(1<<Protobuf3ParserFIXED32)|(1<<Protobuf3ParserFIXED64)|(1<<Protobuf3ParserFLOAT)|(1<<Protobuf3ParserINT32)|(1<<Protobuf3ParserINT64)|(1<<Protobuf3ParserSFIXED32)|(1<<Protobuf3ParserSFIXED64)|(1<<Protobuf3ParserSINT32)|(1<<Protobuf3ParserSINT64)|(1<<Protobuf3ParserSTRING))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(Protobuf3ParserUINT32-32))|(1<<(Protobuf3ParserUINT64-32))|(1<<(Protobuf3ParserIdent-32))|(1<<(Protobuf3ParserSEMI-32))|(1<<(Protobuf3ParserDOT-32)))) != 0) {
		p.SetState(303)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case Protobuf3ParserBOOL, Protobuf3ParserBYTES, Protobuf3ParserDOUBLE, Protobuf3ParserFIXED32, Protobuf3ParserFIXED64, Protobuf3ParserFLOAT, Protobuf3ParserINT32, Protobuf3ParserINT64, Protobuf3ParserSFIXED32, Protobuf3ParserSFIXED64, Protobuf3ParserSINT32, Protobuf3ParserSINT64, Protobuf3ParserSTRING, Protobuf3ParserUINT32, Protobuf3ParserUINT64, Protobuf3ParserIdent, Protobuf3ParserDOT:
			{
				p.SetState(301)
				p.OneofField()
			}

		case Protobuf3ParserSEMI:
			{
				p.SetState(302)
				p.EmptyStatement()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(308)
		p.Match(Protobuf3ParserRBRACE)
	}

	return localctx
}

// IOneofFieldContext is an interface to support dynamic dispatch.
type IOneofFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOneofFieldContext differentiates from other interfaces.
	IsOneofFieldContext()
}

type OneofFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneofFieldContext() *OneofFieldContext {
	var p = new(OneofFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_oneofField
	return p
}

func (*OneofFieldContext) IsOneofFieldContext() {}

func NewOneofFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofFieldContext {
	var p = new(OneofFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_oneofField

	return p
}

func (s *OneofFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *OneofFieldContext) Pbtype() IPbtypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPbtypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPbtypeContext)
}

func (s *OneofFieldContext) FieldName() IFieldNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *OneofFieldContext) FieldNumber() IFieldNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNumberContext)
}

func (s *OneofFieldContext) FieldOptions() IFieldOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *OneofFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneofFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterOneofField(s)
	}
}

func (s *OneofFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitOneofField(s)
	}
}

func (p *Protobuf3Parser) OneofField() (localctx IOneofFieldContext) {
	localctx = NewOneofFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, Protobuf3ParserRULE_oneofField)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Pbtype()
	}
	{
		p.SetState(311)
		p.FieldName()
	}
	{
		p.SetState(312)
		p.Match(Protobuf3ParserASSIGN)
	}
	{
		p.SetState(313)
		p.FieldNumber()
	}
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserLBRACK {
		{
			p.SetState(314)
			p.Match(Protobuf3ParserLBRACK)
		}
		{
			p.SetState(315)
			p.FieldOptions()
		}
		{
			p.SetState(316)
			p.Match(Protobuf3ParserRBRACK)
		}

	}
	{
		p.SetState(320)
		p.Match(Protobuf3ParserSEMI)
	}

	return localctx
}

// IMapFieldContext is an interface to support dynamic dispatch.
type IMapFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapFieldContext differentiates from other interfaces.
	IsMapFieldContext()
}

type MapFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapFieldContext() *MapFieldContext {
	var p = new(MapFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_mapField
	return p
}

func (*MapFieldContext) IsMapFieldContext() {}

func NewMapFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapFieldContext {
	var p = new(MapFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_mapField

	return p
}

func (s *MapFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *MapFieldContext) KeyType() IKeyTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyTypeContext)
}

func (s *MapFieldContext) Pbtype() IPbtypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPbtypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPbtypeContext)
}

func (s *MapFieldContext) MapName() IMapNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapNameContext)
}

func (s *MapFieldContext) FieldNumber() IFieldNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNumberContext)
}

func (s *MapFieldContext) FieldOptions() IFieldOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *MapFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMapField(s)
	}
}

func (s *MapFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMapField(s)
	}
}

func (p *Protobuf3Parser) MapField() (localctx IMapFieldContext) {
	localctx = NewMapFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, Protobuf3ParserRULE_mapField)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(322)
		p.Match(Protobuf3ParserMAP)
	}
	{
		p.SetState(323)
		p.Match(Protobuf3ParserLCHEVR)
	}
	{
		p.SetState(324)
		p.KeyType()
	}
	{
		p.SetState(325)
		p.Match(Protobuf3ParserCOMMA)
	}
	{
		p.SetState(326)
		p.Pbtype()
	}
	{
		p.SetState(327)
		p.Match(Protobuf3ParserRCHEVR)
	}
	{
		p.SetState(328)
		p.MapName()
	}
	{
		p.SetState(329)
		p.Match(Protobuf3ParserASSIGN)
	}
	{
		p.SetState(330)
		p.FieldNumber()
	}
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserLBRACK {
		{
			p.SetState(331)
			p.Match(Protobuf3ParserLBRACK)
		}
		{
			p.SetState(332)
			p.FieldOptions()
		}
		{
			p.SetState(333)
			p.Match(Protobuf3ParserRBRACK)
		}

	}
	{
		p.SetState(337)
		p.Match(Protobuf3ParserSEMI)
	}

	return localctx
}

// IKeyTypeContext is an interface to support dynamic dispatch.
type IKeyTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyTypeContext differentiates from other interfaces.
	IsKeyTypeContext()
}

type KeyTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyTypeContext() *KeyTypeContext {
	var p = new(KeyTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_keyType
	return p
}

func (*KeyTypeContext) IsKeyTypeContext() {}

func NewKeyTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyTypeContext {
	var p = new(KeyTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_keyType

	return p
}

func (s *KeyTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *KeyTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterKeyType(s)
	}
}

func (s *KeyTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitKeyType(s)
	}
}

func (p *Protobuf3Parser) KeyType() (localctx IKeyTypeContext) {
	localctx = NewKeyTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, Protobuf3ParserRULE_keyType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(339)
	_la = p.GetTokenStream().LA(1)

	if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Protobuf3ParserBOOL)|(1<<Protobuf3ParserFIXED32)|(1<<Protobuf3ParserFIXED64)|(1<<Protobuf3ParserINT32)|(1<<Protobuf3ParserINT64)|(1<<Protobuf3ParserSFIXED32)|(1<<Protobuf3ParserSFIXED64)|(1<<Protobuf3ParserSINT32)|(1<<Protobuf3ParserSINT64)|(1<<Protobuf3ParserSTRING))) != 0) || _la == Protobuf3ParserUINT32 || _la == Protobuf3ParserUINT64) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IFullIdentContext is an interface to support dynamic dispatch.
type IFullIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFullIdentContext differentiates from other interfaces.
	IsFullIdentContext()
}

type FullIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFullIdentContext() *FullIdentContext {
	var p = new(FullIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fullIdent
	return p
}

func (*FullIdentContext) IsFullIdentContext() {}

func NewFullIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FullIdentContext {
	var p = new(FullIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fullIdent

	return p
}

func (s *FullIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *FullIdentContext) AllIdent() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserIdent)
}

func (s *FullIdentContext) Ident(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIdent, i)
}

func (s *FullIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FullIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterFullIdent(s)
	}
}

func (s *FullIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitFullIdent(s)
	}
}

func (p *Protobuf3Parser) FullIdent() (localctx IFullIdentContext) {
	localctx = NewFullIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, Protobuf3ParserRULE_fullIdent)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(341)
		p.Match(Protobuf3ParserIdent)
	}
	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserDOT {
		{
			p.SetState(342)
			p.Match(Protobuf3ParserDOT)
		}
		{
			p.SetState(343)
			p.Match(Protobuf3ParserIdent)
		}

		p.SetState(348)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMessageNameContext is an interface to support dynamic dispatch.
type IMessageNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageNameContext differentiates from other interfaces.
	IsMessageNameContext()
}

type MessageNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageNameContext() *MessageNameContext {
	var p = new(MessageNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageName
	return p
}

func (*MessageNameContext) IsMessageNameContext() {}

func NewMessageNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageNameContext {
	var p = new(MessageNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_messageName

	return p
}

func (s *MessageNameContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageNameContext) Ident() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIdent, 0)
}

func (s *MessageNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMessageName(s)
	}
}

func (s *MessageNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMessageName(s)
	}
}

func (p *Protobuf3Parser) MessageName() (localctx IMessageNameContext) {
	localctx = NewMessageNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, Protobuf3ParserRULE_messageName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(349)
		p.Match(Protobuf3ParserIdent)
	}

	return localctx
}

// IEnumNameContext is an interface to support dynamic dispatch.
type IEnumNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumNameContext differentiates from other interfaces.
	IsEnumNameContext()
}

type EnumNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumNameContext() *EnumNameContext {
	var p = new(EnumNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumName
	return p
}

func (*EnumNameContext) IsEnumNameContext() {}

func NewEnumNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumNameContext {
	var p = new(EnumNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumName

	return p
}

func (s *EnumNameContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumNameContext) Ident() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIdent, 0)
}

func (s *EnumNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEnumName(s)
	}
}

func (s *EnumNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEnumName(s)
	}
}

func (p *Protobuf3Parser) EnumName() (localctx IEnumNameContext) {
	localctx = NewEnumNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, Protobuf3ParserRULE_enumName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(351)
		p.Match(Protobuf3ParserIdent)
	}

	return localctx
}

// IMessageOrEnumNameContext is an interface to support dynamic dispatch.
type IMessageOrEnumNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageOrEnumNameContext differentiates from other interfaces.
	IsMessageOrEnumNameContext()
}

type MessageOrEnumNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageOrEnumNameContext() *MessageOrEnumNameContext {
	var p = new(MessageOrEnumNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageOrEnumName
	return p
}

func (*MessageOrEnumNameContext) IsMessageOrEnumNameContext() {}

func NewMessageOrEnumNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageOrEnumNameContext {
	var p = new(MessageOrEnumNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_messageOrEnumName

	return p
}

func (s *MessageOrEnumNameContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageOrEnumNameContext) Ident() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIdent, 0)
}

func (s *MessageOrEnumNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageOrEnumNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageOrEnumNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMessageOrEnumName(s)
	}
}

func (s *MessageOrEnumNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMessageOrEnumName(s)
	}
}

func (p *Protobuf3Parser) MessageOrEnumName() (localctx IMessageOrEnumNameContext) {
	localctx = NewMessageOrEnumNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, Protobuf3ParserRULE_messageOrEnumName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(353)
		p.Match(Protobuf3ParserIdent)
	}

	return localctx
}

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldName
	return p
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) Ident() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIdent, 0)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterFieldName(s)
	}
}

func (s *FieldNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitFieldName(s)
	}
}

func (p *Protobuf3Parser) FieldName() (localctx IFieldNameContext) {
	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, Protobuf3ParserRULE_fieldName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		p.Match(Protobuf3ParserIdent)
	}

	return localctx
}

// IOneofNameContext is an interface to support dynamic dispatch.
type IOneofNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOneofNameContext differentiates from other interfaces.
	IsOneofNameContext()
}

type OneofNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneofNameContext() *OneofNameContext {
	var p = new(OneofNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_oneofName
	return p
}

func (*OneofNameContext) IsOneofNameContext() {}

func NewOneofNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofNameContext {
	var p = new(OneofNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_oneofName

	return p
}

func (s *OneofNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OneofNameContext) Ident() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIdent, 0)
}

func (s *OneofNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneofNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterOneofName(s)
	}
}

func (s *OneofNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitOneofName(s)
	}
}

func (p *Protobuf3Parser) OneofName() (localctx IOneofNameContext) {
	localctx = NewOneofNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, Protobuf3ParserRULE_oneofName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(357)
		p.Match(Protobuf3ParserIdent)
	}

	return localctx
}

// IMapNameContext is an interface to support dynamic dispatch.
type IMapNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapNameContext differentiates from other interfaces.
	IsMapNameContext()
}

type MapNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapNameContext() *MapNameContext {
	var p = new(MapNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_mapName
	return p
}

func (*MapNameContext) IsMapNameContext() {}

func NewMapNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapNameContext {
	var p = new(MapNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_mapName

	return p
}

func (s *MapNameContext) GetParser() antlr.Parser { return s.parser }

func (s *MapNameContext) Ident() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIdent, 0)
}

func (s *MapNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMapName(s)
	}
}

func (s *MapNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMapName(s)
	}
}

func (p *Protobuf3Parser) MapName() (localctx IMapNameContext) {
	localctx = NewMapNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, Protobuf3ParserRULE_mapName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(359)
		p.Match(Protobuf3ParserIdent)
	}

	return localctx
}

// IServiceNameContext is an interface to support dynamic dispatch.
type IServiceNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServiceNameContext differentiates from other interfaces.
	IsServiceNameContext()
}

type ServiceNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceNameContext() *ServiceNameContext {
	var p = new(ServiceNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_serviceName
	return p
}

func (*ServiceNameContext) IsServiceNameContext() {}

func NewServiceNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceNameContext {
	var p = new(ServiceNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_serviceName

	return p
}

func (s *ServiceNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceNameContext) Ident() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIdent, 0)
}

func (s *ServiceNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterServiceName(s)
	}
}

func (s *ServiceNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitServiceName(s)
	}
}

func (p *Protobuf3Parser) ServiceName() (localctx IServiceNameContext) {
	localctx = NewServiceNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, Protobuf3ParserRULE_serviceName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(361)
		p.Match(Protobuf3ParserIdent)
	}

	return localctx
}

// IRpcNameContext is an interface to support dynamic dispatch.
type IRpcNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRpcNameContext differentiates from other interfaces.
	IsRpcNameContext()
}

type RpcNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcNameContext() *RpcNameContext {
	var p = new(RpcNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_rpcName
	return p
}

func (*RpcNameContext) IsRpcNameContext() {}

func NewRpcNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcNameContext {
	var p = new(RpcNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_rpcName

	return p
}

func (s *RpcNameContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcNameContext) Ident() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIdent, 0)
}

func (s *RpcNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterRpcName(s)
	}
}

func (s *RpcNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitRpcName(s)
	}
}

func (p *Protobuf3Parser) RpcName() (localctx IRpcNameContext) {
	localctx = NewRpcNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, Protobuf3ParserRULE_rpcName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		p.Match(Protobuf3ParserIdent)
	}

	return localctx
}

// IMessageTypeContext is an interface to support dynamic dispatch.
type IMessageTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageTypeContext differentiates from other interfaces.
	IsMessageTypeContext()
}

type MessageTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageTypeContext() *MessageTypeContext {
	var p = new(MessageTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageType
	return p
}

func (*MessageTypeContext) IsMessageTypeContext() {}

func NewMessageTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageTypeContext {
	var p = new(MessageTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_messageType

	return p
}

func (s *MessageTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageTypeContext) MessageName() IMessageNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageNameContext)
}

func (s *MessageTypeContext) AllIdent() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserIdent)
}

func (s *MessageTypeContext) Ident(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIdent, i)
}

func (s *MessageTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMessageType(s)
	}
}

func (s *MessageTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMessageType(s)
	}
}

func (p *Protobuf3Parser) MessageType() (localctx IMessageTypeContext) {
	localctx = NewMessageTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, Protobuf3ParserRULE_messageType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserDOT {
		{
			p.SetState(365)
			p.Match(Protobuf3ParserDOT)
		}

	}
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(368)
				p.Match(Protobuf3ParserIdent)
			}
			{
				p.SetState(369)
				p.Match(Protobuf3ParserDOT)
			}

		}
		p.SetState(374)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())
	}
	{
		p.SetState(375)
		p.MessageName()
	}

	return localctx
}

// IMessageOrEnumTypeContext is an interface to support dynamic dispatch.
type IMessageOrEnumTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageOrEnumTypeContext differentiates from other interfaces.
	IsMessageOrEnumTypeContext()
}

type MessageOrEnumTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageOrEnumTypeContext() *MessageOrEnumTypeContext {
	var p = new(MessageOrEnumTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageOrEnumType
	return p
}

func (*MessageOrEnumTypeContext) IsMessageOrEnumTypeContext() {}

func NewMessageOrEnumTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageOrEnumTypeContext {
	var p = new(MessageOrEnumTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_messageOrEnumType

	return p
}

func (s *MessageOrEnumTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageOrEnumTypeContext) MessageOrEnumName() IMessageOrEnumNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageOrEnumNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageOrEnumNameContext)
}

func (s *MessageOrEnumTypeContext) AllIdent() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserIdent)
}

func (s *MessageOrEnumTypeContext) Ident(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIdent, i)
}

func (s *MessageOrEnumTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageOrEnumTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageOrEnumTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMessageOrEnumType(s)
	}
}

func (s *MessageOrEnumTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMessageOrEnumType(s)
	}
}

func (p *Protobuf3Parser) MessageOrEnumType() (localctx IMessageOrEnumTypeContext) {
	localctx = NewMessageOrEnumTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, Protobuf3ParserRULE_messageOrEnumType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserDOT {
		{
			p.SetState(377)
			p.Match(Protobuf3ParserDOT)
		}

	}
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(380)
				p.Match(Protobuf3ParserIdent)
			}
			{
				p.SetState(381)
				p.Match(Protobuf3ParserDOT)
			}

		}
		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())
	}
	{
		p.SetState(387)
		p.MessageOrEnumName()
	}

	return localctx
}

// IEmptyStatementContext is an interface to support dynamic dispatch.
type IEmptyStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEmptyStatementContext differentiates from other interfaces.
	IsEmptyStatementContext()
}

type EmptyStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmptyStatementContext() *EmptyStatementContext {
	var p = new(EmptyStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_emptyStatement
	return p
}

func (*EmptyStatementContext) IsEmptyStatementContext() {}

func NewEmptyStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyStatementContext {
	var p = new(EmptyStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_emptyStatement

	return p
}

func (s *EmptyStatementContext) GetParser() antlr.Parser { return s.parser }
func (s *EmptyStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptyStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEmptyStatement(s)
	}
}

func (s *EmptyStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEmptyStatement(s)
	}
}

func (p *Protobuf3Parser) EmptyStatement() (localctx IEmptyStatementContext) {
	localctx = NewEmptyStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, Protobuf3ParserRULE_emptyStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(389)
		p.Match(Protobuf3ParserSEMI)
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) FullIdent() IFullIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *ConstantContext) IntLit() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIntLit, 0)
}

func (s *ConstantContext) FloatLit() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFloatLit, 0)
}

func (s *ConstantContext) StrLit() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserStrLit, 0)
}

func (s *ConstantContext) BoolLit() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserBoolLit, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *Protobuf3Parser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, Protobuf3ParserRULE_constant)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(401)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(391)
			p.FullIdent()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(393)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == Protobuf3ParserMINUS || _la == Protobuf3ParserPLUS {
			p.SetState(392)
			_la = p.GetTokenStream().LA(1)

			if !(_la == Protobuf3ParserMINUS || _la == Protobuf3ParserPLUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}

		}
		{
			p.SetState(395)
			p.Match(Protobuf3ParserIntLit)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(397)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == Protobuf3ParserMINUS || _la == Protobuf3ParserPLUS {
			p.SetState(396)
			_la = p.GetTokenStream().LA(1)

			if !(_la == Protobuf3ParserMINUS || _la == Protobuf3ParserPLUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}

		}
		{
			p.SetState(399)
			p.Match(Protobuf3ParserFloatLit)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(400)
		_la = p.GetTokenStream().LA(1)

		if !(_la == Protobuf3ParserBoolLit || _la == Protobuf3ParserStrLit) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	}

	return localctx
}
