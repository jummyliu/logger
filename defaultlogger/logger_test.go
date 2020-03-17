package defaultlogger

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/jummyliu/logger"
)

type unitGroup struct {
	level logger.Level
	units []unit
}

type unit struct {
	level  logger.Level
	input  string
	output string
}

var testArr = []unitGroup{
	unitGroup{
		level: logger.LevelDebug,
		units: []unit{
			unit{
				level:  logger.LevelDebug,
				input:  "test debug",
				output: fmt.Sprintf("| [%s]\t| %s\n", logger.LogNameMap[logger.LevelDebug], "test debug"),
			},
			unit{
				level:  logger.LevelInfo,
				input:  "test info",
				output: fmt.Sprintf("| [%s]\t| %s\n", logger.LogNameMap[logger.LevelInfo], "test info"),
			},
		},
	},
	unitGroup{
		level: logger.LevelInfo,
		units: []unit{
			unit{
				level:  logger.LevelDebug,
				input:  "test debug",
				output: "",
			},
			unit{
				level:  logger.LevelInfo,
				input:  "test info",
				output: fmt.Sprintf("| [%s]\t| %s\n", logger.LogNameMap[logger.LevelInfo], "test info"),
			},
		},
	},
}

func Test(t *testing.T) {
	b := make([]byte, 1024)
	buf := bytes.NewBuffer(b)
	for _, group := range testArr {
		log := New(buf, group.level)

		for _, u := range group.units {
			log.Log(u.level, u.input)
			output := buf.String()
			if u.output == "" {
				if output != u.output {
					t.Fatalf("test failure: '%s' != '%s'", output, u.output)
				}
			} else {
				if !strings.Contains(output, u.output) {
					t.Fatalf("test failure: '%s' != '%s'", output, u.output)
				}
			}

			buf.Reset()
		}
	}
	t.Log("pass")
}
