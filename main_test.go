package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"clockapp/config"
	"os"
	"strings"
        "log"
	"bufio"
)

type timerconfig map[string]string

func TestClockLogic(t *testing.T) {
	createdAt := time.Now()
        rounded := time.Date(createdAt.Year(), createdAt.Month(), createdAt.Day(), 0, 0, 10, 0, createdAt.Location())

	outtimestr := ClockLogic(rounded)

	timerenv, err := readEnvFile("timer.env")
    	if err != nil {
        	t.Error("Error while reading timer.env file")
    	}
        config.Loadconfiguration()

	assert.Equal(t, timerenv["second"],outtimestr)
	assert := assert.New(t)
	assert.NotEqual(timerenv["minute"],outtimestr)
	assert.NotEqual(timerenv["hour"],outtimestr)

	rounded = time.Date(createdAt.Year(), createdAt.Month(), createdAt.Day(), 0, 0, 59, 0, createdAt.Location())

	outtimestr = ClockLogic(rounded)

	assert.NotEqual(timerenv["second"],outtimestr)
        assert.Equal(timerenv["minute"],outtimestr)
        assert.NotEqual(timerenv["hour"],outtimestr)

	rounded = time.Date(createdAt.Year(), createdAt.Month(), createdAt.Day(), 0, 59, 59, 0, createdAt.Location())

        outtimestr = ClockLogic(rounded)

        assert.NotEqual(timerenv["second"],outtimestr)
        assert.NotEqual(timerenv["minute"],outtimestr)
        assert.Equal(timerenv["hour"],outtimestr)	
}

func readEnvFile(filename string) (timerconfig, error) {
    config := timerconfig{}

    if len(filename) == 0 {
        return config, nil
    }
    file, err := os.Open(filename)
    if err != nil {
        log.Println("Error here while opening")
        log.Fatal(err)
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if equal := strings.Index(line, "="); equal >= 0 {
            if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
                value := ""
                if len(line) > equal {
                    value = strings.TrimSpace(line[equal+1:])
		    strings.Trim(value, "\"") 
                }
		value = strings.TrimRight(value, "\r\"")
		value = strings.TrimLeft(value, "\r\"")
                config[key] = value
            }
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
        return nil, err
    }

    return config, nil
}




