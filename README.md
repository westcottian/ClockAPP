/*    Clockapp Documentation   */

******************Problem Statement***************

Create a clock application that will print the following values at the following intervals to stdout:
- "tick" every second
- "tock" every minute
- "bong" every hour
 
Only one value should be printed in a given second, i.e. when printing "bong" on the hour, the "tick" and "tock" values should not be printed.
 
It should run for three hours and then exit.
 
A mechanism should exist for the user to alter any of the printed values while the program is running, i.e. after the clock has run for 10 minutes I should, without stopping the program, be able to change it so that it stops printing "tick" every second and starts printing "quack" instead.


**************************************************

*******************Installation******************

-Untar the project under the directory $GOPATH/src/. 
-Go to the dir clockapp: cd clockapp 
-Compile using below command: $ go build

It will create executable binary named clockapp.

*************************************************

*******************Usage*************************

-Run the application: ./clockapp
-It will print the values as described in timer.env file for every second, minute and hour.
- Use the timer.env file in clockapp/ dir to modfiy the string which you want to print at runtime.

*************************************************

*******************Test*************************

-To test the application run command:  go test

************************************************

