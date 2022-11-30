# recruitment-exercise-golang

## Before Starting:
### IMPORTANT: 
Copy this project into your personal git account and commit before making any changes. DO NOT commit your changes in the first commit. The idea with this
is seeing your progress in the history.

### Project Description:
This application emulates a factory that assembles cars. This is a small factory with 5 spots only, meaning it could assemble up to 5 cars at the same time. 
If all spots are busy at a certain moment, no new cars can be assembled until spots get idle. Every single car part takes one second to be assembled so the 
total assemble time could take up to 7s. if you run this app now you will see a console message "Assembling vehicle..." every 7 seconds. Meaning every 7
seconds a new car is assembled, we should assemble 5 cars every second.

### REQUIREMENTS:
This application is syncronous, make it work CONCURRENTLY! You can use goroutines, channels/buffered channels, waitgroups, mutexes, etc at your discression.

Before Starting:
Name the module as "valverdethiago/car-factory-threads", generate the sume and vendor

This project has several go files, YOU MUST FOCUS on these 3 functions, create one commite per function:
1. COMMIT #1: assemblyspot.AssembleVehicle: Make the appropriate changes here to process all assembling parts concurrently. Once done, make one commit with this changes.
2. COMMIT #2: factory.StartAssemblingProcess: Execute each vehicle assembly concurrently, having in mind that the factory has 5 spots only. Once done, make one commit with this changes.
3. COMMIT #3: main.StartAssemblingProcess: Receive each vehicle and display the assembly and testing logs concurrently. Once done, make one commit with this changes.
4. COMMIT #4 (Bonus): If you have time Implement Unit Test (please use factory_test.go as base file)
Have in mind: Each function has a hint comment, read it carefully and if you have any doubt please don't hesitate to ask.

### TO HAVE IN MIND
racing conditions.

At the end of the assemble process of each vehicle, main.go should receive the vehicle and display its TestingLog and AssembleLog. Please have in mind do not wait for all
vehicles to be done to return them all to main, once each single vehicle is assembled send it over to main for log display right away
