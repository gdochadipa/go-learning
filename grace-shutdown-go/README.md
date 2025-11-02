

# Graceful Shutdown on go

there are several things make a program shutdown. but on go some time when you shutdown a process, there a some code are still running, while the main goroutine are shutdown but another process still running and hanging.

so some you must terminate any process when you recive terminate signal from os.
