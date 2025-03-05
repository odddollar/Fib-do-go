# Fib Dot Go

Fib Dot Go is a multithreaded CPU benchmark that aims to assess processor performance in calculating Fibonacci numbers. There is a queue of tasks that worker threads pull from, where each task consists of recursively calculating a large Fibonacci number. 

For example, with a minimum Fibonacci number of 40, a maximum of 45, and 200 tasks, calculating the 40th Fibonacci number is pushed onto a queue. Calculating the 41st number is pushed next, then the 42nd, etc., until the 44th number is pushed. The 40th number is then pushed again and the loop repeats until a total of 200 tasks exist in the queue. The queue would appear as [40 41 42 43 44 40 41 42 43 ...]. Worker threads pull the next available task from this queue and calculate that Fibonacci number. Once this task has been completed, the worker pulls the next task until all tasks have been pulled from the queue. A single shared queue is used by all workers. On exhaustion of all tasks, the time taken to complete the benchmark is reported in seconds.

## Usage

Fib Dot Go has the sub-commands `run`, `about`, and `cpu`.

### 

### Script

A script can be used to run the benchmark with multiple input parameters, specifying the `-m` flag to the `run` command to allow for easy logging of the result to a file. The example script shown below can be used to determine at which point increasing the number of worker threads doesn't result in any speed improvements. The list of arguments can be further modified to test other variations of arguments.

```bat
@echo off
set FILE=output.csv

echo Started: %time%

:: Check if file exists
if not exist %FILE% (
    echo tasks,workers,min,max,time >> %FILE%
)

:: Define list of arguments
set TASKS=200
set WORKERS=5 10 15 20 25 30 35 40
set MIN=40
set MAX=45

:: Loop over all argument sets
for %%T in (%TASKS%) do (
    for %%W in (%WORKERS%) do (
        for %%N in (%MIN%) do (
            for %%X in (%MAX%) do (
                :: Run command with current arguments
                for /f "delims=" %%i in ('fib-dot-go run -t %%T -w %%W -n %%N -x %%X -m') do (
                    echo %%T,%%W,%%N,%%X,%%i >> %FILE%
                    echo Done %%T,%%W,%%N,%%X,%%i
                )
            )
        )
    )
)

echo Completed: %time%
pause
```

## Building

