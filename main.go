package main

import (
    "clockapp/config"
    "time"
    "os"
    "sync"
    "fmt"
)

func init() {

   config.Loadconfiguration()

}

func main() {

    clockApp()
}

func clockApp() {
	// Stop the clock after 3 hours */
    stop := time.After(3 * time.Hour)
    tick := time.NewTicker(1 * time.Second)
    defer tick.Stop()
	
	/* Adding waitgroup */ 
    wg := &sync.WaitGroup{}
    wg.Add(1)
    go func() {
    for {
        select {
        case <-tick.C:
			/*  Logic to print min, hr and sec */
		fmt.Println(ClockLogic(time.Now()))
        case <-stop:
		// Timeout
        	os.Exit(1)
        }
    }
   }()
   wg.Wait()
}


func ClockLogic(Time time.Time) string{
	// Hour
        if(Time.Second() == 59 && Time.Minute() == 59) {
		return config.Hrstr
        } else {
	//Minute
        if(Time.Second() == 59) {
		return config.Minstr
        } else {
		// Second
		return config.Secstr
        }
     }
}

