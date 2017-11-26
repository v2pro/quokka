package runtime

import (
	"fmt"
	goruntime "runtime"
	"strings"
)

const jsExtraHead = 15

func ReportError(jsSrc string, jsOffsets []int, goSrc string, goOffsets []int, recovered interface{}) {
	_, targetFile, _, _ := goruntime.Caller(1)
	pc := make([]uintptr, 256)
	pc = pc[:goruntime.Callers(2, pc)]
	frames := goruntime.CallersFrames(pc)
	line := 0
	for {
		frame, more := frames.Next()
		if frame.File == targetFile {
			line = frame.Line
		}
		if !more {
			break
		}
	}
	targetLine := line - 8
	targetGoSrc, targetStart, targetEnd := searchLine(goSrc, targetLine)
	if targetGoSrc == "" {
		panic(recovered)
	}
	startIndex, endIndex := searchGoOffsets(goOffsets, targetStart, targetEnd)
	jsStart := jsOffsets[startIndex] - jsExtraHead
	var targetJsSrc string
	if endIndex == -1 {
		targetJsSrc = jsSrc[jsStart:]
	} else {
		jsEnd := jsOffsets[endIndex] - jsExtraHead
		targetJsSrc = jsSrc[jsStart:jsEnd]
	}
	targetJsSrc = strings.TrimSpace(targetJsSrc)
	panic(fmt.Sprintf("javascript source:\n%s\n===\ngo source:\n%s\n===\nerror: %v", targetJsSrc, targetGoSrc, recovered))
}

func searchLine(str string, targetLine int) (string, int, int) {
	tmp := str
	var offset int
	var start int
	for i := 0; i < targetLine; i++ {
		start = strings.IndexByte(tmp, '\n')
		if start == -1 {
			return "", 0, 0
		}
		tmp = tmp[start+1:]
		offset += start + 1
	}
	end := strings.IndexByte(tmp, '\n')
	if end == -1 {
		end = len(tmp)
	}
	return tmp[:end], offset, offset + end
}

func searchGoOffsets(goOffsets []int, targetStart int, targetEnd int) (int, int) {
	var startFound bool
	var startIndex int
	for i := 0; i < len(goOffsets)-1; i++ {
		if targetStart >= goOffsets[i] && targetStart < goOffsets[i+1] {
			startFound = true
			startIndex = i
		}
		if startFound && goOffsets[i] >= targetEnd {
			return startIndex, i
		}
	}
	if startFound {
		return startIndex, -1
	}
	return len(goOffsets) - 1, -1
}
