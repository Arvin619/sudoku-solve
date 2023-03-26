package sudoku

import (
	"bytes"
	"errors"
	"io"
	"log"
	"strings"
)

const (
	unit = 9
)

type sudokuSolve struct {
	table     [][]byte
	spaceChar byte
}

func New() *sudokuSolve {
	result := &sudokuSolve{
		table:     make([][]byte, unit),
		spaceChar: '.',
	}
	for i := range result.table {
		result.table[i] = make([]byte, unit)
	}
	return result
}

func (ss *sudokuSolve) SetTableWithReader(r io.Reader) {
	var b bytes.Buffer

	bs := make([]byte, 512)

	for {
		n, err := r.Read(bs)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				panic(err)
			}
		}
		b.Write(bs[:n])
	}

	ss.bytesToTable(b.Bytes())
}

func (ss *sudokuSolve) bytesToTable(b []byte) {
	index := 0

	for _, v := range b {
		if v == ss.spaceChar || (v >= byte('0') && v <= byte('9')) {
			ss.table[index/unit][index%unit] = v
			index++
		}
	}

	if index != unit*unit {
		log.Panicf("string length is %v, not %v", index, unit*unit)
	}

}

func (ss *sudokuSolve) Solve() bool {
	for i := 0; i < unit; i++ {
		for j := 0; j < unit; j++ {
			if ss.table[i][j] != ss.spaceChar {
				continue
			}
			for k := byte('1'); k <= byte('9'); k++ {
				if ss.isValid(i, j, k) {
					ss.table[i][j] = k
					if ss.Solve() {
						return true
					}

					ss.table[i][j] = ss.spaceChar
				}
			}
			return false
		}
	}
	return true
}

func (ss *sudokuSolve) isValid(i, j int, k byte) bool {
	for u := 0; u < unit; u++ {
		if (ss.table[i][u] == k) || (ss.table[u][j] == k) {
			return false
		}
	}
	row := (i / 3) * 3
	col := (j / 3) * 3
	for r := row; r < row+3; r++ {
		for c := col; c < col+3; c++ {
			if ss.table[r][c] == k {
				return false
			}
		}
	}
	return true
}

func (ss *sudokuSolve) String() string {
	var sb strings.Builder
	for i := 0; i < unit; i++ {
		if i%3 == 0 {
			sb.WriteString("+-------+-------+-------+\n")
		}
		for j := 0; j < unit; j++ {
			if j%3 == 0 {
				sb.WriteString("| ")
			}
			sb.WriteByte(ss.table[i][j])
			sb.WriteByte(' ')
		}
		sb.WriteString("|\n")
	}
	sb.WriteString("+-------+-------+-------+")
	return sb.String()
}
